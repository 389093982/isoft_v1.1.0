package iworknode

import (
	"isoft/isoft_iaas_web/core/iworkconst"
	"isoft/isoft_iaas_web/core/iworkdata/block"
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/core/iworkdata/entry"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/models/iwork"
)

type IFNode struct {
	BaseNode
	WorkStep         *iwork.WorkStep
	BlockStep        *block.BlockStep
	BlockStepRunFunc func(trackingId string, blockStep *block.BlockStep, dispatcher *entry.Dispatcher) (receiver *entry.Receiver)
}

func (this *IFNode) Execute(trackingId string, skipFunc func(tmpDataMap map[string]interface{}) bool) {
	// 数据中心
	dataStore := datastore.GetDataStore(trackingId)
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, dataStore)
	if skipFunc(tmpDataMap) {
		return
	} // 跳过当前节点执行
	expression := tmpDataMap[iworkconst.BOOL_PREFIX+"expression"].(bool)
	dataStore.CacheData(this.WorkStep.WorkStepName, iworkconst.BOOL_PREFIX+"expression", expression)

	if expression && this.BlockStep.HasChildren {
		for _, blockStep := range this.BlockStep.ChildBlockSteps {
			this.BlockStepRunFunc(trackingId, blockStep, nil)
		}
	}
}

func (this *IFNode) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	paramMap := map[int][]string{
		1: []string{iworkconst.BOOL_PREFIX + "expression", "if条件表达式,值为 bool 类型!"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *IFNode) GetRuntimeParamInputSchema() *schema.ParamInputSchema {
	return &schema.ParamInputSchema{}
}

func (this *IFNode) GetDefaultParamOutputSchema() *schema.ParamOutputSchema {
	return schema.BuildParamOutputSchemaWithSlice([]string{iworkconst.BOOL_PREFIX + "expression"})
}

func (this *IFNode) GetRuntimeParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}

func (this *IFNode) ValidateCustom() {

}
