package iwork

import (
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
)

type Resource struct {
	Id               int64     `json:"id"`
	ResourceName     string    `json:"resource_name"`
	ResourceType     string    `json:"resource_type"`
	ResourceUrl      string    `json:"resource_url"`
	ResourceDsn      string    `json:"resource_dsn"`
	ResourceUsername string    `json:"resource_username"`
	ResourcePassword string    `json:"resource_password"`
	CreatedBy        string    `json:"created_by"`
	CreatedTime      time.Time `json:"created_time" orm:"auto_now_add;type(datetime)"`
	LastUpdatedBy    string    `json:"last_updated_by"`
	LastUpdatedTime  time.Time `json:"last_updated_time"`
}

func InsertOrUpdateResource(resource *Resource) (id int64, err error) {
	o := orm.NewOrm()
	if resource.Id > 0 {
		id, err = o.Update(resource)
	} else {
		id, err = o.Insert(resource)
	}
	return
}

func QueryResource(condArr map[string]string, page int, offset int) (resources []Resource, counts int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("resource")
	var cond = orm.NewCondition()
	if search, ok := condArr["search"]; ok && strings.TrimSpace(search) != "" {
		subCond := orm.NewCondition()
		subCond = cond.And("resource_name__contains", search).Or("resource_type__contains", search).Or("resource_url__contains", search)
		cond = cond.AndCond(subCond)
	}
	qs = qs.SetCond(cond)
	counts, _ = qs.Count()
	qs = qs.Limit(offset, (page-1)*offset)
	qs.All(&resources)
	return
}

func QueryAllResource() (resources []Resource) {
	o := orm.NewOrm()
	o.QueryTable("resource").All(&resources)
	return
}

func QueryResourceDataSourceNameString(resource_name string) string {
	var resource Resource
	o := orm.NewOrm()
	if err := o.QueryTable("resource").Filter("resource_name", resource_name).One(&resource); err == nil {
		return resource.ResourceDsn
	}
	return ""
}

func DeleteResource(id int64) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("resource").Filter("id", id).Delete()
	return err
}

func QueryResourceById(id int64) (resource Resource, err error) {
	o := orm.NewOrm()
	err = o.QueryTable("resource").Filter("id", id).One(&resource)
	return
}
