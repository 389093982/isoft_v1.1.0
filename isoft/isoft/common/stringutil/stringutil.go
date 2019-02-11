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

