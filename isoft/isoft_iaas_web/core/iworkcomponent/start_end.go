package iworkcomponent

import (
	"isoft/isoft_iaas_web/core/iworkdata"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
)

type WorkStartNode struct {
	WorkStep 		   *iwork.WorkStep
}

func (this *WorkStartNode) Execute(trackingId string) {
	// 存储节点中间数据
	tmpDataMap := make(map[string]interface{})
	paramInputSchema := GetCacheParamInputSchema(this.WorkStep)
	for _, item := range paramInputSchema.ParamInputSchemaItems{
		tmpDataMap[item.ParamName] = item.ParamValue			// 输入数据存临时
	}
	// 获取数据中心
	dataStore := iworkdata.GetDataSource(trackingId)
	// 将执行结果存储到数据中心
	paramOutputSchema := GetCacheParamOutputSchema(this.WorkStep)
	for _,item := range paramOutputSchema.ParamOutputSchemaItems{
		// 将数据数据存储到数据中心
		dataStore.CacheData(this.WorkStep.WorkStepName, item.ParamName, tmpDataMap[item.ParamName])
	}
}

func (this *WorkStartNode) GetDefaultParamInputSchema() *iworkdata.ParamInputSchema {
	return nil
}

func (this *WorkStartNode) GetRuntimeParamOutputSchema() *iworkdata.ParamOutputSchema {
	return nil
}

func (this *WorkStartNode) GetDefaultParamOutputSchema() *iworkdata.ParamOutputSchema {
	return transferParamInputSchemaToParamOutputSchema(this.WorkStep)
}


type WorkEndNode struct {
	DataStore		   *iworkdata.DataStore
	WorkStep 		   *iwork.WorkStep
}


func (this *WorkEndNode) Execute(trackingId string) {
	// 获取数据中心
	dataStore := iworkdata.GetDataSource(trackingId)

	// 存储节点中间数据
	tmpDataMap := make(map[string]interface{})
	paramInputSchema := GetCacheParamInputSchema(this.WorkStep)
	for _, item := range paramInputSchema.ParamInputSchemaItems{
		tmpDataMap[item.ParamName] = this.ParseParamVaule(item.ParamValue, dataStore)			// 输入数据存临时
	}

	// 将执行结果存储到数据中心
	paramOutputSchema := GetCacheParamOutputSchema(this.WorkStep)
	for _,item := range paramOutputSchema.ParamOutputSchemaItems{
		// 将数据数据存储到数据中心
		dataStore.CacheData(this.WorkStep.WorkStepName, item.ParamName, tmpDataMap[item.ParamName])
	}
}

func (this *WorkEndNode) ParseParamVaule(paramVaule string, dataStore *iworkdata.DataStore) interface{} {
	if strings.HasPrefix(paramVaule, "$"){
		resolver := iworkdata.ParamVauleParser{ParamValue:paramVaule}
		return dataStore.GetData(resolver.GetNodeNameFromParamValue(), resolver.GetParamNameFromParamValue())
	}else{
		return paramVaule
	}
}

func (this *WorkEndNode) GetDefaultParamInputSchema() *iworkdata.ParamInputSchema {
	return nil
}

func (this *WorkEndNode) GetRuntimeParamOutputSchema() *iworkdata.ParamOutputSchema {
	return nil
}

func (this *WorkEndNode) GetDefaultParamOutputSchema() *iworkdata.ParamOutputSchema {
	return transferParamInputSchemaToParamOutputSchema(this.WorkStep)
}

// 输入转输出,适用于开始节点和结束节点
func transferParamInputSchemaToParamOutputSchema(step *iwork.WorkStep) *iworkdata.ParamOutputSchema {
	items := []iworkdata.ParamOutputSchemaItem{}
	paramInputSchema := GetCacheParamInputSchema(step)
	for _, paramInputSchemaItem := range paramInputSchema.ParamInputSchemaItems{
		items = append(items, iworkdata.ParamOutputSchemaItem{ParamName: paramInputSchemaItem.ParamName})
	}
	return &iworkdata.ParamOutputSchema{ParamOutputSchemaItems: items}
}