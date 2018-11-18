package share

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"isoft/isoft/common/pageutil"
	"isoft/isoft_iaas_web/models/share"
	"time"
)

type ShareController struct {
	beego.Controller
}

func (this *ShareController) ShowShareDetail() {
	share_id, err := this.GetInt64("share_id")
	if err == nil {
		share.UpdateShareViews(share_id)
		share, err := share.QueryShareById(share_id)
		if err == nil {
			this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "share": &share}
		}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func (this *ShareController) FilterShareList() {
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页
	search_type := this.GetString("search_type")
	userName := this.GetSession("UserName").(string)
	shares, count, err := share.FilterShareList(map[string]string{"search_type": search_type}, current_page, offset, userName)
	paginator := pagination.SetPaginator(this.Ctx, offset, count)
	//初始化
	if err != nil {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	} else {
		paginatorMap := pageutil.Paginator(paginator.Page(), paginator.PerPageNums, paginator.Nums())
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "shares": &shares, "paginator": &paginatorMap}
	}
	this.ServeJSON()
}

func (this *ShareController) AddNewShare() {
	share_type := this.GetString("share_type")
	share_desc := this.GetString("share_desc")
	link_href := this.GetString("link_href")
	content := this.GetString("content")
	userName := this.GetSession("UserName").(string)
	newShare := share.Share{
		ShareType:       share_type,
		ShareDesc:       share_desc,
		Author:          userName,
		LinkHref:        link_href,
		Content:         content,
		CreatedBy:       userName,
		CreatedTime:     time.Now(),
		LastUpdatedBy:   userName,
		LastUpdatedTime: time.Now(),
	}
	_, err := share.AddNewShare(&newShare)
	//初始化
	if err != nil {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	}
	this.ServeJSON()
}
