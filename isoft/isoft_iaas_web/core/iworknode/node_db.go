package iworknode

import (
	"fmt"
	"isoft/isoft_iaas_web/core/iworkconst"
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/core/iworkdata/param"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworkutil/sqlutil"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
	"time"
)

type DBParserNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *DBParserNode) Execute(trackingId string) {
	// 数据中心
	dataStore := datastore.GetDataStore(trackingId)
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, dataStore)
	dataSourceName := param.GetStaticParamValue(iworkconst.STRING_PREFIX+"db_conn", this.WorkStep).(string)
	_, _, rowDatas := sqlutil.Query("show tables;", []interface{}{}, dataSourceName)
	tableNames := make([]string, 0)
	for _, rowData := range rowDatas {
		for _, tableName := range rowData {
			tableNames = append(tableNames, tableName.(string))
		}
	}
	tablecolsmap := make(map[string]string, 0)
	for _, tableName := range tableNames {
		cols := sqlutil.GetMetaDatas(fmt.Sprintf("select * from %s where 1=0", tableName), dataSourceName)
		tablecolsmap[tableName] = strings.Join(cols, ",")
	}
	// 将其自动存为实体类
	saveEntity(tmpDataMap, tablecolsmap)
	// 存进 dataStore
	dataStore.CacheData(this.WorkStep.WorkStepName, iworkconst.MULTI_PREFIX+"tablecolsmap", tablecolsmap)
	// 数组对象整体存储在 rows 里面
	dataStore.CacheData(this.WorkStep.WorkStepName, iworkconst.STRING_PREFIX+"tables", strings.Join(tableNames, ","))
}

func saveEntity(tmpDataMap map[string]interface{}, tablecolsmap map[string]string) {
	if save_entity, ok := tmpDataMap[iworkconst.BOOL_PREFIX+"save_entity?"].(string); !ok || strings.TrimSpace(save_entity) == "" {
		return
	}
	for tableName, tablecols := range tablecolsmap {
		entity := &iwork.Entity{
			EntityName:      tableName,
			EntityFieldStr:  tablecols,
			CreatedBy:       "SYSTEM",
			CreatedTime:     time.Now(),
			LastUpdatedBy:   "SYSTEM",
			LastUpdatedTime: time.Now(),
		}
		if _entity, err := iwork.QueryEntityByName(tableName); err == nil {
			entity.Id = _entity.Id
		}
		iwork.InsertOrUpdateEntity(entity)
	}
}

func (this *DBParserNode) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "db_conn", "数据库连接信息,需要使用 $RESOURCE 全局参数"},
		2: {iworkconst.BOOL_PREFIX + "save_entity?", "是否将分析的结果映射成实体类?"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *DBParserNode) GetRuntimeParamInputSchema() *schema.ParamInputSchema {
	return &schema.ParamInputSchema{}
}

func (this *DBParserNode) GetDefaultParamOutputSchema() *schema.ParamOutputSchema {
	return schema.BuildParamOutputSchemaWithSlice([]string{iworkconst.STRING_PREFIX + "tables", iworkconst.MULTI_PREFIX + "tablecolsmap"})
}

func (this *DBParserNode) GetRuntimeParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}

func (this *DBParserNode) ValidateCustom() (checkResult []string) {
	validateAndGetDataStoreName(this.WorkStep)
	return []string{}
}
