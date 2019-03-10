package stringutil

import (
	"github.com/satori/go.uuid"
	"reflect"
	"regexp"
)

func RandomUUID() string {
	return uuid.NewV4().String()
}

func GetTypeOfInterface(v interface{}) string {
	return reflect.TypeOf(v).String()
}

func CheckContains(s string, slice []string) bool {
	if len(slice) == 0 {
		return false
	}
	for _, _s := range slice {
		if _s == s {
			return true
		}
	}
	return false
}

func ChangeStringsToInterfaces(ss []string) []interface{} {
	result := make([]interface{}, 0)
	for _, s := range ss {
		result = append(result, s)
	}
	return result
}

func GetSubStringWithRegexp(s, regex string) []string {
	reg := regexp.MustCompile(regex)
	return reg.FindAllString(s, -1)
}

func GetNoRepeatSubStringWithRegexp(s, regex string) []string {
	ss := GetSubStringWithRegexp(s, regex)
	return RemoveRepeatForSlice(ss)
}

// 通过map主键唯一的特性过滤重复元素
func RemoveRepeatForSlice(slc []string) []string {
	result := []string{}
	// 存放不重复主键
	tempMap := map[string]byte{}
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l {
			// 加入map后,map长度变化,则元素不重复
			result = append(result, e)
		}
	}
	return result
}
