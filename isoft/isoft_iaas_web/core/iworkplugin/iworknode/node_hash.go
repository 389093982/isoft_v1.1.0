package iworknode

import (
	"isoft/isoft/common/hashutil"
	"isoft/isoft_iaas_web/core/iworkconst"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworkmodels"
	"isoft/isoft_iaas_web/models/iwork"
)

type CalHashNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *CalHashNode) Execute(trackingId string) {
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, this.DataStore)
	str_data := tmpDataMap[iworkconst.STRING_PREFIX+"str_data"].(string)
	hash := hashutil.CalculateHashWithString(str_data)
	this.DataStore.CacheDatas(this.WorkStep.WorkStepName, map[string]interface{}{iworkconst.STRING_PREFIX + "hash": hash})
}

func (this *CalHashNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "str_data", "需要计算hash值的字符串"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *CalHashNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return schema.BuildParamOutputSchemaWithSlice([]string{iworkconst.STRING_PREFIX + "hash"})
}
