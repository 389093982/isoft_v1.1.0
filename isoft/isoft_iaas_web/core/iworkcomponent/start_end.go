package iworkcomponent

import (
	"isoft/isoft_iaas_web/core/iworkdata"
	"isoft/isoft_iaas_web/models/iwork"
)

type WorkStartNode struct {
	WorkStep 		   *iwork.WorkStep
}

type WorkEndNode struct {
	WorkStep 		   *iwork.WorkStep
}

func (this *WorkStartNode) GetDefaultParamOutputSchema() *iworkdata.ParamOutputSchema {
	return transferParamInputSchemaToParamOutputSchema(this.WorkStep)
}

func (this *WorkEndNode) GetDefaultParamOutputSchema() *iworkdata.ParamOutputSchema {
	return transferParamInputSchemaToParamOutputSchema(this.WorkStep)
}

// 输入转输出,适用于开始节点和结束节点
func transferParamInputSchemaToParamOutputSchema(step *iwork.WorkStep) *iworkdata.ParamOutputSchema {
	items := []iworkdata.ParamOutputSchemaItem{}
	paramInputSchema := GetCacheParamInputSchema(step)
	for _, paramInputSchemaItem := range paramInputSchema.ParamInputSchemaItems{
		items = append(items, iworkdata.ParamOutputSchemaItem{ParamName: paramInputSchemaItem.ParamName})
	}
	return &iworkdata.ParamOutputSchema{ParamOutputSchemaItems: items}
}