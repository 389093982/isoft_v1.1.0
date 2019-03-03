package iworkservice

import "isoft/isoft_iaas_web/models/iwork"

func LoadValidateResultService(serviceArgs map[string]interface{}) (result map[string]interface{}, err error) {
	result = make(map[string]interface{}, 0)
	details, err := iwork.QueryLastValidateLogDetail()
	if err != nil {
		return nil, err
	}
	result["details"] = details
	return
}
