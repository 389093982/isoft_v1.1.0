package iwork

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"isoft/isoft/common/pageutil"
	"isoft/isoft_iaas_web/models/iwork"
	"time"
)

type WorkController struct {
	beego.Controller
}

func (this *WorkController) AddWork()  {
	var work iwork.Work
	work.WorkName = this.GetString("work_name")
	work.CreatedBy = "SYSTEM"
	work.CreatedTime = time.Now()
	work.LastUpdatedBy = "SYSTEM"
	work.LastUpdatedTime = time.Now()
	if _, err := iwork.InsertOrUpdateWork(&work);err == nil{
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	}else{
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func (this *WorkController) FilterPageWork()  {
	condArr := make(map[string]string)
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页
	if search := this.GetString("search");search != "" {
		condArr["search"] = search
	}
	works, count, err := iwork.QueryWork(condArr, current_page, offset)
	paginator := pagination.SetPaginator(this.Ctx, offset, count)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status":"SUCCESS", "works": works,
			"paginator": pageutil.Paginator(paginator.Page(), paginator.PerPageNums, paginator.Nums())}
	}else{
		this.Data["json"] = &map[string]interface{}{"status":"ERROR"}
	}
	this.ServeJSON()
}

func (this *WorkController) DeleteWorkById()  {
	id,_ := this.GetInt64("id")
	if err := iwork.DeleteWorkById(id); err == nil{
		this.Data["json"] = &map[string]interface{}{"status":"SUCCESS"}
	}else{
		this.Data["json"] = &map[string]interface{}{"status":"ERROR"}
	}
	this.ServeJSON()
}

func (this *WorkController) AddWorkStep()  {
	work_id := this.GetString("work_id")
	step := &iwork.WorkStep{
		WorkId:work_id,
		WorkStepId:iwork.GetNextWorkStepId(work_id),
		CreatedBy:"SYSTEM",
		CreatedTime:time.Now(),
		LastUpdatedBy:"SYSTEM",
		LastUpdatedTime:time.Now(),
	}
	if _, err := iwork.InsertOrUpdateWorkStep(step);err == nil{
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	}else{
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}


func (this *WorkController) EditWorkStep()  {
	work_id := this.GetString("work_id")
	work_step_id,_ := this.GetInt64("work_step_id", -1)
	this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	if step, err := iwork.GetOneWorkStep(work_id, work_step_id); err == nil{
		step.WorkStepInput = this.GetString("work_step_input")
		step.WorkStepName = this.GetString("work_step_name")
		step.WorkStepType = this.GetString("work_step_type")
		step.WorkStepOutput = this.GetString("work_step_output")
		step.CreatedBy = "SYSTEM"
		step.CreatedTime = time.Now()
		step.LastUpdatedBy = "SYSTEM"
		step.LastUpdatedTime = time.Now()
		if _, err := iwork.InsertOrUpdateWorkStep(&step);err == nil{
			this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
		}
	}
	this.ServeJSON()
}

func (this *WorkController) FilterPageWorkStep()  {
	condArr := make(map[string]string)
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页
	condArr["work_id"] = this.GetString("work_id")
	worksteps, count, err := iwork.QueryWorkStep(condArr, current_page, offset)
	paginator := pagination.SetPaginator(this.Ctx, offset, count)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status":"SUCCESS", "worksteps": worksteps,
			"paginator": pageutil.Paginator(paginator.Page(), paginator.PerPageNums, paginator.Nums())}
	}else{
		this.Data["json"] = &map[string]interface{}{"status":"ERROR"}
	}
	this.ServeJSON()
}

func (this *WorkController) DeleteWorkStepById()  {
	id,_ := this.GetInt64("id")
	if err := iwork.DeleteWorkStepById(id); err == nil{
		this.Data["json"] = &map[string]interface{}{"status":"SUCCESS"}
	}else{
		this.Data["json"] = &map[string]interface{}{"status":"ERROR"}
	}
	this.ServeJSON()
}

func (this *WorkController) LoadWorkStepInfo()  {
	work_id := this.GetString("work_id")
	work_step_id,_ := this.GetInt8("work_step_id")
	if step, err := iwork.LoadWorkStepInfo(work_id, work_step_id); err == nil{
		this.Data["json"] = &map[string]interface{}{"status":"SUCCESS", "step":step}
	}else{
		this.Data["json"] = &map[string]interface{}{"status":"ERROR"}
	}
	this.ServeJSON()
}

func (this *WorkController) GetAllWorkStepInfo() {
	work_id := this.GetString("work_id")
	if steps, err := iwork.GetAllWorkStepInfo(work_id); err == nil{
		this.Data["json"] = &map[string]interface{}{"status":"SUCCESS", "steps":steps}
	}else{
		this.Data["json"] = &map[string]interface{}{"status":"ERROR"}
	}
	this.ServeJSON()
}

func (this *WorkController) ChangeWorkStepOrder()  {
	this.Data["json"] = &map[string]interface{}{"status":"ERROR"}
	work_id := this.GetString("work_id")
	work_step_id,_ := this.GetInt64("work_step_id")
	_type := this.GetString("type")
	// 获取当前步骤
	step, _ := iwork.GetOneWorkStep(work_id, work_step_id)
	if _type == "up"{
		if prevStep, err := iwork.GetOneWorkStep(work_id, work_step_id - 1); err == nil{
			prevStep.WorkStepId = prevStep.WorkStepId + 1
			step.WorkStepId = step.WorkStepId - 1
			iwork.InsertOrUpdateWorkStep(&prevStep)
			iwork.InsertOrUpdateWorkStep(&step)
			this.Data["json"] = &map[string]interface{}{"status":"SUCCESS"}
		}
	}else{
		if nextStep, err := iwork.GetOneWorkStep(work_id, work_step_id + 1); err == nil{
			nextStep.WorkStepId = nextStep.WorkStepId + 1
			step.WorkStepId = step.WorkStepId + 1
			iwork.InsertOrUpdateWorkStep(&nextStep)
			iwork.InsertOrUpdateWorkStep(&step)
			this.Data["json"] = &map[string]interface{}{"status":"SUCCESS"}
		}
	}
	this.ServeJSON()
}

