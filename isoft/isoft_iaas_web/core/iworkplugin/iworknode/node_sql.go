package iworknode

import (
	"fmt"
	"isoft/isoft/common/pageutil"
	"isoft/isoft_iaas_web/core/iworkconst"
	"isoft/isoft_iaas_web/core/iworkdata/param"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworkfunc"
	"isoft/isoft_iaas_web/core/iworkmodels"
	"isoft/isoft_iaas_web/core/iworkutil/sqlutil"
	"isoft/isoft_iaas_web/models/iwork"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type SQLQueryNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *SQLQueryNode) Execute(trackingId string) {
	paramMap := make(map[string]interface{}, 0)
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, this.DataStore)
	sql := tmpDataMap[iworkconst.STRING_PREFIX+"sql"].(string)
	dataSourceName := tmpDataMap[iworkconst.STRING_PREFIX+"db_conn"].(string)
	// sql_binding 参数获取
	sql_binding := getSqlBinding(tmpDataMap)
	sql = strings.ReplaceAll(sql, "{{", "")
	sql = strings.ReplaceAll(sql, "}}", "")
	datacounts, rowDatas := sqlutil.Query(sql, sql_binding, dataSourceName)
	// 存储 datacounts
	paramMap[iworkconst.NUMBER_PREFIX+"datacounts"] = datacounts
	// 数组对象整体存储在 rows 里面
	paramMap[iworkconst.MULTI_PREFIX+"rows"] = rowDatas
	this.DataStore.CacheDatas(this.WorkStep.WorkStepName, paramMap)
}

func (this *SQLQueryNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "sql", "查询sql语句,可以使用 {{}} 标识出 metadata 字段名和 tableName 表名,用于自动推断 metadatasql"},
		2: {iworkconst.STRING_PREFIX + "metadata_sql?", "元数据sql语句,针对复杂查询sql,需要提供类似于select * from blog where 1=0的辅助sql用来构建节点输出"},
		3: {iworkconst.MULTI_PREFIX + "sql_binding?", "sql绑定数据,个数必须和当前执行sql语句中的占位符参数个数相同"},
		4: {iworkconst.STRING_PREFIX + "db_conn", "数据库连接信息,需要使用 $RESOURCE 全局参数"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *SQLQueryNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return schema.BuildParamOutputSchemaWithSlice([]string{iworkconst.NUMBER_PREFIX + "datacounts"})
}

func (this *SQLQueryNode) GetRuntimeParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return getMetaDataQuietlyForQuery(this.WorkStep)
}

func (this *SQLQueryNode) ValidateCustom() (checkResult []string) {
	validateAndGetDataStoreName(this.WorkStep)
	validateAndGetMetaDataSql(this.WorkStep)
	validateSqlBindingParamCount(this.WorkStep)
	return []string{}
}

// 从 tmpDataMap 获取 sql_binding 数据
func getSqlBinding(tmpDataMap map[string]interface{}) []interface{} {
	result := make([]interface{}, 0)
	sql_binding := tmpDataMap[iworkconst.MULTI_PREFIX+"sql_binding?"]
	t1 := reflect.TypeOf(sql_binding)
	v1 := reflect.ValueOf(sql_binding)
	// 支持单层切片和双层切片
	if t1.Kind() == reflect.Slice {
		for i := 0; i < v1.Len(); i++ {
			v2 := v1.Index(i)
			if v2.Kind() == reflect.Slice {
				for j := 0; j < v2.Len(); j++ {
					result = append(result, v2.Index(j))
				}
			} else if v2.Kind() == reflect.Interface {
				if datas, ok := v2.Interface().([]interface{}); ok {
					for _, data := range datas {
						result = append(result, data)
					}
				} else {
					result = append(result, v2.Interface())
				}
			} else {
				result = append(result, v2)
			}
		}
	} else if t1.Kind() == reflect.Interface {
		if datas, ok := sql_binding.([]interface{}); ok {
			for _, data := range datas {
				result = append(result, data)
			}
		} else {
			result = append(result, sql_binding)
		}
	} else {
		result = append(result, sql_binding)
	}
	return result
}

type SQLExecuteNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *SQLExecuteNode) Execute(trackingId string) {
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, this.DataStore)
	sql := tmpDataMap[iworkconst.STRING_PREFIX+"sql"].(string)
	dataSourceName := tmpDataMap[iworkconst.STRING_PREFIX+"db_conn"].(string)
	// insert 语句且有批量操作时整改 sql 语句
	sql = this.modifySqlInsertWithBatch(tmpDataMap, sql)
	// sql_binding 参数获取
	sql_binding := getSqlBinding(tmpDataMap)
	affected := sqlutil.Execute(sql, sql_binding, dataSourceName)
	// 将数据数据存储到数据中心
	// 存储 affected
	this.DataStore.CacheDatas(this.WorkStep.WorkStepName, map[string]interface{}{iworkconst.NUMBER_PREFIX + "affected": affected})
}

func (this *SQLExecuteNode) modifySqlInsertWithBatch(tmpDataMap map[string]interface{}, sql string) string {
	if batch_data := tmpDataMap[iworkconst.FOREACH_PREFIX+"batch_data?"]; batch_data != nil {
		t := reflect.TypeOf(batch_data)
		v := reflect.ValueOf(batch_data)
		if t.Kind() == reflect.Slice {
			if v.Len() > 1 && strings.HasPrefix(strings.ToUpper(strings.TrimSpace(sql)), "INSERT") {
				// 最后一个左括号索引
				index1 := strings.LastIndex(sql, "(")
				// 最后一个右括号索引
				index2 := strings.LastIndex(sql, ")")
				// value 填充子句
				valueSql := sql[index1:(index2 + 1)]
				// newValueArr 等于 value 填充子句复制 _batch_number 份
				newValueArr := make([]string, 0)
				for i := 0; i < v.Len(); i++ {
					newValueArr = append(newValueArr, valueSql)
				}
				newValueSql := strings.Join(newValueArr, ",")
				// 进行替换,相当于 () 替换成 (),(),(),()...
				sql = strings.Replace(sql, valueSql, newValueSql, -1)
			}
		}

	}
	return sql
}

func (this *SQLExecuteNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.FOREACH_PREFIX + "batch_data?", "仅供批量插入数据时使用"},
		2: {iworkconst.STRING_PREFIX + "sql", "执行sql语句"},
		3: {iworkconst.MULTI_PREFIX + "sql_binding?", "sql绑定数据,个数必须和当前执行sql语句中的占位符参数个数相同", "repeatable__" + iworkconst.FOREACH_PREFIX + "batch_data?"},
		4: {iworkconst.STRING_PREFIX + "db_conn", "数据库连接信息,需要使用 $RESOURCE 全局参数"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *SQLExecuteNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return schema.BuildParamOutputSchemaWithSlice([]string{iworkconst.NUMBER_PREFIX + "affected"})
}

func (this *SQLExecuteNode) ValidateCustom() (checkResult []string) {
	validateAndGetDataStoreName(this.WorkStep)
	return []string{}
}

type SQLQueryPageNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *SQLQueryPageNode) Execute(trackingId string) {
	// 需要存储的中间数据
	paramMap := make(map[string]interface{}, 0)
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, this.DataStore)
	sql := tmpDataMap[iworkconst.STRING_PREFIX+"sql"].(string)
	var total_sql string
	if _total_sql, ok := tmpDataMap[iworkconst.STRING_PREFIX+"total_sql?"].(string); ok && strings.TrimSpace(_total_sql) != "" {
		total_sql = _total_sql
	} else {
		total_sql = getTotalSqlFromQuery(sql)
	}
	dataSourceName := tmpDataMap[iworkconst.STRING_PREFIX+"db_conn"].(string)
	// sql_binding 参数获取
	sql_binding := getSqlBinding(tmpDataMap)
	totalcount := sqlutil.QuerySelectCount(total_sql, sql_binding, dataSourceName)
	sql = fmt.Sprintf(`%s limit ?,?`, sql) // 追加 limit 语句
	current_page := tmpDataMap[iworkconst.NUMBER_PREFIX+"current_page"].(string)
	page_size := tmpDataMap[iworkconst.NUMBER_PREFIX+"page_size"].(string)
	_current_page, _ := strconv.Atoi(current_page)
	_page_size, _ := strconv.Atoi(page_size)
	sql_binding = append(sql_binding, (_current_page-1)*_page_size, _page_size)
	sql = strings.ReplaceAll(sql, "{{", "")
	sql = strings.ReplaceAll(sql, "}}", "")
	datacounts, rowDatas := sqlutil.Query(sql, sql_binding, dataSourceName)
	// 将数据数据存储到数据中心
	// 存储 datacounts
	paramMap[iworkconst.NUMBER_PREFIX+"datacounts"] = datacounts
	// 数组对象整体存储在 rows 里面
	paramMap[iworkconst.MULTI_PREFIX+"rows"] = rowDatas
	// 存储分页信息
	pageIndex, pageSize := getPageIndexAndPageSize(tmpDataMap)
	paginator := pageutil.Paginator(pageIndex, pageSize, totalcount)
	paramMap[iworkconst.COMPLEX_PREFIX+"paginator"] = paginator

	for key, value := range paginator {
		paramMap[iworkconst.FIELD_PREFIX+"paginator."+key] = value
	}
	this.DataStore.CacheDatas(this.WorkStep.WorkStepName, paramMap)
}

func getPageIndexAndPageSize(tmpDataMap map[string]interface{}) (currentPage int, pageSize int) {
	var convert = func(data interface{}) (result int) {
		if _data, ok := data.(string); ok {
			result, _ = strconv.Atoi(_data)
		} else if _data, ok := data.(int); ok {
			result = _data
		} else if _data, ok := data.(int64); ok {
			result = int(_data)
		}
		return
	}
	currentPage = convert(tmpDataMap[iworkconst.NUMBER_PREFIX+"current_page"])
	pageSize = convert(tmpDataMap[iworkconst.NUMBER_PREFIX+"page_size"])
	return
}

func (this *SQLQueryPageNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "sql", "查询sql语句,可以使用 {{}} 标识出 metadata 字段名和 tableName 表名,用于自动推断 metadatasql, 带分页条件的sql,等价于 ${total_sql} limit ?,?"},
		2: {iworkconst.STRING_PREFIX + "metadata_sql?", "元数据sql语句,针对复杂查询sql,需要提供类似于select * from blog where 1=0的辅助sql用来构建节点输出"},
		3: {iworkconst.STRING_PREFIX + "total_sql?", "统计总数sql,返回N页总数据量,格式参考select count(*) as count from blog where xxx"},
		4: {iworkconst.MULTI_PREFIX + "sql_binding?", "sql绑定数据,个数和sql中的?数量相同,前N-2位参数和total_sql中的?数量相同"},
		5: {iworkconst.NUMBER_PREFIX + "current_page", "当前页数"},
		6: {iworkconst.NUMBER_PREFIX + "page_size", "每页数据量"},
		7: {iworkconst.STRING_PREFIX + "db_conn", "数据库连接信息,需要使用 $RESOURCE 全局参数"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *SQLQueryPageNode) GetRuntimeParamInputSchema() *iworkmodels.ParamInputSchema {
	return &iworkmodels.ParamInputSchema{}
}

func (this *SQLQueryPageNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	items := make([]iworkmodels.ParamOutputSchemaItem, 0)
	for _, paginatorField := range pageutil.GetPaginatorFields() {
		items = append(items, iworkmodels.ParamOutputSchemaItem{
			ParentPath: iworkconst.COMPLEX_PREFIX + "paginator",
			ParamName:  paginatorField,
		})
	}
	return &iworkmodels.ParamOutputSchema{ParamOutputSchemaItems: items}
}

func (this *SQLQueryPageNode) GetRuntimeParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return getMetaDataQuietlyForQuery(this.WorkStep)
}

func (this *SQLQueryPageNode) ValidateCustom() (checkResult []string) {
	validateAndGetDataStoreName(this.WorkStep)
	validateAndGetMetaDataSql(this.WorkStep)
	validateSqlBindingParamCount(this.WorkStep)
	validateSqlBindingParamCount(this.WorkStep)
	validateTotalSqlBindingParamCount(this.WorkStep)
	return
}

func validateTotalSqlBindingParamCount(step *iwork.WorkStep) {
	total_sql := param.GetStaticParamValueWithStep(iworkconst.STRING_PREFIX+"total_sql?", step).(string)
	sql_binding := param.GetStaticParamValueWithStep(iworkconst.MULTI_PREFIX+"sql_binding?", step).(string)
	if strings.Count(total_sql, "?")+2 != strings.Count(iworkfunc.EncodeSpecialForParamVaule(sql_binding), ";") {
		panic("Number of ? in total_sql and number of ; in sql_binding is mismatch!")
	}
}

func validateSqlBindingParamCount(step *iwork.WorkStep) {
	sql := param.GetStaticParamValueWithStep(iworkconst.STRING_PREFIX+"sql", step).(string)
	sql_binding := param.GetStaticParamValueWithStep(iworkconst.MULTI_PREFIX+"sql_binding?", step).(string)
	if strings.Count(sql, "?") != strings.Count(iworkfunc.EncodeSpecialForParamVaule(sql_binding), ";") {
		panic("Number of ? in SQL and number of ; in sql_binding is unequal!")
	}
}

func getTotalSqlFromQuery(querySql string) string {
	reg := regexp.MustCompile("^.+{{.+}}.+{{.+}}.*$")
	if match := reg.Match([]byte(querySql)); match == false {
		panic(fmt.Sprintf(`sql 格式不正确,必须使用 {{}} 标识出字段名和表名！`))
	}
	totalSql := querySql[:strings.Index(querySql, "{{")] + " count(*) as count " + querySql[strings.Index(querySql, "}}")+2:]
	totalSql = strings.ReplaceAll(totalSql, "{{", "")
	totalSql = strings.ReplaceAll(totalSql, "}}", "")
	return totalSql
}

func getMetaDataSqlFromQuery(querySql string) string {
	reg := regexp.MustCompile("^.+{{.+}}.+{{.+}}.*$")
	if match := reg.Match([]byte(querySql)); match == false {
		panic(fmt.Sprintf(`sql 格式不正确,必须使用 {{}} 标识出字段名和表名！`))
	}
	fieldNames := querySql[strings.Index(querySql, "{{")+2 : strings.Index(querySql, "}}")]
	tableName := querySql[strings.LastIndex(querySql, "{{")+2 : strings.LastIndex(querySql, "}}")]
	return fmt.Sprintf(`select %s from %s where 1=0`, fieldNames, tableName)
}

func validateAndGetMetaDataSql(step *iwork.WorkStep) string {
	metadata_sql := param.GetStaticParamValueWithStep(iworkconst.STRING_PREFIX+"metadata_sql", step).(string)
	if strings.TrimSpace(metadata_sql) == "" {
		metadata_sql = getMetaDataSqlFromQuery(param.GetStaticParamValueWithStep(iworkconst.STRING_PREFIX+"sql", step).(string))
	}
	if strings.TrimSpace(metadata_sql) == "" {
		panic("Empty paramValue for metadata_sql was found!")
	}
	if strings.Contains(metadata_sql, "?") {
		panic("Invalid paramValue form metadata_sql was found!")
	}
	return strings.TrimSpace(metadata_sql)
}

func validateAndGetDataStoreName(step *iwork.WorkStep) string {
	dataSourceName := param.GetStaticParamValueWithStep(iworkconst.STRING_PREFIX+"db_conn", step).(string)
	if strings.TrimSpace(dataSourceName) == "" {
		panic("Invalid param for db_conn! Can't resolve it!")
	}
	db, err := sqlutil.GetConnForMysql("mysql", dataSourceName)
	if err != nil {
		panic(fmt.Sprintf("Can't get DB connection for %s!", dataSourceName))
	}
	defer db.Close()
	return dataSourceName
}

func getMetaDataForQuery(step *iwork.WorkStep) *iworkmodels.ParamOutputSchema {
	metadataSql := validateAndGetMetaDataSql(step)
	dataSourceName := validateAndGetDataStoreName(step)
	paramNames := sqlutil.GetMetaDatas(metadataSql, dataSourceName)
	items := make([]iworkmodels.ParamOutputSchemaItem, 0)
	for _, paramName := range paramNames {
		items = append(items, iworkmodels.ParamOutputSchemaItem{
			ParentPath: iworkconst.MULTI_PREFIX + "rows",
			ParamName:  paramName,
		})
	}
	return &iworkmodels.ParamOutputSchema{ParamOutputSchemaItems: items}
}

func getMetaDataQuietlyForQuery(step *iwork.WorkStep) *iworkmodels.ParamOutputSchema {
	return getMetaDataForQuery(step)
}
