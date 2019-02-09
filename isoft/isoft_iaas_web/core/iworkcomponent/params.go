package iworkcomponent

import (
	"encoding/xml"
	"isoft/isoft_iaas_web/core/iworkdata"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
)

// 获取缓存的出参 schema,即从 DB 中读取
func GetCacheParamOutputSchema(step *iwork.WorkStep) *iworkdata.ParamOutputSchema {
	// 从缓存(数据库字段)中获取
	if strings.TrimSpace(step.WorkStepOutput) != "" {
		var paramOutputSchema *iworkdata.ParamOutputSchema
		if err := xml.Unmarshal([]byte(step.WorkStepOutput), &paramOutputSchema); err == nil {
			return paramOutputSchema
		}
	}
	return &iworkdata.ParamOutputSchema{}
}

// 获取出参 schema
func GetRuntimeParamOutputSchema(step *iwork.WorkStep) *iworkdata.ParamOutputSchema {
	// 获取当前 work_step 对应的 paramOutputSchema
	factory := &WorkStepFactory{WorkStep: step}
	paramOutputSchema := factory.GetDefaultParamOutputSchema()
	paramOutputSchema2 := factory.GetRuntimeParamOutputSchema()
	// 合并
	paramOutputSchema.ParamOutputSchemaItems =
		append(paramOutputSchema.ParamOutputSchemaItems, paramOutputSchema2.ParamOutputSchemaItems...)
	return paramOutputSchema
}

// 获取入参 schema
func GetCacheParamInputSchema(step *iwork.WorkStep) *iworkdata.ParamInputSchema {
	// 从缓存(数据库字段)中获取
	if strings.TrimSpace(step.WorkStepInput) != "" {
		var paramInputSchema *iworkdata.ParamInputSchema
		if err := xml.Unmarshal([]byte(step.WorkStepInput), &paramInputSchema); err == nil {
			return paramInputSchema
		}
	}
	// 获取当前 work_step 对应的 paramInputSchema
	factory := &WorkStepFactory{WorkStep: step}
	paramInputSchema := factory.GetDefaultParamInputSchema()
	return paramInputSchema
}



