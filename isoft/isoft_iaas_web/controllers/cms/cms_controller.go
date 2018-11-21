package cms

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"isoft/isoft/common/pageutil"
	"isoft/isoft_iaas_web/models/cms"
	"time"
)

type CMSController struct {
	beego.Controller
}

func (this *CMSController) FilterConfigurations() {
	condArr := make(map[string]string)
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页
	search := this.GetString("search")
	if search != "" {
		condArr["search"] = search
	}
	configurations, count, err := cms.FilterConfigurations(condArr, current_page, offset)
	paginator := pagination.SetPaginator(this.Ctx, offset, count)
	if err != nil {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	} else {
		paginatorMap := pageutil.Paginator(paginator.Page(), paginator.PerPageNums, paginator.Nums())
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "configurations": &configurations, "paginator": &paginatorMap}
	}
	this.ServeJSON()
}

func (this *CMSController) QueryAllConfigurations() {
	configuration_name := this.GetString("configuration_name")
	configurations, err := cms.QueryAllConfigurations(configuration_name, 0)
	if err != nil {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "configurations": &configurations}
	}
	this.ServeJSON()
}

func (this *CMSController) AddConfiguration() {
	user_name := this.Ctx.Input.Session("UserName").(string)
	parent_id, _ := this.GetInt64("parent_id", 0)
	configuration_name := this.GetString("configuration_name")
	configuration_value := this.GetString("configuration_value")
	configuration := &cms.Configuration{
		ParentId:           parent_id,
		ConfigurationName:  configuration_name,
		ConfigurationValue: configuration_value,
		CreatedBy:          user_name,
		CreatedTime:        time.Now(),
		LastUpdatedBy:      user_name,
		LastUpdatedTime:    time.Now(),
	}
	_, err := cms.AddConfiguration(configuration)
	if err != nil {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	}
	this.ServeJSON()
}

func (this *CMSController) QueryRandomFrinkLink()  {
	frind_links, err := cms.QueryRandomFrinkLink()
	if err != nil {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "frind_links": &frind_links}
	}
	this.ServeJSON()
}