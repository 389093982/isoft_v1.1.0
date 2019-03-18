package iwork

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"isoft/isoft_iaas_web/models/iwork"
	"time"
)

func (this *WorkController) SaveHistory() {
	work_id, _ := this.GetInt64("work_id")
	work, _ := iwork.QueryWorkById(work_id, orm.NewOrm())
	steps, _ := iwork.QueryAllWorkStepInfo(work_id, orm.NewOrm())

	historyMap := make(map[string]interface{})
	historyMap["work"] = work
	historyMap["steps"] = steps

	var err error
	workHistory, err := json.MarshalIndent(historyMap, "", "\t")
	if err == nil {
		history := &iwork.WorkHistory{
			WorkId:          work.Id,
			WorkName:        work.WorkName,
			WorkDesc:        work.WorkDesc,
			WorkHistory:     string(workHistory),
			CreatedBy:       "SYSTEM",
			CreatedTime:     time.Now(),
			LastUpdatedBy:   "SYSTEM",
			LastUpdatedTime: time.Now(),
		}
		_, err = iwork.InsertOrUpdateWorkHistory(history)
	}
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}
