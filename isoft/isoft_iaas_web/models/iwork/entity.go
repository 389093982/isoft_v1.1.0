package iwork

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Entity struct {
	Id              int64     `json:"id"`
	EntityName      string    `json:"entity_name"`
	EntityFieldStr  string    `json:"entity_field_str"`
	CreatedBy       string    `json:"created_by"`
	CreatedTime     time.Time `json:"created_time" orm:"auto_now_add;type(datetime)"`
	LastUpdatedBy   string    `json:"last_updated_by"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}

func QueryEntity(page int, offset int) (entities []Entity, counts int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("entity")
	counts, _ = qs.Count()
	qs = qs.Limit(offset, (page-1)*offset)
	qs.All(&entities)
	return
}