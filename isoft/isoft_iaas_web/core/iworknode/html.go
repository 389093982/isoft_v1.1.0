package iworknode

import (
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworkutil/htmlutil"
	"isoft/isoft_iaas_web/models/iwork"
)

type HrefParserNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *HrefParserNode) Execute(trackingId string) {
	// 数据中心
	dataStore := datastore.GetDataSource(trackingId)
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, dataStore)
	if url, ok := tmpDataMap["url"].(string); ok {
		if hrefs := htmlutil.GetAllHref(url); len(hrefs) > 0 {
			dataStore.CacheData(this.WorkStep.WorkStepName, "hrefs", hrefs)
		}
	}
}

func (this *HrefParserNode) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	paramMap := map[string]string{
		"url":"需要分析资源的url地址",
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *HrefParserNode) GetRuntimeParamInputSchema() *schema.ParamInputSchema {
	return &schema.ParamInputSchema{}
}

func (this *HrefParserNode) GetDefaultParamOutputSchema() *schema.ParamOutputSchema {
	return schema.BuildParamOutputSchemaWithSlice([]string{"hrefs"})
}

func (this *HrefParserNode) GetRuntimeParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}
