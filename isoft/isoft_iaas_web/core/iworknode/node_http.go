package iworknode

import (
	"fmt"
	"github.com/pkg/errors"
	"isoft/isoft_iaas_web/core/iworkconst"
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
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, this.DataStore)
	// 参数准备
	var request_url, request_method string
	if _request_url, ok := tmpDataMap[iworkconst.STRING_PREFIX+"request_url"].(string); ok {
		request_url = _request_url
	}
	if _request_method, ok := tmpDataMap[iworkconst.STRING_PREFIX+"request_method?"].(string); ok {
		request_method = _request_method
	}
	paramMap := fillParamMapData(tmpDataMap, iworkconst.MULTI_PREFIX+"request_params?")
	headerMap := fillParamMapData(tmpDataMap, iworkconst.MULTI_PREFIX+"request_headers?")

	responsebytes := httputil.DoHttpRequestWithParserFunc(request_url, request_method, paramMap, headerMap, func(resp *http.Response) {
		this.DataStore.CacheData(this.WorkStep.WorkStepName, iworkconst.NUMBER_PREFIX+"StatusCode", resp.StatusCode)
		this.DataStore.CacheData(this.WorkStep.WorkStepName, iworkconst.STRING_PREFIX+"ContentType", resp.Header.Get("content-type"))
	})
	this.DataStore.CacheByteData(this.WorkStep.WorkStepName, iworkconst.STRING_PREFIX+"response_data", string(responsebytes))
	this.DataStore.CacheByteData(this.WorkStep.WorkStepName, iworkconst.BYTE_ARRAY_PREFIX+"response_data", responsebytes)
	this.DataStore.CacheByteData(this.WorkStep.WorkStepName, iworkconst.BASE64STRING_PREFIX+"response_data", iworkutil.EncodeToBase64String(responsebytes))
}

func (this *HttpRequestNode) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	paramMap := map[int][]string{
		1: []string{iworkconst.STRING_PREFIX + "request_url", "请求资源的url地址"},
		2: []string{iworkconst.STRING_PREFIX + "request_method?", "可选参数,请求方式,默认是GET请求,支持GET、POST"},
		3: []string{iworkconst.MULTI_PREFIX + "request_params?", "可选参数,请求参数,格式参考：key=value"},
		4: []string{iworkconst.MULTI_PREFIX + "request_headers?", "可选参数,请求头参数,格式参考：key=value"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *HttpRequestNode) GetRuntimeParamInputSchema() *schema.ParamInputSchema {
	return &schema.ParamInputSchema{}
}

func (this *HttpRequestNode) GetDefaultParamOutputSchema() *schema.ParamOutputSchema {
	return schema.BuildParamOutputSchemaWithSlice([]string{
		iworkconst.STRING_PREFIX + "response_data",
		iworkconst.BYTE_ARRAY_PREFIX + "response_data",
		iworkconst.BASE64STRING_PREFIX + "response_data",
		iworkconst.NUMBER_PREFIX + "StatusCode",
		iworkconst.STRING_PREFIX + "ContentType"})
}

func (this *HttpRequestNode) GetRuntimeParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}

func (this *HttpRequestNode) ValidateCustom() (checkResult []string) {
	return []string{}
}

func fillParamMapData(tmpDataMap map[string]interface{}, paramName string) map[string]interface{} {
	paramMap := make(map[string]interface{})
	if _paramName, ok := tmpDataMap[paramName].(string); ok {
		if paramName, paramValue := checkParameter(_paramName); strings.TrimSpace(paramName) != "" {
			paramMap[strings.TrimSpace(paramName)] = strings.TrimSpace(paramValue)
		}
	} else if _paramNames, ok := tmpDataMap[paramName].([]string); ok {
		for _, _paramName := range _paramNames {
			if paramName, paramValue := checkParameter(_paramName); strings.TrimSpace(paramName) != "" {
				paramMap[strings.TrimSpace(paramName)] = strings.TrimSpace(paramValue)
			}
		}
	}
	return paramMap
}

func checkParameter(s string) (paramName, paramValue string) {
	s = strings.TrimSpace(s)
	if !strings.Contains(s, "=") {
		panic(errors.New(fmt.Sprint("invalid parameter for %s", s)))
	}
	index := strings.Index(s, "=")
	paramName = strings.TrimSpace(s[:index])
	paramValue = strings.TrimSpace(s[index+1:])
	return
}
