package reflectutil

import (
	"fmt"
	"reflect"
)

// 将结构体里的成员按照字段名字来赋值
func FillFieldValueToStruct(ptr interface{}, fields map[string]interface{}) {
	v := reflect.ValueOf(ptr).Elem() // the struct variable
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		if value, ok := fields[fieldInfo.Name]; ok {
			//给结构体赋值
			//保证赋值时数据类型一致
			if reflect.ValueOf(value).Type() == v.FieldByName(fieldInfo.Name).Type() {
				v.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(value))
			} else {
				fmt.Println("mismatch field for struct....")
			}
		}
	}
	return
}
