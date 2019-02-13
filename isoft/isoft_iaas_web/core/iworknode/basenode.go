package iworknode

import (
	"fmt"
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/core/iworkdata/param"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/models/iresource"
	"isoft/isoft_iaas_web/models/iwork"
	"strconv"
	"strings"
)

// 所有 node 的基类
type BaseNode struct {
}

// paramValue 来源于 iresource 模块
func (this *BaseNode) parseAndFillParamVauleWithResource(paramVaule string) interface{} {
	return iresource.GetResourceDataSourceNameString(strings.Replace(paramVaule, "$RESOURCE.", "", -1))
}

// paramValue 来源于前置节点
func (this *BaseNode) parseAndFillParamVauleWithNode(paramVaule string, dataStore *datastore.DataStore) interface{} {
	if strings.HasPrefix(paramVaule, "$") {
		resolver := param.ParamVauleParser{ParamValue: paramVaule}
		return dataStore.GetData(resolver.GetNodeNameFromParamValue(), resolver.GetParamNameFromParamValue())
	} else {
		return paramVaule
	}
}

// 解析 paramVaule 并从 dataStore 中获取实际值
func (this *BaseNode) ParseAndFillParamVaule(paramVaule string, dataStore *datastore.DataStore) interface{} {
	values := this.parseParamValueToMulti(paramVaule)
	if len(values) == 1 {
		// 单值
		return this.parseAndFillSingleParamVaule(values[0], dataStore)
	} else {
		// 多值
		results := make([]interface{}, 0)
		for _, value := range values {
			result := this.parseAndFillSingleParamVaule(value, dataStore)
			results = append(results, result)
		}
		return results
	}
}

func (this *BaseNode) parseParamValueToMulti(paramVaule string) []string {
	results := []string{}
	vaules := strings.Split(paramVaule, "__sep__")
	for _, value := range vaules {
		if _value := this.removeUnsupportChars(value); strings.TrimSpace(_value) != "" {
			results = append(results, strings.TrimSpace(_value))
		}
	}
	return results
}

func (this *BaseNode) parseAndFillSingleParamVaule(paramVaule string, dataStore *datastore.DataStore) interface{} {
	if strings.HasPrefix(strings.ToUpper(paramVaule), "$RESOURCE.") {
		return this.parseAndFillParamVauleWithResource(paramVaule)
	}
	return this.parseAndFillParamVauleWithNode(paramVaule, dataStore)
}

// 将 ParamInputSchema 填充数据并返回临时的数据中心 tmpDataMap
func (this *BaseNode) FillParamInputSchemaDataToTmp(workStep *iwork.WorkStep, dataStore *datastore.DataStore) map[string]interface{} {
	// 存储节点中间数据
	tmpDataMap := make(map[string]interface{})
	paramInputSchema := schema.GetCacheParamInputSchema(workStep, &WorkStepFactory{WorkStep: workStep})
	for _, item := range paramInputSchema.ParamInputSchemaItems {
		// 个性化重写操作
		this.modifySqlBindingParamValueWithBatchNumber(&item, tmpDataMap)
		tmpDataMap[item.ParamName] = this.ParseAndFillParamVaule(item.ParamValue, dataStore) // 输入数据存临时
	}
	return tmpDataMap
}

func (this *BaseNode) modifySqlBindingParamValueWithBatchNumber(item *schema.ParamInputSchemaItem, tmpDataMap map[string]interface{}) {
	// 当前填充的字段为 sql_binding? 时,检测到批量操作数据大于 1
	if item.ParamName == "sql_binding?" && GetBatchNumber(tmpDataMap) > 1 {
		var newParamValue string
		for i := 0; i < GetBatchNumber(tmpDataMap); i++ {
			newParamValue += strings.Replace(item.ParamValue, ".rows.", fmt.Sprintf(".rows[%v].", i), -1)
		}
		item.ParamValue = newParamValue
	}
}

// 从 tmpDataMap 获取 batch_number? 数据
func GetBatchNumber(tmpDataMap map[string]interface{}) int {
	if _, ok := tmpDataMap["batch_number?"]; !ok {
		return 0
	}
	if batch_number, ok := tmpDataMap["batch_number?"].(int64); ok {
		return int(batch_number)
	}
	if batch_number, ok := tmpDataMap["batch_number?"].(string); ok {
		if _batch_number, err := strconv.Atoi(batch_number); err == nil {
			return _batch_number
		}
	}
	return 0
}

// 提交输出数据至数据中心,此类数据能直接从 tmpDataMap 中获取,而不依赖于计算,只适用于 WORK_START、WORK_END 节点
func (this *BaseNode) SubmitParamOutputSchemaDataToDataStore(workStep *iwork.WorkStep, dataStore *datastore.DataStore, tmpDataMap map[string]interface{}) {
	paramOutputSchema := schema.GetCacheParamOutputSchema(workStep)
	for _, item := range paramOutputSchema.ParamOutputSchemaItems {
		// 将数据数据存储到数据中心
		dataStore.CacheData(workStep.WorkStepName, item.ParamName, tmpDataMap[item.ParamName])
	}
}

// 去除不合理的字符
func (this *BaseNode) removeUnsupportChars(paramValue string) string {
	// 先进行初次的 trim
	paramValue = strings.TrimSpace(paramValue)
	// 去除前后的 \n
	paramValue = strings.TrimPrefix(paramValue, "\n")
	paramValue = strings.TrimSuffix(paramValue, "\n")
	// 再进行二次 trim
	paramValue = strings.TrimSpace(paramValue)
	return paramValue
}
