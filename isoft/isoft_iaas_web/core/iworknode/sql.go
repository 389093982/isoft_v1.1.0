package iworknode

import (
	"fmt"
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/core/iworkdata/param"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworkutil/sqlutil"
	"isoft/isoft_iaas_web/models/iwork"
)

type SQLQueryNode struct {
	BaseNode
	WorkStep 		    *iwork.WorkStep
}

func (this *SQLQueryNode) Execute(trackingId string) {
	// 数据中心
	dataStore := datastore.GetDataSource(trackingId)
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep,dataStore)
	sql := tmpDataMap["sql"].(string) 				  // 等价于 param.GetStaticParamValue("sql",this.WorkStep)
	dataSourceName := tmpDataMap["db_conn"].(string)  // 等价于 param.GetStaticParamValue("db_conn", this.WorkStep)
	// sql_binding 参数获取
	_sql_binding := getSqlBinding(tmpDataMap)
	datacounts, rowDatas := sqlutil.Query(sql, _sql_binding, dataSourceName)
	// 将数据数据存储到数据中心
	// 存储 datacounts
	dataStore.CacheData(this.WorkStep.WorkStepName, fmt.Sprintf("$%s.datacounts", this.WorkStep.WorkStepName), datacounts)
	for key, value := range rowDatas {
		// 存储具体字段值
		dataStore.CacheData(this.WorkStep.WorkStepName, fmt.Sprintf("$%s.%s", this.WorkStep.WorkStepName, key), value)
	}
}

// 从 tmpDataMap 获取 sql_binding 数据
func getSqlBinding(tmpDataMap map[string]interface{}) []interface{} {
	_sql_binding := []interface{}{}
	if sql_binding, ok := tmpDataMap["sql_binding?"].([]interface{}); ok {
		_sql_binding = sql_binding
	} else if sql_binding, ok := tmpDataMap["sql_binding?"].(interface{}); ok {
		_sql_binding = append(_sql_binding, sql_binding)
	}
	return _sql_binding
}

func (this *SQLQueryNode) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	return schema.BuildParamInputSchemaWithSlice([]string{"sql", "sql_binding?", "db_conn"})
}

func (this *SQLQueryNode) GetDefaultParamOutputSchema() *schema.ParamOutputSchema {
	return schema.BuildParamOutputSchemaWithSlice([]string{"datacounts"})
}

func (this *SQLQueryNode) GetRuntimeParamOutputSchema() *schema.ParamOutputSchema {
	sql := param.GetStaticParamValue("sql",this.WorkStep)
	dataSourceName := param.GetStaticParamValue("db_conn", this.WorkStep)
	paramNames := sqlutil.GetMetaDatas(sql, dataSourceName)
	items := []schema.ParamOutputSchemaItem{}
	for _, paramName := range paramNames {
		items = append(items, schema.ParamOutputSchemaItem{
			ParentPath:"rows",
			ParamName: paramName,
		})
	}
	return &schema.ParamOutputSchema{ParamOutputSchemaItems: items}
}



type SQLExecuteNode struct {
	BaseNode
	WorkStep 		    *iwork.WorkStep
}

func (this *SQLExecuteNode) Execute(trackingId string) {
	// 数据中心
	dataStore := datastore.GetDataSource(trackingId)
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep,dataStore)
	sql := tmpDataMap["sql"].(string) 				  // 等价于 param.GetStaticParamValue("sql",this.WorkStep)
	dataSourceName := tmpDataMap["db_conn"].(string)  // 等价于 param.GetStaticParamValue("db_conn", this.WorkStep)
	// sql_binding 参数获取
	_sql_binding := getSqlBinding(tmpDataMap)
	affected := sqlutil.Execute(sql, _sql_binding, dataSourceName)
	// 将数据数据存储到数据中心
	// 存储 affected
	dataStore.CacheData(this.WorkStep.WorkStepName, fmt.Sprintf("$%s.affected", this.WorkStep.WorkStepName), affected)
}

func (this *SQLExecuteNode) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	return schema.BuildParamInputSchemaWithSlice([]string{"sql", "sql_binding?", "db_conn"})
}

func (this *SQLExecuteNode) GetDefaultParamOutputSchema() *schema.ParamOutputSchema {
	return schema.BuildParamOutputSchemaWithSlice([]string{"affected"})
}

func (this *SQLExecuteNode) GetRuntimeParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}