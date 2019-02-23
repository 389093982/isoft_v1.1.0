package iwork

import (
	"fmt"
	"isoft/isoft/common/stringutil"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworknode"
	"isoft/isoft_iaas_web/core/iworkvalid"
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

	logCh := make(chan *iwork.ValidateLogDetail)
	workChan := make(chan int)
	works := iwork.QueryAllWorkInfo()

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
		step, _ := iwork.QueryOneWorkStep(work.Id, log.WorkStepId)
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
	steps, _ := iwork.QueryAllWorkStepInfo(work.Id)
	// 验证流程必须以 work_start 开始,以 work_end 结束
	validateWorkStartAndEnd(steps, logCh, work)

	for _, step := range steps {
		go func(step iwork.WorkStep) {
			validateStep(&step, logCh, stepChan)
		}(step)
	}

	for i := 0; i < len(steps); i++ {
		<-stepChan
	}
	// 所有 step 执行完成后就往 workChan 里面发送完成通知
	workChan <- 1
}

func validateWorkStartAndEnd(steps []iwork.WorkStep, logCh chan *iwork.ValidateLogDetail, work *iwork.Work) {
	if steps[0].WorkStepType != "work_start" {
		logCh <- &iwork.ValidateLogDetail{
			WorkId:     work.Id,
			WorkStepId: steps[0].WorkStepId,
			Detail:     "work must start with a work_start node!",
		}
	}
	if steps[len(steps)-1].WorkStepType != "work_end" {
		logCh <- &iwork.ValidateLogDetail{
			WorkId:     work.Id,
			WorkStepId: steps[len(steps)-1].WorkStepId,
			Detail:     "work must end with a work_end node!",
		}
	}
	return
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
	CheckGeneral(step, logCh)
	// 定制化校验
	CheckCustom(step)
}

func CheckGeneral(step *iwork.WorkStep,logCh chan *iwork.ValidateLogDetail)  {
	// 校验 step 中的参数是否为空
	iworkvalid.CheckEmpty(step, &iworknode.WorkStepFactory{WorkStep:step})
	checkVariableRelationShip(step, logCh)
}

func CheckCustom(step *iwork.WorkStep)  {
	factory := &iworknode.WorkStepFactory{WorkStep:step}
	factory.ValidateCustom()
}

// 校验变量的引用关系
func checkVariableRelationShip(step *iwork.WorkStep, logCh chan *iwork.ValidateLogDetail)  {
	inputSchema := schema.GetCacheParamInputSchema(step, &iworknode.WorkStepFactory{WorkStep:step})
	for _, item := range inputSchema.ParamInputSchemaItems{
		checkVariableRelationShipDetail(item, step.WorkId, step.WorkStepId, logCh)
	}
}

func checkVariableRelationShipDetail(item schema.ParamInputSchemaItem,work_id, work_step_id int64, logCh chan *iwork.ValidateLogDetail)  {
	// 根据正则找到关联的节点名称
	referNodeNames := stringutil.GetNoRepeatSubStringWithRegexp(item.ParamValue, `\$[a-zA-Z0-9_]+`)
	if len(referNodeNames) == 0{
		return
	}
	preStepNodeNames := getAllPreStepNodeName(work_id, work_step_id)
	preStepNodeNames = append(preStepNodeNames, []string{"RESOURCE"}...)
	for _, referNodeName := range referNodeNames{
		if !stringutil.CheckContains(strings.Replace(referNodeName, "$.", "", -1), preStepNodeNames){
			logCh <- &iwork.ValidateLogDetail{
				WorkId:work_id,
				WorkStepId:work_step_id,
				Detail:fmt.Sprintf("Invalid variable relationship for %s was found!", referNodeName),
			}
		}
	}
}

func getAllPreStepNodeName(work_id, work_step_id int64) []string {
	result := make([]string,0)
	steps, err := iwork.QueryAllPreStepInfo(work_id, work_step_id)
	if err == nil{
		for _,step := range steps{
			result = append(result, step.WorkStepName)
		}
	}
	return result
}