package iworknode

import (
	"isoft/isoft/common/hashutil"
	"isoft/isoft_iaas_web/core/iworkconst"
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/models/iwork"
)

type CalHashNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *CalHashNode) Execute(trackingId string) {
	// 数据中心
	dataStore := datastore.GetDataStore(trackingId)
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, dataStore)
	str_data := tmpDataMap[iworkconst.STRING_PREFIX+"str_data"].(string)
	hash := hashutil.CalculateHashWithString(str_data)
	dataStore.CacheData(this.WorkStep.WorkStepName, iworkconst.STRING_PREFIX+"hash", hash)
}

func (this *CalHashNode) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "str_data", "需要计算hash值的字符串"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *CalHashNode) GetRuntimeParamInputSchema() *schema.ParamInputSchema {
	return &schema.ParamInputSchema{}
}

func (this *CalHashNode) GetDefaultParamOutputSchema() *schema.ParamOutputSchema {
	return schema.BuildParamOutputSchemaWithSlice([]string{iworkconst.STRING_PREFIX + "hash"})
}

func (this *CalHashNode) GetRuntimeParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}

func (this *CalHashNode) ValidateCustom() (checkResult []string) {
	return
}
