package iworkcomponent

import (
	"isoft/isoft_iaas_web/core/iworkdata"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
)

type IWorkHelper struct {
	Work      iwork.Work
	WorkSteps []*iwork.WorkStep
}

func (this *IWorkHelper) Execute() {
	for _, workStep := range this.WorkSteps {
		workStepHelper := &IWorkStepHelper{Work: this.Work, WorkStep: workStep}
		workStepHelper.Execute()
	}
}

type IWorkStepHelper struct {
	Work     iwork.Work
	WorkStep *iwork.WorkStep
}

func (this *IWorkStepHelper) Execute() {
	factory := &WorkStepTypeFactory{Executor: this}
	if executable := factory.GetExecutor(); executable != nil {
		executable.Execute()
	}
}

func (this *IWorkStepHelper) GetDefaultParamInputSchema() *iworkdata.ParamInputSchema {
	factory := &WorkStepTypeFactory{Executor: this}
	if schema := factory.GetDefaultParamInputSchema(); schema != nil {
		return schema
	}
	return &iworkdata.ParamInputSchema{}
}

func (this *IWorkStepHelper) GetDefaultParamOutputSchema() *iworkdata.ParamOutputSchema {
	factory := &WorkStepTypeFactory{Executor: this}
	if schema := factory.GetDefaultParamOutputSchema(); schema != nil {
		return schema
	}
	return &iworkdata.ParamOutputSchema{}
}

func (this *IWorkStepHelper) GetRuntimeParamOutputSchema() *iworkdata.ParamOutputSchema {
	factory := &WorkStepTypeFactory{Executor: this}
	if schema := factory.GetRuntimeParamOutputSchema(); schema != nil {
		return schema
	}
	return &iworkdata.ParamOutputSchema{}
}

type WorkStepTypeFactory struct {
	Executor *IWorkStepHelper
}

type Executable interface {
	Execute()
}

func (this *WorkStepTypeFactory) GetExecutor() Executable {
	switch strings.ToUpper(this.Executor.WorkStep.WorkStepType) {
	case "SQL_INSERT":
		return &SQLInsert{}
	case "SQL_QUERY":
		return &SQLQuery{WorkStep: this.Executor.WorkStep}
	}
	return nil
}

func (this *WorkStepTypeFactory) GetDefaultParamInputSchema() *iworkdata.ParamInputSchema {
	switch strings.ToUpper(this.Executor.WorkStep.WorkStepType) {
	case "SQL_INSERT":
		helper := &SQLQuery{WorkStep: this.Executor.WorkStep}
		return helper.GetDefaultParamInputSchema()
	case "SQL_QUERY":
		helper := &SQLQuery{WorkStep: this.Executor.WorkStep}
		return helper.GetDefaultParamInputSchema()
	}
	return nil
}

func (this *WorkStepTypeFactory) GetDefaultParamOutputSchema() *iworkdata.ParamOutputSchema {
	switch strings.ToUpper(this.Executor.WorkStep.WorkStepType) {
	case "WORK_START":
		node := &WorkStartNode{WorkStep: this.Executor.WorkStep}
		return node.GetDefaultParamOutputSchema()
	case "WORK_END":
		node := &WorkEndNode{WorkStep: this.Executor.WorkStep}
		return node.GetDefaultParamOutputSchema()
	case "SQL_INSERT":
		node := &SQLQuery{WorkStep: this.Executor.WorkStep}
		return node.GetDefaultParamOutputSchema()
	case "SQL_QUERY":
		node := &SQLQuery{WorkStep: this.Executor.WorkStep}
		return node.GetDefaultParamOutputSchema()
	}
	return &iworkdata.ParamOutputSchema{}
}

func (this *WorkStepTypeFactory) GetRuntimeParamOutputSchema() *iworkdata.ParamOutputSchema {
	switch strings.ToUpper(this.Executor.WorkStep.WorkStepType) {
	case "SQL_INSERT":
		node := &SQLQuery{WorkStep: this.Executor.WorkStep}
		return node.GetRuntimeParamOutputSchema()
	case "SQL_QUERY":
		node := &SQLQuery{WorkStep: this.Executor.WorkStep}
		return node.GetRuntimeParamOutputSchema()
	}
	return &iworkdata.ParamOutputSchema{}
}
