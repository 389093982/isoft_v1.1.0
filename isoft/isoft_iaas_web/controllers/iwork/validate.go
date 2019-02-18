package iwork

import (
	"fmt"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworknode"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
)

func (this *WorkController) ValidateAllWork()  {
	defer func() {
		if err := recover(); err != nil{
			fmt.Println(err)
		}
	}()
	works := iwork.GetAllWorkInfo()
	for _, work := range works{
		go func(work iwork.Work) {
			validateWork(&work)
		}(work)
	}
	this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	this.ServeJSON()
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
		fmt.Println("check param :" + item.ParamName)
		if !strings.HasSuffix(item.ParamName,"?") && strings.TrimSpace(item.ParamValue) == ""{
			fmt.Println("found empty param......................")
		}
	}
}