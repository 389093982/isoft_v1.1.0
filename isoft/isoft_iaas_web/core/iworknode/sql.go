package iworknode

import (
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/core/iworkdata/param"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworkutil/sqlutil"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
)

type SQLQueryNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *SQLQueryNode) Execute(trackingId string) {
	// 数据中心
	dataStore := datastore.GetDataSource(trackingId)
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, dataStore)
	sql := tmpDataMap["sql"].(string)                // 等价于 param.GetStaticParamValue("sql",this.WorkStep)
	dataSourceName := tmpDataMap["db_conn"].(string) // 等价于 param.GetStaticParamValue("db_conn", this.WorkStep)
	// sql_binding 参数获取
	_sql_binding := getSqlBinding(tmpDataMap)
	datacounts, rowDetailDatas, rowDatas := sqlutil.Query(sql, _sql_binding, dataSourceName)
	// 将数据数据存储到数据中心
	// 存储 datacounts
	dataStore.CacheData(this.WorkStep.WorkStepName, "datacounts", datacounts)
	for paramName, paramValue := range rowDetailDatas {
		// 存储具体字段值
		dataStore.CacheData(this.WorkStep.WorkStepName, paramName, paramValue)
	}
	// 数组对象整体存储在 rows 里面
	dataStore.CacheData(this.WorkStep.WorkStepName, "rows", rowDatas)
}

func (this *SQLQueryNode) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	paramMap := map[int][]string{
		1:[]string{"metadata_sql?","元数据sql语句,可选参数,针对复杂查询sql,需要提供类似于select * from blog where 1=0的辅助sql用来构建节点输出"},
		2:[]string{"sql","查询sql语句"},
		3:[]string{"sql_binding?","sql绑定数据,个数必须和当前执行sql语句中的占位符参数个数相同"},
		4:[]string{"db_conn","数据库连接信息,需要使用 $RESOURCE 全局参数"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *SQLQueryNode) GetRuntimeParamInputSchema() *schema.ParamInputSchema {
	return &schema.ParamInputSchema{}
}

func (this *SQLQueryNode) GetDefaultParamOutputSchema() *schema.ParamOutputSchema {
	return schema.BuildParamOutputSchemaWithSlice([]string{"datacounts"})
}

func (this *SQLQueryNode) GetRuntimeParamOutputSchema() *schema.ParamOutputSchema {
	var metadataSql string
	if sql := param.GetStaticParamValue("metadata_sql?", this.WorkStep); strings.TrimSpace(sql) != "" {
		metadataSql = sql
	} else {
		metadataSql = param.GetStaticParamValue("sql", this.WorkStep)
	}
	dataSourceName := param.GetStaticParamValue("db_conn", this.WorkStep)
	paramNames := sqlutil.GetMetaDatas(metadataSql, dataSourceName)
	items := []schema.ParamOutputSchemaItem{}
	for _, paramName := range paramNames {
		items = append(items, schema.ParamOutputSchemaItem{
			ParentPath: "rows",
			ParamName:  paramName,
		})
	}
	return &schema.ParamOutputSchema{ParamOutputSchemaItems: items}
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

type SQLExecuteNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *SQLExecuteNode) Execute(trackingId string) {
	// 数据中心
	dataStore := datastore.GetDataSource(trackingId)
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, dataStore)
	sql := tmpDataMap["sql"].(string)                // 等价于 param.GetStaticParamValue("sql",this.WorkStep)
	dataSourceName := tmpDataMap["db_conn"].(string) // 等价于 param.GetStaticParamValue("db_conn", this.WorkStep)
	// insert 语句且有批量操作时整改 sql 语句
	sql = this.modifySqlInsertWithBatchNumber(tmpDataMap, sql)
	// sql_binding 参数获取
	_sql_binding := getSqlBinding(tmpDataMap)
	affected := sqlutil.Execute(sql, _sql_binding, dataSourceName)
	// 将数据数据存储到数据中心
	// 存储 affected
	dataStore.CacheData(this.WorkStep.WorkStepName, "affected", affected)
}

func (this *SQLExecuteNode) modifySqlInsertWithBatchNumber(tmpDataMap map[string]interface{}, sql string) string {
	_batch_number := GetBatchNumber(tmpDataMap)
	if _batch_number > 1 && strings.HasPrefix(strings.ToUpper(strings.TrimSpace(sql)), "INSERT") {
		// 最后一个左括号索引
		index1 := strings.LastIndex(sql, "(")
		// 最后一个右括号索引
		index2 := strings.LastIndex(sql, ")")
		// value 填充子句
		valueSql := sql[index1:(index2 + 1)]
		// newValueArr 等于 value 填充子句复制 _batch_number 份
		newValueArr := make([]string, 0)
		for i := 0; i < _batch_number; i++ {
			newValueArr = append(newValueArr, valueSql)
		}
		newValueSql := strings.Join(newValueArr, ",")
		// 进行替换,相当于 () 替换成 (),(),(),()...
		sql = strings.Replace(sql, valueSql, newValueSql, -1)
	}
	return sql
}

func (this *SQLExecuteNode) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	paramMap := map[int][]string{
		1:[]string{"batch_number?","仅供批量插入数据时使用"},
		2:[]string{"sql","执行sql语句"},
		3:[]string{"sql_binding?","sql绑定数据,个数必须和当前执行sql语句中的占位符参数个数相同"},
		4:[]string{"db_conn","数据库连接信息,需要使用 $RESOURCE 全局参数"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *SQLExecuteNode) GetRuntimeParamInputSchema() *schema.ParamInputSchema {
	return &schema.ParamInputSchema{}
}

func (this *SQLExecuteNode) GetDefaultParamOutputSchema() *schema.ParamOutputSchema {
	return schema.BuildParamOutputSchemaWithSlice([]string{"affected"})
}

func (this *SQLExecuteNode) GetRuntimeParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}
