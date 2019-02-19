package iworknode

import (
	"fmt"
	"isoft/isoft/common/pageutil"
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/core/iworkdata/param"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworkutil/funcutil"
	"isoft/isoft_iaas_web/core/iworkutil/sqlutil"
	"isoft/isoft_iaas_web/models/iwork"
	"strconv"
	"strings"
)

type SQLQueryNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *SQLQueryNode) Execute(trackingId string) {
	// 跳过解析和填充的数据
	skips := []string{"sql","db_conn"}
	// 数据中心
	dataStore := datastore.GetDataSource(trackingId)
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, dataStore, skips...)
	sql := param.GetStaticParamValue("sql", this.WorkStep)
	dataSourceName := param.GetStaticParamValue("db_conn", this.WorkStep)
	// sql_binding 参数获取
	sql_binding := getSqlBinding(tmpDataMap)
	datacounts, rowDetailDatas, rowDatas := sqlutil.Query(sql, sql_binding, dataSourceName)
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
		1:[]string{"metadata_sql","元数据sql语句,针对复杂查询sql,需要提供类似于select * from blog where 1=0的辅助sql用来构建节点输出"},
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
	return getMetaDataForQuery(this.WorkStep)
}
func (this *SQLQueryNode) ValidateCustom() {
	validateAndGetDataSourceName(this.WorkStep)
	validateAndGetMetaDataSql(this.WorkStep)
	validateSqlBindingParamCount(this.WorkStep)
}

// 从 tmpDataMap 获取 sql_binding 数据
func getSqlBinding(tmpDataMap map[string]interface{}) []interface{} {
	result := make([]interface{},0)
	if sql_binding, ok := tmpDataMap["sql_binding?"].([]interface{}); ok {
		result = sql_binding
	} else if sql_binding, ok := tmpDataMap["sql_binding?"].(interface{}); ok {
		result = append(result, sql_binding)
	}
	return result
}

type SQLExecuteNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *SQLExecuteNode) Execute(trackingId string) {
	// 跳过解析和填充的数据
	skips := []string{"sql","db_conn"}
	// 数据中心
	dataStore := datastore.GetDataSource(trackingId)
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, dataStore, skips...)
	sql := param.GetStaticParamValue("sql", this.WorkStep)
	dataSourceName := param.GetStaticParamValue("db_conn", this.WorkStep)
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
	batch_number := GetBatchNumber(tmpDataMap)
	if batch_number > 1 && strings.HasPrefix(strings.ToUpper(strings.TrimSpace(sql)), "INSERT") {
		// 最后一个左括号索引
		index1 := strings.LastIndex(sql, "(")
		// 最后一个右括号索引
		index2 := strings.LastIndex(sql, ")")
		// value 填充子句
		valueSql := sql[index1:(index2 + 1)]
		// newValueArr 等于 value 填充子句复制 _batch_number 份
		newValueArr := make([]string, 0)
		for i := 0; i < batch_number; i++ {
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

func (this *SQLExecuteNode) ValidateCustom() {
	validateAndGetDataSourceName(this.WorkStep)
}

type SQLQueryPageNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *SQLQueryPageNode) Execute(trackingId string) {
	// 跳过解析和填充的数据
	skips := []string{"total_sql","sql","db_conn"}
	// 数据中心
	dataStore := datastore.GetDataSource(trackingId)
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, dataStore, skips...)
	total_sql := param.GetStaticParamValue("total_sql",this.WorkStep)
	sql := param.GetStaticParamValue("sql",this.WorkStep)
	dataSourceName := param.GetStaticParamValue("db_conn",this.WorkStep)
	// sql_binding 参数获取
	sql_binding := getSqlBinding(tmpDataMap)
	totalcount := sqlutil.QuerySelectCount(total_sql, sql_binding[:len(sql_binding) - 2], dataSourceName)
	datacounts, rowDetailDatas, rowDatas := sqlutil.Query(sql, sql_binding, dataSourceName)
	// 将数据数据存储到数据中心
	// 存储 datacounts
	dataStore.CacheData(this.WorkStep.WorkStepName, "datacounts", datacounts)
	for paramName, paramValue := range rowDetailDatas {
		// 存储具体字段值
		dataStore.CacheData(this.WorkStep.WorkStepName, paramName, paramValue)
	}
	// 数组对象整体存储在 rows 里面
	dataStore.CacheData(this.WorkStep.WorkStepName, "rows", rowDatas)
	// 存储分页信息
	pageIndex,pageSize := getPageIndexAndPageSize(tmpDataMap)
	paginator := pageutil.Paginator(pageIndex, pageSize, totalcount)
	dataStore.CacheData(this.WorkStep.WorkStepName, "paginator", paginator)
	for key,value := range paginator{
		dataStore.CacheData(this.WorkStep.WorkStepName, "paginator." + key, value)
	}
}

func getPageIndexAndPageSize(tmpDataMap map[string]interface{}) (currentPage int,pageSize int) {
	if current_page, ok := tmpDataMap["current_page"].(string); ok{
		currentPage, _ = strconv.Atoi(current_page)
	}else if current_page, ok := tmpDataMap["current_page"].(int); ok{
		currentPage = current_page
	}else if current_page, ok := tmpDataMap["current_page"].(int64); ok{
		currentPage = int(current_page)
	}
	if page_size, ok := tmpDataMap["page_size"].(string); ok{
		pageSize, _ = strconv.Atoi(page_size)
	}else if page_size, ok := tmpDataMap["page_size"].(int); ok{
		pageSize = page_size
	}else if page_size, ok := tmpDataMap["page_size"].(int64); ok{
		pageSize = int(page_size)
	}
	return
}

func (this *SQLQueryPageNode) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	paramMap := map[int][]string{
		1:[]string{"metadata_sql","元数据sql语句,针对复杂查询sql,需要提供类似于select * from blog where 1=0的辅助sql用来构建节点输出"},
		2:[]string{"total_sql","统计总数sql,返回N页总数据量,格式参考select count(*) as count from blog where xxx"},
		3:[]string{"sql","带分页条件的sql,等价于 ${total_sql} limit ?,?"},
		4:[]string{"current_page","当前页数"},
		5:[]string{"page_size","每页数据量"},
		6:[]string{"sql_binding?","sql绑定数据,个数和sql中的?数量相同,前N-2位参数和total_sql中的?数量相同"},
		7:[]string{"db_conn","数据库连接信息,需要使用 $RESOURCE 全局参数"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *SQLQueryPageNode) GetRuntimeParamInputSchema() *schema.ParamInputSchema {
	return &schema.ParamInputSchema{}
}

func (this *SQLQueryPageNode) GetDefaultParamOutputSchema() *schema.ParamOutputSchema {
	items := make([]schema.ParamOutputSchemaItem,0)
	for _, paginatorField := range pageutil.GetPaginatorFields() {
		items = append(items, schema.ParamOutputSchemaItem{
			ParentPath: "paginator",
			ParamName:  paginatorField,
		})
	}
	return &schema.ParamOutputSchema{ParamOutputSchemaItems: items}
}

func (this *SQLQueryPageNode) GetRuntimeParamOutputSchema() *schema.ParamOutputSchema {
	return getMetaDataForQuery(this.WorkStep)
}

func (this *SQLQueryPageNode) ValidateCustom() {
	validateAndGetDataSourceName(this.WorkStep)
	validateAndGetMetaDataSql(this.WorkStep)
	validateSqlBindingParamCount(this.WorkStep)
	validateSqlBindingParamCount(this.WorkStep)
	validateTotalSqlBindingParamCount(this.WorkStep)
}

func validateTotalSqlBindingParamCount(step *iwork.WorkStep) {
	total_sql := param.GetStaticParamValue("total_sql", step)
	sql_binding := param.GetStaticParamValue("sql_binding?", step)
	if strings.Count(total_sql, "?") + 2 != strings.Count(funcutil.EncodeSpecialForParamVaule(sql_binding), ";"){
		panic("Number of ? in total_sql and number of ; in sql_binding is mismatch!")
	}
}

func validateSqlBindingParamCount(step *iwork.WorkStep) {
	sql := param.GetStaticParamValue("sql", step)
	sql_binding := param.GetStaticParamValue("sql_binding?", step)
	if strings.Count(sql, "?") != strings.Count(funcutil.EncodeSpecialForParamVaule(sql_binding), ";"){
		panic("Number of ? in SQL and number of ; in sql_binding is unequal!")
	}
}

func validateAndGetMetaDataSql(step *iwork.WorkStep) string {
	metadata_sql := param.GetStaticParamValue("metadata_sql", step)
	if strings.TrimSpace(metadata_sql) == ""{
		panic("Empty paramValue for metadata_sql was found!")
	}
	if strings.Contains(metadata_sql, "?"){
		panic("Invalid paramValue form metadata_sql was found!")
	}
	return strings.TrimSpace(metadata_sql)
}

func validateAndGetDataSourceName(step *iwork.WorkStep) string {
	dataSourceName := param.GetStaticParamValue("db_conn", step)
	if strings.TrimSpace(dataSourceName) == ""{
		panic("Invalid param for db_conn! Can't resolve it!")
	}
	db, err := sqlutil.GetConnForMysql("mysql", dataSourceName)
	if err != nil {
		panic(fmt.Sprintf("Can't get DB connection for %s!", dataSourceName))
	}
	defer db.Close()
	return dataSourceName
}

func getMetaDataForQuery(step *iwork.WorkStep) *schema.ParamOutputSchema {
	metadataSql := validateAndGetMetaDataSql(step)
	dataSourceName := validateAndGetDataSourceName(step)
	paramNames := sqlutil.GetMetaDatas(metadataSql, dataSourceName)
	items := make([]schema.ParamOutputSchemaItem,0)
	for _, paramName := range paramNames {
		items = append(items, schema.ParamOutputSchemaItem{
			ParentPath: "rows",
			ParamName:  paramName,
		})
	}
	return &schema.ParamOutputSchema{ParamOutputSchemaItems: items}
}
