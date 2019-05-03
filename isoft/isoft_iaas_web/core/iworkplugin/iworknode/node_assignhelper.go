package iworknode

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
)

type NodeAssign struct {
	AssignVar     interface{}
	AssignOperate string
	AssignData    interface{}
}

func (this *NodeAssign) Calculate() (interface{}, error) {
	t := reflect.TypeOf(this.AssignVar)
	switch t.Kind() {
	case reflect.String:
		return this.CalculateString()
	case reflect.Interface, reflect.Ptr:
		return this.CalculateInterface()
	case reflect.Slice:
		return this.CalculateSlice()
	case reflect.Map:
		return this.CalculateMap()
	default:
		return nil, errors.New(fmt.Sprintf(`unsupport type for %s`, t.Kind().String()))
	}
}

func (this *NodeAssign) CalculateInterface() (interface{}, error) {
	if this.AssignOperate != `interface{}Assign` {
		return nil, errors.New(fmt.Sprintf("unsupport operateType for %s", this.AssignOperate))
	}
	if data, ok := this.AssignData.(interface{}); ok {
		return data, nil
	} else {
		return nil, errors.New(fmt.Sprintf(`unsupport data for %v`, this.AssignData))
	}
}

func (this *NodeAssign) CalculateMap() (interface{}, error) {
	if this.AssignOperate != `map[string]interface{}Assign` {
		return nil, errors.New(fmt.Sprintf("unsupport operateType for %s", this.AssignOperate))
	}
	if data, ok := this.AssignData.([]interface{}); ok {
		mapData := this.AssignData.(map[string]interface{})
		mapData[data[0].(string)] = data[1]
		return mapData, nil
	} else {
		return nil, errors.New(fmt.Sprintf(`unsupport data for %v`, this.AssignData))
	}
}

func (this *NodeAssign) CalculateSlice() (interface{}, error) {
	if this.AssignOperate != `[]interface{}Assign` {
		return nil, errors.New(fmt.Sprintf("unsupport operateType for %s", this.AssignOperate))
	}
	sliceData := this.AssignData.([]interface{})
	if data, ok := this.AssignData.([]interface{}); ok {
		for _, _data := range data {
			sliceData = append(sliceData, _data)
		}
		return sliceData, nil
	} else {
		return nil, errors.New(fmt.Sprintf(`unsupport data for %v`, this.AssignData))
	}
}

func (this *NodeAssign) CalculateString() (interface{}, error) {
	if this.AssignOperate != `stringAssign` {
		return nil, errors.New(fmt.Sprintf("unsupport operateType for %s", this.AssignOperate))
	}
	if data, ok := this.AssignData.(string); ok {
		return data, nil
	} else {
		return nil, errors.New(fmt.Sprintf(`unsupport data for %v`, this.AssignData))
	}
}
