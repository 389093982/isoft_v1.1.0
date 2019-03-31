package iwork

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type TableMigrate struct {
	Id              int64     `json:"id"`
	TableName       string    `json:"table_name"`
	TableColumns    string    `json:"table_columns"`
	CreatedBy       string    `json:"created_by"`
	CreatedTime     time.Time `json:"created_time" orm:"auto_now_add;type(datetime)"`
	LastUpdatedBy   string    `json:"last_updated_by"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}

func InsertOrUpdateTableMigrate(tm *TableMigrate) (id int64, err error) {
	o := orm.NewOrm()
	if tm.Id > 0 {
		id, err = o.Update(tm)
	} else {
		id, err = o.Insert(tm)
	}
	return
}
