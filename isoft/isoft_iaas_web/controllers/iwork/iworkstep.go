package iwork

import (
	"encoding/json"
	"isoft/isoft/common/stringutil"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworknode"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
	"time"
)

func (this *WorkController) AddWorkStep() {
	work_id,_ := this.GetInt64("work_id")
	work_step_id,_ := this.GetInt64("work_step_id")
	this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	// 将 work_step_id 之后的所有节点后移一位
	err := iwork.BatchChangeWorkStepIdOrder(work_id, work_step_id, "+")
	if err == nil{
		work_step_type := this.GetString("default_work_step_type")
		step := &iwork.WorkStep{
			WorkId:          work_id,
			WorkStepName:    "random_" + stringutil.RandomUUID(),
			WorkStepType:    work_step_type,
			WorkStepId:      work_step_id + 1,
			CreatedBy:       "SYSTEM",
			CreatedTime:     time.Now(),
			LastUpdatedBy:   "SYSTEM",
			LastUpdatedTime: time.Now(),
		}
		if _, err := iwork.InsertOrUpdateWorkStep(step); err == nil {
			this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
		}
	}
	this.ServeJSON()
}

func (this *WorkController) EditWorkStepColorInfo()  {
	work_id,_ := this.GetInt64("work_id")
	work_step_id, _ := this.GetInt64("work_step_id", -1)
	work_step_color := this.GetString("work_step_color")
	this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	if step, err := iwork.GetOneWorkStep(work_id, work_step_id); err == nil {
		step.WorkStepColor = work_step_color
		if _, err := iwork.InsertOrUpdateWorkStep(&step); err == nil {
			this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
		}
	}
	this.ServeJSON()
}

func (this *WorkController) EditWorkStepBaseInfo() {
	defer func() {
		if err := recover(); err != nil{
			this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
			this.ServeJSON()
		}
	}()

	work_id,_ := this.GetInt64("work_id")
	work_step_id, _ := this.GetInt64("work_step_id", -1)
	work_step_name := this.GetString("work_step_name")
	work_step_desc := this.GetString("work_step_desc")
	work_step_type := this.GetString("work_step_type")
	step, err := iwork.GetOneWorkStep(work_id, work_step_id)
	if err != nil{
		panic(err)
	}
	oldWorkStepName := step.WorkStepName
	step.WorkStepName = work_step_name
	step.WorkStepDesc = work_step_desc
	// 变更类型需要置空 input 和 output 参数
	if step.WorkStepType != work_step_type {
		step.WorkStepType = this.GetString("work_step_type")
		step.WorkStepInput = ""
		step.WorkStepOutput = ""
	}
	step.CreatedBy = "SYSTEM"
	step.CreatedTime = time.Now()
	step.LastUpdatedBy = "SYSTEM"
	step.LastUpdatedTime = time.Now()
	if _, err := iwork.InsertOrUpdateWorkStep(&step); err == nil {
		// 级联更改相关联的步骤名称
		changeReferencesWorkStepName(work_id, oldWorkStepName, work_step_name)
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	}
	this.ServeJSON()
}

func changeReferencesWorkStepName(work_id int64, oldWorkStepName, workStepName string) error {
	if oldWorkStepName == workStepName{
		return nil
	}
	steps, err := iwork.GetAllWorkStepInfo(work_id)
	if err != nil{
		return err
	}
	for _, step := range steps{
		step.WorkStepInput = strings.Replace(step.WorkStepInput, "$" + oldWorkStepName, "$" + workStepName, -1)
		_, err := iwork.InsertOrUpdateWorkStep(&step)
		if err != nil{
			return err
		}
	}
	return nil
}

func (this *WorkController) FilterWorkStep() {
	condArr := make(map[string]interface{})
	condArr["work_id"],_ = this.GetInt64("work_id")
	worksteps, err := iwork.QueryWorkStep(condArr)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "worksteps": worksteps}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func (this *WorkController) DeleteWorkStepByWorkStepId() {
	work_id, _ := this.GetInt64("work_id")
	work_step_id, _ := this.GetInt64("work_step_id")
	if err := iwork.DeleteWorkStepByWorkStepId(work_id, work_step_id); err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func (this *WorkController) LoadWorkStepInfo() {
	work_id,_ := this.GetInt64("work_id")
	work_step_id, _ := this.GetInt64("work_step_id")
	// 读取 work_step 信息
	if step, err := iwork.LoadWorkStepInfo(work_id, work_step_id); err == nil {
		var paramMappingsArr []string
		json.Unmarshal([]byte(step.WorkStepParamMapping), &paramMappingsArr)
		// 返回结果
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "step": step,
			"paramInputSchema":          schema.GetCacheParamInputSchema(&step, &iworknode.WorkStepFactory{WorkStep: &step}),
			"paramOutputSchema":         schema.GetCacheParamOutputSchema(&step),
			"paramOutputSchemaTreeNode": schema.GetCacheParamOutputSchema(&step).RenderToTreeNodes("output"),
			"paramMappings":             paramMappingsArr,
		}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func (this *WorkController) GetAllWorkStepInfo() {
	work_id,_ := this.GetInt64("work_id")
	if steps, err := iwork.GetAllWorkStepInfo(work_id); err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "steps": steps}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func (this *WorkController) ChangeWorkStepOrder() {
	this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	work_id,_ := this.GetInt64("work_id")
	work_step_id, _ := this.GetInt64("work_step_id")
	_type := this.GetString("type")
	// 获取当前步骤
	step, _ := iwork.GetOneWorkStep(work_id, work_step_id)
	if _type == "up" {
		if prevStep, err := iwork.GetOneWorkStep(work_id, work_step_id-1); err == nil {
			prevStep.WorkStepId = prevStep.WorkStepId + 1
			step.WorkStepId = step.WorkStepId - 1
			iwork.InsertOrUpdateWorkStep(&prevStep)
			iwork.InsertOrUpdateWorkStep(&step)
			this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
		}
	} else {
		if nextStep, err := iwork.GetOneWorkStep(work_id, work_step_id+1); err == nil {
			nextStep.WorkStepId = nextStep.WorkStepId + 1
			step.WorkStepId = step.WorkStepId + 1
			iwork.InsertOrUpdateWorkStep(&nextStep)
			iwork.InsertOrUpdateWorkStep(&step)
			this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
		}
	}
	this.ServeJSON()
}


func (this *WorkController) LoadPreNodeOutput() {
	work_id,_ := this.GetInt64("work_id")
	work_step_id, _ := this.GetInt64("work_step_id")

	preParamOutputSchemaTreeNodeArr := []*schema.TreeNode{}
	// 加载 resource 参数
	pos := LoadResourceInfo()
	preParamOutputSchemaTreeNodeArr = append(preParamOutputSchemaTreeNodeArr, pos.RenderToTreeNodes("$RESOURCE"))
	// 加载 work 参数
	pos = LoadWorkInfo()
	preParamOutputSchemaTreeNodeArr = append(preParamOutputSchemaTreeNodeArr, pos.RenderToTreeNodes("$WORK"))
	// 加载 entity 参数
	pos = LoadEntityInfo()
	preParamOutputSchemaTreeNodeArr = append(preParamOutputSchemaTreeNodeArr, pos.RenderToTreeNodes("$Entity"))
	// 加载前置步骤输出
	if steps, err := iwork.GetAllPreStepInfo(work_id, work_step_id); err == nil {
		for _, step := range steps {
			pos := schema.GetCacheParamOutputSchema(&step)
			preParamOutputSchemaTreeNodeArr = append(preParamOutputSchemaTreeNodeArr, pos.RenderToTreeNodes("$"+step.WorkStepName))
		}
	}
	// 返回结果
	this.Data["json"] = &map[string]interface{}{"status": "SUCCESS",
		"preParamOutputSchemaTreeNodeArr": preParamOutputSchemaTreeNodeArr,
	}
	this.ServeJSON()
}

func LoadResourceInfo() *schema.ParamOutputSchema {
	pos := &schema.ParamOutputSchema{
		ParamOutputSchemaItems: []schema.ParamOutputSchemaItem{},
	}
	resources := iwork.GetAllResource()
	for _, resource := range resources {
		pos.ParamOutputSchemaItems = append(pos.ParamOutputSchemaItems, schema.ParamOutputSchemaItem{
			ParamName: resource.ResourceName,
		})
	}
	return pos
}

func LoadWorkInfo() *schema.ParamOutputSchema {
	pos := &schema.ParamOutputSchema{
		ParamOutputSchemaItems: []schema.ParamOutputSchemaItem{},
	}
	works := iwork.GetAllWorkInfo()
	for _, work := range works {
		pos.ParamOutputSchemaItems = append(pos.ParamOutputSchemaItems, schema.ParamOutputSchemaItem{
			ParamName: work.WorkName,
		})
	}
	return pos
}

func LoadEntityInfo() *schema.ParamOutputSchema {
	pos := &schema.ParamOutputSchema{
		ParamOutputSchemaItems: []schema.ParamOutputSchemaItem{},
	}
	entities := iwork.GetAllEntityInfo()
	for _, entity := range entities {
		pos.ParamOutputSchemaItems = append(pos.ParamOutputSchemaItems, schema.ParamOutputSchemaItem{
			ParamName: entity.EntityName,
		})
	}
	return pos
}
