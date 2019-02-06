package iworkdata

import (
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

func (this *IWorkStepHelper) GetDefaultParamInputSchema() *ParamInputSchema {
	factory := &WorkStepTypeFactory{Executor: this}
	if schema := factory.GetDefaultParamInputSchema(); schema != nil {
		return schema
	}
	return &ParamInputSchema{}
}

func (this *IWorkStepHelper) GetDefaultParamOutputSchema() *ParamOutputSchema {
	factory := &WorkStepTypeFactory{Executor: this}
	if schema := factory.GetDefaultParamOutputSchema(); schema != nil {
		return schema
	}
	return &ParamOutputSchema{}
}

func (this *IWorkStepHelper) GetRuntimeParamOutputSchema() *ParamOutputSchema {
	factory := &WorkStepTypeFactory{Executor: this}
	if schema := factory.GetRuntimeParamOutputSchema(); schema != nil {
		return schema
	}
	return &ParamOutputSchema{}
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
		return &SQLQuery{Executor: this.Executor}
	}
	return nil
}

func (this *WorkStepTypeFactory) GetDefaultParamInputSchema() *ParamInputSchema {
	switch strings.ToUpper(this.Executor.WorkStep.WorkStepType) {
	case "SQL_INSERT":
		helper := &SQLQuery{Executor: this.Executor}
		return helper.GetDefaultParamInputSchema()
	case "SQL_QUERY":
		helper := &SQLQuery{Executor: this.Executor}
		return helper.GetDefaultParamInputSchema()
	}
	return nil
}

func (this *WorkStepTypeFactory) GetDefaultParamOutputSchema() *ParamOutputSchema {
	switch strings.ToUpper(this.Executor.WorkStep.WorkStepType) {
	case "SQL_INSERT":
		helper := &SQLQuery{Executor: this.Executor}
		return helper.GetDefaultParamOutputSchema()
	case "SQL_QUERY":
		helper := &SQLQuery{Executor: this.Executor}
		return helper.GetDefaultParamOutputSchema()
	}
	return nil
}

func (this *WorkStepTypeFactory) GetRuntimeParamOutputSchema() *ParamOutputSchema {
	switch strings.ToUpper(this.Executor.WorkStep.WorkStepType) {
	case "SQL_INSERT":
		helper := &SQLQuery{Executor: this.Executor}
		return helper.GetRuntimeParamOutputSchema()
	case "SQL_QUERY":
		helper := &SQLQuery{Executor: this.Executor}
		return helper.GetRuntimeParamOutputSchema()
	}
	return nil
}
