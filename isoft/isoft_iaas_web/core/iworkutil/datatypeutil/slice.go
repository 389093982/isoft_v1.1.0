package datatypeutil

import (
	"reflect"
	"strings"
)

func ReverseSlice(s interface{}) interface{} {
	t := reflect.TypeOf(s)
	if t.Kind() != reflect.Slice {
		return s
	}
	v := reflect.MakeSlice(reflect.TypeOf(s), 0, 0)
	for i := reflect.ValueOf(s).Len() - 1; i >= 0; i-- {
		v = reflect.Append(v, reflect.ValueOf(s).Index(i))
	}
	return v.Interface()
}

func FilterSlice(ss []string, filterFunc func(s string) bool) []string {
	result := make([]string, 0)
	for _, str := range ss {
		if filterFunc(str) {
			result = append(result, str)
		}
	}
	return result
}

func CheckEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}

func CheckNotEmpty(s string) bool {
	return !CheckEmpty(s)
}
