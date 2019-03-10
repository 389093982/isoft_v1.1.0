package iwork

import (
	"github.com/astaxie/beego/utils/pagination"
	"isoft/isoft/common/pageutil"
	"isoft/isoft_iaas_web/models/iwork"
	"time"
)

func (this *WorkController) DeleteEntity() {
	entity_id, _ := this.GetInt64("entity_id", -1)
	if err := iwork.DeleteEntityById(entity_id); err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func (this *WorkController) EditEntity() {
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

func (this *WorkController) FilterPageEntity() {
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
