package iworknode

import (
	"isoft/isoft/common/stringutil"
	"isoft/isoft_iaas_web/core/iworkconst"
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworkutil/htmlutil"
	"isoft/isoft_iaas_web/models/iwork"
)

type HrefParserNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *HrefParserNode) Execute(trackingId string, skipFunc func(tmpDataMap map[string]interface{}) bool) {
	// 数据中心
	dataStore := datastore.GetDataSource(trackingId)
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, dataStore)
	if skipFunc(tmpDataMap){return}			// 跳过当前节点执行

	hrefs := make([]interface{},0)
	if url, ok := tmpDataMap[iworkconst.STRING_PREFIX + "url"].(string); ok {
		if _hrefs := htmlutil.GetAllHref(url); len(_hrefs) > 0 {
			// 将 []string 转换成 []interface{}
			hrefs = stringutil.ChangeStringsToInterfaces(_hrefs)
		}
	}
	// 放在外面保证条件不满足时也是零值,不报空指针异常
	dataStore.CacheData(this.WorkStep.WorkStepName, iworkconst.MULTI_PREFIX + "hrefs", hrefs)
	dataStore.CacheData(this.WorkStep.WorkStepName, iworkconst.NUMBER_PREFIX + "href_amounts", len(hrefs))
}

func (this *HrefParserNode) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	paramMap := map[int][]string{
		1:[]string{iworkconst.STRING_PREFIX + "url","需要分析资源的url地址"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *HrefParserNode) GetRuntimeParamInputSchema() *schema.ParamInputSchema {
	return &schema.ParamInputSchema{}
}

func (this *HrefParserNode) GetDefaultParamOutputSchema() *schema.ParamOutputSchema {
	return schema.BuildParamOutputSchemaWithSlice([]string{iworkconst.MULTI_PREFIX + "hrefs",iworkconst.NUMBER_PREFIX + "href_amounts"})
}

func (this *HrefParserNode) GetRuntimeParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}

func (this *HrefParserNode) ValidateCustom() {

}