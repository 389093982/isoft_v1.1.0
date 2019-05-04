package iworknode

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworkmodels"
	"isoft/isoft_iaas_web/models/iwork"
)

type DefineVarNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *DefineVarNode) Execute(trackingId string) {
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep)
	dataMap := make(map[string]interface{}, 0)
	for paramName, paramType := range tmpDataMap {
		var paramValue interface{}
		switch paramType {
		case `string`:
			paramValue = ""
		case `interface{}`:
			paramValue = new(interface{})
		case `[]interface{}`:
			paramValue = make([]interface{}, 0)
		case `map[string]interface{}`:
			paramValue = make(map[string]interface{}, 0)
		default:
			panic(errors.New(fmt.Sprintf("unsupport paramType for %s", paramType)))
		}
		dataMap[paramName] = paramValue
	}
	this.DataStore.CacheDatas(this.WorkStep.WorkStepName, dataMap)
}

func (this *DefineVarNode) GetRuntimeParamInputSchema() *iworkmodels.ParamInputSchema {
	var paramMappingsArr []string
	json.Unmarshal([]byte(this.WorkStep.WorkStepParamMapping), &paramMappingsArr)
	items := make([]iworkmodels.ParamInputSchemaItem, 0)
	for _, paramMapping := range paramMappingsArr {
		items = append(items, iworkmodels.ParamInputSchemaItem{
			ParamName: paramMapping,
			ParamChoices: []string{
				"`string`",
				"`interface{}`",
				"`[]interface{}`",
				"`map[string]interface{}`",
			},
		})
	}
	return &iworkmodels.ParamInputSchema{ParamInputSchemaItems: items}
}

func (this *DefineVarNode) GetRuntimeParamOutputSchema() *iworkmodels.ParamOutputSchema {
	items := make([]iworkmodels.ParamOutputSchemaItem, 0)
	inputSchema := schema.GetCacheParamInputSchema(this.WorkStep, &WorkStepFactory{WorkStep: this.WorkStep})
	for _, item := range inputSchema.ParamInputSchemaItems {
		items = append(items, iworkmodels.ParamOutputSchemaItem{ParamName: item.ParamName})
	}
	return &iworkmodels.ParamOutputSchema{ParamOutputSchemaItems: items}
}
