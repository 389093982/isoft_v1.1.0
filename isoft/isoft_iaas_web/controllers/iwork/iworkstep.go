package iwork

import (
	"isoft/isoft_iaas_web/models/iwork"
	"isoft/isoft_iaas_web/service"
	"isoft/isoft_iaas_web/service/iworkservice"
)

func (this *WorkController) AddWorkStep() {
	serviceArgs := make(map[string]interface{}, 0)
	serviceArgs["work_id"], _ = this.GetInt64("work_id")
	serviceArgs["work_step_id"], _ = this.GetInt64("work_step_id")
	serviceArgs["default_work_step_type"] = this.GetString("default_work_step_type")
	if err := service.ExecuteServiceWithTx(serviceArgs, iworkservice.AddWorkStepService); err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

func (this *WorkController) EditWorkStepBaseInfo() {
	step := new(iwork.WorkStep)
	work_id, _ := this.GetInt64("work_id", -1)
	work_step_id, _ := this.GetInt64("work_step_id", -1)
	step.WorkId = work_id
	step.WorkStepId = work_step_id
	step.WorkStepName = this.GetString("work_step_name")
	step.WorkStepType = this.GetString("work_step_type")
	step.WorkStepDesc = this.GetString("work_step_desc")
	serviceArgs := map[string]interface{}{"step": step}
	if err := service.ExecuteServiceWithTx(serviceArgs, iworkservice.EditWorkStepBaseInfoService); err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

func (this *WorkController) FilterWorkStep() {
	work_id, _ := this.GetInt64("work_id")
	serviceArgs := map[string]interface{}{"work_id": work_id}
	if result, err := service.ExecuteResultServiceWithTx(serviceArgs, iworkservice.FilterWorkStepService); err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "worksteps": result["worksteps"]}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

func (this *WorkController) DeleteWorkStepByWorkStepId() {
	work_id, _ := this.GetInt64("work_id")
	work_step_id, _ := this.GetInt64("work_step_id")
	serviceArgs := map[string]interface{}{"work_id": work_id, "work_step_id": work_step_id}
	if err := service.ExecuteServiceWithTx(serviceArgs, iworkservice.DeleteWorkStepByWorkStepIdService); err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

func (this *WorkController) LoadWorkStepInfo() {
	work_id, _ := this.GetInt64("work_id")
	work_step_id, _ := this.GetInt64("work_step_id")
	serviceArgs := map[string]interface{}{"work_id": work_id, "work_step_id": work_step_id}
	if result, err := service.ExecuteResultServiceWithTx(serviceArgs, iworkservice.LoadWorkStepInfoService); err == nil {
		this.Data["json"] = &map[string]interface{}{
			"status":                    "SUCCESS",
			"step":                      result["step"],
			"paramInputSchema":          result["paramInputSchema"],
			"paramOutputSchema":         result["paramOutputSchema"],
			"paramOutputSchemaTreeNode": result["paramOutputSchemaTreeNode"],
			"paramMappings":             result["paramMappings"],
		}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

func (this *WorkController) GetAllWorkStepInfo() {
	work_id, _ := this.GetInt64("work_id")
	serviceArgs := map[string]interface{}{"work_id": work_id}
	if result, err := service.ExecuteResultServiceWithTx(serviceArgs, iworkservice.GetAllWorkStepInfoService); err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "steps": result["steps"]}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

func (this *WorkController) ChangeWorkStepOrder() {
	work_id, _ := this.GetInt64("work_id")
	work_step_id, _ := this.GetInt64("work_step_id")
	_type := this.GetString("type")
	serviceArgs := map[string]interface{}{"work_id": work_id, "work_step_id": work_step_id, "_type": _type}
	if err := service.ExecuteServiceWithTx(serviceArgs, iworkservice.ChangeWorkStepOrderService); err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

func (this *WorkController) LoadPreNodeOutput() {
	work_id, _ := this.GetInt64("work_id")
	work_step_id, _ := this.GetInt64("work_step_id")
	serviceArgs := map[string]interface{}{"work_id": work_id, "work_step_id": work_step_id}
	if result, err := service.ExecuteResultServiceWithTx(serviceArgs, iworkservice.LoadPreNodeOutputService); err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "preParamOutputSchemaTreeNodeArr": result["preParamOutputSchemaTreeNodeArr"]}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

func (this *WorkController) RefactorWorkStepInfo() {
	serviceArgs := make(map[string]interface{}, 0)
	serviceArgs["work_id"], _ = this.GetInt64("work_id")
	serviceArgs["refactor_worksub_name"] = this.GetString("refactor_worksub_name")
	serviceArgs["refactor_work_step_ids"] = this.GetString("refactor_work_step_ids")
	if err := service.ExecuteServiceWithTx(serviceArgs, iworkservice.RefactorWorkStepInfoService); err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

func (this *WorkController) EditWorkStepParamInfo() {
	work_id, _ := this.GetInt64("work_id")
	work_step_id, _ := this.GetInt64("work_step_id", -1)
	paramInputSchemaStr := this.GetString("paramInputSchemaStr")
	paramMappingsStr := this.GetString("paramMappingsStr")
	serviceArgs := map[string]interface{}{
		"work_id":             work_id,
		"work_step_id":        work_step_id,
		"paramInputSchemaStr": paramInputSchemaStr,
		"paramMappingsStr":    paramMappingsStr,
	}
	if err := service.ExecuteServiceWithTx(serviceArgs, iworkservice.EditWorkStepParamInfoService); err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}
