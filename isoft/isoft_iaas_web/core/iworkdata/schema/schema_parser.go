package schema

import (
	"encoding/xml"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
)

// 获取缓存的出参 schema,即从 DB 中读取
func GetCacheParamOutputSchema(step *iwork.WorkStep) *ParamOutputSchema {
	// 从缓存(数据库字段)中获取
	if strings.TrimSpace(step.WorkStepOutput) != "" {
		var paramOutputSchema *ParamOutputSchema
		if err := xml.Unmarshal([]byte(step.WorkStepOutput), &paramOutputSchema); err == nil {
			return paramOutputSchema
		}
	}
	return &ParamOutputSchema{}
}

// 使用接口来解决循环依赖问题
// 使用接口的好处是传入实现类,而不是引用实现类并创建实例（这样就解决了引用导致的循环依赖问题）
type IParamSchemaParser interface {
	GetDefaultParamOutputSchema() *ParamOutputSchema
	GetRuntimeParamOutputSchema() *ParamOutputSchema
	GetDefaultParamInputSchema() *ParamInputSchema
}

// 获取出参 schema
func GetRuntimeParamOutputSchema(paramSchemaParser IParamSchemaParser) *ParamOutputSchema {
	paramOutputSchema := paramSchemaParser.GetDefaultParamOutputSchema()
	paramOutputSchema2 := paramSchemaParser.GetRuntimeParamOutputSchema()
	// 合并
	paramOutputSchema.ParamOutputSchemaItems =
		append(paramOutputSchema.ParamOutputSchemaItems, paramOutputSchema2.ParamOutputSchemaItems...)
	return paramOutputSchema
}

// 获取入参 schema
func GetCacheParamInputSchema(step *iwork.WorkStep, paramSchemaParser IParamSchemaParser) *ParamInputSchema {
	// 从缓存(数据库字段)中获取
	if strings.TrimSpace(step.WorkStepInput) != "" {
		var paramInputSchema *ParamInputSchema
		if err := xml.Unmarshal([]byte(step.WorkStepInput), &paramInputSchema); err == nil {
			return paramInputSchema
		}
	}
	// 获取当前 work_step 对应的 paramInputSchema
	paramInputSchema := paramSchemaParser.GetDefaultParamInputSchema()
	return paramInputSchema
}



