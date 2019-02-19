package iworkvalid

import (
	"fmt"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworknode"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
)

// 对必须参数进行非空校验
func CheckEmpty(step *iwork.WorkStep)  {
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
