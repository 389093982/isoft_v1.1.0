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

}

func (this *WorkVarAssignNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "workVarName", "流程变量名称"},
		2: {iworkconst.BOOL_PREFIX + "workVarValue", "流程变量值"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}
