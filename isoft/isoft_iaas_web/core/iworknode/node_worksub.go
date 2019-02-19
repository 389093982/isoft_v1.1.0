package iworknode

import (
	"github.com/pkg/errors"
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

func (this *WorkSub) Execute(trackingId string) {
	// 获取子流程流程名称
	workSubName := iworkutil.GetWorkSubNameForWorkSubNode(
		schema.GetCacheParamInputSchema(this.WorkStep, &WorkStepFactory{WorkStep: this.WorkStep}))
	if strings.TrimSpace(workSubName) == "" {
		panic(errors.New("invalid workSubName"))
	}
	// 数据中心
	dataStore := datastore.GetDataSource(trackingId)
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, dataStore)

	// 运行子流程
	work, _ := iwork.QueryWorkByName(workSubName)
	steps, _ := iwork.GetAllWorkStepByWorkName(workSubName)

	if foreachDatas, ok := tmpDataMap["foreach_data?"].([]map[string]interface{}); ok{
		itemKey := this.getForeachItemKey(tmpDataMap)
		// work_sub 节点支持 foreach 循环功能,此处循环 foreach 次数
		for _, foreachData := range foreachDatas {
			if itemKey != ""{
				// 找到 tmpDataMap 中的迭代元素 __item__,将其替换成需要迭代的元素
				tmpDataMap[itemKey] = foreachData
			}
			this.RunOnceSubWork(work, steps, trackingId, tmpDataMap, dataStore)
		}
	} else {
		this.RunOnceSubWork(work, steps, trackingId, tmpDataMap, dataStore)
	}
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
		1:[]string{"work_sub","子流程信息"},
		2:[]string{"foreach_data?","可选参数,当有值时表示迭代流程,该节点会执行多次,并将当前迭代元素放入 __item__ 变量中,其它参数需要引用 __item__ 即可"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

// 获取动态输入值
func (this *WorkSub) GetRuntimeParamInputSchema() *schema.ParamInputSchema {
	items := make([]schema.ParamInputSchemaItem,0)
	// 读取历史输入值
	paramInputSchema := schema.GetCacheParamInputSchema(this.WorkStep, &WorkStepFactory{WorkStep: this.WorkStep})
	// 从历史输入值中获取子流程名称
	workSubName := iworkutil.GetWorkSubNameForWorkSubNode(paramInputSchema)
	if strings.TrimSpace(workSubName) != "" {
		// 获取子流程所有步骤
		subSteps, err := iwork.GetAllWorkStepByWorkName(workSubName)
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

func (this *WorkSub) GetDefaultParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}

func (this *WorkSub) GetRuntimeParamOutputSchema() *schema.ParamOutputSchema {
	items := []schema.ParamOutputSchemaItem{}
	// 读取静态输入值
	paramInputSchema := schema.GetCacheParamInputSchema(this.WorkStep, &WorkStepFactory{WorkStep: this.WorkStep})
	// 从静态输入值中获取子流程名称
	workSubName := iworkutil.GetWorkSubNameForWorkSubNode(paramInputSchema)
	if strings.TrimSpace(workSubName) != "" {
		// 获取子流程所有步骤
		subSteps, err := iwork.GetAllWorkStepByWorkName(workSubName)
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