package param

import (
	"encoding/xml"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
)

type ParamVauleParser struct {
	ParamValue string
}

func (this *ParamVauleParser) CheckParamValueFormat() bool {
	this.removeUnsupportChars()
	if strings.HasPrefix(this.ParamValue, "$") && strings.Contains(this.ParamValue, ".") {
		return true
	}
	return false
}

func (this *ParamVauleParser) GetNodeNameFromParamValue() string {
	if this.CheckParamValueFormat() {
		return this.ParamValue[1:strings.Index(this.ParamValue, ".")]
	}
	return this.ParamValue
}

func (this *ParamVauleParser) GetParamNameFromParamValue() string {
	if this.CheckParamValueFormat() {
		return this.ParamValue[strings.Index(this.ParamValue, ".")+1:]
	}
	return this.ParamValue
}

// 去除不合理的字符
func (this *ParamVauleParser) removeUnsupportChars() {
	this.ParamValue = strings.TrimSpace(this.ParamValue)
	this.ParamValue = strings.Replace(this.ParamValue, "\n", "", -1)
}

func (this *ParamVauleParser) GetStaticParamValue() string {
	this.removeUnsupportChars()
	if strings.HasPrefix(this.ParamValue, "$RESOURCE.") {
		resource_name := strings.TrimSpace(this.ParamValue)
		resource_name = strings.Replace(resource_name, "$RESOURCE.", "", -1)
		resource_name = strings.Replace(resource_name, ";", "", -1)
		resource_name = strings.TrimSpace(resource_name)
		return iwork.GetResourceDataSourceNameString(resource_name)
	}
	return this.ParamValue
}

type ParamNameParser struct {
	ParamName string
	Step      *iwork.WorkStep
}

// 根据 ParamName 获取相对值,真值可能需要 ParamVauleParser 处理一下
func (this *ParamNameParser) ParseAndGetRelativeParamValue() string {
	var paramInputSchema schema.ParamInputSchema
	if err := xml.Unmarshal([]byte(this.Step.WorkStepInput), &paramInputSchema); err != nil {
		return ""
	}
	for _, item := range paramInputSchema.ParamInputSchemaItems {
		if item.ParamName == this.ParamName {
			// 非必须参数不得为空
			if !strings.HasSuffix(item.ParamName, "?") && strings.TrimSpace(item.ParamValue) == "" {
				//panic(errors.New(fmt.Sprint("it is a mast parameter for %s", item.ParamName)))
				return ""
			}
			return item.ParamValue
		}
	}
	return ""
}

// 根据步骤和参数名称获取静态参数值
func GetStaticParamValue(paramName string, step *iwork.WorkStep) string {
	paramNameParser := &ParamNameParser{
		ParamName: paramName,
		Step:      step,
	}
	paramValueParser := &ParamVauleParser{
		ParamValue: paramNameParser.ParseAndGetRelativeParamValue(),
	}
	return paramValueParser.GetStaticParamValue()
}
