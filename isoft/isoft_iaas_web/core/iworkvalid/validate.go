package iworkvalid

import (
	"fmt"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
)

// 对必须参数进行非空校验
func CheckEmpty(step *iwork.WorkStep, paramSchemaParser schema.IParamSchemaParser) (checkResult []string) {
	if strings.TrimSpace(step.WorkStepName) == "" || strings.TrimSpace(step.WorkStepType) == "" {
		checkResult = append(checkResult, fmt.Sprintf("Empty workStepName or empty workStepType was found!"))
		return
	}
	paramInputSchema := schema.GetCacheParamInputSchema(step, paramSchemaParser)
	for _, item := range paramInputSchema.ParamInputSchemaItems {
		// work_start 节点参数由调度者提供,不做非空校验
		if step.WorkStepType != "work_start" {
			_, result := CheckEmptyForItem(item)
			checkResult = append(checkResult, result...)
		}
	}
	return
}

// 对输入参数做非空校验
func CheckEmptyForItem(item schema.ParamInputSchemaItem) (ok bool, checkResult []string) {
	if !strings.HasSuffix(item.ParamName, "?") && strings.TrimSpace(item.ParamValue) == "" {
		checkResult = append(checkResult, fmt.Sprintf("Empty paramValue for %s was found!", item.ParamName))
	}
	return len(checkResult) == 0, checkResult
}
