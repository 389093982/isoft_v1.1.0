package iworknode

import (
	"encoding/json"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/models/iwork"
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
	return &schema.ParamOutputSchema{}
}



