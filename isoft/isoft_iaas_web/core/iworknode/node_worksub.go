package iworknode

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
	"isoft/isoft_iaas_web/core/iworkconst"
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/core/iworkdata/entry"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworkutil"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
)

type WorkSub struct {
	BaseNode
	WorkStep *iwork.WorkStep
	RunFunc  func(work iwork.Work, steps []iwork.WorkStep, dispatcher *entry.Dispatcher) (receiver *entry.Receiver)
}

func (this *WorkSub) Execute(trackingId string, skipFunc func(tmpDataMap map[string]interface{}) bool) {
	// 获取子流程流程名称
	workSubName := this.checkAndGetWorkSubName()
	// 数据中心
	dataStore := datastore.GetDataStore(trackingId)
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, dataStore)
	if skipFunc(tmpDataMap) {
		return
	} // 跳过当前节点执行
	// 运行子流程
	work, _ := iwork.QueryWorkByName(workSubName, orm.NewOrm())
	steps, _ := iwork.QueryAllWorkStepByWorkName(workSubName, orm.NewOrm())
	// 获取 foreach_data 数据
	foreachDatas := getConvertedForEachData(tmpDataMap)
	if len(foreachDatas) > 0 {
		itemKey := this.getForeachItemKey(tmpDataMap)
		// work_sub 节点支持 foreach 循环功能,此处循环 foreach 次数
		for _, foreachData := range foreachDatas {
			if itemKey != "" {
				// 找到 tmpDataMap 中的迭代元素 __item__,将其替换成需要迭代的元素
				tmpDataMap[itemKey] = foreachData
			}
			this.RunOnceSubWork(work, steps, trackingId, tmpDataMap, dataStore)
		}
	} else {
		this.RunOnceSubWork(work, steps, trackingId, tmpDataMap, dataStore)
	}
}

// 获取转换后的 foreach_data 数据
// 任意类型切片对象,目前仅限制于 []interface{} 和 []map[string]interface{},需要由前置节点标准化成这种类型才可
func getConvertedForEachData(tmpDataMap map[string]interface{}) []interface{} {
	foreachDatas := make([]interface{}, 0)
	if _foreachDatas, ok := tmpDataMap[iworkconst.FOREACH_PREFIX+"data?"].([]interface{}); ok {
		foreachDatas = append(foreachDatas, _foreachDatas...)
	} else if _foreachDatas, ok := tmpDataMap[iworkconst.FOREACH_PREFIX+"data?"].([]map[string]interface{}); ok {
		for _, _foreachData := range _foreachDatas {
			foreachDatas = append(foreachDatas, _foreachData)
		}
	}
	return foreachDatas
}

func (this *WorkSub) checkAndGetWorkSubName() string {
	workSubName := iworkutil.GetWorkSubNameForWorkSubNode(
		schema.GetCacheParamInputSchema(this.WorkStep, &WorkStepFactory{WorkStep: this.WorkStep}))
	if strings.TrimSpace(workSubName) == "" {
		panic(errors.New("invalid workSubName"))
	}
	return workSubName
}

func (this *WorkSub) getForeachItemKey(tmpDataMap map[string]interface{}) string {
	var itemKey string
	for key, value := range tmpDataMap {
		if _value, ok := value.(string); ok && strings.TrimSpace(_value) == "__item__" {
			itemKey = key
		}
	}
	return itemKey
}

func (this *WorkSub) RunOnceSubWork(work iwork.Work, steps []iwork.WorkStep, trackingId string,
	tmpDataMap map[string]interface{}, dataStore *datastore.DataStore) {
	receiver := this.RunFunc(work, steps, &entry.Dispatcher{TrackingId: trackingId, TmpDataMap: tmpDataMap})
	// 接收子流程数据存入 dataStore
	for paramName, paramValue := range receiver.TmpDataMap {
		dataStore.CacheData(this.WorkStep.WorkStepName, paramName, paramValue)
	}
}

func (this *WorkSub) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	paramMap := map[int][]string{
		1: []string{iworkconst.STRING_PREFIX + "work_sub", "子流程信息"},
		2: []string{iworkconst.FOREACH_PREFIX + "data?", "可选参数,当有值时表示迭代流程,该节点会执行多次,并将当前迭代元素放入 __item__ 变量中,其它参数需要引用 __item__ 即可"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

// 获取动态输入值
func (this *WorkSub) GetRuntimeParamInputSchema() *schema.ParamInputSchema {
	items := make([]schema.ParamInputSchemaItem, 0)
	// 获取子流程信息
	workSubName := this.getWorkSubName()
	if strings.TrimSpace(workSubName) != "" {
		// 获取子流程所有步骤
		subSteps, err := iwork.QueryAllWorkStepByWorkName(workSubName, orm.NewOrm())
		if err != nil {
			panic(err)
		}
		for _, subStep := range subSteps {
			// 找到子流程起始节点
			if strings.ToUpper(subStep.WorkStepType) == "WORK_START" {
				// 子流程起始节点输入参数
				subItems := schema.GetCacheParamInputSchema(&subStep, &WorkStepFactory{WorkStep: &subStep})
				for _, subItem := range subItems.ParamInputSchemaItems {
					items = append(items, schema.ParamInputSchemaItem{ParamName: subItem.ParamName})
				}
			}
		}
	}
	return &schema.ParamInputSchema{ParamInputSchemaItems: items}
}

func (this *WorkSub) getWorkSubName() string {
	// 读取历史输入值
	paramInputSchema := schema.GetCacheParamInputSchema(this.WorkStep, &WorkStepFactory{WorkStep: this.WorkStep})
	// 从历史输入值中获取子流程名称
	workSubName := iworkutil.GetWorkSubNameForWorkSubNode(paramInputSchema)
	return workSubName
}

func (this *WorkSub) GetDefaultParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}

func (this *WorkSub) GetRuntimeParamOutputSchema() *schema.ParamOutputSchema {
	items := make([]schema.ParamOutputSchemaItem, 0)
	// 读取静态输入值
	paramInputSchema := schema.GetCacheParamInputSchema(this.WorkStep, &WorkStepFactory{WorkStep: this.WorkStep})
	// 从静态输入值中获取子流程名称
	workSubName := iworkutil.GetWorkSubNameForWorkSubNode(paramInputSchema)
	if strings.TrimSpace(workSubName) != "" {
		// 获取子流程所有步骤
		subSteps, err := iwork.QueryAllWorkStepByWorkName(workSubName, orm.NewOrm())
		if err != nil {
			panic(err)
		}
		for _, subStep := range subSteps {
			// 找到子流程结束节点
			if strings.ToUpper(subStep.WorkStepType) == "WORK_END" {
				// 子流程结束节点输出参数
				subItems := schema.GetCacheParamOutputSchema(&subStep)
				for _, subItem := range subItems.ParamOutputSchemaItems {
					items = append(items, schema.ParamOutputSchemaItem{ParamName: subItem.ParamName})
				}
			}
		}
	}
	return &schema.ParamOutputSchema{ParamOutputSchemaItems: items}
}

func (this *WorkSub) ValidateCustom() {
	if workSubName := this.getWorkSubName(); workSubName == "" {
		panic("Empty workSubName was found!")
	} else {
		if _, err := iwork.QueryWorkByName(workSubName, orm.NewOrm()); err != nil {
			panic(fmt.Sprintf("WorkSubName for %s was not found!", workSubName))
		}
	}
}
