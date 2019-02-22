package iwork

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"isoft/isoft/common/pageutil"
	"isoft/isoft/common/stringutil"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworknode"
	"isoft/isoft_iaas_web/core/iworkrun"
	"isoft/isoft_iaas_web/models/iwork"
	"time"
)

type WorkController struct {
	beego.Controller
}

func (this *WorkController) BuildIWorkDL()  {
	//dls := make([]*IWorkDL,0)
	//works := iwork.GetAllWorkInfo()
	//for _, work := range works{
	//	dl := &IWorkDL{}
	//	steps, _ := iwork.GetAllWorkStepInfo(work.Id)
	//	for _, step := range steps{
	//		if step.WorkStepType == "work_start"{
	//			dl.RequestInfo = step.WorkStepInput
	//		}
	//		if step.WorkStepType == "work_end"{
	//			dl.ResponseInfo = step.WorkStepOutput
	//		}
	//	}
	//	dls = append(dls, dl)
	//}
}

func (this *WorkController) DeleteEntity()  {
	entity_id, _ := this.GetInt64("entity_id", -1)
	if err := iwork.DeleteEntityById(entity_id); err == nil{
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	}else{
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func (this *WorkController) EditEntity()  {
	var entity iwork.Entity
	entity_id, err := this.GetInt64("entity_id", -1)
	if err == nil && entity_id > 0 {
		entity.Id = entity_id
	}
	entity.EntityName = this.GetString("entity_name")
	entity.EntityFieldStr = this.GetString("entity_field_str")
	entity.CreatedBy = "SYSTEM"
	entity.CreatedTime = time.Now()
	entity.LastUpdatedBy = "SYSTEM"
	entity.LastUpdatedTime = time.Now()
	if _, err := iwork.InsertOrUpdateEntity(&entity); err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func (this *WorkController) FilterPageEntity()  {
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页
	entities, count, err := iwork.QueryEntity(current_page, offset)
	paginator := pagination.SetPaginator(this.Ctx, offset, count)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "entities": entities,
			"paginator": pageutil.Paginator(paginator.Page(), paginator.PerPageNums, paginator.Nums())}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func (this *WorkController) GetRelativeWork() {
	work_id,_ := this.GetInt64("work_id")
	subWorks := make([]iwork.Work,0)
	parentWorks, _,_ := iwork.QueryParentWorks(work_id)
	steps, _ := iwork.GetAllWorkStepInfo(work_id)
	for _, step := range steps{
		if step.WorkSubId > 0{
			subwork, _ := iwork.QueryWorkById(step.WorkSubId)
			subWorks = append(subWorks, subwork)
		}
	}
	this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "parentWorks":parentWorks,"subwork": subWorks}
	this.ServeJSON()
}

func (this *WorkController) GetLastRunLogDetail() {
	tracking_id := this.GetString("tracking_id")
	runLogDetails, err := iwork.GetLastRunLogDetail(tracking_id)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "runLogDetails": runLogDetails}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func (this *WorkController) FilterPageLogRecord() {
	work_id,_ := this.GetInt64("work_id")
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页
	runLogRecords, count, err := iwork.QueryRunLogRecord(work_id, current_page, offset)
	paginator := pagination.SetPaginator(this.Ctx, offset, count)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "runLogRecords": runLogRecords,
			"paginator": pageutil.Paginator(paginator.Page(), paginator.PerPageNums, paginator.Nums())}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func (this *WorkController) RunWork() {
	work_id,_ := this.GetInt64("work_id")
	work, _ := iwork.QueryWorkById(work_id)
	steps, _ := iwork.GetAllWorkStepInfo(work_id)
	go iworkrun.Run(work, steps, nil)
	this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	this.ServeJSON()
}

func (this *WorkController) EditWork() {
	var work iwork.Work
	work_id, err := this.GetInt64("work_id", -1);
	if err == nil && work_id > 0 {
		work.Id = work_id
	}
	work.WorkName = this.GetString("work_name")
	work.WorkDesc = this.GetString("work_desc")
	work.CreatedBy = "SYSTEM"
	work.CreatedTime = time.Now()
	work.LastUpdatedBy = "SYSTEM"
	work.LastUpdatedTime = time.Now()
	if _, err := iwork.InsertOrUpdateWork(&work); err == nil {
		if work_id <= 0 {
			// 新增 work 场景,自动添加开始和结束节点
			insertStartEndWorkStepNode(work.Id)
		}
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func insertStartEndWorkStepNode(work_id int64) {
	insertDefaultWorkStepNodeFunc := func(nodeName string) {
		step := &iwork.WorkStep{
			WorkId:          work_id,
			WorkStepId:      iwork.GetNextWorkStepId(work_id),
			WorkStepName:    nodeName,
			WorkStepDesc:    fmt.Sprintf("%s节点", nodeName),
			WorkStepType:    fmt.Sprintf("work_%s", nodeName),
			CreatedBy:       "SYSTEM",
			CreatedTime:     time.Now(),
			LastUpdatedBy:   "SYSTEM",
			LastUpdatedTime: time.Now(),
		}
		iwork.InsertOrUpdateWorkStep(step)
	}
	insertDefaultWorkStepNodeFunc("start")
	insertDefaultWorkStepNodeFunc("end")
}

func (this *WorkController) FilterPageWork() {
	condArr := make(map[string]string)
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页
	if search := this.GetString("search"); search != "" {
		condArr["search"] = search
	}
	works, count, err := iwork.QueryWork(condArr, current_page, offset)
	paginator := pagination.SetPaginator(this.Ctx, offset, count)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "works": works,
			"paginator": pageutil.Paginator(paginator.Page(), paginator.PerPageNums, paginator.Nums())}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func (this *WorkController) DeleteWorkById() {
	id, _ := this.GetInt64("id")
	if err := iwork.DeleteWorkById(id); err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func (this *WorkController) AddWorkStep() {
	work_id,_ := this.GetInt64("work_id")
	work_step_type := this.GetString("default_work_step_type")
	step := &iwork.WorkStep{
		WorkId:          work_id,
		WorkStepName:    "random_" + stringutil.RandomUUID(),
		WorkStepType:    work_step_type,
		WorkStepId:      iwork.GetNextWorkStepId(work_id),
		CreatedBy:       "SYSTEM",
		CreatedTime:     time.Now(),
		LastUpdatedBy:   "SYSTEM",
		LastUpdatedTime: time.Now(),
	}
	if _, err := iwork.InsertOrUpdateWorkStep(step); err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
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
	work_id,_ := this.GetInt64("work_id")
	work_step_id, _ := this.GetInt64("work_step_id", -1)
	work_step_name := this.GetString("work_step_name")
	work_step_desc := this.GetString("work_step_desc")
	work_step_type := this.GetString("work_step_type")
	this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	// 变更类型需要置空 input 和 output 参数
	if step, err := iwork.GetOneWorkStep(work_id, work_step_id); err == nil {
		step.WorkStepName = work_step_name
		step.WorkStepDesc = work_step_desc
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
			this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
		}
	}
	this.ServeJSON()
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
	work_step_id, _ := this.GetInt64("work_step_id")
	if err := iwork.DeleteWorkStepByWorkStepId(work_step_id); err == nil {
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

