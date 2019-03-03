package iwork

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"isoft/isoft/common/pageutil"
	"isoft/isoft_iaas_web/core/iworkrun"
	"isoft/isoft_iaas_web/models/iwork"
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
	subWorks := make([]iwork.Work, 0)
	parentWorks, _, _ := iwork.QueryParentWorks(work_id)
	steps, _ := iwork.QueryAllWorkStepInfo(work_id)
	for _, step := range steps {
		if step.WorkSubId > 0 {
			subwork, _ := iwork.QueryWorkById(step.WorkSubId)
			subWorks = append(subWorks, subwork)
		}
	}
	this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "parentWorks": parentWorks, "subwork": subWorks}
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
	work_id, _ := this.GetInt64("work_id")
	work, _ := iwork.QueryWorkById(work_id)
	steps, _ := iwork.QueryAllWorkStepInfo(work_id)
	go iworkrun.Run(work, steps, nil)
	this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
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

	if err := iworkservice.EditWorkService(work); err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
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
	if err := iworkservice.DeleteWorkByIdService(id); err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func (this *WorkController) RefactorWorkStepInfo() {
	this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	work_id, _ := this.GetInt64("work_id")
	refactor_worksub_name := this.GetString("refactor_worksub_name")
	refactor_work_step_ids := this.GetString("refactor_work_step_ids")
	var refactor_work_step_id_arr []int
	json.Unmarshal([]byte(refactor_work_step_ids), &refactor_work_step_id_arr)
	// 校验 refactor_work_step_id_arr 是否连续
	if refactor_work_step_id_arr[len(refactor_work_step_id_arr)-1]-refactor_work_step_id_arr[0] != len(refactor_work_step_id_arr)-1 {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": "refactor workStepId 必须是连续的!"}
	} else {
		// 创建子流程
		subWork := &iwork.Work{
			WorkName:        refactor_worksub_name,
			WorkDesc:        "",
			CreatedBy:       "SYSTEM",
			CreatedTime:     time.Now(),
			LastUpdatedBy:   "SYSTEM",
			LastUpdatedTime: time.Now(),
		}
		iwork.InsertOrUpdateWork(subWork)
		// 循环移动子步骤
		for index, work_step_id := range refactor_work_step_id_arr {
			step, err := iwork.QueryWorkStepInfo(work_id, int64(work_step_id))
			if err == nil {
				if step.WorkStepType == "work_start" || step.WorkStepType == "work_end" {
					this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
					break
				}
				iworkservice.InsertStartEndWorkStepNode(subWork.Id)
				var newStep *iwork.WorkStep
				deepCopy(newStep, step)
				newStep.WorkId = subWork.Id
				newStep.WorkStepId = int64(index + 2)
				newStep.Id = 0
				iwork.InsertOrUpdateWorkStep(newStep)

				// 当前流程循环删除该节点
				iworkservice.DeleteWorkStepByWorkStepIdService(work_id, int64(work_step_id))
			}
		}
	}
	this.ServeJSON()
}

// 深拷贝对象
func deepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}
