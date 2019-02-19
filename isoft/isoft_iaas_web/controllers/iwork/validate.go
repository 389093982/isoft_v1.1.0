package iwork

import (
	"isoft/isoft/common/stringutil"
	"isoft/isoft_iaas_web/core/iworknode"
	"isoft/isoft_iaas_web/core/iworkvalid"
	"isoft/isoft_iaas_web/models/iwork"
	"time"
)

func (this *WorkController) LoadValidateResult() {
	if details, err := iwork.QueryLastValidateLogDetail(); err == nil{
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "details":details}
	}else{
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func (this *WorkController) ValidateAllWork()  {
	validateAll()
	this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	this.ServeJSON()
}

func validateAll()  {
	trackingId := stringutil.RandomUUID()
	// 记录日志
	iwork.InsertValidateLogRecord(&iwork.ValidateLogRecord{
		TrackingId:      trackingId,
		CreatedBy:       "SYSTEM",
		CreatedTime:     time.Now(),
		LastUpdatedBy:   "SYSTEM",
		LastUpdatedTime: time.Now(),
	})

	logCh := make(chan *iwork.ValidateLogDetail)
	workChan := make(chan int)
	works := iwork.GetAllWorkInfo()

	for _, work := range works{
		go func(work iwork.Work) {
			validateWork(&work, logCh, workChan)
		}(work)
	}

	go func() {
		for i :=0 ;i<len(works); i++{
			<- workChan
		}
		// 所有 work 执行完成后关闭 logCh
		close(logCh)
	}()

	// 从 logCh 中循环读取校验不通过的信息,并将其写入日志表中去
	for log := range logCh{
		work, _ := iwork.QueryWorkById(log.WorkId)
		step, _ := iwork.GetOneWorkStep(work.Id, log.WorkStepId)
		log.TrackingId = trackingId
		log.WorkName = work.WorkName
		log.WorkStepName = step.WorkStepName
		log.CreatedBy = "SYSTEM"
		log.LastUpdatedBy = "SYSTEM"
		log.CreatedTime = time.Now()
		log.LastUpdatedTime = time.Now()
		iwork.InsertValidateLogDetail(log)
	}
	iwork.InsertValidateLogDetail(&iwork.ValidateLogDetail{
		TrackingId:trackingId,
		Detail:"校验完成！",
		CreatedBy:"SYSTEM",
		LastUpdatedBy:"SYSTEM",
		CreatedTime:time.Now(),
		LastUpdatedTime:time.Now(),
	})
}

// 校验单个 work
func validateWork(work *iwork.Work, logCh chan *iwork.ValidateLogDetail, workChan chan int)  {
	stepChan := make(chan int)
	steps, _ := iwork.GetAllWorkStepInfo(work.Id)
	for _, step := range steps{
		go func(step iwork.WorkStep) {
			validateStep(&step, logCh, stepChan)
		}(step)
	}

	for i := 0; i<len(steps); i++{
		<- stepChan
	}
	// 所有 step 执行完成后就往 workChan 里面发送完成通知
	workChan <- 1
}

// 校验单个 step,并将校验不通过的信息放入 logCh 中
func validateStep(step *iwork.WorkStep, logCh chan *iwork.ValidateLogDetail, stepChan chan int)  {
	defer func() {
		if err := recover(); err != nil{
			if _err,ok := err.(error); ok {
				logCh <- &iwork.ValidateLogDetail{
					WorkId:step.WorkId,
					WorkStepId:step.WorkStepId,
					Detail:_err.Error(),
				}
			} else if _err,ok := err.(string); ok {
				logCh <- &iwork.ValidateLogDetail{
					WorkId:step.WorkId,
					WorkStepId:step.WorkStepId,
					Detail:_err,
				}
			} else if _err,ok := err.(iwork.ValidateLogDetail); ok {
				logCh <- &_err
			}
		}
		// 每执行完一个 step 就往 stepChan 里面发送完成通知
		stepChan <- 1
	}()

	// 通用校验
	CheckGeneral(step)
	// 定制化校验
	CheckCustom(step)
}

func CheckGeneral(step *iwork.WorkStep)  {
	// 校验 step 中的参数是否为空
	iworkvalid.CheckEmpty(step, &iworknode.WorkStepFactory{WorkStep:step})
}

func CheckCustom(step *iwork.WorkStep)  {
	factory := &iworknode.WorkStepFactory{WorkStep:step}
	factory.ValidateCustom()
}

