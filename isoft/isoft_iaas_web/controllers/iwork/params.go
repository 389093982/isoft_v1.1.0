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

			//// 如果是 start 或者 end 类型的节点,则通知其做一些事后适配
			//NoticeWorkStartEndAdjust(work_id, work_step_id)
			//// 如果是 work_sub 类型的节点,则通知其做一些事后适配
			//NoticeWorkSubAdjust(work_id, work_step_id)
			// 保存完静态参数后自动构建获动态参数并保存
			BuildDynamic(work_id, work_step_id)
		}
	}
	this.ServeJSON()
}

//// paramMappings 只有起始和结束节点才有,而且起始和结束节点的 paramMappings 也是 paramInput 和 paramOutput
//func NoticeWorkStartEndAdjust(work_id string, work_step_id int64) {
//	// 读取 step 记录
//	if step, err := iwork.GetOneWorkStep(work_id, work_step_id); err == nil &&
//		(step.WorkStepType == "work_start" || step.WorkStepType == "work_end") {
//		adjustWorkStartEndNodeParamSchema(step.WorkStepParamMapping, &step)
//	}
//}

//func NoticeWorkSubAdjust(work_id string, work_step_id int64) {
//	// 读取 step 记录
//	if step, err := iwork.GetOneWorkStep(work_id, work_step_id); err == nil && step.WorkStepType == "work_sub" {
//		// 从 db 中读取 paramInputSchema
//		paramInputSchema := schema.GetCacheParamInputSchema(&step, &iworknode.WorkStepFactory{WorkStep: &step})
//		workSubName := iworkutil.GetWorkSubNameForWorkSubNode(paramInputSchema)
//		if strings.TrimSpace(workSubName) != ""{
//			adjustWorkSubNodeParamSchema(workSubName, *paramInputSchema, step)
//		}
//	}
//}

//func adjustWorkStartEndNodeParamSchema(paramMappingsStr string, step *iwork.WorkStep) {
//	paramInputSchema := schema.GetCacheParamInputSchema(step, &iworknode.WorkStepFactory{WorkStep: step})
//	var paramMappingsArr []string
//	json.Unmarshal([]byte(paramMappingsStr), &paramMappingsArr)
//	// 沿用旧值,添加新值,去除无效的值,即以 paramMapping 为准
//	items := []schema.ParamInputSchemaItem{}
//	for _, paramMapping := range paramMappingsArr {
//		var oldValue string // 旧值默认为空
//		for _, _item := range paramInputSchema.ParamInputSchemaItems {
//			if _item.ParamName == paramMapping {
//				oldValue = _item.ParamValue
//			}
//		}
//		items = append(items, schema.ParamInputSchemaItem{ParamName: paramMapping, ParamValue: oldValue})
//	}
//	paramInputSchema.ParamInputSchemaItems = items
//	step.WorkStepInput = paramInputSchema.RenderToXml()
//	iwork.InsertOrUpdateWorkStep(step)
//}

//// 调整 work_sub 类型节点参数
//func adjustWorkSubNodeParamSchema(workSubName string, paramInputSchema schema.ParamInputSchema, step iwork.WorkStep) {
//	subSteps, err := iwork.GetAllWorkStepByWorkName(workSubName)
//	if err != nil {
//		return
//	}
//	for _, subStep := range subSteps {
//		if strings.ToUpper(subStep.WorkStepType) == "WORK_START" {
//			adjustWorkSubNodeParamSchemaByUsingWorkStart(subStep, paramInputSchema, &step)
//		}
//		if strings.ToUpper(subStep.WorkStepType) == "WORK_END" {
//			adjustWorkSubNodeParamSchemaByUsingWorkEnd(subStep, &step)
//		}
//	}
//	iwork.InsertOrUpdateWorkStep(&step)
//}

//func adjustWorkSubNodeParamSchemaByUsingWorkEnd(subStep iwork.WorkStep, step *iwork.WorkStep) {
//	// 拷贝子节点输出
//	outputSchema := schema.GetCacheParamOutputSchema(&subStep)
//	step.WorkStepOutput = outputSchema.RenderToXml()
//}
//
//func adjustWorkSubNodeParamSchemaByUsingWorkStart(subStep iwork.WorkStep, paramInputSchema schema.ParamInputSchema, step *iwork.WorkStep) {
//	// 获取子流程 start 节点所有输出参数名
//	inputSchema := schema.GetCacheParamInputSchema(&subStep, &iworknode.WorkStepFactory{WorkStep: &subStep})
//	inputParamNames := GetInputSchemaNameArr(inputSchema)
//	items := []schema.ParamInputSchemaItem{}
//	for _, name := range inputParamNames {
//		// 存在则不添加且沿用旧值
//		if exist, paramValue := CheckAndGetParamValueByInputSchemaParamName(paramInputSchema.ParamInputSchemaItems, name); exist {
//			ModifyParamValue(paramInputSchema.ParamInputSchemaItems, name, paramValue)
//		} else {
//			items = append(items, schema.ParamInputSchemaItem{ParamName: name})
//		}
//	}
//	paramInputSchema.ParamInputSchemaItems = append(paramInputSchema.ParamInputSchemaItems, items...)
//	step.WorkStepInput = paramInputSchema.RenderToXml()
//}

func GetInputSchemaNameArr(schema *schema.ParamInputSchema) []string {
	paramNameArr := []string{}
	for _, item := range schema.ParamInputSchemaItems {
		paramNameArr = append(paramNameArr, item.ParamName)
	}
	return paramNameArr
}

func CheckAndGetParamValueByInputSchemaParamName(items []schema.ParamInputSchemaItem, paramName string) (exist bool, paramValue string) {
	for _, item := range items {
		if item.ParamName == paramName {
			return true, item.ParamValue
		}
	}
	return false, ""
}

func ModifyParamValue(items []schema.ParamInputSchemaItem, paramName string, paramValue string) {
	for _, item := range items {
		if item.ParamName == paramName {
			item.ParamValue = paramValue
		}
	}
}
