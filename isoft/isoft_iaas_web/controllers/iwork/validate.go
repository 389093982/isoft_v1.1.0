package iwork

import (
	"fmt"
	"isoft/isoft/common/stringutil"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworknode"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
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

	logCh := make(chan string)
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
		iwork.InsertValidateLogDetail(trackingId, fmt.Sprintf("internal error:%s", log))
	}
	iwork.InsertValidateLogDetail(trackingId, "校验完成!")
}

// 校验单个 work
func validateWork(work *iwork.Work, logCh chan string, workChan chan int)  {
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
func validateStep(step *iwork.WorkStep, logCh chan string, stepChan chan int)  {
	defer func() {
		if err := recover(); err != nil{
			if _err,ok := err.(error); ok {
				logCh <- _err.Error()
			}
			if _err,ok := err.(string); ok {
				logCh <- _err
			}
		}
		// 每执行完一个 step 就往 stepChan 里面发送完成通知
		stepChan <- 1
	}()
	checkEmpty(step)
}

// 对必须参数进行非空校验
func checkEmpty(step *iwork.WorkStep)  {
	if strings.TrimSpace(step.WorkStepName) == "" || strings.TrimSpace(step.WorkStepType) == ""{
		panic(fmt.Sprintf("[%v-%v]found empty step for %v", step.WorkId, step.WorkStepId, step.WorkStepId))
	}
	paramInputSchema := schema.GetCacheParamInputSchema(step, &iworknode.WorkStepFactory{WorkStep: step})
	for _, item := range paramInputSchema.ParamInputSchemaItems{
		if !strings.HasSuffix(item.ParamName,"?") && strings.TrimSpace(item.ParamValue) == ""{
			panic(fmt.Sprintf("[%v-%v]found empty paramValue from %s", step.WorkId, step.WorkStepId, item.ParamName))
		}
	}
}