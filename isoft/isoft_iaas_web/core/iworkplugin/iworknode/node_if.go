package iworknode

import (
	"fmt"
	"isoft/isoft_iaas_web/core/iworkconst"
	"isoft/isoft_iaas_web/core/iworkdata/block"
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/core/iworkdata/entry"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworkmodels"
	"isoft/isoft_iaas_web/models/iwork"
)

type IFNode struct {
	BaseNode
	WorkStep         *iwork.WorkStep
	BlockStep        *block.BlockStep
	BlockStepRunFunc func(trackingId string, blockStep *block.BlockStep, datastore *datastore.DataStore, dispatcher *entry.Dispatcher) (receiver *entry.Receiver)
}

func (this *IFNode) Execute(trackingId string) {
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, this.DataStore)
	expression := tmpDataMap[iworkconst.BOOL_PREFIX+"expression"].(bool)
	this.DataStore.CacheData(this.WorkStep.WorkStepName, iworkconst.BOOL_PREFIX+"expression", expression)

	if expression && this.BlockStep.HasChildren {
		// 需要延迟执行的 BlockSteps
		deferBlockSteps := make([]*block.BlockStep, 0)
		for _, blockStep := range this.BlockStep.ChildBlockSteps {
			if blockStep.Step.IsDefer == "true" {
				// 加入切片中
				deferBlockSteps = append(deferBlockSteps, blockStep)
			} else {
				// 直接执行
				this.BlockStepRunFunc(trackingId, blockStep, this.DataStore, nil)
			}
		}
		// 倒叙执行
		for i := len(deferBlockSteps) - 1; i >= 0; i-- {
			this.BlockStepRunFunc(trackingId, deferBlockSteps[i], this.DataStore, nil)
		}
	} else {
		iwork.InsertRunLogDetail(trackingId, fmt.Sprintf("The blockStep for %s was skipped!", this.WorkStep.WorkStepName))
	}
}

func (this *IFNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: []string{iworkconst.BOOL_PREFIX + "expression", "if条件表达式,值为 bool 类型!"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *IFNode) GetRuntimeParamInputSchema() *iworkmodels.ParamInputSchema {
	return &iworkmodels.ParamInputSchema{}
}

func (this *IFNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return schema.BuildParamOutputSchemaWithSlice([]string{iworkconst.BOOL_PREFIX + "expression"})
}

func (this *IFNode) GetRuntimeParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return &iworkmodels.ParamOutputSchema{}
}

func (this *IFNode) ValidateCustom() (checkResult []string) {
	return []string{}
}