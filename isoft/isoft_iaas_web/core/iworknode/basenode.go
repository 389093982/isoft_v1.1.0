package iworknode

import (
	"errors"
	"fmt"
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/core/iworkdata/param"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworkutil"
	"isoft/isoft_iaas_web/core/iworkutil/funcutil"
	"isoft/isoft_iaas_web/models/iwork"
	"strconv"
	"strings"
)

// 所有 node 的基类
type BaseNode struct {}

// paramValue 来源于 iwork 模块
func (this *BaseNode) parseAndFillParamVauleWithResource(paramVaule string) interface{} {
	return iwork.GetResourceDataSourceNameString(strings.Replace(paramVaule, "$RESOURCE.", "", -1))
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

// 判断是否需要跳过解析
func checkSkipParse(paramName string) bool {
	names := []string{"sql","count_sql","metadata_sql?"}
	for _,name := range names{
		if name == paramName{
			return true
		}
	}
	return false
}

// 解析 paramVaule 并从 dataStore 中获取实际值
func (this *BaseNode) ParseAndGetParamVaule(paramName, paramVaule string, dataStore *datastore.DataStore) interface{} {
	if checkSkipParse(paramName){
		return paramVaule
	}
	values := this.parseParamValueToMulti(paramVaule)
	if len(values) == 1 {
		// 单值
		return this.parseAndGetSingleParamVaule(values[0], dataStore)
	} else {
		// 多值
		results := make([]interface{}, 0)
		for _, value := range values {
			result := this.parseAndGetSingleParamVaule(value, dataStore)
			results = append(results, result)
		}
		return results
	}
}

func (this *BaseNode) parseParamValueToMulti(paramVaule string) []string {
	results := []string{}
	// 对转义字符 \, \; \( \) 等进行编码
	paramVaule = funcutil.EncodeSpecialForParamVaule(paramVaule)
	vaules := strings.Split(paramVaule, ";")
	for _, value := range vaules {
		if _value := this.removeUnsupportChars(value); strings.TrimSpace(_value) != "" {
			results = append(results, strings.TrimSpace(_value))
		}
	}
	return results
}

func (this *BaseNode) _parseAndGetSingleParamVaule(paramVaule string, dataStore *datastore.DataStore) interface{} {
	paramVaule = funcutil.DncodeSpecialForParamVaule(paramVaule)
	if strings.HasPrefix(strings.ToUpper(paramVaule), "$RESOURCE.") {
		return this.parseAndFillParamVauleWithResource(paramVaule)
	} else if strings.HasPrefix(strings.ToUpper(paramVaule), "$WORK.") {
		return iworkutil.GetWorkSubNameFromParamValue(paramVaule)
	} else if strings.HasPrefix(strings.ToUpper(paramVaule), "$ENTITY.") {
		return iworkutil.GetParamValueForEntity(paramVaule)
	}
	return this.parseAndFillParamVauleWithNode(paramVaule, dataStore)
}

func (this *BaseNode) parseAndGetSingleParamVaule(paramVaule string, dataStore *datastore.DataStore) interface{} {
	// 对单个 paramVaule 进行特殊字符编码
	paramVaule = funcutil.EncodeSpecialForParamVaule(paramVaule)
	executors := funcutil.GetAllFuncExecutor(paramVaule)
	if executors == nil || len(executors) == 0 {
		// 是直接参数,不需要函数进行特殊处理
		return this._parseAndGetSingleParamVaule(paramVaule, dataStore)
	} else {
		historyFuncResultMap := make(map[string]interface{}, 0)
		var lastFuncResult interface{}
		// 按照顺序依次执行函数
		for _, executor := range executors {
			// executor 所有参数进行 trim 操作
			funcutil.GetTrimFuncExecutor(executor)
			args := make([]interface{}, 0)
			// 函数参数替换成实际意义上的值
			for _, arg := range executor.FuncArgs {
				// 判断参数是否来源于 historyFuncResultMap
				if _arg, ok := historyFuncResultMap[arg]; ok {
					args = append(args, _arg)
				} else {
					args = append(args, this._parseAndGetSingleParamVaule(arg, dataStore))
				}
			}
			// 执行函数并记录结果,供下一个函数执行使用
			result := funcutil.CallFuncExecutor(executor, args)
			historyFuncResultMap[executor.FuncUUID] = result
			lastFuncResult = result
		}
		return lastFuncResult
	}
}

// 对输入参数做非空校验
func checkEmptyForParam(item schema.ParamInputSchemaItem) {
	if !strings.HasSuffix(item.ParamName, "?") && strings.TrimSpace(item.ParamValue) == ""{
		panic(errors.New(fmt.Sprintf("empty param for %s", item.ParamName)))
	}
}

// 将 ParamInputSchema 填充数据并返回临时的数据中心 tmpDataMap
func (this *BaseNode) FillParamInputSchemaDataToTmp(workStep *iwork.WorkStep, dataStore *datastore.DataStore) map[string]interface{} {
	// 存储节点中间数据
	tmpDataMap := make(map[string]interface{})
	paramInputSchema := schema.GetCacheParamInputSchema(workStep, &WorkStepFactory{WorkStep: workStep})
	for _, item := range paramInputSchema.ParamInputSchemaItems {
		// 对参数进行非空校验
		checkEmptyForParam(item)
		// 个性化重写操作
		this.modifySqlBindingParamValueWithBatchNumber(&item, tmpDataMap)
		tmpDataMap[item.ParamName] = this.ParseAndGetParamVaule(item.ParamName, item.ParamValue, dataStore) // 输入数据存临时
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
