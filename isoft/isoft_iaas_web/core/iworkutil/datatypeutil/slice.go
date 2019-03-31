package datatypeutil

import (
	"reflect"
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
