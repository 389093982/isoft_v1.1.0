package deploy_core

import "strings"

// 获取真实的命令脚本执行类型
func GetRealCommandType(serviceType, operate_type string) string {
	if !strings.HasPrefix(operate_type, serviceType+"_") {
		return serviceType + "_" + operate_type
	}
	return operate_type
}
