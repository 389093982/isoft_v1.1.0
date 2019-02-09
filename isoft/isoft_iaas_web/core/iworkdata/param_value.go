package iworkdata

import "strings"

type ParamVauleParser struct {
	ParamValue string
}

func (this *ParamVauleParser) CheckParamValueFormat() bool {
	if strings.HasPrefix(this.ParamValue, "$") && strings.Contains(this.ParamValue, "."){
		return true
	}
	return false
}

func (this *ParamVauleParser) GetNodeNameFromParamValue() string {
	if this.CheckParamValueFormat(){
		return string([]rune(this.ParamValue)[1:strings.Index(this.ParamValue, ".")])
	}
	return this.ParamValue
}

func (this *ParamVauleParser) GetParamNameFromParamValue() string {
	if this.CheckParamValueFormat(){
		return string([]rune(this.ParamValue)[strings.Index(this.ParamValue, ".") + 1:])
	}
	return this.ParamValue
}


