package iwork

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"isoft/isoft/common/pageutil"
	"isoft/isoft_iaas_web/core/iworkconst"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworknode"
	"isoft/isoft_iaas_web/core/iworkrun"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
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

func (this *WorkController) GetRelativeWork() {
	work_id,_ := this.GetInt64("work_id")
	subWorks := make([]iwork.Work,0)
	parentWorks, _,_ := iwork.QueryParentWorks(work_id)
	steps, _ := iwork.QueryAllWorkStepInfo(work_id)
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
	runLogDetails, err := iwork.QueryLastRunLogDetail(tracking_id)
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
	steps, _ := iwork.QueryAllWorkStepInfo(work_id)
	go iworkrun.Run(work, steps, nil)
	this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	this.ServeJSON()
}


func (this *WorkController) EditWork() {
	defer func() {
		if err := recover(); err != nil{
			this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
			this.ServeJSON()
		}
	}()
	// 将请求参数封装成 work
	var work iwork.Work
	work.WorkName = this.GetString("work_name")
	work.WorkDesc = this.GetString("work_desc")
	work.CreatedBy = "SYSTEM"
	work.CreatedTime = time.Now()
	work.LastUpdatedBy = "SYSTEM"
	work.LastUpdatedTime = time.Now()

	var oldWorkName string
	work_id, err := this.GetInt64("work_id", -1)
	if err == nil && work_id > 0 {
		work, err = iwork.QueryWorkById(work_id)
		oldWorkName = work.WorkName
	}
	if _, err := iwork.InsertOrUpdateWork(&work); err == nil {
		if work_id <= 0 {
			// 新增 work 场景,自动添加开始和结束节点
			insertStartEndWorkStepNode(work.Id)
			iwork.InsertOrUpdateCronMeta(&iwork.CronMeta{
				TaskName:work.WorkName,
				TaskType:"iwork_quartz",
				CronStr:"0 * * * * ?",
				CreatedBy:"SYSTEM",
				CreatedTime:time.Now(),
				LastUpdatedBy:"SYSTEM",
				LastUpdatedTime:time.Now(),
			})
		}else{
			// 修改 work 场景
			changeReferencesWorkName(work_id, oldWorkName, work.WorkName)
			meta, _ := iwork.QueryCronMetaByName(oldWorkName)
			iwork.InsertOrUpdateCronMeta(&iwork.CronMeta{
				Id:meta.Id,
				TaskName:work.WorkName,
				TaskType:"iwork_quartz",
				CronStr:"0 * * * * ?",
				CreatedBy:"SYSTEM",
				CreatedTime:time.Now(),
				LastUpdatedBy:"SYSTEM",
				LastUpdatedTime:time.Now(),
			})
		}
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func changeReferencesWorkName(work_id int64, oldWorkName,workName string) error {
	if oldWorkName == workName{
		return nil
	}
	parentWorks, _,err := iwork.QueryParentWorks(work_id)
	if err != nil {
		return nil
	}
	for _, parentWork := range parentWorks{
		steps, _ := iwork.QueryAllWorkStepInfo(parentWork.Id)
		for _, step := range steps{
			if step.WorkStepType != "work_sub"{
				continue
			}
			inputSchema := schema.GetCacheParamInputSchema(&step, &iworknode.WorkStepFactory{WorkStep:&step})
			for index, item := range inputSchema.ParamInputSchemaItems{
				if item.ParamName == iworkconst.STRING_PREFIX + "work_sub" && strings.Contains(item.ParamValue, oldWorkName){
					inputSchema.ParamInputSchemaItems[index].ParamValue = strings.Replace(item.ParamValue, oldWorkName, workName, -1)
				}
			}
			step.WorkStepInput = inputSchema.RenderToXml()
			iwork.InsertOrUpdateWorkStep(&step)
		}
	}
	return nil
}

func insertStartEndWorkStepNode(work_id int64) {
	insertDefaultWorkStepNodeFunc := func(nodeName string, work_step_id int64) {
		step := &iwork.WorkStep{
			WorkId:          work_id,
			WorkStepId:      work_step_id,
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
	insertDefaultWorkStepNodeFunc("start", 1)
	insertDefaultWorkStepNodeFunc("end", 2)
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



