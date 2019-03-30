package iwork

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"isoft/isoft/common/stringutil"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworkplugin/iworknode"
	"isoft/isoft_iaas_web/core/iworkutil"
	"isoft/isoft_iaas_web/core/iworkutil/errorutil"
	"isoft/isoft_iaas_web/core/iworkvalid"
	"isoft/isoft_iaas_web/models/iwork"
	"isoft/isoft_iaas_web/service"
	"isoft/isoft_iaas_web/service/iworkservice"
	"strings"
	"time"
)

func (this *WorkController) LoadValidateResult() {
	if result, err := service.ExecuteResultServiceWithTx(map[string]interface{}{}, iworkservice.LoadValidateResultService); err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "details": result["details"]}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

func (this *WorkController) ValidateAllWork() {
	validateAll()
	this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	this.ServeJSON()
}

// 统计操作所花费的时间方法
func recordCostTimeLog(trackingId string, start time.Time) {
	iwork.InsertValidateLogDetail(getValidateLogDetail(trackingId,
		fmt.Sprintf("validate complete! total cost %d ms!", time.Now().Sub(start).Nanoseconds()/1e6)))
}

func validateAll() {
	trackingId := stringutil.RandomUUID()
	// 记录校验耗费时间
	defer recordCostTimeLog(trackingId, time.Now())
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
	works := iwork.QueryAllWorkInfo(orm.NewOrm())

	for _, work := range works {
		go func(work iwork.Work) {
			validateWork(&work, logCh, workChan)
		}(work)
	}

	go func() {
		for i := 0; i < len(works); i++ {
			<-workChan
		}
		// 所有 work 执行完成后关闭 logCh
		close(logCh)
	}()

	// 从 logCh 中循环读取校验不通过的信息,并将其写入日志表中去
	for log := range logCh {
		work, _ := iwork.QueryWorkById(log.WorkId, orm.NewOrm())
		step, _ := iwork.QueryOneWorkStep(work.Id, log.WorkStepId, orm.NewOrm())
		log.TrackingId = trackingId
		log.WorkName = work.WorkName
		log.WorkStepName = step.WorkStepName
		log.CreatedBy = "SYSTEM"
		log.LastUpdatedBy = "SYSTEM"
		log.CreatedTime = time.Now()
		log.LastUpdatedTime = time.Now()
		iwork.InsertValidateLogDetail(log)
	}
}

// 校验单个 work
func validateWork(work *iwork.Work, logCh chan *iwork.ValidateLogDetail, workChan chan int) {
	stepChan := make(chan int)
	steps, _ := iwork.QueryAllWorkStepInfo(work.Id, orm.NewOrm())
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
func validateStep(step *iwork.WorkStep, logCh chan *iwork.ValidateLogDetail, stepChan chan int) {
	defer func() {
		if err := recover(); err != nil {
			if _, ok := err.(error); ok {
				logCh <- &iwork.ValidateLogDetail{
					WorkId:     step.WorkId,
					WorkStepId: step.WorkStepId,
					Detail:     string(errorutil.PanicTrace(4)),
				}
			} else if _err, ok := err.(string); ok {
				logCh <- &iwork.ValidateLogDetail{
					WorkId:     step.WorkId,
					WorkStepId: step.WorkStepId,
					Detail:     _err,
				}
			} else if _err, ok := err.(iwork.ValidateLogDetail); ok {
				logCh <- &_err
			}
		}
		// 每执行完一个 step 就往 stepChan 里面发送完成通知
		stepChan <- 1
	}()

	checkResults := CheckStep(step)
	for _, checkResult := range checkResults {
		logCh <- &iwork.ValidateLogDetail{
			WorkId:     step.WorkId,
			WorkStepId: step.WorkStepId,
			Detail:     checkResult,
		}
	}
}

func CheckStep(step *iwork.WorkStep) (checkResult []string) {
	// 校验 step 中的参数是否为空
	checkResults1 := iworkvalid.CheckEmpty(step, &iworknode.WorkStepFactory{WorkStep: step})
	checkResults2 := checkVariableRelationShip(step)
	// 定制化校验
	checkResults3 := CheckCustom(step)
	checkResult = append(checkResult, checkResults1...)
	checkResult = append(checkResult, checkResults2...)
	checkResult = append(checkResult, checkResults3...)
	return
}

func CheckCustom(step *iwork.WorkStep) (checkResult []string) {
	factory := &iworknode.WorkStepFactory{WorkStep: step}
	return factory.ValidateCustom()
}

// 校验变量的引用关系
func checkVariableRelationShip(step *iwork.WorkStep) (checkResult []string) {
	inputSchema := schema.GetCacheParamInputSchema(step, &iworknode.WorkStepFactory{WorkStep: step})
	for _, item := range inputSchema.ParamInputSchemaItems {
		result := checkVariableRelationShipDetail(item, step.WorkId, step.WorkStepId)
		checkResult = append(checkResult, result...)
	}
	return
}

func checkVariableRelationShipDetail(item schema.ParamInputSchemaItem, work_id, work_step_id int64) (checkResult []string) {
	// 根据正则找到关联的节点名和字段名
	refers := iworkutil.GetRelativeValueWithReg(item.ParamValue)
	if len(refers) == 0 {
		return
	}
	preStepNodeNames := iworkutil.GetAllPreStepNodeName(work_id, work_step_id)
	skipNodeNames := []string{"RESOURCE", "WORK"}
	for _, refer := range refers {
		referNodeName := refer[1:strings.Index(refer, ".")]
		referFileName := refer[strings.Index(refer, ".")+1:]
		// 非节点类型直接跳过
		if stringutil.CheckContains(referNodeName, skipNodeNames) {
			break
		}
		// 判断节点名称是否有效
		if !stringutil.CheckContains(referNodeName, preStepNodeNames) {
			checkResult = append(checkResult, fmt.Sprintf("Invalid referNodeName relationship for %s was found!", referNodeName))
			continue
		}
		// 判断字段名称是否有效
		if step, err := iwork.QueryWorkStepByStepName(work_id, referNodeName, orm.NewOrm()); err == nil {
			outputSchema := schema.GetCacheParamOutputSchema(&step)
			exist := false
			for _, item := range outputSchema.ParamOutputSchemaItems {
				if item.ParamName == referFileName {
					exist = true
					break
				}
			}
			if !exist {
				checkResult = append(checkResult, fmt.Sprintf("Invalid referFileName relationship for %s was found!", referFileName))
			}
		}
	}
	return
}

func getValidateLogDetail(trackingId, detail string) *iwork.ValidateLogDetail {
	return &iwork.ValidateLogDetail{
		TrackingId:      trackingId,
		Detail:          detail,
		CreatedBy:       "SYSTEM",
		LastUpdatedBy:   "SYSTEM",
		CreatedTime:     time.Now(),
		LastUpdatedTime: time.Now(),
	}
}
