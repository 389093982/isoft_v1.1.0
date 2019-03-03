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

func QueryRunLogRecord(work_id int64, page int, offset int) (runLogRecords []RunLogRecord, counts int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("run_log_record")
	if work_id > 0 {
		work, _ := QueryWorkById(work_id, o)
		qs = qs.Filter("work_name", work.WorkName)
	}
	// Exclude 非顶级流程日志不查出来
	qs = qs.Exclude("tracking_id__contains", ".").OrderBy("-last_updated_time")
	counts, _ = qs.Count()
	qs = qs.Limit(offset, (page-1)*offset)
	qs.All(&runLogRecords)
	return
}

func QueryLastRunLogDetail(tracking_id string) (runLogDetails []RunLogDetail, err error) {
	o := orm.NewOrm()
	// __startswith 多级 tracking_id 也查出来
	_, err = o.QueryTable("run_log_detail").Filter("tracking_id__startswith", tracking_id).All(&runLogDetails)
	return
}

type ValidateLogRecord struct {
	Id              int64     `json:"id"`
	TrackingId      string    `json:"tracking_id"`
	CreatedBy       string    `json:"created_by"`
	CreatedTime     time.Time `json:"created_time" orm:"auto_now_add;type(datetime)"`
	LastUpdatedBy   string    `json:"last_updated_by"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}

type ValidateLogDetail struct {
	Id              int64     `json:"id"`
	TrackingId      string    `json:"tracking_id"`
	WorkId          int64     `json:"work_id"`
	WorkStepId      int64     `json:"work_step_id"`
	WorkName        string    `json:"work_name"`
	WorkStepName    string    `json:"work_step_name"`
	Detail          string    `json:"detail" orm:"type(text)"`
	CreatedBy       string    `json:"created_by"`
	CreatedTime     time.Time `json:"created_time" orm:"auto_now_add;type(datetime)"`
	LastUpdatedBy   string    `json:"last_updated_by"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}

func InsertValidateLogRecord(record *ValidateLogRecord) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(record)
	return
}

func InsertValidateLogDetail(detail *ValidateLogDetail) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(detail)
	return
}

func QueryLastValidateLogRecord() (record ValidateLogRecord, err error) {
	o := orm.NewOrm()
	err = o.QueryTable("validate_log_record").OrderBy("-last_updated_time").One(&record)
	return
}

func QueryLastValidateLogDetail() (details []ValidateLogDetail, err error) {
	if record, err := QueryLastValidateLogRecord(); err == nil {
		o := orm.NewOrm()
		_, err = o.QueryTable("validate_log_detail").Filter("tracking_id", record.TrackingId).All(&details)
	}
	return
}
