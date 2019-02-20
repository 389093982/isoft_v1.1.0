package iwork

import (
	"encoding/json"
	"fmt"
	"isoft/isoft_iaas_web/core/iworkconst"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworknode"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
	"time"
)

// 构建动态输入值
func BuildDynamicInput(work_id int64, work_step_id int64) {
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
	paramInputSchema := &schema.ParamInputSchema{ParamInputSchemaItems: newInputSchemaItems}
	step.WorkStepInput = paramInputSchema.RenderToXml()
	if _, err = iwork.InsertOrUpdateWorkStep(&step); err != nil {
		panic(err)
	}
}

// 构建动态输出值
func BuildDynamicOutput(work_id int64, work_step_id int64) {
	// 读取 work_step 信息
	step, err := iwork.LoadWorkStepInfo(work_id, work_step_id)
	if err != nil {
		panic(err)
	}
	runtimeParamOutputSchema := schema.GetRuntimeParamOutputSchema(&iworknode.WorkStepFactory{WorkStep: &step})
	defaultParamOutputSchema := schema.GetDefaultParamOutputSchema(&iworknode.WorkStepFactory{WorkStep: &step})
	defaultParamOutputSchema.ParamOutputSchemaItems = append(defaultParamOutputSchema.ParamOutputSchemaItems, runtimeParamOutputSchema.ParamOutputSchemaItems...)
	// 构建输出参数,使用全新值
	step.WorkStepOutput = defaultParamOutputSchema.RenderToXml()
	if _, err = iwork.InsertOrUpdateWorkStep(&step); err != nil {
		panic(err)
	}
}

func checkAndCreateSubWork(work_name string) {
	if _, err := iwork.QueryWorkByName(work_name); err != nil {
		// 不存在 work 则直接创建
		work := &iwork.Work{
			WorkName:        work_name,
			WorkDesc:        fmt.Sprintf("自动创建子流程:%s", work_name),
			CreatedBy:       "SYSTEM",
			CreatedTime:     time.Now(),
			LastUpdatedBy:   "SYSTEM",
			LastUpdatedTime: time.Now(),
		}
		if _, err := iwork.InsertOrUpdateWork(work); err == nil {
			// 写入 DB 并自动创建开始和结束节点
			insertStartEndWorkStepNode(work.Id)
		}
	}
}

func BuildAutoCreateSubWork(work_id int64, work_step_id int64) {
	// 读取 work_step 信息
	step, err := iwork.LoadWorkStepInfo(work_id, work_step_id)
	if err != nil {
		panic(err)
	}
	if step.WorkStepType != "work_sub" {
		return
	}
	paramInputSchema := schema.GetCacheParamInputSchema(&step, &iworknode.WorkStepFactory{WorkStep: &step})
	for index, item := range paramInputSchema.ParamInputSchemaItems {
		if item.ParamName == iworkconst.STRING_PREFIX + "work_sub" {
			paramValue := strings.TrimSpace(item.ParamValue)
			if !strings.HasPrefix(paramValue, "$WORK.") {
				// 修改值并同步到数据库
				paramInputSchema.ParamInputSchemaItems[index] = schema.ParamInputSchemaItem{
					ParamName:  item.ParamName,
					ParamValue: strings.Join([]string{"$WORK.", paramValue}, ""),
				}
				step.WorkStepInput = paramInputSchema.RenderToXml()
				// 自动创建子流程
				checkAndCreateSubWork(paramValue)
			}
			// 维护 work 的 WorkSubId 属性
			subWork,_ := iwork.QueryWorkByName(strings.Replace(paramValue, "$WORK.", "", -1))
			step.WorkSubId = subWork.Id
			break
		}
	}
	iwork.InsertOrUpdateWorkStep(&step)
}

// 构建动态值
func BuildDynamic(work_id int64, work_step_id int64) {
	// 自动创建子流程
	BuildAutoCreateSubWork(work_id, work_step_id)
	// 构建动态输入值
	BuildDynamicInput(work_id, work_step_id)
	// 构建动态输出值
	BuildDynamicOutput(work_id, work_step_id)
}

func (this *WorkController) EditWorkStepParamInfo() {
	work_id,_ := this.GetInt64("work_id")
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
