package iworknode

import (
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworkutil"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
)

type WorkSub struct {
	BaseNode
	WorkStep *iwork.WorkStep
	RunFunc  func(work iwork.Work, steps []iwork.WorkStep, args ...interface{})
}

func (this *WorkSub) Execute(trackingId string) {
	// 从 db 中读取 paramInputSchema
	paramInputSchema := schema.GetCacheParamInputSchema(this.WorkStep, &WorkStepFactory{WorkStep: this.WorkStep})
	for _, item := range paramInputSchema.ParamInputSchemaItems {
		if item.ParamName == "work_sub" && strings.HasPrefix(item.ParamValue, "$WORK.") {
			// 找到 work_sub 字段值
			workSubName := iworkutil.GetWorkSubNameFromParamValue(item.ParamValue)
			work, _ := iwork.QueryWorkByName(workSubName)
			steps, _ := iwork.GetAllWorkStepByWorkName(workSubName)
			// 运行子流程
			this.RunFunc(work, steps, trackingId)
		}
	}
}

func (this *WorkSub) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	return schema.BuildParamInputSchemaWithSlice([]string{"work_sub"})
}
func (this *WorkSub) GetDefaultParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}
func (this *WorkSub) GetRuntimeParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}
