package task

import (
	"github.com/astaxie/beego/logs"
	"isoft/isoft_deploy_web/models"
	"net/http"
	"time"
)

func RunServiceMonitorTask() error {
	serviceMonitors, err := models.QueryAllServiceMonitor()
	if err != nil {
		logs.Error("start serviceMonitorTask err: %s", err.Error())
		return err
	}

	for _, serviceMonitor := range serviceMonitors {
		go func() {
			response, err := http.Get(serviceMonitor.Url)
			if err != nil {
				logs.Error("check service err: %s", err.Error())
				return
			}
			defer response.Body.Close()

			serviceMonitor.StatusCode = int64(response.StatusCode)
			serviceMonitor.LastUpdatedTime = time.Now()
			models.InsertOrUpdateServiceMonitor(&serviceMonitor)

			serviceMonitorDetail := &models.ServiceMonitorDetail{
				Url:             serviceMonitor.Url,
				Method:          serviceMonitor.Method,
				StatusCode:      serviceMonitor.StatusCode,
				CreatedBy:       serviceMonitor.CreatedBy,
				CreatedTime:     time.Now(),
				LastUpdatedBy:   serviceMonitor.LastUpdatedBy,
				LastUpdatedTime: time.Now(),
			}
			models.InsertOrUpdateServiceMonitorDetail(serviceMonitorDetail)
		}()
	}

	return nil
}
