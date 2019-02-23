package task

import (
	"errors"
	"github.com/astaxie/beego/toolbox"
	"isoft/isoft_iaas_web/imodules"
	"isoft/isoft_iaas_web/models/monitor"
	"net/http"
)

func startILearningCronTask() {
	if task, err := getHeartBeatTask(); err == nil{
		task.Run()
		toolbox.AddTask("mytask", task)
	}
	toolbox.StartTask()
}


func getHeartBeatTask() (*toolbox.Task, error) {
	if imodules.CheckModule("ilearning"){
		return toolbox.NewTask("heartBeatTask", "0 * * * * *", heartBeatTaskFunc), nil
	}
	return nil, errors.New("ilearning moudle is not open!")
}

func heartBeatTaskFunc() error {
	heartBeats, err := monitor.GetAllHeartBeat()
	if err != nil {
		return err
	}
	for _, heartBeat := range heartBeats {
		var statusCode int
		resp, err := http.Get(heartBeat.Addr)
		if err != nil {
			statusCode = -1
		} else {
			statusCode = resp.StatusCode
		}
		heartBeat.StatusCode = statusCode
		monitor.InsertOrUpdateHeartBeat(&heartBeat)
		heartBeatDetail := &monitor.HeartBeatDetail{
			Addr:            heartBeat.Addr,
			StatusCode:      heartBeat.StatusCode,
			CreatedBy:       heartBeat.CreatedBy,
			CreatedTime:     heartBeat.CreatedTime,
			LastUpdatedBy:   heartBeat.LastUpdatedBy,
			LastUpdatedTime: heartBeat.LastUpdatedTime,
		}
		monitor.InsertHeartBeatDetail(heartBeatDetail)
	}
	return nil
}
