package iworknode

import (
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworkutil/httputil"
	"isoft/isoft_iaas_web/models/iwork"
)

type HttpRequestNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *HttpRequestNode) Execute(trackingId string) {
	// 数据中心
	dataStore := datastore.GetDataSource(trackingId)
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, dataStore)
	var request_url, request_method string
	if _request_url, ok := tmpDataMap["request_url"].(string); ok{
		request_url = _request_url
	}
	if _request_method, ok := tmpDataMap["request_method?"].(string); ok{
		request_method = _request_method
	}
	paramMap := make(map[string]interface{})
	response := httputil.DoHttpRequest(request_url, request_method, paramMap)
	dataStore.CacheData(this.WorkStep.WorkStepName, "response_str", string(response))
}

func (this *HttpRequestNode) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	return schema.BuildParamInputSchemaWithSlice([]string{"request_url", "request_method?"})
}

func (this *HttpRequestNode) GetRuntimeParamInputSchema() *schema.ParamInputSchema {
	return &schema.ParamInputSchema{}
}

func (this *HttpRequestNode) GetDefaultParamOutputSchema() *schema.ParamOutputSchema {
	return schema.BuildParamOutputSchemaWithSlice([]string{"response_str"})
}

func (this *HttpRequestNode) GetRuntimeParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}

