package iworkutil

import (
	"encoding/base64"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"strings"
)

func GetWorkSubNameFromParamValue(paramValue string) string {
	value := strings.TrimSpace(paramValue)
	value = strings.Replace(value, "$WORK.", "", -1)
	value = strings.Replace(value, ";", "", -1)
	value = strings.Replace(value, "\n", "", -1)
	value = strings.TrimSpace(value)
	return value
}

func GetWorkSubNameForWorkSubNode(paramInputSchema *schema.ParamInputSchema) string {
	for _, item := range paramInputSchema.ParamInputSchemaItems {
		if item.ParamName == "work_sub" && strings.HasPrefix(strings.TrimSpace(item.ParamValue), "$WORK.") {
			// 找到 work_sub 字段值
			return GetWorkSubNameFromParamValue(strings.TrimSpace(item.ParamValue))

		}
	}
	return ""
}

func EncodeToBase64String(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

func DecodeBase64String(encodeString string) (bytes []byte) {
	if bytes, err := base64.StdEncoding.DecodeString(encodeString); err != nil {
		return bytes
	}
	return
}
