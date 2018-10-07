package httputil

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetParamReader(paramMap map[string]interface{}) *strings.Reader {
	if paramMap == nil || len(paramMap) == 0 {
		return nil
	}
	var s []string = make([]string, 3)
	for k, v := range paramMap {
		s = append(s, k+"="+v.(string))
	}
	paramStr := strings.Join(s, "&")
	return strings.NewReader(paramStr)
}

func DoPost(url string, paramMap map[string]interface{}) (map[string]interface{}, error) {
	resp, err := http.Post(url, "application/x-www-form-urlencoded", GetParamReader(paramMap))
	if err != nil {
		return nil, err
	}
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resultmap := make(map[string]interface{}, 3)
	err = json.Unmarshal(responseBody, &resultmap)
	if err != nil {
		return nil, err
	}
	return resultmap, nil
}
