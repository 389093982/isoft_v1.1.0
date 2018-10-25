package easyshare

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"isoft/isoft/common/pageutil"
	"isoft/isoft_iaas_web/models/easyshare"
)

type ShareLinkController struct {
	beego.Controller
}

func (this *ShareLinkController) FilterShareLinkList()  {
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页
	shareLinks, count, err := easyshare.FilterShareLinkList(map[string]string{}, current_page, offset)
	paginator := pagination.SetPaginator(this.Ctx, offset, count)
	//初始化
	if err != nil {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	} else {
		paginatorMap := pageutil.Paginator(paginator.Page(), paginator.PerPageNums, paginator.Nums())
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "shareLinks": &shareLinks, "paginator": &paginatorMap}
	}
	this.ServeJSON()
}

