package iworkdata

type SQLQuery struct {
	Executor *IWorkStepHelper
	paramInputMap *map[string]interface{}
	paramOutputMap *map[string]interface{}
}

func (this *SQLQuery) Execute()  {
	workStepInput := this.Executor.WorkStep.WorkStepInput
	workStepOutput := this.Executor.WorkStep.WorkStepOutput
	inputResolver := &ParamResolver{ParamStr:workStepInput}
	this.paramInputMap = inputResolver.ParseParamStrToMap()
	outputResolver := &ParamResolver{ParamStr:workStepOutput}
	this.paramOutputMap = outputResolver.ParseParamStrToMap()
	this.ExecuteWithParams()
}

func (this *SQLQuery) GetDefaultParamDefinition() *ParamDefinition {
	items := []ParamDefinitionItem{
		{
			ParamName:"sql",
			ParamValue:"",
		},
		{
			ParamName:"sql_binding?",
			ParamValue:"",
		},
	}
	return &ParamDefinition{ParamDefinitionItems:items}
}

func (this *SQLQuery) ExecuteWithParams()  {
	
}

type SQLInsert struct {

}

func (this *SQLInsert) Execute()  {

}