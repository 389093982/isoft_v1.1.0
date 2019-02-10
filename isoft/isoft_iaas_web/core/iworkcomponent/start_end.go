package iworkcomponent

import (
	"isoft/isoft_iaas_web/core/iworkdata"
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/models/iwork"
)

type WorkStartNode struct {
	BaseNode
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
	dataStore := datastore.GetDataSource(trackingId)
	// 提交输出数据至数据中心,此类数据能直接从 tmpDataMap 中获取,而不依赖于计算,只适用于 WORK_START、WORK_END 节点
	this.SubmitParamOutputSchemaDataToDataStore(this.WorkStep, dataStore, tmpDataMap)
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
	BaseNode
	WorkStep 		   *iwork.WorkStep
}


func (this *WorkEndNode) Execute(trackingId string) {
	// 数据中心
	dataStore := datastore.GetDataSource(trackingId)
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep,dataStore)
	// 提交输出数据至数据中心,此类数据能直接从 tmpDataMap 中获取,而不依赖于计算,只适用于 WORK_START、WORK_END 节点
	this.SubmitParamOutputSchemaDataToDataStore(this.WorkStep, dataStore, tmpDataMap)
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