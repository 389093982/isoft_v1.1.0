package iworkutil

import (
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"strings"
)

func GetWorkSubNameFromParamValue(paramValue string) string {
	value := strings.TrimSpace(paramValue)
	value = strings.Replace(value, "$WORK.", "", -1)
	value = strings.Replace(value, "__sep__", "", -1)
	value = strings.Replace(value, "\n", "", -1)
	value = strings.TrimSpace(value)
	return value
}


func GetWorkSubNameForWorkSubNode(paramInputSchema *schema.ParamInputSchema) string {
	for _, item := range paramInputSchema.ParamInputSchemaItems {
		if item.ParamName == "work_sub" && strings.HasPrefix(item.ParamValue, "$WORK.") {
			// 找到 work_sub 字段值
			return GetWorkSubNameFromParamValue(item.ParamValue)

		}
	}
	return ""
}