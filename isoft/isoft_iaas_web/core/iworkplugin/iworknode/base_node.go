package iworknode

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
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
	"reflect"
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
		if resource.ResourceType == "db" {
			return resource.ResourceDsn
		} else if resource.ResourceType == "sftp" || resource.ResourceType == "ssh" {
			return resource
		}
	}
	return ""
}

// paramValue 来源于前置节点
func (this *BaseNode) parseAndFillParamVauleWithPrefixNode(paramName, paramVaule string, dataStore *datastore.DataStore) interface{} {
	// 格式校验
	if !strings.HasPrefix(paramVaule, "$") {
		panic(errors.New(fmt.Sprintf("%s ~ %s is not start with $", paramName, paramVaule)))
	}
	resolver := param.ParamVauleParser{ParamValue: paramVaule}
	nodeName := resolver.GetNodeNameFromParamValue()
	paramName = resolver.GetParamNameFromParamValue()
	paramValue := dataStore.GetData(nodeName, paramName) // 作为直接对象, dataStore 里面可以直接获取
	if paramValue != nil {
		return paramValue
	}
	_paramName := paramName[:strings.LastIndex(paramName, ".")]
	datas := dataStore.GetData(nodeName, _paramName) // 作为对象属性
	attr := paramName[strings.LastIndex(paramName, ".")+1:]
	if reflect.TypeOf(datas).Kind() == reflect.Slice {
		return reflect.ValueOf(datas).Index(0).Interface().(map[string]interface{})[attr]
	}
	return datas.(map[string]interface{})[attr]
}

// 解析 paramVaule 并从 dataStore 中获取实际值
func (this *BaseNode) ParseAndGetParamVaule(paramName, paramVaule string, dataStore *datastore.DataStore, replaceMap ...map[string]interface{}) interface{} {
	values := this.parseParamValueToMulti(paramVaule)
	// 单值
	if len(values) == 1 {
		return this.parseAndGetSingleParamVaule(paramName, values[0], dataStore, replaceMap...)
	}
	// 多值
	results := make([]interface{}, 0)
	for _, value := range values {
		result := this.parseAndGetSingleParamVaule(paramName, value, dataStore, replaceMap...)
		results = append(results, result)
	}
	return results
}

func (this *BaseNode) parseParamValueToMulti(paramVaule string) []string {
	results := make([]string, 0)
	// 对转义字符 \, \; \( \) 等进行编码
	paramVaule = iworkfunc.EncodeSpecialForParamVaule(paramVaule)
	multiVals, err := iworkfunc.SplitWithLexerAnalysis(paramVaule)
	if err != nil {
		panic(err)
	}
	for _, value := range multiVals {
		if _value := this.trim(value); strings.TrimSpace(_value) != "" {
			results = append(results, strings.TrimSpace(_value))
		}
	}
	return results
}

func (this *BaseNode) callParseAndGetSingleParamVaule(paramName, paramVaule string, dataStore *datastore.DataStore, replaceMap ...map[string]interface{}) interface{} {
	paramVaule = iworkfunc.DncodeSpecialForParamVaule(paramVaule)
	// 变量
	if strings.HasPrefix(strings.ToUpper(paramVaule), "$RESOURCE.") {
		return this.parseAndFillParamVauleWithResource(paramVaule)
	} else if strings.HasPrefix(strings.ToUpper(paramVaule), "$WORK.") {
		return iworkutil.GetWorkSubNameFromParamValue(paramVaule)
	} else if strings.HasPrefix(strings.ToUpper(paramVaule), "$ENTITY.") {
		return iworkutil.GetParamValueForEntity(paramVaule)
	} else if strings.HasPrefix(strings.ToUpper(paramVaule), "$") {
		if len(replaceMap) > 0 {
			if paramVaule := this.parseAndFillParamVauleWithReplaceProviderNode(paramVaule, replaceMap...); paramVaule != nil {
				return paramVaule
			}
		}
		return this.parseAndFillParamVauleWithPrefixNode(paramName, paramVaule, dataStore)
	} else if strings.HasPrefix(paramVaule, "`") && strings.HasSuffix(paramVaule, "`") {
		// 字符串
		return paramVaule[1 : len(paramVaule)-1]
	} else {
		// 数字
		return paramVaule
	}
}

func (this *BaseNode) parseAndFillParamVauleWithReplaceProviderNode(paramVaule string, replaceMap ...map[string]interface{}) interface{} {
	for replaceProviderNodeName, replaceProviderData := range replaceMap[0] {
		replaceProviderNodeName = strings.ReplaceAll(replaceProviderNodeName, ";", "")
		if strings.HasPrefix(paramVaule, replaceProviderNodeName) {
			attr := strings.Replace(paramVaule, replaceProviderNodeName+".", "", 1)
			attr = strings.ReplaceAll(attr, ";", "")
			return replaceProviderData.(map[string]interface{})[attr]
		}
	}
	return nil
}

func (this *BaseNode) parseAndGetSingleParamVaule(paramName, paramVaule string, dataStore *datastore.DataStore, replaceMap ...map[string]interface{}) interface{} {
	defer func() {
		if err := recover(); err != nil {
			panic(fmt.Sprintf("<span style='color:red;'>execute func with expression is %s, error msg is :%s</span>", paramVaule, err.(error).Error()))
		}
	}()
	// 对单个 paramVaule 进行特殊字符编码
	paramVaule = iworkfunc.EncodeSpecialForParamVaule(paramVaule)
	callers, err := iworkfunc.ParseToFuncCallers(paramVaule)
	if err != nil {
		panic(err)
	}
	if callers == nil || len(callers) == 0 {
		// 是直接参数,不需要函数进行特殊处理
		return this.callParseAndGetSingleParamVaule(paramName, paramVaule, dataStore, replaceMap...)
	}
	historyFuncResultMap := make(map[string]interface{}, 0)
	var lastFuncResult interface{}
	// 按照顺序依次执行函数
	for _, caller := range callers {
		args := make([]interface{}, 0)
		// 函数参数替换成实际意义上的值
		for _, arg := range caller.FuncArgs {
			// 判断参数是否来源于 historyFuncResultMap
			if _arg, ok := historyFuncResultMap[arg]; ok {
				args = append(args, _arg)
			} else {
				args = append(args, this.callParseAndGetSingleParamVaule(paramName, arg, dataStore, replaceMap...))
			}
		}
		// 执行函数并记录结果,供下一个函数执行使用
		result := iworkfunc.ExecuteFuncCaller(caller, args)
		historyFuncResultMap["$func."+caller.FuncUUID] = result
		lastFuncResult = result
	}
	return lastFuncResult
}

// 存储 pureText 值
func (this *BaseNode) FillPureTextParamInputSchemaDataToTmp(workStep *iwork.WorkStep, dataStore *datastore.DataStore) map[string]interface{} {
	// 存储节点中间数据
	tmpDataMap := make(map[string]interface{})
	paramInputSchema := schema.GetCacheParamInputSchema(workStep, &WorkStepFactory{WorkStep: workStep})
	for _, item := range paramInputSchema.ParamInputSchemaItems {
		// tmpDataMap 存储引用值 pureText
		tmpDataMap[item.ParamName] = item.ParamValue
	}
	return tmpDataMap
}

// 将 ParamInputSchema 填充数据并返回临时的数据中心 tmpDataMap
func (this *BaseNode) FillParamInputSchemaDataToTmp(workStep *iwork.WorkStep, dataStore *datastore.DataStore) map[string]interface{} {
	// 存储节点中间数据
	tmpDataMap := make(map[string]interface{})
	pureTextTmpDataMap := make(map[string]string)
	paramInputSchema := schema.GetCacheParamInputSchema(workStep, &WorkStepFactory{WorkStep: workStep})
	for _, item := range paramInputSchema.ParamInputSchemaItems {
		this.FillParamInputSchemaItemDataToTmp(pureTextTmpDataMap, tmpDataMap, item, dataStore)
	}
	return tmpDataMap
}

func (this *BaseNode) FillParamInputSchemaItemDataToTmp(pureTextTmpDataMap map[string]string, tmpDataMap map[string]interface{}, item iworkmodels.ParamInputSchemaItem, dataStore *datastore.DataStore) {
	pureTextTmpDataMap[item.ParamName] = item.ParamValue
	// tmpDataMap 存储解析值
	if item.PureText {
		tmpDataMap[item.ParamName] = item.ParamValue
		return
	}
	// 对参数进行非空校验
	if ok, checkResults := iworkvalid.CheckEmptyForItem(item); !ok {
		panic(strings.Join(checkResults, ";"))
	}
	// 判断当前参数是否是 repeat 参数
	if !item.Repeatable {
		tmpDataMap[item.ParamName] = this.ParseAndGetParamVaule(item.ParamName, item.ParamValue, dataStore) // 输入数据存临时
		return
	}
	repeatDatas := this.getRepeatDatas(tmpDataMap, item)
	if len(repeatDatas) > 0 {
		paramValues := make([]interface{}, 0)
		for _, repeatData := range repeatDatas {
			// 替代的节点名称
			replaceProviderNodeName := strings.ReplaceAll(strings.TrimSpace(pureTextTmpDataMap[item.RepeatRefer]), ";", "")
			// 替代的对象
			replaceProviderData := repeatData
			replaceMap := map[string]interface{}{replaceProviderNodeName: replaceProviderData}
			paramValue := this.ParseAndGetParamVaule(item.ParamName, item.ParamValue, dataStore, replaceMap) // 输入数据存临时
			paramValues = append(paramValues, paramValue)
		}
		tmpDataMap[item.ParamName] = paramValues // 所得值则是个切片
	} else {
		tmpDataMap[item.ParamName] = this.ParseAndGetParamVaule(item.ParamName, item.ParamValue, dataStore) // 输入数据存临时
	}
}

func (this *BaseNode) getRepeatDatas(tmpDataMap map[string]interface{}, item iworkmodels.ParamInputSchemaItem) []interface{} {
	repeatDatas := make([]interface{}, 0)
	// 获取 item.RepeatRefer 对应的 repeat 切片数据,作为迭代参数,而不再从前置节点输出获取
	t := reflect.TypeOf(tmpDataMap[item.RepeatRefer])
	v := reflect.ValueOf(tmpDataMap[item.RepeatRefer])
	if t.Kind() == reflect.Slice {
		for i := 0; i < v.Len(); i++ {
			repeatDatas = append(repeatDatas, v.Index(i).Interface())
		}
	}
	return repeatDatas
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
func (this *BaseNode) trim(paramValue string) string {
	// 先进行初次的 trim
	paramValue = strings.TrimSpace(paramValue)
	// 去除前后的 \n
	paramValue = strings.TrimPrefix(paramValue, "\n")
	paramValue = strings.TrimSuffix(paramValue, "\n")
	// 再进行二次 trim
	paramValue = strings.TrimSpace(paramValue)
	return paramValue
}
