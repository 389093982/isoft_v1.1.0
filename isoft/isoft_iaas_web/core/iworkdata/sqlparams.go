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

func (this *SQLQuery) GetDefaultParamInputSchema() *ParamInputSchema {
	paramNames := []string{"sql","sql_binding?","db_conn"}
	items := []ParamInputSchemaItem{}
	for _, paramName := range paramNames{
		items = append(items, ParamInputSchemaItem{ParamName:paramName})
	}
	return &ParamInputSchema{ParamInputSchemaItems:items}
}

func (this *SQLQuery) ExecuteWithParams()  {
	
}

type SQLInsert struct {

}

func (this *SQLInsert) Execute()  {

}