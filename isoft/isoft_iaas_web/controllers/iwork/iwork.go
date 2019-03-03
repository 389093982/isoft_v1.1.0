package iwork

import (
	"github.com/astaxie/beego"
	"isoft/isoft_iaas_web/models/iwork"
	"isoft/isoft_iaas_web/service"
	"isoft/isoft_iaas_web/service/iworkservice"
	"time"
)

type WorkController struct {
	beego.Controller
}

func (this *WorkController) BuildIWorkDL() {
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

func (this *WorkController) GetRelativeWork() {
	work_id, _ := this.GetInt64("work_id")
	serviceArgs := map[string]interface{}{"work_id": work_id}
	if result, err := service.ExecuteResultServiceWithTx(serviceArgs, iworkservice.GetRelativeWorkService); err == nil {
		this.Data["json"] = &map[string]interface{}{
			"status":      "SUCCESS",
			"parentWorks": result["parentWorks"],
			"subworks":    result["subworks"],
		}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

func (this *WorkController) GetLastRunLogDetail() {
	tracking_id := this.GetString("tracking_id")
	runLogDetails, err := iwork.QueryLastRunLogDetail(tracking_id)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "runLogDetails": runLogDetails}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func (this *WorkController) FilterPageLogRecord() {
	work_id, _ := this.GetInt64("work_id")
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页
	serviceArgs := map[string]interface{}{"work_id": work_id, "offset": offset, "current_page": current_page, "ctx": this.Ctx}
	if result, err := service.ExecuteResultServiceWithTx(serviceArgs, iworkservice.FilterPageLogRecord); err == nil {
		this.Data["json"] = &map[string]interface{}{
			"status":        "SUCCESS",
			"runLogRecords": result["runLogRecords"],
			"paginator":     result["paginator"],
		}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

func (this *WorkController) RunWork() {
	work_id, _ := this.GetInt64("work_id")
	serviceArgs := map[string]interface{}{"work_id": work_id}
	if err := service.ExecuteServiceWithTx(serviceArgs, iworkservice.RunWorkService); err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func (this *WorkController) EditWork() {
	// 将请求参数封装成 work
	var work iwork.Work
	work_id, err := this.GetInt64("work_id", -1)
	if err == nil && work_id > 0 {
		work.Id = work_id
	}
	work.WorkName = this.GetString("work_name")
	work.WorkDesc = this.GetString("work_desc")
	work.CreatedBy = "SYSTEM"
	work.CreatedTime = time.Now()
	work.LastUpdatedBy = "SYSTEM"
	work.LastUpdatedTime = time.Now()
	serviceArgs := map[string]interface{}{"work": work}
	if err := service.ExecuteServiceWithTx(serviceArgs, iworkservice.EditWorkService); err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}

	this.ServeJSON()
}

func (this *WorkController) FilterPageWork() {
	condArr := make(map[string]string)
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页
	if search := this.GetString("search"); search != "" {
		condArr["search"] = search
	}
	serviceArgs := map[string]interface{}{"condArr": condArr, "offset": offset, "current_page": current_page, "ctx": this.Ctx}
	if result, err := service.ExecuteResultServiceWithTx(serviceArgs, iworkservice.FilterPageWorkService); err == nil {
		this.Data["json"] = &map[string]interface{}{
			"status":    "SUCCESS",
			"works":     result["works"],
			"paginator": result["paginator"],
		}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

func (this *WorkController) DeleteWorkById() {
	id, _ := this.GetInt64("id")
	serviceArgs := map[string]interface{}{"id": id}
	if err := service.ExecuteServiceWithTx(serviceArgs, iworkservice.DeleteWorkByIdService); err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}
