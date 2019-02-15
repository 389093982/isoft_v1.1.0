package iworknode

import (
	"fmt"
	"github.com/pkg/errors"
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworkutil"
	"isoft/isoft_iaas_web/core/iworkutil/httputil"
	"isoft/isoft_iaas_web/models/iwork"
	"net/http"
	"strings"
)

type HttpRequestNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *HttpRequestNode) Execute(trackingId string) {
	// 数据中心
	_dataStore := datastore.GetDataSource(trackingId)
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, _dataStore)

	// 参数准备
	var request_url, request_method string
	if _request_url, ok := tmpDataMap["request_url"].(string); ok{
		request_url = _request_url
	}
	if _request_method, ok := tmpDataMap["request_method?"].(string); ok{
		request_method = _request_method
	}
	paramMap := fillParamMapData(tmpDataMap, "request_params?")
	headerMap := fillParamMapData(tmpDataMap, "request_headers?")

	responsebytes := httputil.DoHttpRequestWithParserFunc(request_url, request_method, paramMap, headerMap, func(resp *http.Response) {
		_dataStore.CacheData(this.WorkStep.WorkStepName, "StatusCode", resp.StatusCode)
		_dataStore.CacheData(this.WorkStep.WorkStepName, "ContentType", resp.Header.Get("content-type"))
	})
	_dataStore.CacheData(this.WorkStep.WorkStepName, "response_str", string(responsebytes))
	_dataStore.CacheData(this.WorkStep.WorkStepName, "response_bytes", responsebytes)
	_dataStore.CacheData(this.WorkStep.WorkStepName, "base64res_str", iworkutil.EncodeToBase64String(responsebytes))
}

func (this *HttpRequestNode) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	return schema.BuildParamInputSchemaWithSlice([]string{"request_url", "request_method?", "request_params?", "request_headers?"})
}

func (this *HttpRequestNode) GetRuntimeParamInputSchema() *schema.ParamInputSchema {
	return &schema.ParamInputSchema{}
}

func (this *HttpRequestNode) GetDefaultParamOutputSchema() *schema.ParamOutputSchema {
	return schema.BuildParamOutputSchemaWithSlice([]string{"response_str", "response_bytes", "base64res_str", "StatusCode", "ContentType"})
}

func (this *HttpRequestNode) GetRuntimeParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}

func fillParamMapData(tmpDataMap map[string]interface{}, paramName string) map[string]interface{} {
	paramMap := make(map[string]interface{})
	if _paramName, ok := tmpDataMap[paramName].(string); ok {
		if paramName, paramValue := checkParameter(_paramName); strings.TrimSpace(paramName) != ""{
			paramMap[strings.TrimSpace(paramName)] = strings.TrimSpace(paramValue)
		}
	} else if _paramNames, ok := tmpDataMap[paramName].([]string); ok {
		for _, _paramName := range _paramNames {
			if paramName, paramValue := checkParameter(_paramName); strings.TrimSpace(paramName) != ""{
				paramMap[strings.TrimSpace(paramName)] = strings.TrimSpace(paramValue)
			}
		}
	}
	return paramMap
}

func checkParameter(s string) (paramName, paramValue string) {
	s = strings.TrimSpace(s)
	if !strings.Contains(s, "="){
		panic(errors.New(fmt.Sprint("invalid parameter for %s", s)))
	}
	index := strings.Index(s, "=")
	paramName = strings.TrimSpace(s[:index])
	paramValue = strings.TrimSpace(s[index+1:])
	return
}