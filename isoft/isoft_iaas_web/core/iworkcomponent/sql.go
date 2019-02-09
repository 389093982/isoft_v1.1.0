package iworkcomponent

import (
	"isoft/isoft_iaas_web/core/iworkcomponent/sqlutil"
	"isoft/isoft_iaas_web/core/iworkdata"
	"isoft/isoft_iaas_web/models/iwork"
)

type SQLQueryNode struct {
	WorkStep 		   *iwork.WorkStep
	paramInputMap  *map[string]interface{}
	paramOutputMap *map[string]interface{}
}

func (this *SQLQueryNode) Execute(trackingId string) {
	workStepInput := this.WorkStep.WorkStepInput
	workStepOutput := this.WorkStep.WorkStepOutput
	inputResolver := &ParamResolver{ParamStr: workStepInput}
	this.paramInputMap = inputResolver.ParseParamStrToMap()
	outputResolver := &ParamResolver{ParamStr: workStepOutput}
	this.paramOutputMap = outputResolver.ParseParamStrToMap()
	this.ExecuteWithParams()
}

func (this *SQLQueryNode) GetDefaultParamInputSchema() *iworkdata.ParamInputSchema {
	paramNames := []string{"sql", "sql_binding?", "db_conn"}
	items := []iworkdata.ParamInputSchemaItem{}
	for _, paramName := range paramNames {
		items = append(items, iworkdata.ParamInputSchemaItem{ParamName: paramName})
	}
	return &iworkdata.ParamInputSchema{ParamInputSchemaItems: items}
}

func (this *SQLQueryNode) GetDefaultParamOutputSchema() *iworkdata.ParamOutputSchema {
	paramNames := []string{"datacounts"}
	items := []iworkdata.ParamOutputSchemaItem{}
	for _, paramName := range paramNames {
		items = append(items, iworkdata.ParamOutputSchemaItem{ParamName: paramName})
	}
	return &iworkdata.ParamOutputSchema{ParamOutputSchemaItems: items}
}

func (this *SQLQueryNode) GetRuntimeParamOutputSchema() *iworkdata.ParamOutputSchema {
	paramNames := sqlutil.GetMetaDatas(GetParamValue(*this.WorkStep, "sql"),
		GetParamValue(*this.WorkStep, "db_conn"))
	items := []iworkdata.ParamOutputSchemaItem{}
	for _, paramName := range paramNames {
		items = append(items, iworkdata.ParamOutputSchemaItem{
			ParentPath:"rows",
			ParamName: paramName,
		})
	}
	return &iworkdata.ParamOutputSchema{ParamOutputSchemaItems: items}
}

func (this *SQLQueryNode) ExecuteWithParams() {

}

type SQLInsert struct {
}

func (this *SQLInsert) Execute() {

}


//func SQLQueryNodeRun(work iwork.Work, step iwork.WorkStep) {
//	db, err := sqlutil.GetConnForMysql("mysql", iworkcomponent.GetParamValue(step, "db_conn"))
//	if err != nil {
//		panic(err)
//	}
//	rows, err := db.Query(iworkcomponent.GetParamValue(step, "sql"))
//	if err != nil {
//		panic(err)
//	}
//	defer rows.Close()
//
//	colNames, _ := rows.Columns()
//	for rows.Next() {
//		colValues := make([]sql.RawBytes, len(colNames))
//		scanArgs := make([]interface{}, len(colValues))
//		for i := range colValues {
//			scanArgs[i] = &colValues[i]
//		}
//		rows.Scan(scanArgs...)
//		for _, colValue := range colValues {
//			fmt.Print(colValue)
//		}
//	}
//}