package iwork

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils/pagination"
	"isoft/isoft/common/pageutil"
	"isoft/isoft_iaas_web/models/iwork"
	"time"
)

func (this *WorkController) AddQuartz() {
	var meta iwork.CronMeta
	meta.TaskName = this.Input().Get("task_name")
	meta.TaskType = this.Input().Get("task_type")
	meta.CronStr = this.Input().Get("cron_str")
	meta.Enable = false
	meta.CreatedBy = "SYSTEM"
	meta.CreatedTime = time.Now()
	meta.LastUpdatedBy = "SYSTEM"
	meta.LastUpdatedTime = time.Now()
	if _, err := iwork.InsertOrUpdateCronMeta(&meta, orm.NewOrm()); err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func (this *WorkController) FilterPageQuartz() {
	condArr := make(map[string]string)
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页
	if search := this.GetString("search"); search != "" {
		condArr["search"] = search
	}
	quartzs, count, err := iwork.QueryCronMeta(condArr, current_page, offset)
	paginator := pagination.SetPaginator(this.Ctx, offset, count)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "quartzs": quartzs,
			"paginator": pageutil.Paginator(paginator.Page(), paginator.PerPageNums, paginator.Nums())}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}
