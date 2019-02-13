package iwork

import (
	"encoding/json"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworknode"
	"isoft/isoft_iaas_web/models/iwork"
	"time"
)

// 构建动态输入值
func BuildDynamicInput(work_id string, work_step_id int64) {
	// 读取 work_step 信息
	step, err := iwork.LoadWorkStepInfo(work_id, work_step_id)
	if err != nil {
		panic(err)
	}
	// 获取默认数据
	defaultParamInputSchema := schema.GetDefaultParamInputSchema(&iworknode.WorkStepFactory{WorkStep: &step})
	// 获取动态数据
	runtimeParamInputSchema := schema.GetRuntimeParamInputSchema(&iworknode.WorkStepFactory{WorkStep: &step})
	// 合并默认数据和动态数据作为新数据
	newInputSchemaItems := append(defaultParamInputSchema.ParamInputSchemaItems, runtimeParamInputSchema.ParamInputSchemaItems...)
	// 获取历史数据
	historyParamInputSchema := schema.GetCacheParamInputSchema(&step, &iworknode.WorkStepFactory{WorkStep: &step})
	for index, newInputSchemaItem := range newInputSchemaItems {
		// 存在则不添加且沿用旧值
		if exist, paramValue := CheckAndGetParamValueByInputSchemaParamName(historyParamInputSchema.ParamInputSchemaItems, newInputSchemaItem.ParamName); exist {
			newInputSchemaItems[index].ParamValue = paramValue
		}
	}
	paramInputSchema := &schema.ParamInputSchema{ParamInputSchemaItems:newInputSchemaItems}
	step.WorkStepInput = paramInputSchema.RenderToXml()
	if _, err = iwork.InsertOrUpdateWorkStep(&step); err != nil {
		panic(err)
	}
}

// 构建动态输出值
func BuildDynamicOutput(work_id string, work_step_id int64) {
	// 读取 work_step 信息
	step, err := iwork.LoadWorkStepInfo(work_id, work_step_id)
	if err != nil {
		panic(err)
	}
	// 构建输出参数,使用全新值
	step.WorkStepOutput = schema.GetRuntimeParamOutputSchema(&iworknode.WorkStepFactory{WorkStep: &step}).RenderToXml()
	if _, err = iwork.InsertOrUpdateWorkStep(&step); err != nil {
		panic(err)
	}
}

// 构建动态值
func BuildDynamic(work_id string, work_step_id int64) {
	// 构建动态输入值
	BuildDynamicInput(work_id, work_step_id)
	// 构建动态输出值
	BuildDynamicOutput(work_id, work_step_id)
}

func (this *WorkController) EditWorkStepParamInfo() {
	work_id := this.GetString("work_id")
	work_step_id, _ := this.GetInt64("work_step_id", -1)
	paramInputSchemaStr := this.GetString("paramInputSchemaStr")
	paramMappingsStr := this.GetString("paramMappingsStr")
	var paramInputSchema schema.ParamInputSchema
	json.Unmarshal([]byte(paramInputSchemaStr), &paramInputSchema)
	this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	if step, err := iwork.GetOneWorkStep(work_id, work_step_id); err == nil {
		step.WorkStepInput = paramInputSchema.RenderToXml()
		step.WorkStepParamMapping = paramMappingsStr
		step.CreatedBy = "SYSTEM"
		step.CreatedTime = time.Now()
		step.LastUpdatedBy = "SYSTEM"
		step.LastUpdatedTime = time.Now()
		if _, err := iwork.InsertOrUpdateWorkStep(&step); err == nil {
			this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
			// 保存完静态参数后自动构建获动态参数并保存
			BuildDynamic(work_id, work_step_id)
		}
	}
	this.ServeJSON()
}

func CheckAndGetParamValueByInputSchemaParamName(items []schema.ParamInputSchemaItem, paramName string) (exist bool, paramValue string) {
	for _, item := range items {
		if item.ParamName == paramName {
			return true, item.ParamValue
		}
	}
	return false, ""
}

