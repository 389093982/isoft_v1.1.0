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

func (this *WorkController) AddWorkStep()  {
	var step iwork.WorkStep
	step.WorkId = this.GetString("work_id")
	step.WorkStepId,_ = this.GetInt8("work_step_id", -1)
	step.WorkStepInput = this.GetString("work_step_input")
	step.WorkStepOutput = this.GetString("work_step_output")
	step.CreatedBy = "SYSTEM"
	step.CreatedTime = time.Now()
	step.LastUpdatedBy = "SYSTEM"
	step.LastUpdatedTime = time.Now()
	if _, err := iwork.InsertOrUpdateWorkStep(&step);err == nil{
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	}else{
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func (this *WorkController) FilterPageWorkStep()  {
	condArr := make(map[string]string)
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页
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
