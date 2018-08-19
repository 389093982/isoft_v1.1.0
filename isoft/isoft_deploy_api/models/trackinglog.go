package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

const TASK_STATUS_BEGIN = "BEGIN"
const TASK_STATUS_END = "END"

type TrackingTask struct {
	Id              int64     `json:"id"`
	TrackingId      string    `json:"tracking_id"`
	TaskName        string    `json:"task_name"`
	TaskStatus      string    `json:"task_status"` // BEGIN/END
	EnvId           int64     `json:"env_id"`
	ServiceId       int64     `json:"service_id"`
	CreatedBy       string    `json:"created_by"`
	CreatedTime     time.Time `json:"created_time"`
	LastUpdatedBy   string    `json:"last_updated_by"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}

type TrackingLog struct {
	Id              int64         `json:"id"`
	TrackingId      string        `json:"tracking_id"`
	TaskName        string        `json:"task_name"`
	TrackingDetail  orm.TextField `json:"tracking_detail"`
	EnvId           int64         `json:"env_id"`
	TrackingKey     string        `json:"tracking_key"`
	TrackingValue   string        `json:"tracking_value"`
	ServiceId       int64         `json:"service_id"`
	CreatedBy       string        `json:"created_by"`
	CreatedTime     time.Time     `json:"created_time"`
	LastUpdatedBy   string        `json:"last_updated_by"`
	LastUpdatedTime time.Time     `json:"last_updated_time"`
}

func InsertTrackingLog(trackingLog *TrackingLog) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(trackingLog)
	return
}

func InsertOrUpdateTrackingTask(trackingTask *TrackingTask) (id int64, err error) {
	o := orm.NewOrm()
	if trackingTask.Id > 0 {
		id, err = o.Update(trackingTask)
	} else {
		id, err = o.Insert(trackingTask)
	}
	return
}

func QueryLastRunTrackingLog(tracking_id string) (trackingLogs []*TrackingLog, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("tracking_log")
	// 根据主键 id 倒序排序
	_, err = qs.Filter("tracking_id", tracking_id).OrderBy("-id").All(&trackingLogs)
	return
}

// 判断任务是否完成
func IsFinishTask(tracking_id string) bool {
	var trackingTask TrackingTask
	o := orm.NewOrm()
	qs := o.QueryTable("tracking_task")
	err := qs.Filter("tracking_id", tracking_id).One(&trackingTask)
	if err == nil && trackingTask.TaskStatus != "END" {
		return false
	} else {
		return true
	}
}

func QueryLastDeployTrackings(service_id int64) (trackingLogs []TrackingLog, err error) {
	tracking_id, err := QueryLastDeployTrackingId(service_id)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	qs := o.QueryTable("tracking_log")
	_, err = qs.Filter("tracking_id", tracking_id).OrderBy("last_updated_time").All(&trackingLogs)
	if err != nil {
		return nil, err
	}
	return trackingLogs, nil
}

func QueryLastDeployTrackingId(service_id int64) (tracking_id string, err error) {
	var trackingLog TrackingLog
	o := orm.NewOrm()
	qs := o.QueryTable("tracking_log")
	err = qs.Filter("service_id", service_id).OrderBy("-last_updated_time").One(&trackingLog)
	if err != nil {
		return "", err
	}
	return trackingLog.TrackingId, nil
}
