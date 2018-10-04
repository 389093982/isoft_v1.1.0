package ifile

import (
	"github.com/astaxie/beego"
	"isoft/isoft_iaas_web/models/ifile"
	"time"
)

type HeartBeatController struct {
	beego.Controller
}

func (this *HeartBeatController) SendHeartBeat()  {
	addr := this.GetString("addr")
	heartBeat := ifile.HeartBeat{
		Addr:addr,
		CreatedBy:"AutoInsert",
		CreatedTime:time.Now(),
		LastUpdatedBy:"AutoInsert",
		LastUpdatedTime:time.Now(),
	}
	_, err := ifile.InsertOrUpdateHeartBeat(&heartBeat)
	if err != nil{
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}else{
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	}
	this.ServeJSON()
}

func (this *HeartBeatController) QueryAllAliveHeartBeat()  {
	heartbeats,err := ifile.QueryAllAliveHeartBeat()
	if err != nil{
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}else{
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "heartbeats":heartbeats}
	}
	this.ServeJSON()
}