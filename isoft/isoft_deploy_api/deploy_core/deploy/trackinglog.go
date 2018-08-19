package deploy

import (
	"github.com/astaxie/beego/orm"
	"isoft/isoft_deploy_api/models"
	"log"
	"strings"
	"time"
)

type TrackingLogResolver struct {
	ServiceInfo *models.ServiceInfo
	Task        *models.TrackingTask
}

// 开启一个新的任务,并标记状态为 BEGIN
func (this *TrackingLogResolver) StartRecordNewTask(tracking_id string, task_name string) {
	this.Task = &models.TrackingTask{
		TrackingId:      tracking_id,
		TaskName:        task_name,
		TaskStatus:      models.TASK_STATUS_BEGIN,
		EnvId:           this.ServiceInfo.EnvInfo.Id,
		ServiceId:       this.ServiceInfo.Id,
		CreatedBy:       "AutoInsert",
		CreatedTime:     time.Now(),
		LastUpdatedBy:   "AutoInsert",
		LastUpdatedTime: time.Now(),
	}

	_, err := models.InsertOrUpdateTrackingTask(this.Task)
	if err != nil {
		log.Panicln("insert trackingtask err")
	}
}

// 结束任务,并标记任务状态为 END
func (this *TrackingLogResolver) EndRecordTask() {
	this.Task.TaskStatus = models.TASK_STATUS_END
	this.Task.LastUpdatedTime = time.Now()
	_, err := models.InsertOrUpdateTrackingTask(this.Task)
	if err != nil {
		log.Panicln("insert trackingtask err")
	}
}

// 记录日志
func (this *TrackingLogResolver) WriteSuccessLog(message string) {
	trackingLog := &models.TrackingLog{
		TrackingId:      this.Task.TrackingId,
		TaskName:        this.Task.TaskName,
		TrackingDetail:  orm.TextField(message),
		EnvId:           this.ServiceInfo.EnvInfo.Id,
		ServiceId:       this.ServiceInfo.Id,
		CreatedBy:       "AutoInsert",
		CreatedTime:     time.Now(),
		LastUpdatedBy:   "AutoInsert",
		LastUpdatedTime: time.Now(),
	}

	if strings.Contains(message, "__") {
		slice := strings.Split(message, "__")
		trackingLog.TrackingKey = slice[0]
		trackingLog.TrackingValue = slice[1]
	}

	_, err := models.InsertTrackingLog(trackingLog)

	if err != nil {
		log.Panicln("insert trackinglog err")
	}
}

func (this *TrackingLogResolver) WriteErrorLog(message string) {
	this.WriteSuccessLog(message)
}
