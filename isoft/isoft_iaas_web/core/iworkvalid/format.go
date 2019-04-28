package iworkvalid

import (
	"isoft/isoft_iaas_web/core/iworkfunc"
)

type ParamValueFormatChecker struct {
	ParamValue string
}

// 定义规则：函数中所有的值都必须是变量,不能是字符串
// 变量定义区域
func (this *ParamValueFormatChecker) Check() (bool, error) {
	// 暂不支持 ` 反引号转义
	_, err := iworkfunc.GetAllFuncExecutor(this.ParamValue)
	if err != nil {
		return false, err
	}
	return true, nil
}
