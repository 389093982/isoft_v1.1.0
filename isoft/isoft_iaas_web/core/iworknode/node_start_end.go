package iworknode

import (
	"encoding/json"
	"fmt"
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
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, this.DataStore)
	// dispatcher 非空时替换成父流程参数
	if this.Dispatcher != nil && len(this.Dispatcher.TmpDataMap) > 0 {
		// 从父流程中获取值,即从 Dispatcher 中获取值
		for key, value := range this.Dispatcher.TmpDataMap {
			if value != "__default__" { // __default__ 则表示不用替换,还是使用子流程默认值参数
				tmpDataMap[key] = value
			}
		}
	}
	for key, value := range tmpDataMap {
		iwork.InsertRunLogDetail(trackingId, fmt.Sprintf("fill param with for %s:%s", key, value))
	}
	// 提交输出数据至数据中心,此类数据能直接从 tmpDataMap 中获取,而不依赖于计算,只适用于 WORK_START、WORK_END、Mapper 等节点
	this.SubmitParamOutputSchemaDataToDataStore(this.WorkStep, this.DataStore, tmpDataMap)
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

func (this *WorkStartNode) ValidateCustom() (checkResult []string) {
	return []string{}
}

type WorkEndNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
	Receiver *entry.Receiver
}

func (this *WorkEndNode) Execute(trackingId string) {
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, this.DataStore)
	// 提交输出数据至数据中心,此类数据能直接从 tmpDataMap 中获取,而不依赖于计算,只适用于 WORK_START、WORK_END 节点
	this.SubmitParamOutputSchemaDataToDataStore(this.WorkStep, this.DataStore, tmpDataMap)
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

func (this *WorkEndNode) ValidateCustom() (checkResult []string) {
	return []string{}
}
