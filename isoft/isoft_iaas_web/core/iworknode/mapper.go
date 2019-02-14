package iworknode

import (
	"encoding/json"
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/models/iwork"
)

type MapperNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}


func (this *MapperNode) Execute(trackingId string) {
	// 获取数据中心
	dataStore := datastore.GetDataSource(trackingId)
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, dataStore)
	// 提交输出数据至数据中心,此类数据能直接从 tmpDataMap 中获取,而不依赖于计算,只适用于 WORK_START、WORK_END、Mapper 等节点
	this.SubmitParamOutputSchemaDataToDataStore(this.WorkStep, dataStore, tmpDataMap)
}

func (this *MapperNode) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	return &schema.ParamInputSchema{}
}

func (this *MapperNode) GetRuntimeParamInputSchema() *schema.ParamInputSchema {
	var paramMappingsArr []string
	json.Unmarshal([]byte(this.WorkStep.WorkStepParamMapping), &paramMappingsArr)
	items := make([]schema.ParamInputSchemaItem,0)
	for _, paramMapping := range paramMappingsArr {
		items = append(items, schema.ParamInputSchemaItem{ParamName: paramMapping})
	}
	return &schema.ParamInputSchema{ParamInputSchemaItems:items}
}

func (this *MapperNode) GetDefaultParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}

func (this *MapperNode) GetRuntimeParamOutputSchema() *schema.ParamOutputSchema {
	items := make([]schema.ParamOutputSchemaItem,0)
	inputSchema := schema.GetCacheParamInputSchema(this.WorkStep, &WorkStepFactory{WorkStep:this.WorkStep})
	for _, item := range inputSchema.ParamInputSchemaItems{
		items = append(items, schema.ParamOutputSchemaItem{ParamName:item.ParamName})
	}
	return &schema.ParamOutputSchema{ParamOutputSchemaItems:items}
}
