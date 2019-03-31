package iworknode

import (
	"isoft/isoft_iaas_web/core/iworkmodels"
	"isoft/isoft_iaas_web/models/iwork"
)

type EmptyNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *EmptyNode) Execute(trackingId string) {

}

func (this *EmptyNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	return &iworkmodels.ParamInputSchema{}
}

func (this *EmptyNode) GetRuntimeParamInputSchema() *iworkmodels.ParamInputSchema {
	return &iworkmodels.ParamInputSchema{}
}

func (this *EmptyNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return &iworkmodels.ParamOutputSchema{}
}

func (this *EmptyNode) GetRuntimeParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return &iworkmodels.ParamOutputSchema{}
}

func (this *EmptyNode) ValidateCustom() (checkResult []string) {
	return []string{}
}
