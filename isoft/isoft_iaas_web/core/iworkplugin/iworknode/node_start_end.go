package iworknode

import (
	"encoding/json"
	"fmt"
	"isoft/isoft_iaas_web/core/iworkdata/entry"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworkmodels"
	"isoft/isoft_iaas_web/models/iwork"
)

type WorkStartNode struct {
	BaseNode
	WorkStep   *iwork.WorkStep
	Dispatcher *entry.Dispatcher
}

func (this *WorkStartNode) Execute(trackingId string) {
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep)
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
		this.LogWriter.Write(trackingId, fmt.Sprintf("fill param with for %s:%s", key, value))
	}
	// 提交输出数据至数据中心,此类数据能直接从 tmpDataMap 中获取,而不依赖于计算,只适用于 WORK_START、WORK_END、Mapper 等节点
	this.SubmitParamOutputSchemaDataToDataStore(this.WorkStep, this.DataStore, tmpDataMap)
}

func (this *WorkStartNode) GetRuntimeParamInputSchema() *iworkmodels.ParamInputSchema {
	var paramMappingsArr []string
	json.Unmarshal([]byte(this.WorkStep.WorkStepParamMapping), &paramMappingsArr)
	items := make([]iworkmodels.ParamInputSchemaItem, 0)
	for _, paramMapping := range paramMappingsArr {
		items = append(items, iworkmodels.ParamInputSchemaItem{ParamName: paramMapping})
	}
	return &iworkmodels.ParamInputSchema{ParamInputSchemaItems: items}
}

func (this *WorkStartNode) GetRuntimeParamOutputSchema() *iworkmodels.ParamOutputSchema {
	items := make([]iworkmodels.ParamOutputSchemaItem, 0)
	inputSchema := schema.GetCacheParamInputSchema(this.WorkStep, &WorkStepFactory{WorkStep: this.WorkStep})
	for _, item := range inputSchema.ParamInputSchemaItems {
		items = append(items, iworkmodels.ParamOutputSchemaItem{ParamName: item.ParamName})
	}
	return &iworkmodels.ParamOutputSchema{ParamOutputSchemaItems: items}
}

type WorkEndNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
	Receiver *entry.Receiver
}

func (this *WorkEndNode) Execute(trackingId string) {
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep)
	// 提交输出数据至数据中心,此类数据能直接从 tmpDataMap 中获取,而不依赖于计算,只适用于 WORK_START、WORK_END 节点
	this.SubmitParamOutputSchemaDataToDataStore(this.WorkStep, this.DataStore, tmpDataMap)
	// 同时需要将数据提交到 Receiver
	this.Receiver = &entry.Receiver{TmpDataMap: tmpDataMap}
}

func (this *WorkEndNode) GetRuntimeParamInputSchema() *iworkmodels.ParamInputSchema {
	var paramMappingsArr []string
	json.Unmarshal([]byte(this.WorkStep.WorkStepParamMapping), &paramMappingsArr)
	items := make([]iworkmodels.ParamInputSchemaItem, 0)
	for _, paramMapping := range paramMappingsArr {
		items = append(items, iworkmodels.ParamInputSchemaItem{ParamName: paramMapping})
	}
	return &iworkmodels.ParamInputSchema{ParamInputSchemaItems: items}
}

func (this *WorkEndNode) GetRuntimeParamOutputSchema() *iworkmodels.ParamOutputSchema {
	items := make([]iworkmodels.ParamOutputSchemaItem, 0)
	inputSchema := schema.GetCacheParamInputSchema(this.WorkStep, &WorkStepFactory{WorkStep: this.WorkStep})
	for _, item := range inputSchema.ParamInputSchemaItems {
		items = append(items, iworkmodels.ParamOutputSchemaItem{ParamName: item.ParamName})
	}
	return &iworkmodels.ParamOutputSchema{ParamOutputSchemaItems: items}
}
