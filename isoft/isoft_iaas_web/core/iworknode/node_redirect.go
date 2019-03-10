package iworknode

import (
	"encoding/json"
	"fmt"
	"isoft/isoft_iaas_web/core/iworkconst"
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
)

type GotoConditionNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func getMappingInfoWithRemovePrefixAndSuffix(paramName string) string {
	paramName = strings.Replace(paramName, iworkconst.BOOL_PREFIX, "", -1)
	paramName = strings.Replace(paramName, "_condition", "", -1)
	return paramName
}

func (this *GotoConditionNode) Execute(trackingId string) {
	// 获取数据中心
	dataStore := datastore.GetDataStore(trackingId)
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, dataStore)
	inputSchema := schema.GetCacheParamInputSchema(this.WorkStep, &WorkStepFactory{WorkStep: this.WorkStep})
	for _, item := range inputSchema.ParamInputSchemaItems {
		if item.ParamName == iworkconst.BOOL_PREFIX+"goto_end_condition?" {
			if bol, ok := tmpDataMap[item.ParamName].(bool); ok && bol == true {
				// 往 dataStore 中发送一条 redirect 指令
				dataStore.CacheData("__goto_condition__", "__redirect__", "end")
				return
			}
		} else if item.ParamName == iworkconst.BOOL_PREFIX+"goto_out_condition?" {
			if bol, ok := tmpDataMap[item.ParamName].(bool); ok && bol == true {
				// 往 dataStore 中发送一条 redirect 指令
				dataStore.CacheData("__goto_condition__", "__redirect__", "__out__")
				return
			}
		} else if strings.HasSuffix(item.ParamName, "_condition") {
			if bol, ok := tmpDataMap[item.ParamName].(bool); ok && bol == true {
				// 表达式为 true 时执行跳转动作
				mappingInfo := getMappingInfoWithRemovePrefixAndSuffix(item.ParamName)
				redirectNodeName := tmpDataMap[iworkconst.STRING_PREFIX+mappingInfo+"_redirect"].(string)
				// 往 dataStore 中发送一条 redirect 指令
				dataStore.CacheData("__goto_condition__", "__redirect__", redirectNodeName)
				return
			}
		}
	}
}

func (this *GotoConditionNode) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	paramMap := map[int][]string{
		1: []string{iworkconst.BOOL_PREFIX + "goto_end_condition?", "表达式为真时直接跳往 end 节点!"},
		2: []string{iworkconst.BOOL_PREFIX + "goto_out_condition?", "表达式为真时直接结束当前流程,返回上级流程继续执行!"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *GotoConditionNode) GetRuntimeParamInputSchema() *schema.ParamInputSchema {
	var paramMappingsArr []string
	json.Unmarshal([]byte(this.WorkStep.WorkStepParamMapping), &paramMappingsArr)
	items := make([]schema.ParamInputSchemaItem, 0)
	for _, paramMapping := range paramMappingsArr {
		items = append(items, schema.ParamInputSchemaItem{
			ParamName: iworkconst.BOOL_PREFIX + paramMapping + "_condition",
			ParamDesc: fmt.Sprintf("条件满足时跳往对应节点继续执行！"),
		})
		items = append(items, schema.ParamInputSchemaItem{
			ParamName: iworkconst.STRING_PREFIX + paramMapping + "_redirect",
			ParamDesc: fmt.Sprintf("满足条件时跳往的节点！"),
		})
	}
	return &schema.ParamInputSchema{ParamInputSchemaItems: items}
}

func (this *GotoConditionNode) GetDefaultParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}

func (this *GotoConditionNode) GetRuntimeParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}

func (this *GotoConditionNode) ValidateCustom() {

}
