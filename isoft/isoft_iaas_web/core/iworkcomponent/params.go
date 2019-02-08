package iworkcomponent

import (
	"encoding/xml"
	"isoft/isoft_iaas_web/core/iworkdata"
	"isoft/isoft_iaas_web/models/iresource"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
)

type ParamResolver struct {
	ParamStr string
}

func (this *ParamResolver) ParseParamStrToMap() *map[string]interface{} {
	return &map[string]interface{}{}
}

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

// 去除不合理的字符
func removeUnsupportChars(paramName string) string {
	paramName = strings.TrimSpace(paramName)
	paramName = strings.Replace(paramName, "\n","",-1)
	return paramName
}

// 获取参数值,支持获取动态参数
func GetParamValue(step iwork.WorkStep, paramName string) string {
	paramValueString := removeUnsupportChars(GetParamValueString(step, removeUnsupportChars(paramName)))
	if IsDynamicParam(paramValueString){
		if strings.HasPrefix(paramValueString, "$RESOURCE."){
			return iresource.GetResourceDataSourceNameString(strings.Replace(paramValueString, "$RESOURCE.", "", -1))
		}
	}
	return paramValueString
}

func GetParamValueString(step iwork.WorkStep, paramName string) string {
	var paramInputSchema iworkdata.ParamInputSchema
	if err := xml.Unmarshal([]byte(step.WorkStepInput), &paramInputSchema); err != nil {
		return ""
	}
	for _, item := range paramInputSchema.ParamInputSchemaItems {
		if item.ParamName == paramName {
			// 非必须参数不得为空
			if !strings.HasSuffix(item.ParamName, "?") && strings.TrimSpace(item.ParamValue) == "" {
				//panic(errors.New(fmt.Sprint("it is a mast parameter for %s", item.ParamName)))
				return ""
			}
			return item.ParamValue
		}
	}
	return ""
}

// 判断参数是否是动态参数
func IsDynamicParam(param string) bool {
	return strings.HasPrefix(param, "$")
}