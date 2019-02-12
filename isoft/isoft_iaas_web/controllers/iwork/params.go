package iwork

import (
	"encoding/json"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworknode"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
	"time"
)

func (this *WorkController) EditWorkStepParamInfo() {
	work_id := this.GetString("work_id")
	work_step_id, _ := this.GetInt64("work_step_id", -1)
	paramInputSchemaStr := this.GetString("paramInputSchemaStr")
	paramMappingsStr := this.GetString("paramMappingsStr")
	var paramInputSchema schema.ParamInputSchema
	json.Unmarshal([]byte(paramInputSchemaStr), &paramInputSchema)
	this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	if step, err := iwork.GetOneWorkStep(work_id, work_step_id); err == nil {
		// paramMappings 只有起始和结束节点才有,而且起始和结束节点的 paramMappings 也是 paramInput 和 paramOutput
		if strings.TrimSpace(paramMappingsStr) != "" && (step.WorkStepType == "work_start" || step.WorkStepType == "work_end"){
			this.adjustWorkStartEndNodeParamSchema(paramMappingsStr, paramInputSchema)
		}
		if strings.TrimSpace(paramMappingsStr) != "" && step.WorkStepType == "work_sub" {
			for _, item := range paramInputSchema.ParamInputSchemaItems {
				if item.ParamName == "work_sub" && strings.HasPrefix(item.ParamValue, "$WORK.") {
					this.adjustWorkSubNodeParamSchema(item, paramInputSchema, step)
				}
			}
		}

		step.WorkStepInput = paramInputSchema.RenderToXml()
		step.WorkStepParamMapping = paramMappingsStr
		step.CreatedBy = "SYSTEM"
		step.CreatedTime = time.Now()
		step.LastUpdatedBy = "SYSTEM"
		step.LastUpdatedTime = time.Now()
		if _, err := iwork.InsertOrUpdateWorkStep(&step); err == nil {
			this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
		}
	}
	this.ServeJSON()
}

func (this *WorkController) adjustWorkStartEndNodeParamSchema(paramMappingsStr string, paramInputSchema schema.ParamInputSchema) {
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
}

// 调整 work_sub 类型节点参数
func (this *WorkController) adjustWorkSubNodeParamSchema(item schema.ParamInputSchemaItem, paramInputSchema schema.ParamInputSchema, step iwork.WorkStep) {
	value := strings.TrimSpace(item.ParamValue)
	value = strings.Replace(value, "$WORK.", "", -1)
	value = strings.Replace(value, "__sep__", "", -1)
	value = strings.Replace(value, "\n", "", -1)
	value = strings.TrimSpace(value)
	steps, err := iwork.GetAllWorkStepByWorkName(value)
	if err == nil {
		for _, _step := range steps {
			if strings.ToUpper(_step.WorkStepType) == "WORK_START" {
				inputSchema := schema.GetCacheParamInputSchema(&_step, &iworknode.WorkStepFactory{WorkStep: &_step})
				inputParamNames := GetInputSchemaNameArr(inputSchema)
				items := []schema.ParamInputSchemaItem{}
				for _, name := range inputParamNames {
					// 存在则不添加且沿用旧值
					if exist, paramValue := CheckAndGetParamValueByInputSchemaParamName(paramInputSchema.ParamInputSchemaItems, name); exist {
						ModifyParamValue(items, name, paramValue)
					} else {
						items = append(items, schema.ParamInputSchemaItem{ParamName: name})
					}
				}
				paramInputSchema.ParamInputSchemaItems = append(paramInputSchema.ParamInputSchemaItems, items...)
			}
			if strings.ToUpper(_step.WorkStepType) == "WORK_END" {
				outputSchema := schema.GetCacheParamOutputSchema(&_step)
				outputParamNames := GetOutputSchemaNameArr(outputSchema)
				items := []schema.ParamOutputSchemaItem{}
				for _, name := range outputParamNames {
					items = append(items, schema.ParamOutputSchemaItem{ParamName: name})
				}
				schema := schema.ParamOutputSchema{ParamOutputSchemaItems: items}
				step.WorkStepOutput = schema.RenderToXml()
			}
		}
	}
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
