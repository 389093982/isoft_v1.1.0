package iworknode

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"isoft/isoft/common/stringutil"
	"isoft/isoft_iaas_web/core/iworkconst"
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/core/iworkdata/param"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworkfunc"
	"isoft/isoft_iaas_web/core/iworklog"
	"isoft/isoft_iaas_web/core/iworkmodels"
	"isoft/isoft_iaas_web/core/iworkplugin/iworkprotocol"
	"isoft/isoft_iaas_web/core/iworkutil"
	"isoft/isoft_iaas_web/core/iworkvalid"
	"isoft/isoft_iaas_web/models/iwork"
	"strconv"
	"strings"
)

// 所有 node 的基类
type BaseNode struct {
	iworkprotocol.IWorkStep
	DataStore *datastore.DataStore
	o         orm.Ormer
	LogWriter *iworklog.CacheLoggerWriter
}

func (this *BaseNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	fmt.Println("execute default GetDefaultParamInputSchema method...")
	return &iworkmodels.ParamInputSchema{}
}

func (this *BaseNode) GetRuntimeParamInputSchema() *iworkmodels.ParamInputSchema {
	fmt.Println("execute default GetRuntimeParamInputSchema method...")
	return &iworkmodels.ParamInputSchema{}
}

func (this *BaseNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	fmt.Println("execute default GetDefaultParamOutputSchema method...")
	return &iworkmodels.ParamOutputSchema{}
}

func (this *BaseNode) GetRuntimeParamOutputSchema() *iworkmodels.ParamOutputSchema {
	fmt.Println("execute default GetRuntimeParamOutputSchema method...")
	return &iworkmodels.ParamOutputSchema{}
}

func (this *BaseNode) ValidateCustom() (checkResult []string) {
	fmt.Println("execute default ValidateCustom method...")
	return
}

func (this *BaseNode) GetOrmer() orm.Ormer {
	if this.o == nil {
		this.o = orm.NewOrm()
	}
	return this.o
}

// paramValue 来源于 iwork 模块
func (this *BaseNode) parseAndFillParamVauleWithResource(paramVaule string) interface{} {
	resource, err := iwork.QueryResourceByName(strings.Replace(paramVaule, "$RESOURCE.", "", -1))
	if err == nil {
		return resource.ResourceDsn
	}
	return ""
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
	names := []string{"sql", "count_sql", "metadata_sql?"}
	for _, name := range names {
		if name == paramName {
			return true
		}
	}
	return false
}

// 解析 paramVaule 并从 dataStore 中获取实际值
func (this *BaseNode) ParseAndGetParamVaule(paramName, paramVaule string, dataStore *datastore.DataStore) interface{} {
	if checkSkipParse(paramName) {
		return paramVaule
	}
	values := this.parseParamValueToMulti(paramVaule)
	// 单值
	if len(values) == 1 {
		return this.parseAndGetSingleParamVaule(values[0], dataStore)
	}
	// 多值
	results := make([]interface{}, 0)
	for _, value := range values {
		result := this.parseAndGetSingleParamVaule(value, dataStore)
		results = append(results, result)
	}
	return results
}

func (this *BaseNode) parseParamValueToMulti(paramVaule string) []string {
	results := []string{}
	// 对转义字符 \, \; \( \) 等进行编码
	paramVaule = iworkfunc.EncodeSpecialForParamVaule(paramVaule)
	vaules := strings.Split(paramVaule, ";")
	for _, value := range vaules {
		if _value := this.removeUnsupportChars(value); strings.TrimSpace(_value) != "" {
			results = append(results, strings.TrimSpace(_value))
		}
	}
	return results
}

func (this *BaseNode) _parseAndGetSingleParamVaule(paramVaule string, dataStore *datastore.DataStore) interface{} {
	paramVaule = iworkfunc.DncodeSpecialForParamVaule(paramVaule)
	if strings.HasPrefix(strings.ToUpper(paramVaule), "$RESOURCE.") {
		return this.parseAndFillParamVauleWithResource(paramVaule)
	} else if strings.HasPrefix(strings.ToUpper(paramVaule), "$WORK.") {
		return iworkutil.GetWorkSubNameFromParamValue(paramVaule)
	} else if strings.HasPrefix(strings.ToUpper(paramVaule), "$ENTITY.") {
		return iworkutil.GetParamValueForEntity(paramVaule)
	} else if strings.HasPrefix(strings.ToUpper(paramVaule), "$WORKVARS.") {
		return iworkutil.GetParamValueForWorkVars(paramVaule, this.DataStore)
	}
	return this.parseAndFillParamVauleWithNode(paramVaule, dataStore)
}

func (this *BaseNode) parseAndGetSingleParamVaule(paramVaule string, dataStore *datastore.DataStore) interface{} {
	defer func() {
		if err := recover(); err != nil {
			panic(fmt.Sprintf("execute func with expression is %s, error msg is :%s", paramVaule, err.(error).Error()))
		}
	}()
	// 对单个 paramVaule 进行特殊字符编码
	paramVaule = iworkfunc.EncodeSpecialForParamVaule(paramVaule)
	executors := iworkfunc.GetAllFuncExecutor(paramVaule)
	if executors == nil || len(executors) == 0 {
		// 是直接参数,不需要函数进行特殊处理
		return this._parseAndGetSingleParamVaule(paramVaule, dataStore)
	}
	historyFuncResultMap := make(map[string]interface{}, 0)
	var lastFuncResult interface{}
	// 按照顺序依次执行函数
	for _, executor := range executors {
		// executor 所有参数进行 trim 操作
		iworkfunc.GetTrimFuncExecutor(executor)
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
		result := iworkfunc.CallFuncExecutor(executor, args)
		historyFuncResultMap[executor.FuncUUID] = result
		lastFuncResult = result
	}
	return lastFuncResult
}

// 将 ParamInputSchema 填充数据并返回临时的数据中心 tmpDataMap
// skips 表示可以跳过填充的参数
func (this *BaseNode) FillParamInputSchemaDataToTmp(workStep *iwork.WorkStep, dataStore *datastore.DataStore, skips ...string) map[string]interface{} {
	// 存储节点中间数据
	tmpDataMap := make(map[string]interface{})
	paramInputSchema := schema.GetCacheParamInputSchema(workStep, &WorkStepFactory{WorkStep: workStep})
	for _, item := range paramInputSchema.ParamInputSchemaItems {
		// 跳过校验
		if stringutil.CheckContains(item.ParamName, skips) {
			continue
		}
		// 对参数进行非空校验
		if ok, checkResults := iworkvalid.CheckEmptyForItem(item); !ok {
			panic(strings.Join(checkResults, ";"))
		}
		// 个性化重写操作
		this.modifySqlBindingParamValueWithBatchNumber(&item, tmpDataMap)
		tmpDataMap[item.ParamName] = this.ParseAndGetParamVaule(item.ParamName, item.ParamValue, dataStore) // 输入数据存临时
	}
	return tmpDataMap
}

func (this *BaseNode) modifySqlBindingParamValueWithBatchNumber(item *iworkmodels.ParamInputSchemaItem, tmpDataMap map[string]interface{}) {
	// 当前填充的字段为 sql_binding? 时,检测到批量操作数据大于 1
	if item.ParamName == iworkconst.MULTI_PREFIX+"sql_binding?" && GetBatchNumber(tmpDataMap) > 1 {
		var newParamValue string
		for i := 0; i < GetBatchNumber(tmpDataMap); i++ {
			newParamValue += strings.Replace(item.ParamValue, iworkconst.MULTI_PREFIX+"rows.", fmt.Sprintf(iworkconst.MULTI_PREFIX+"rows[%v].", i), -1)
		}
		item.ParamValue = newParamValue
	}
}

// 从 tmpDataMap 获取 batch_number? 数据
func GetBatchNumber(tmpDataMap map[string]interface{}) int {
	if batch_number, ok := tmpDataMap[iworkconst.NUMBER_PREFIX+"batch_number?"].(int64); ok {
		return int(batch_number)
	}
	if batch_number, ok := tmpDataMap[iworkconst.NUMBER_PREFIX+"batch_number?"].(string); ok {
		if _batch_number, err := strconv.Atoi(batch_number); err == nil {
			return _batch_number
		}
	}
	return 0
}

// 提交输出数据至数据中心,此类数据能直接从 tmpDataMap 中获取,而不依赖于计算,只适用于 WORK_START、WORK_END 节点
func (this *BaseNode) SubmitParamOutputSchemaDataToDataStore(workStep *iwork.WorkStep, dataStore *datastore.DataStore, tmpDataMap map[string]interface{}) {
	paramOutputSchema := schema.GetCacheParamOutputSchema(workStep)
	paramMap := make(map[string]interface{})
	for _, item := range paramOutputSchema.ParamOutputSchemaItems {
		paramMap[item.ParamName] = tmpDataMap[item.ParamName]
	}
	// 将数据数据存储到数据中心
	dataStore.CacheDatas(workStep.WorkStepName, paramMap)
}

// 去除不合理的字符
func (this *BaseNode) removeUnsupportChars(paramValue string) string {
	// 先进行初次的 trim
	paramValue = strings.TrimSpace(paramValue)
	// 去除前后的 \n9
	paramValue = strings.TrimPrefix(paramValue, "\n")
	paramValue = strings.TrimSuffix(paramValue, "\n")
	// 再进行二次 trim
	paramValue = strings.TrimSpace(paramValue)
	return paramValue
}
