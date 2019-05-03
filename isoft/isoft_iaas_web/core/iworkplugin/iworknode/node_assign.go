package iworknode

import (
	"encoding/json"
	"isoft/isoft_iaas_web/core/iworkmodels"
	"isoft/isoft_iaas_web/models/iwork"
)

type AssignVarNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *AssignVarNode) Execute(trackingId string) {

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
				"`assign`",
				"`arrayassign`",
				"`mapassign`",
			},
		})
		items = append(items, iworkmodels.ParamInputSchemaItem{ParamName: paramMapping + "_value"})
	}
	return &iworkmodels.ParamInputSchema{ParamInputSchemaItems: items}
}
