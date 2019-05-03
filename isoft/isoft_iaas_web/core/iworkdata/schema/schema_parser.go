package schema

import (
	"encoding/json"
	"isoft/isoft_iaas_web/core/iworkmodels"
	"isoft/isoft_iaas_web/core/iworkplugin/iworkprotocol"
	"isoft/isoft_iaas_web/core/iworkutil/datatypeutil"
	"isoft/isoft_iaas_web/models/iwork"
	"sort"
	"strings"
)

// 获取缓存的出参 schema,即从 DB 中读取
func GetCacheParamOutputSchema(step *iwork.WorkStep) *iworkmodels.ParamOutputSchema {
	// 从缓存(数据库字段)中获取
	if strings.TrimSpace(step.WorkStepOutput) != "" {
		var paramOutputSchema *iworkmodels.ParamOutputSchema
		if err := json.Unmarshal([]byte(step.WorkStepOutput), &paramOutputSchema); err == nil {
			return paramOutputSchema
		}
	}
	return &iworkmodels.ParamOutputSchema{}
}

// 获取出参 schema
func GetRuntimeParamOutputSchema(paramSchemaParser iworkprotocol.IParamSchemaParser) *iworkmodels.ParamOutputSchema {
	return paramSchemaParser.GetRuntimeParamOutputSchema()
}

func GetDefaultParamOutputSchema(paramSchemaParser iworkprotocol.IParamSchemaParser) *iworkmodels.ParamOutputSchema {
	return paramSchemaParser.GetDefaultParamOutputSchema()
}

// 获取入参 schema
func GetCacheParamInputSchema(step *iwork.WorkStep, paramSchemaParser iworkprotocol.IParamSchemaParser) *iworkmodels.ParamInputSchema {
	// 从缓存(数据库字段)中获取
	if strings.TrimSpace(step.WorkStepInput) != "" {
		var paramInputSchema *iworkmodels.ParamInputSchema
		if err := json.Unmarshal([]byte(step.WorkStepInput), &paramInputSchema); err == nil {
			return paramInputSchema
		}
	}
	// 获取当前 work_step 对应的 paramInputSchema
	return paramSchemaParser.GetDefaultParamInputSchema()
}

// 获取默认入参 schema
func GetDefaultParamInputSchema(paramSchemaParser iworkprotocol.IParamSchemaParser) *iworkmodels.ParamInputSchema {
	return paramSchemaParser.GetDefaultParamInputSchema()
}

// 获取入参 schema
func GetRuntimeParamInputSchema(paramSchemaParser iworkprotocol.IParamSchemaParser) *iworkmodels.ParamInputSchema {
	return paramSchemaParser.GetRuntimeParamInputSchema()
}

// 根据传入的 paramMap 构建 ParamInputSchema 对象
func BuildParamInputSchemaWithDefaultMap(paramMap map[int][]string) *iworkmodels.ParamInputSchema {
	keys := datatypeutil.GetMapKeySlice(paramMap, []int{}).([]int)
	sort.Ints(keys)
	items := make([]iworkmodels.ParamInputSchemaItem, 0)
	for _, key := range keys {
		_paramMap := paramMap[key]
		// 前两位分别是名称和描述
		paramName := _paramMap[0]
		paramDesc := _paramMap[1]
		item := iworkmodels.ParamInputSchemaItem{ParamName: paramName, ParamDesc: paramDesc}
		// 后面字段为 extra 字段
		for _, paramExtra := range _paramMap[1:] {
			if strings.HasPrefix(paramExtra, "repeatable__") {
				item.Repeatable = true
				item.RepeatRefer = strings.Replace(paramExtra, "repeatable__", "", 1)
			}
		}
		items = append(items, item)
	}
	return &iworkmodels.ParamInputSchema{ParamInputSchemaItems: items}
}

// 根据传入的 paramNames 构建 ParamOutputSchema 对象
func BuildParamOutputSchemaWithSlice(paramNames []string) *iworkmodels.ParamOutputSchema {
	items := make([]iworkmodels.ParamOutputSchemaItem, 0)
	for _, paramName := range paramNames {
		items = append(items, iworkmodels.ParamOutputSchemaItem{ParamName: paramName})
	}
	return &iworkmodels.ParamOutputSchema{ParamOutputSchemaItems: items}
}
