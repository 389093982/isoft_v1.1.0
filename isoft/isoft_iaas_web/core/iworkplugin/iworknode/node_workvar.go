package iworknode

import (
	"isoft/isoft_iaas_web/core/iworkconst"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworkmodels"
	"isoft/isoft_iaas_web/models/iwork"
)

type WorkVarAssignNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *WorkVarAssignNode) Execute(trackingId string) {
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, this.DataStore)
	workVarName := tmpDataMap[iworkconst.STRING_PREFIX+"workVarName"].(string)
	workVarValue := tmpDataMap[iworkconst.STRING_PREFIX+"workVarValue"].(string)
	this.DataStore.CacheDatas("__workVars__", map[string]interface{}{workVarName: workVarValue})
}

func (this *WorkVarAssignNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "workVarName", "流程变量名称"},
		2: {iworkconst.STRING_PREFIX + "workVarValue", "流程变量值"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}
