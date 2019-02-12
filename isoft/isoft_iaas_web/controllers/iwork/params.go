package iwork

import (
	"encoding/json"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworknode"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
	"time"
)

func BuildOutput(work_id string, work_step_id int64)  {
	// 读取 work_step 信息
	step, err := iwork.LoadWorkStepInfo(work_id, work_step_id)
	if err != nil {
		panic(err)
	}
	step.WorkStepOutput = schema.GetRuntimeParamOutputSchema(&iworknode.WorkStepFactory{WorkStep:&step}).RenderToXml()
	if _, err = iwork.InsertOrUpdateWorkStep(&step); err != nil{
		panic(err)
	}
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

			// 如果是 start 或者 end 类型的节点,则通知其做一些事后适配
			NoticeWorkStartEndAdjust(work_id, work_step_id)
			// 如果是 work_sub 类型的节点,则通知其做一些事后适配
			NoticeWorkSubAdjust(work_id, work_step_id)
			// 所有操作完成后自动构建输出
			BuildOutput(work_id, work_step_id)
		}
	}
	this.ServeJSON()
}

// paramMappings 只有起始和结束节点才有,而且起始和结束节点的 paramMappings 也是 paramInput 和 paramOutput
func NoticeWorkStartEndAdjust(work_id string, work_step_id int64) {
	// 读取 step 记录
	if step, err := iwork.GetOneWorkStep(work_id, work_step_id); err == nil &&
			(step.WorkStepType == "work_start" || step.WorkStepType == "work_end"){
		adjustWorkStartEndNodeParamSchema(step.WorkStepParamMapping, &step)
	}
}

func NoticeWorkSubAdjust(work_id string, work_step_id int64) {
	// 读取 step 记录
	if step, err := iwork.GetOneWorkStep(work_id, work_step_id); err == nil && step.WorkStepType == "work_sub"{
		// 从 db 中读取 paramInputSchema
		paramInputSchema := schema.GetCacheParamInputSchema(&step, &iworknode.WorkStepFactory{WorkStep:&step})
		for _, item := range paramInputSchema.ParamInputSchemaItems {
			if item.ParamName == "work_sub" && strings.HasPrefix(item.ParamValue, "$WORK.") {
				// 找到 work_sub 字段值
				workSubName := getWorkSubNameFromParamValue(item.ParamValue)
				adjustWorkSubNodeParamSchema(workSubName, *paramInputSchema, step)
			}
		}
	}
}

func adjustWorkStartEndNodeParamSchema(paramMappingsStr string, step *iwork.WorkStep) {
	paramInputSchema := schema.GetCacheParamInputSchema(step, &iworknode.WorkStepFactory{WorkStep:step})
	var paramMappingsArr []string
	json.Unmarshal([]byte(paramMappingsStr), &paramMappingsArr)
	// 沿用旧值,添加新值,去除无效的值,即以 paramMapping 为准
	items := []schema.ParamInputSchemaItem{}
	for _, paramMapping := range paramMappingsArr {
		var oldValue string // 旧值默认为空
		for _, _item := range paramInputSchema.ParamInputSchemaItems {
			if _item.ParamName == paramMapping {
				oldValue = _item.ParamValue
			}
		}
		items = append(items, schema.ParamInputSchemaItem{ParamName: paramMapping, ParamValue: oldValue})
	}
	paramInputSchema.ParamInputSchemaItems = items
	step.WorkStepInput = paramInputSchema.RenderToXml()
	iwork.InsertOrUpdateWorkStep(step)
}

// 调整 work_sub 类型节点参数
func adjustWorkSubNodeParamSchema(workSubName string, paramInputSchema schema.ParamInputSchema, step iwork.WorkStep) {
	subSteps, err := iwork.GetAllWorkStepByWorkName(workSubName)
	if err != nil{
		return
	}
	for _, subStep := range subSteps {
		if strings.ToUpper(subStep.WorkStepType) == "WORK_START" {
			adjustWorkSubNodeParamSchemaByUsingWorkStart(subStep, paramInputSchema, &step)
		}
		if strings.ToUpper(subStep.WorkStepType) == "WORK_END" {
			adjustWorkSubNodeParamSchemaByUsingWorkEnd(subStep, &step)
		}
	}
	iwork.InsertOrUpdateWorkStep(&step)
}

func adjustWorkSubNodeParamSchemaByUsingWorkEnd(subStep iwork.WorkStep, step *iwork.WorkStep) {
	// 拷贝子节点输出
	outputSchema := schema.GetCacheParamOutputSchema(&subStep)
	step.WorkStepOutput = outputSchema.RenderToXml()
}

func adjustWorkSubNodeParamSchemaByUsingWorkStart(subStep iwork.WorkStep, paramInputSchema schema.ParamInputSchema, step *iwork.WorkStep) {
	// 获取子流程 start 节点所有输出参数名
	inputSchema := schema.GetCacheParamInputSchema(&subStep, &iworknode.WorkStepFactory{WorkStep: &subStep})
	inputParamNames := GetInputSchemaNameArr(inputSchema)
	items := []schema.ParamInputSchemaItem{}
	for _, name := range inputParamNames {
		// 存在则不添加且沿用旧值
		if exist, paramValue := CheckAndGetParamValueByInputSchemaParamName(paramInputSchema.ParamInputSchemaItems, name); exist {
			ModifyParamValue(paramInputSchema.ParamInputSchemaItems, name, paramValue)
		} else {
			items = append(items, schema.ParamInputSchemaItem{ParamName: name})
		}
	}
	paramInputSchema.ParamInputSchemaItems = append(paramInputSchema.ParamInputSchemaItems, items...)
	step.WorkStepInput = paramInputSchema.RenderToXml()
}

func getWorkSubNameFromParamValue(paramValue string) string {
	value := strings.TrimSpace(paramValue)
	value = strings.Replace(value, "$WORK.", "", -1)
	value = strings.Replace(value, "__sep__", "", -1)
	value = strings.Replace(value, "\n", "", -1)
	value = strings.TrimSpace(value)
	return value
}

func GetInputSchemaNameArr(schema *schema.ParamInputSchema) []string {
	paramNameArr := []string{}
	for _,item := range schema.ParamInputSchemaItems{
		paramNameArr = append(paramNameArr, item.ParamName)
	}
	return paramNameArr
}

func GetOutputSchemaNameArr(schema *schema.ParamOutputSchema) []string {
	paramNameArr := []string{}
	for _,item := range schema.ParamOutputSchemaItems{
		paramNameArr = append(paramNameArr, item.ParamName)
	}
	return paramNameArr
}

func CheckAndGetParamValueByInputSchemaParamName(items []schema.ParamInputSchemaItem, paramName string) (exist bool, paramValue string) {
	for _,item := range items{
		if item.ParamName == paramName{
			return true, item.ParamValue
		}
	}
	return false, ""
}

func CheckAndGetParamValueByOutputSchemaParamName(items []schema.ParamOutputSchemaItem, paramName string) (exist bool, paramValue string) {
	for _,item := range items{
		if item.ParamName == paramName{
			return true, item.ParamValue
		}
	}
	return false, ""
}

func ModifyParamValue(items []schema.ParamInputSchemaItem, paramName string, paramValue string)  {
	for _, item := range items {
		if item.ParamName == paramName{
			item.ParamValue = paramValue
		}
	}
}
