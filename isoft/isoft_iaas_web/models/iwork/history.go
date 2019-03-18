package iwork

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type WorkHistory struct {
	Id              int64     `json:"id"`
	WorkId          int64     `json:"work_id"`
	WorkName        string    `json:"work_name"`
	WorkDesc        string    `json:"work_desc"`
	WorkHistory     string    `json:"work_history" orm:"type(text)"`
	CreatedBy       string    `json:"created_by"`
	CreatedTime     time.Time `json:"created_time" orm:"auto_now_add;type(datetime)"`
	LastUpdatedBy   string    `json:"last_updated_by"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}

func InsertOrUpdateWorkHistory(history *WorkHistory) (id int64, err error) {
	o := orm.NewOrm()
	if history.Id > 0 {
		id, err = o.Update(history)
	} else {
		id, err = o.Insert(history)
	}
	return
}
