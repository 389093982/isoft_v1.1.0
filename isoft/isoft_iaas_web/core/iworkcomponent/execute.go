package iworkcomponent

import (
	"isoft/isoft_iaas_web/core/iworkdata"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
)

type WorkStepFactory struct {
	Work     iwork.Work
	WorkStep *iwork.WorkStep
}

type IStandardWorkStep interface {
	Execute()
	GetDefaultParamInputSchema() *iworkdata.ParamInputSchema
	GetDefaultParamOutputSchema() *iworkdata.ParamOutputSchema
	GetRuntimeParamOutputSchema() *iworkdata.ParamOutputSchema
}

func (this *WorkStepFactory) Execute() {
	this.getProxy().Execute()
}

func (this *WorkStepFactory) getProxy() IStandardWorkStep {
	switch strings.ToUpper(this.WorkStep.WorkStepType) {
	case "SQL_INSERT":
		return &SQLQuery{WorkStep: this.WorkStep}
	case "SQL_QUERY":
		return &SQLQuery{WorkStep: this.WorkStep}
	}
	return nil
}

func (this *WorkStepFactory) GetDefaultParamInputSchema() *iworkdata.ParamInputSchema {
	if schema := this.getProxy().GetDefaultParamInputSchema(); schema != nil{
		return schema
	}
	return &iworkdata.ParamInputSchema{}
}

func (this *WorkStepFactory) GetDefaultParamOutputSchema() *iworkdata.ParamOutputSchema {
	if schema := this.getProxy().GetDefaultParamOutputSchema(); schema != nil{
		return schema
	}
	return &iworkdata.ParamOutputSchema{}
}

func (this *WorkStepFactory) GetRuntimeParamOutputSchema() *iworkdata.ParamOutputSchema {
	if schema := this.getProxy().GetRuntimeParamOutputSchema(); schema != nil{
		return schema
	}
	return &iworkdata.ParamOutputSchema{}
}
