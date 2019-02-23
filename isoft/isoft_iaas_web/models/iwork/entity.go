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

func InsertOrUpdateEntity(entity *Entity) (id int64, err error) {
	o := orm.NewOrm()
	if entity.Id > 0 {
		id, err = o.Update(entity)
	} else {
		id, err = o.Insert(entity)
	}
	return
}

func QueryEntityByName(entity_name string) (entity Entity, err error) {
	o := orm.NewOrm()
	err = o.QueryTable("entity").Filter("entity_name", entity_name).One(&entity)
	return
}

func DeleteEntityById(entity_id int64) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("entity").Filter("id", entity_id).Delete()
	return err
}

func QueryAllEntityInfo() (entities []Entity) {
	o := orm.NewOrm()
	o.QueryTable("entity").OrderBy("id").All(&entities)
	return
}

func QueryEntityByEntityName(entity_name string) (entity Entity, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("entity").Filter("entity_name",entity_name).All(&entity)
	return
}