package iworknode

import (
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/models/iwork"
)

type EmptyNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *EmptyNode) Execute(trackingId string, skipFunc func(tmpDataMap map[string]interface{}) bool) {

}

func (this *EmptyNode) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	return &schema.ParamInputSchema{}
}

func (this *EmptyNode) GetRuntimeParamInputSchema() *schema.ParamInputSchema {
	return &schema.ParamInputSchema{}
}

func (this *EmptyNode) GetDefaultParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}

func (this *EmptyNode) GetRuntimeParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}

func (this *EmptyNode) ValidateCustom() {

}
