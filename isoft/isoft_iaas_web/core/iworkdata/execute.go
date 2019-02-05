package iworkdata

import (
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
)

type IWorkHelper struct {
	Work iwork.Work
	WorkSteps []*iwork.WorkStep
}

func (this *IWorkHelper) Execute()  {
	for _, workStep := range this.WorkSteps{
		workStepHelper := &IWorkStepHelper{Work:this.Work, WorkStep:workStep}
		workStepHelper.Execute()
	}
}

type IWorkStepHelper struct {
	Work iwork.Work
	WorkStep *iwork.WorkStep
}

func (this *IWorkStepHelper) Execute()  {
	factory := &WorkStepTypeFactory{Executor:this}
	if executable := factory.GetExecutor(); executable != nil{
		executable.Execute()
	}
}

func (this *IWorkStepHelper) GetDefaultParamDefinition() *ParamDefinition {
	factory := &WorkStepTypeFactory{Executor:this}
	if definition := factory.GetDefaultParamDefinition(); definition != nil{
		return definition
	}
	return nil
}

type WorkStepTypeFactory struct {
	Executor *IWorkStepHelper
}

type Executable interface {
	Execute()
}

type DefaultParamDefinition interface {
	GetDefaultParamDefinition()
}

func (this *WorkStepTypeFactory) GetExecutor() Executable {
	switch strings.ToUpper(this.Executor.WorkStep.WorkStepType) {
	case "SQL_INSERT":
		return &SQLInsert{}
	case "SQL_QUERY":
		return &SQLQuery{Executor:this.Executor}
	}
	return nil
}

func (this *WorkStepTypeFactory) GetDefaultParamDefinition() *ParamDefinition {
	switch strings.ToUpper(this.Executor.WorkStep.WorkStepType) {
	case "SQL_INSERT":
		helper := &SQLQuery{Executor:this.Executor}
		return helper.GetDefaultParamDefinition()
	case "SQL_QUERY":
		helper := &SQLQuery{Executor:this.Executor}
		return helper.GetDefaultParamDefinition()
	}
	return nil
}