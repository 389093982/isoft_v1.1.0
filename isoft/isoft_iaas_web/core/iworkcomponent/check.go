package iworkcomponent

import "strings"

// 判断参数是否是动态参数
func IsDynamicParam(param string) bool {
	return strings.HasPrefix(param, "$")
}
