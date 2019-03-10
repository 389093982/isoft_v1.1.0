package iworknode

import (
	"encoding/json"
	"fmt"
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/core/iworkdata/entry"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/models/iwork"
)

type WorkStartNode struct {
	BaseNode
	WorkStep   *iwork.WorkStep
	Dispatcher *entry.Dispatcher
}

func (this *WorkStartNode) Execute(trackingId string) {
	// 存储节点中间数据
	tmpDataMap := make(map[string]interface{})
	if this.Dispatcher != nil && len(this.Dispatcher.TmpDataMap) > 0 {
		// 从父流程中获取值,即从 Dispatcher 中获取值
		tmpDataMap = this.Dispatcher.TmpDataMap
	} else {
		// 使用节点默认值
		paramInputSchema := schema.GetCacheParamInputSchema(this.WorkStep, &WorkStepFactory{WorkStep: this.WorkStep})
		for _, item := range paramInputSchema.ParamInputSchemaItems {
			iwork.InsertRunLogDetail(trackingId, fmt.Sprintf("fill param with default for %s:%s", item.ParamName, item.ParamValue))
			tmpDataMap[item.ParamName] = item.ParamValue // 输入数据存临时
		}
	}
	// 获取数据中心
	dataStore := datastore.GetDataStore(trackingId)
	// 提交输出数据至数据中心,此类数据能直接从 tmpDataMap 中获取,而不依赖于计算,只适用于 WORK_START、WORK_END、Mapper 等节点
	this.SubmitParamOutputSchemaDataToDataStore(this.WorkStep, dataStore, tmpDataMap)
}

func (this *WorkStartNode) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	return &schema.ParamInputSchema{}
}

func (this *WorkStartNode) GetRuntimeParamInputSchema() *schema.ParamInputSchema {
	var paramMappingsArr []string
	json.Unmarshal([]byte(this.WorkStep.WorkStepParamMapping), &paramMappingsArr)
	items := make([]schema.ParamInputSchemaItem, 0)
	for _, paramMapping := range paramMappingsArr {
		items = append(items, schema.ParamInputSchemaItem{ParamName: paramMapping})
	}
	return &schema.ParamInputSchema{ParamInputSchemaItems: items}
}

func (this *WorkStartNode) GetRuntimeParamOutputSchema() *schema.ParamOutputSchema {
	items := make([]schema.ParamOutputSchemaItem, 0)
	inputSchema := schema.GetCacheParamInputSchema(this.WorkStep, &WorkStepFactory{WorkStep: this.WorkStep})
	for _, item := range inputSchema.ParamInputSchemaItems {
		items = append(items, schema.ParamOutputSchemaItem{ParamName: item.ParamName})
	}
	return &schema.ParamOutputSchema{ParamOutputSchemaItems: items}
}

func (this *WorkStartNode) GetDefaultParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}

func (this *WorkStartNode) ValidateCustom() {

}

type WorkEndNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
	Receiver *entry.Receiver
}

func (this *WorkEndNode) Execute(trackingId string) {
	// 数据中心
	dataStore := datastore.GetDataStore(trackingId)
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, dataStore)
	// 提交输出数据至数据中心,此类数据能直接从 tmpDataMap 中获取,而不依赖于计算,只适用于 WORK_START、WORK_END 节点
	this.SubmitParamOutputSchemaDataToDataStore(this.WorkStep, dataStore, tmpDataMap)
	// 同时需要将数据提交到 Receiver
	this.Receiver = &entry.Receiver{TmpDataMap: tmpDataMap}
}

func (this *WorkEndNode) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	return &schema.ParamInputSchema{}
}

func (this *WorkEndNode) GetRuntimeParamInputSchema() *schema.ParamInputSchema {
	var paramMappingsArr []string
	json.Unmarshal([]byte(this.WorkStep.WorkStepParamMapping), &paramMappingsArr)
	items := make([]schema.ParamInputSchemaItem, 0)
	for _, paramMapping := range paramMappingsArr {
		items = append(items, schema.ParamInputSchemaItem{ParamName: paramMapping})
	}
	return &schema.ParamInputSchema{ParamInputSchemaItems: items}
}

func (this *WorkEndNode) GetRuntimeParamOutputSchema() *schema.ParamOutputSchema {
	items := make([]schema.ParamOutputSchemaItem, 0)
	inputSchema := schema.GetCacheParamInputSchema(this.WorkStep, &WorkStepFactory{WorkStep: this.WorkStep})
	for _, item := range inputSchema.ParamInputSchemaItems {
		items = append(items, schema.ParamOutputSchemaItem{ParamName: item.ParamName})
	}
	return &schema.ParamOutputSchema{ParamOutputSchemaItems: items}
}

func (this *WorkEndNode) GetDefaultParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}

func (this *WorkEndNode) ValidateCustom() {

}
