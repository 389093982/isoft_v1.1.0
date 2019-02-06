package iworkdata

import (
	"encoding/xml"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
)

type ParamResolver struct {
	ParamStr string
} 

func (this *ParamResolver) ParseParamStrToMap() *map[string]interface{}{
	return &map[string]interface{}{}
}

// 获取出参 schema
func GetParamOutputSchema(step *iwork.WorkStep) *ParamOutputSchema {
	// 从缓存(数据库字段)中获取
	if strings.TrimSpace(step.WorkStepOutput) != ""{
		var paramOutputSchema *ParamOutputSchema
		if err := xml.Unmarshal([]byte(step.WorkStepOutput), &paramOutputSchema); err == nil{
			return paramOutputSchema
		}
	}
	// 获取当前 work_step 对应的 paramOutputSchema
	helper := &IWorkStepHelper{WorkStep:step}
	paramOutputSchema := helper.GetDefaultParamOutputSchema()
	return paramOutputSchema
}

// 获取入参 schema
func GetParamInputSchema(step *iwork.WorkStep) *ParamInputSchema {
	// 从缓存(数据库字段)中获取
	if strings.TrimSpace(step.WorkStepInput) != ""{
		var paramInputSchema *ParamInputSchema
		if err := xml.Unmarshal([]byte(step.WorkStepInput), &paramInputSchema); err == nil{
			return paramInputSchema
		}
	}
	// 获取当前 work_step 对应的 paramInputSchema
	helper := &IWorkStepHelper{WorkStep:step}
	paramInputSchema := helper.GetDefaultParamInputSchema()
	return paramInputSchema
}