package iworknode

import (
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/models/iwork"
)

type WorkSub struct {
	BaseNode
	WorkStep 		   *iwork.WorkStep
}

func (this *WorkSub) Execute(trackingId string) {}
func (this *WorkSub) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	return schema.BuildParamInputSchemaWithSlice([]string{"work_sub"})
}
func (this *WorkSub) GetDefaultParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}
func (this *WorkSub) GetRuntimeParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}