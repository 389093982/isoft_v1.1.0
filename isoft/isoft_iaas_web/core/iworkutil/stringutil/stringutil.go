package stringutil

import "strings"

// 将 v 转成非空字符串,转换失败时使用默认值
func GetString(v interface{}, def string, trim ...bool) string {
	if _, ok := v.(string); !ok {
		return def
	}
	s := v.(string)
	if len(trim) > 0 && trim[0] == true {
		s = strings.TrimSpace(s)
	}
	if s != "" {
		return s
	}
	return def
}
