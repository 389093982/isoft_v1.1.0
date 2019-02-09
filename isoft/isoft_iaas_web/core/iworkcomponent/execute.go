package iworkcomponent

import (
	"fmt"
	"isoft/isoft_iaas_web/core/iworkdata"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
)

type WorkStepFactory struct {
	Work     iwork.Work
	WorkStep *iwork.WorkStep
}

type IStandardWorkStep interface {
	Execute(trackingId string)
	GetDefaultParamInputSchema() *iworkdata.ParamInputSchema
	GetDefaultParamOutputSchema() *iworkdata.ParamOutputSchema
	GetRuntimeParamOutputSchema() *iworkdata.ParamOutputSchema
}

func (this *WorkStepFactory) Execute(trackingId string) {
	this.getProxy().Execute(trackingId)
}

func (this *WorkStepFactory) getProxy() IStandardWorkStep {
	switch strings.ToUpper(this.WorkStep.WorkStepType) {
	case "WORK_START":
		return &WorkStartNode{WorkStep: this.WorkStep}
	case "WORK_END":
		return &WorkEndNode{WorkStep: this.WorkStep}
	case "SQL_INSERT":
		return &SQLQueryNode{WorkStep: this.WorkStep}
	case "SQL_QUERY":
		return &SQLQueryNode{WorkStep: this.WorkStep}
	}
	panic(fmt.Sprintf("unsupport workStepType:%s",this.WorkStep.WorkStepType))
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
