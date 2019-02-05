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

func (this *IWorkStepHelper) GetDefaultParamSchema() *ParamSchema {
	factory := &WorkStepTypeFactory{Executor:this}
	if schema := factory.GetDefaultParamSchema(); schema != nil{
		return schema
	}
	return nil
}

type WorkStepTypeFactory struct {
	Executor *IWorkStepHelper
}

type Executable interface {
	Execute()
}

type DefaultParamSchema interface {
	GetDefaultParamSchema()
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

func (this *WorkStepTypeFactory) GetDefaultParamSchema() *ParamSchema {
	switch strings.ToUpper(this.Executor.WorkStep.WorkStepType) {
	case "SQL_INSERT":
		helper := &SQLQuery{Executor:this.Executor}
		return helper.GetDefaultParamSchema()
	case "SQL_QUERY":
		helper := &SQLQuery{Executor:this.Executor}
		return helper.GetDefaultParamSchema()
	}
	return nil
}