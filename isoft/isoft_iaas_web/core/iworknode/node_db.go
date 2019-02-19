package iworknode

import (
	"fmt"
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/core/iworkdata/param"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworkutil/sqlutil"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
)

type DBParserNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *DBParserNode) Execute(trackingId string) {
	// 数据中心
	dataStore := datastore.GetDataSource(trackingId)
	dataSourceName := param.GetStaticParamValue("db_conn", this.WorkStep)
	_, _, rowDatas := sqlutil.Query("show tables;", []interface{}{}, dataSourceName)
	tableNames := make([]string,0)
	for _,rowData := range rowDatas{
		for _, tableName := range rowData{
			tableNames = append(tableNames, tableName.(string))
		}
	}
	tablecolsmap := make(map[string]string,0)
	for _, tableName := range tableNames{
		cols := sqlutil.GetMetaDatas(fmt.Sprintf("select * from %s where 1=0",tableName),dataSourceName)
		tablecolsmap[tableName] = strings.Join(cols, ",")
	}
	dataStore.CacheData(this.WorkStep.WorkStepName, "tablecolsmap", tablecolsmap)
	// 数组对象整体存储在 rows 里面
	dataStore.CacheData(this.WorkStep.WorkStepName, "tables", strings.Join(tableNames, ","))
}

func (this *DBParserNode) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	paramMap := map[int][]string{
		1:[]string{"db_conn","数据库连接信息,需要使用 $RESOURCE 全局参数"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *DBParserNode) GetRuntimeParamInputSchema() *schema.ParamInputSchema {
	return &schema.ParamInputSchema{}
}

func (this *DBParserNode) GetDefaultParamOutputSchema() *schema.ParamOutputSchema {
	return schema.BuildParamOutputSchemaWithSlice([]string{"tables","tablecolsmap"})
}

func (this *DBParserNode) GetRuntimeParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}

func (this *DBParserNode) ValidateCustom() {
	validateAndGetDataSourceName(this.WorkStep)
}
