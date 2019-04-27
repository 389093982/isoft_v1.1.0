package iworknode

import (
	"isoft/isoft/common/stringutil"
	"isoft/isoft_iaas_web/core/iworkconst"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworkmodels"
	"isoft/isoft_iaas_web/core/iworkutil/htmlutil"
	"isoft/isoft_iaas_web/models/iwork"
)

type HrefParserNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *HrefParserNode) Execute(trackingId string) {
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, this.DataStore)
	hrefs := make([]interface{}, 0)
	if url, ok := tmpDataMap[iworkconst.STRING_PREFIX+"url"].(string); ok {
		if _hrefs := htmlutil.GetAllHref(url); len(_hrefs) > 0 {
			// 将 []string 转换成 []interface{}
			hrefs = stringutil.ChangeStringsToInterfaces(_hrefs)
		}
	}
	// 放在外面保证条件不满足时也是零值,不报空指针异常
	this.DataStore.CacheDatas(this.WorkStep.WorkStepName, map[string]interface{}{
		iworkconst.MULTI_PREFIX + "hrefs":         hrefs,
		iworkconst.NUMBER_PREFIX + "href_amounts": len(hrefs),
	})
}

func (this *HrefParserNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: []string{iworkconst.STRING_PREFIX + "url", "需要分析资源的url地址"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *HrefParserNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return schema.BuildParamOutputSchemaWithSlice([]string{iworkconst.MULTI_PREFIX + "hrefs", iworkconst.NUMBER_PREFIX + "href_amounts"})
}
