package iworkcomponent

import (
	"fmt"
	"isoft/isoft_iaas_web/core/iworkcomponent/sqlutil"
	"isoft/isoft_iaas_web/core/iworkdata"
	"isoft/isoft_iaas_web/models/iwork"
)

type SQLQueryNode struct {
	BaseNode
	WorkStep 		    *iwork.WorkStep
}

func (this *SQLQueryNode) Execute(trackingId string) {
	// 数据中心
	dataStore := iworkdata.GetDataSource(trackingId)
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep,dataStore)
	sql := tmpDataMap["sql"].(string) 				  // 等价于 iworkdata.GetStaticParamValue("sql",this.WorkStep)
	dataSourceName := tmpDataMap["db_conn"].(string)  // 等价于 iworkdata.GetStaticParamValue("db_conn", this.WorkStep)
	datacounts, rowDatas := sqlutil.ExcuteQuery(sql, dataSourceName)
	// 将数据数据存储到数据中心
	// 存储 datacounts
	dataStore.CacheData(this.WorkStep.WorkStepName, fmt.Sprintf("$%s.datacounts", this.WorkStep.WorkStepName), datacounts)
	for key,value := range rowDatas{
		// 存储具体字段值
		dataStore.CacheData(this.WorkStep.WorkStepName, fmt.Sprintf("$%s.%s", this.WorkStep.WorkStepName,key), value)
	}
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
	sql := iworkdata.GetStaticParamValue("sql",this.WorkStep)
	dataSourceName := iworkdata.GetStaticParamValue("db_conn", this.WorkStep)
	paramNames := sqlutil.GetMetaDatas(sql, dataSourceName)
	items := []iworkdata.ParamOutputSchemaItem{}
	for _, paramName := range paramNames {
		items = append(items, iworkdata.ParamOutputSchemaItem{
			ParentPath:"rows",
			ParamName: paramName,
		})
	}
	return &iworkdata.ParamOutputSchema{ParamOutputSchemaItems: items}
}

