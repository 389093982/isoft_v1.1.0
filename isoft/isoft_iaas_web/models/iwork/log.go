package iwork

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type RunLogRecord struct {
	Id              int64     `json:"id"`
	TrackingId      string    `json:"tracking_id"`
	WorkName        string    `json:"work_name"`
	CreatedBy       string    `json:"created_by"`
	CreatedTime     time.Time `json:"created_time" orm:"auto_now_add;type(datetime)"`
	LastUpdatedBy   string    `json:"last_updated_by"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}

type RunLogDetail struct {
	Id              int64     `json:"id"`
	TrackingId      string    `json:"tracking_id"`
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

func InsertRunLogDetail(trackingId, detail string) {
	insertRunLogDetailData(&RunLogDetail{
		TrackingId:      trackingId,
		Detail:          detail,
		CreatedBy:       "SYSTEM",
		CreatedTime:     time.Now(),
		LastUpdatedBy:   "SYSTEM",
		LastUpdatedTime: time.Now(),
	})
}

func QueryRunLogRecord(work_id string, page int, offset int) (runLogRecords []RunLogRecord, counts int64, err error) {
	work, _ := QueryWorkById(work_id)
	o := orm.NewOrm()
	qs := o.QueryTable("run_log_record").Filter("work_name", work.WorkName).OrderBy("-last_updated_time")
	counts, _ = qs.Count()
	qs = qs.Limit(offset, (page-1)*offset)
	qs.All(&runLogRecords)
	return
}

func GetLastRunLogDetail(tracking_id string) (runLogDetails []RunLogDetail, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("run_log_detail").Filter("tracking_id", tracking_id).All(&runLogDetails)
	return
}
