package iresource

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"isoft/isoft/common/pageutil"
	"isoft/isoft_iaas_web/models/iresource"
	"time"
)

type ResourceController struct {
	beego.Controller
}

func (this *ResourceController) AddResource() {
	var resource iresource.Resource
	resource.ResourceName = this.Input().Get("resource_name")
	resource.ResourceType = this.Input().Get("resource_type")
	resource.ResourceUrl = this.Input().Get("resource_url")
	resource.ResourceDsn = this.Input().Get("resource_dsn")
	resource.ResourceUsername = this.Input().Get("resource_username")
	resource.ResourcePassword = this.Input().Get("resource_password")
	resource.EnvName = this.Input().Get("env_name")
	resource.CreatedBy = "SYSTEM"
	resource.CreatedTime = time.Now()
	resource.LastUpdatedBy = "SYSTEM"
	resource.LastUpdatedTime = time.Now()
	if _, err := iresource.InsertOrUpdateResource(&resource); err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func (this *ResourceController) FilterPageResource() {
	condArr := make(map[string]string)
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页
	if search := this.GetString("search"); search != "" {
		condArr["search"] = search
	}
	resources, count, err := iresource.QueryResource(condArr, current_page, offset)
	paginator := pagination.SetPaginator(this.Ctx, offset, count)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "resources": resources,
			"paginator": pageutil.Paginator(paginator.Page(), paginator.PerPageNums, paginator.Nums())}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}
