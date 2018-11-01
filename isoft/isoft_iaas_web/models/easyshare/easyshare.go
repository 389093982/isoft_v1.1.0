package easyshare

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// 共享链接模型类
type ShareLink struct {
	Id              int       `json:"id"`
	ShareType       string    `json:"share_type"`        // 分享类型
	ShareDesc		string	  `json:"share_desc"`		 // 分享描述
	Author          string    `json:"author"`            // 作者
	LinkHref        string    `json:"link_href"`         // 分享链接
	Content			string	  `json:"content" orm:"type(text)"`			 // 内容
	CreatedBy       string    `json:"created_by"`        // 创建人
	CreatedTime     time.Time `json:"created_time"`      // 创建时间
	LastUpdatedBy   string    `json:"last_updated_by"`   // 修改人
	LastUpdatedTime time.Time `json:"last_updated_time"` // 修改时间
}

func FilterShareLinkList(condArr map[string]string, page int, offset int, userName string) (shareLink []ShareLink, counts int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("share_link")
	if search_type, ok := condArr["search_type"]; ok {
		if search_type == "_hot"{

		}else if search_type == "_personal"{
			qs = qs.Filter("created_by",userName)
		}else if search_type != "_all"{
			qs = qs.Filter("share_type",search_type)
		}
	}
	qs = qs.OrderBy("-last_updated_time")
	counts, _ = qs.Count()
	qs = qs.Limit(offset, (page-1)*offset)
	qs.All(&shareLink)
	return
}

func AddNewShareLink(shareLink *ShareLink) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(shareLink)
	return id, err
}
