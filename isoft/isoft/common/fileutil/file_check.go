package fileutil

import "strings"

func CheckFilePathValid(filepath string) bool {
	return CheckFilePathPrefixValid(filepath)
}

// 校验文件路径前缀是否合法
func CheckFilePathPrefixValid(filepath string) bool {
	filepath = strings.TrimSpace(filepath)
	if filepath == "" {
		return false
	}
	// windows 路径前缀判断 :
	if strings.Index(filepath, ":") == 1 {
		return true
	}
	// linux 路径前缀判断 /
	if strings.HasPrefix(filepath, "/") {
		return true
	}
	return false
}
