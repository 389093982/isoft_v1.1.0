package iworknode

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"isoft/isoft_iaas_web/core/iworkmodels"
	"isoft/isoft_iaas_web/models/iwork"
	"regexp"
	"strings"
)

type AssignVarNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

// 检测赋值引用字段格式是否满足正则
func checkAssignReferFormat(paramName, assignVar_pureText string) {
	reg := regexp.MustCompile("^\\$[a-zA-Z_0-9]+\\.[a-zA-Z0-9\\-]+$")
	if !reg.MatchString(assignVar_pureText) {
		panic(errors.New(fmt.Sprintf(`error format assign for %s`, paramName)))
	}
}

func parseAssignRefer(paramName, assignVar_pureText string) (string, string) {
	assignVar_pureText = strings.TrimSpace(strings.ReplaceAll(assignVar_pureText, ";", ""))
	checkAssignReferFormat(paramName, assignVar_pureText)
	// 去除 $ 和 . 后面的字符得到 assignNode 名称
	assignNodeName := assignVar_pureText[1:strings.Index(assignVar_pureText, ".")]
	// 截取 assignData 名称
	assignDataName := assignVar_pureText[strings.Index(assignVar_pureText, ".")+1:]
	return assignNodeName, assignDataName
}

func (this *AssignVarNode) Execute(trackingId string) {
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, this.DataStore)
	// pureText 节点中间数据
	pureTextTmpDataMap := this.FillPureTextParamInputSchemaDataToTmp(this.WorkStep, this.DataStore)
	for paramName, paramValue := range tmpDataMap {
		if strings.HasSuffix(paramName, "_operate") {
			_paramName := paramName[:strings.LastIndex(paramName, "_operate")]
			assignVar := tmpDataMap[_paramName]
			assignOperate := paramValue.(string)
			assignVal := tmpDataMap[_paramName+"_value"]

			// pureText
			assignVar_pureText := pureTextTmpDataMap[_paramName].(string)
			assignNodeName, assignDataName := parseAssignRefer(_paramName, assignVar_pureText)

			assign := NodeAssign{
				AssignVar:     assignVar,
				AssignOperate: assignOperate,
				AssignData:    assignVal,
			}
			assignVar, err := assign.Calculate()
			if err != nil {
				panic(err)
			}
			// 重新将值绑定到对应的 assign 节点
			this.DataStore.CacheDatas(assignNodeName, map[string]interface{}{
				assignDataName: assignVar,
			})
		}
	}
}

func (this *AssignVarNode) GetRuntimeParamInputSchema() *iworkmodels.ParamInputSchema {
	var paramMappingsArr []string
	json.Unmarshal([]byte(this.WorkStep.WorkStepParamMapping), &paramMappingsArr)
	items := make([]iworkmodels.ParamInputSchemaItem, 0)
	for _, paramMapping := range paramMappingsArr {
		items = append(items, iworkmodels.ParamInputSchemaItem{ParamName: paramMapping})
		items = append(items, iworkmodels.ParamInputSchemaItem{
			ParamName: paramMapping + "_operate",
			ParamChoices: []string{
				"`stringAssign`",
				"`interface{}Assign`",
				"`[]interface{}Assign`",
				"`map[string]interface{}Assign`",
			},
		})
		items = append(items, iworkmodels.ParamInputSchemaItem{ParamName: paramMapping + "_value"})
	}
	return &iworkmodels.ParamInputSchema{ParamInputSchemaItems: items}
}
