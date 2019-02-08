package iwork

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type RunLogRecord struct {
	Id              int64     `json:"id"`
	TrackingId		string	  `json:"tracking_id"`
	WorkName 		string	  `json:"work_name"`
	CreatedBy       string    `json:"created_by"`
	CreatedTime     time.Time `json:"created_time" orm:"auto_now_add;type(datetime)"`
	LastUpdatedBy   string    `json:"last_updated_by"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}

type RunLogDetail struct {
	Id              int64     `json:"id"`
	TrackingId		string	  `json:"tracking_id"`
	Detail          string    `json:"detail" orm:"type(text)"`
	CreatedBy       string    `json:"created_by"`
	CreatedTime     time.Time `json:"created_time" orm:"auto_now_add;type(datetime)"`
	LastUpdatedBy   string    `json:"last_updated_by"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}

func InsertRunLogRecord(record *RunLogRecord) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(record)
	return
}

func insertRunLogDetailData(detail *RunLogDetail) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(detail)
	return
}

func InsertRunLogDetail(trackingId, detail string)  {
	insertRunLogDetailData(&RunLogDetail{
		TrackingId:trackingId,
		Detail:detail,
		CreatedBy:"SYSTEM",
		CreatedTime:time.Now(),
		LastUpdatedBy:"SYSTEM",
		LastUpdatedTime:time.Now(),
	})
}