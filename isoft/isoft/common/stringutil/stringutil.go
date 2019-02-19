package stringutil

import (
	"github.com/satori/go.uuid"
	"reflect"
)

func RandomUUID() string {
	return uuid.NewV4().String()
}

func GetTypeOfInterface(v interface{}) string {
	return reflect.TypeOf(v).String()
}

func CheckContains(s string, slice []string) bool {
	if len(slice) == 0{
		return false
	}
	for _, _s := range slice{
		if _s == s{
			return true
		}
	}
	return false
}