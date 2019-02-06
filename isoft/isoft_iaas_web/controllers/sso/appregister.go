package sso

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"isoft/isoft/common/pageutil"
	"isoft/isoft_iaas_web/models/sso"
	"time"
)

type AppRegisterController struct {
	beego.Controller
}

func (this *AppRegisterController) AddAppRegister() {
	appAddress := this.GetString("app_address")
	var appRegister sso.AppRegister
	appRegister.AppAddress = appAddress
	appRegister.CreatedBy = "SYSTEM"
	appRegister.LastUpdatedBy = "SYSTEM"
	appRegister.CreatedTime = time.Now()
	appRegister.LastUpdatedTime = time.Now()

	count, err := sso.QueryRegisterCount(appAddress)
	if err == nil && count == 0 {
		_, err = sso.AddRegister(&appRegister)
		if err == nil {
			this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
		} else {
			this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": "保存失败!"}
		}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": "该系统已注册,请不要重复注册!"}
	}
	this.ServeJSON()
}

func (this *AppRegisterController) AppRegisterList() {
	condArr := make(map[string]string)
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页

	search := this.GetString("search")
	if search != "" {
		condArr["search"] = search
	}
	appregisters, count, err := sso.QueryRegister(condArr, current_page, offset)
	paginator := pagination.SetPaginator(this.Ctx, offset, count)

	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "appRegisters": appregisters,
			"paginator": pageutil.Paginator(paginator.Page(), paginator.PerPageNums, paginator.Nums())}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}
