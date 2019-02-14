package httputil

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func DoHttpRequestWithParserFunc(url string, method string, paramMap map[string]interface{}, parseFunc func(resp *http.Response)) (responsebody []byte) {
	client := &http.Client{}
	req, err := http.NewRequest(checkMethod(method), url, GetParamReader(paramMap))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	parseFunc(resp)
	responsebody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return
}

func DoHttpRequest(url string, method string, paramMap map[string]interface{}) (responsebody []byte) {
	return DoHttpRequestWithParserFunc(url, method, paramMap, func(resp *http.Response) {})
}

func GetParamReader(paramMap map[string]interface{}) *strings.Reader {
	if paramMap == nil || len(paramMap) == 0 {
		return strings.NewReader("")
	}
	s := make([]string, 0)
	for k, v := range paramMap {
		s = append(s, k+"="+v.(string))
	}
	paramStr := strings.Join(s, "&")
	return strings.NewReader(paramStr)
}

func checkMethod(method string) string {
	method = strings.TrimSpace(method)
	if method == ""{
		return "GET"
	}
	defaultMethods := []string{"GET","POST","PUT","DELETE"}
	for _, dm := range defaultMethods{
		if strings.ToUpper(method) == dm{
			return dm
		}
	}
	panic(fmt.Sprintf("unsupport method : %s", method))
}

