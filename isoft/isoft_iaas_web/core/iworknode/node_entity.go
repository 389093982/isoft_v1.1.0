package iworknode

import (
	"encoding/json"
	"fmt"
	"isoft/isoft_iaas_web/core/iworkconst"
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworkutil"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
)

type EntityParserNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *EntityParserNode) Execute(trackingId string, skipFunc func(tmpDataMap map[string]interface{}) bool) {
	// 获取数据中心
	dataStore := datastore.GetDataSource(trackingId)
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, dataStore)
	if skipFunc(tmpDataMap){return}			// 跳过当前节点执行
	inputSchema := schema.GetCacheParamInputSchema(this.WorkStep, &WorkStepFactory{WorkStep: this.WorkStep})
	for _, item := range inputSchema.ParamInputSchemaItems {
		if strings.HasSuffix(item.ParamName, "_entity"){
			entityName := getEntityNameWithRemovePrefixAndSuffix(item.ParamName)
			// 从 tmpDataMap 中获取入参实体类数据
			entityDataMap := make(map[string]interface{})
			if dataMap, ok := tmpDataMap[iworkconst.COMPLEX_PREFIX + entityName + "_data"].([]map[string]interface{});ok && len(dataMap) > 0{
				entityDataMap = dataMap[0]
			}else if dataMap, ok := tmpDataMap[iworkconst.COMPLEX_PREFIX + entityName + "_data"].(map[string]interface{});ok{
				entityDataMap = dataMap
			}
			entityFieldStr := tmpDataMap[iworkconst.STRING_PREFIX + entityName + "_entity"].(string)
			for _, entityField := range strings.Split(entityFieldStr, ","){
				// 将数据数据存储到数据中心
				dataStore.CacheData(this.WorkStep.WorkStepName,
					fmt.Sprintf("%s.%s", iworkconst.COMPLEX_PREFIX + entityName,
						strings.TrimSpace(entityField)), entityDataMap[entityField])
			}
		}
	}
}

func (this *EntityParserNode) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	return &schema.ParamInputSchema{}
}

func (this *EntityParserNode) GetRuntimeParamInputSchema() *schema.ParamInputSchema {
	var paramMappingsArr []string
	json.Unmarshal([]byte(this.WorkStep.WorkStepParamMapping), &paramMappingsArr)
	items := make([]schema.ParamInputSchemaItem, 0)
	for _, paramMapping := range paramMappingsArr {
		// paramMapping 存放实体类定义 $Entity
		items = append(items, schema.ParamInputSchemaItem{ParamName: fmt.Sprintf(iworkconst.STRING_PREFIX + "%s_entity",paramMapping)})
		// paramMapping_data 存放实体类数据
		items = append(items, schema.ParamInputSchemaItem{ParamName: fmt.Sprintf(iworkconst.COMPLEX_PREFIX + "%s_data",paramMapping)})
	}
	return &schema.ParamInputSchema{ParamInputSchemaItems: items}
}

func (this *EntityParserNode) GetDefaultParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}

func (this *EntityParserNode) GetRuntimeParamOutputSchema() *schema.ParamOutputSchema {
	items := make([]schema.ParamOutputSchemaItem, 0)
	inputSchema := schema.GetCacheParamInputSchema(this.WorkStep, &WorkStepFactory{WorkStep: this.WorkStep})
	for _, item := range inputSchema.ParamInputSchemaItems {
		if !strings.HasSuffix(item.ParamName, "_data"){			// _data 需要排除
			// 从用户输入值中提取实体类字段详细信息
			if entityFieldStr := iworkutil.GetParamValueForEntity(item.ParamValue); strings.TrimSpace(entityFieldStr) != ""{
				for _, entityField := range strings.Split(entityFieldStr, ","){
					// 每个字段放入 items 中
					items = append(items, schema.ParamOutputSchemaItem{
						ParamName: strings.TrimSpace(entityField), ParentPath: iworkconst.COMPLEX_PREFIX + getEntityNameWithRemovePrefixAndSuffix(item.ParamName),
					})
				}
			}
		}
	}
	return &schema.ParamOutputSchema{ParamOutputSchemaItems: items}
}

func getEntityNameWithRemovePrefixAndSuffix(paramName string) string {
	// 去除 str_ 和 _entity
	paramName = strings.Replace(paramName, iworkconst.STRING_PREFIX, "", -1)
	paramName = strings.Replace(paramName, "_entity", "", -1)
	return paramName
}

func (this *EntityParserNode) ValidateCustom() {

}
