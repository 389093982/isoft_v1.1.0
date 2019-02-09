package iworkdata

import "strings"

type ParamVauleParser struct {
	ParamValue string
}

func (this *ParamVauleParser) CheckParamValueFormat() bool {
	this.removeUnsupportChars()
	if strings.HasPrefix(this.ParamValue, "$") && strings.Contains(this.ParamValue, "."){
		return true
	}
	return false
}

func (this *ParamVauleParser) GetNodeNameFromParamValue() string {
	if this.CheckParamValueFormat(){
		return this.ParamValue[1:strings.Index(this.ParamValue, ".")]
	}
	return this.ParamValue
}

func (this *ParamVauleParser) GetParamNameFromParamValue() string {
	if this.CheckParamValueFormat(){
		return this.ParamValue[strings.Index(this.ParamValue, ".") + 1:]
	}
	return this.ParamValue
}

// 去除不合理的字符
func (this *ParamVauleParser) removeUnsupportChars() {
	this.ParamValue = strings.TrimSpace(this.ParamValue)
	this.ParamValue = strings.Replace(this.ParamValue, "\n","",-1)
}