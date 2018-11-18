package common

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"isoft/isoft/common/pageutil"
	"isoft/isoft_iaas_web/models/common"
)

type HistoryController struct {
	beego.Controller
}

func (this *HistoryController) ShowCourseHistory() {
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页
	user_name := this.Ctx.Input.Session("UserName").(string)
	historys, count, err := common.FilterHistoryByName("show_course_history", current_page, offset, user_name)
	//初始化
	if err != nil {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	} else {
		paginator := pagination.SetPaginator(this.Ctx, offset, count)
		paginatorMap := pageutil.Paginator(paginator.Page(), paginator.PerPageNums, paginator.Nums())
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "historys": &historys, "paginator": &paginatorMap}
	}
	this.ServeJSON()
}
