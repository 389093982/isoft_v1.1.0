package iworkvalid

import (
	"fmt"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
)

// 对必须参数进行非空校验
func CheckEmpty(step *iwork.WorkStep, paramSchemaParser schema.IParamSchemaParser) {
	if strings.TrimSpace(step.WorkStepName) == "" || strings.TrimSpace(step.WorkStepType) == ""{
		panic(fmt.Sprintf("Empty workStepName or empty workStepType was found!"))
	}
	paramInputSchema := schema.GetCacheParamInputSchema(step, paramSchemaParser)
	for _, item := range paramInputSchema.ParamInputSchemaItems{
		CheckEmptyForItem(item)
	}
}

// 对输入参数做非空校验
func CheckEmptyForItem(item schema.ParamInputSchemaItem) {
	if !strings.HasSuffix(item.ParamName,"?") && strings.TrimSpace(item.ParamValue) == ""{
		panic(fmt.Sprintf("Empty paramValue for %s was found!", item.ParamName))
	}
}
