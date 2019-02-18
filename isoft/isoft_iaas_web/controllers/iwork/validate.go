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
	defer func() {
		if err := recover(); err != nil{
			if _err,ok := err.(error); ok {
				iwork.InsertValidateLogDetail(trackingId, fmt.Sprintf("internal error:%s", _err.Error()))
			}
		}
	}()
	works := iwork.GetAllWorkInfo()
	for _, work := range works{
		go func(work iwork.Work) {
			validateWork(&work)
		}(work)
	}
}

func validateWork(work *iwork.Work)  {
	steps, _ := iwork.GetAllWorkStepInfo(work.Id)
	for _, step := range steps{
		go func(step iwork.WorkStep) {
			validateStep(&step)
		}(step)
	}
}

func validateStep(step *iwork.WorkStep)  {
	checkEmpty(step)
}

func checkEmpty(step *iwork.WorkStep)  {
	paramInputSchema := schema.GetCacheParamInputSchema(step, &iworknode.WorkStepFactory{WorkStep: step})
	for _, item := range paramInputSchema.ParamInputSchemaItems{
		if !strings.HasSuffix(item.ParamName,"?") && strings.TrimSpace(item.ParamValue) == ""{
			fmt.Println("found empty param......................")
		}
	}
}