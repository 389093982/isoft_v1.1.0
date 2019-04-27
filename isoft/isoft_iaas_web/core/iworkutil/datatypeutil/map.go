package datatypeutil

import "reflect"

// m 对应的 map
// s 返回的 key 切片
func GetMapKeySlice(m interface{}, s interface{}) interface{} {
	t := reflect.TypeOf(m)
	if t.Kind() != reflect.Map {
		return s
	}
	for _, k := range reflect.ValueOf(m).MapKeys() {
		s = reflect.Append(reflect.ValueOf(s), k).Interface()
	}
	return s
}

func GetMapValueSlice(m interface{}, s interface{}) interface{} {
	t := reflect.TypeOf(m)
	if t.Kind() != reflect.Map {
		return s
	}
	for r := reflect.ValueOf(m).MapRange(); r.Next(); {
		s = reflect.Append(reflect.ValueOf(s), r.Value()).Interface()
	}
	return s
}

func CombineMap(m, m1 map[string]interface{}) map[string]interface{} {
	for k, v := range m1 {
		m[k] = v
	}
	return m
}
