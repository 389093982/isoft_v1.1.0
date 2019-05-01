package iworkvalid

import (
	"isoft/isoft_iaas_web/core/iworkfunc"
)

type ParamValueFormatChecker struct {
	ParamName  string
	PureText   bool
	ParamValue string
}

// 定义规则：函数中所有的值都必须是变量,不能是字符串
// 变量定义区域
func (this *ParamValueFormatChecker) Check() (bool, error) {
	if this.PureText {
		return true, nil
	}
	_, err := iworkfunc.ParseToFuncCallers(this.ParamValue)
	if err != nil {
		return false, err
	}
	return true, nil
}
