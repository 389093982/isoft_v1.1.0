package iworknode

import (
	"encoding/json"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
)

type EntityParserNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *EntityParserNode) Execute(trackingId string) {

}

func (this *EntityParserNode) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	return &schema.ParamInputSchema{}
}

func (this *EntityParserNode) GetRuntimeParamInputSchema() *schema.ParamInputSchema {
	var paramMappingsArr []string
	json.Unmarshal([]byte(this.WorkStep.WorkStepParamMapping), &paramMappingsArr)
	items := make([]schema.ParamInputSchemaItem, 0)
	for _, paramMapping := range paramMappingsArr {
		items = append(items, schema.ParamInputSchemaItem{ParamName: paramMapping})
	}
	return &schema.ParamInputSchema{ParamInputSchemaItems: items}
}

func (this *EntityParserNode) GetDefaultParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}

func (this *EntityParserNode) GetRuntimeParamOutputSchema() *schema.ParamOutputSchema {
	items := make([]schema.ParamOutputSchemaItem, 0)
	inputSchema := schema.GetCacheParamInputSchema(this.WorkStep, &WorkStepFactory{WorkStep: this.WorkStep})
	for _, item := range inputSchema.ParamInputSchemaItems {
		// 从用户输入值中提取实体类字段详细信息
		entityFieldStr := getParamValueForEntity(item.ParamValue)
		for _, entityField := range strings.Split(entityFieldStr, ","){
			// 每个字段放入 items 中
			items = append(items, schema.ParamOutputSchemaItem{ParamName: strings.TrimSpace(entityField), ParentPath: item.ParamName})
		}
	}
	return &schema.ParamOutputSchema{ParamOutputSchemaItems: items}
}

func getParamValueForEntity(paramValue string) string {
	paramValue = strings.TrimSpace(paramValue)
	paramValue = strings.Replace(paramValue, ";", "", -1)
	if !strings.HasPrefix(paramValue, "$Entity."){
		return paramValue
	}
	entity_name := strings.Replace(paramValue, "$Entity.", "", -1)
	if entity, err := iwork.QueryEntityByEntityName(entity_name); err == nil{
		return entity.EntityFieldStr
	}
	return ""
}

