package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"isoft/isoft/common/pageutil"
	"isoft/isoft_deploy_web/models"
	"time"
)

type MonitorHeartBeatController struct {
	beego.Controller
}

func (this *MonitorHeartBeatController) SendMonitorHeartBeat() {
	addr := this.GetString("addr")
	monitorHeartBeat := models.MonitorHeartBeat{
		Addr:            addr,
		CreatedBy:       "AutoInsert",
		CreatedTime:     time.Now(),
		LastUpdatedBy:   "AutoInsert",
		LastUpdatedTime: time.Now(),
	}
	_, err := models.InsertOrUpdateMonitorHeartBeat(&monitorHeartBeat)
	if err != nil {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	}
	this.ServeJSON()
}

func (this *MonitorHeartBeatController) FilterPageMonitorHeartBeat() {
	page_size, _ := this.GetInt("page_size", 10)      // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页

	condArr := make(map[string]interface{})
	monitorHeartBeats, count, err := models.FilterPageMonitorHeartBeat(condArr, current_page, page_size)
	paginator := pagination.SetPaginator(this.Ctx, page_size, count)
	//初始化
	dataMap := make(map[string]interface{}, 1)
	if err == nil {
		dataMap["status"] = "SUCCESS"
		dataMap["monitorHeartBeats"] = monitorHeartBeats
		dataMap["paginator"] = pageutil.Paginator(paginator.Page(), paginator.PerPageNums, paginator.Nums())
		dataMap["alives"] = getAliveMonitorHeartBeats(monitorHeartBeats)
	} else {
		dataMap["status"] = "ERROR"
		dataMap["errorMsg"] = err.Error()
	}
	this.Data["json"] = &dataMap
	this.ServeJSON()
}

func getAliveMonitorHeartBeats(monitorHeartBeats []models.MonitorHeartBeat) (alives []string) {
	for _, monitorHeartBeat := range monitorHeartBeats {
		s, _ := time.ParseDuration("-1s") // 1s 前
		if time.Now().Add(s * 10).Before(monitorHeartBeat.LastUpdatedTime) {
			alives = append(alives, monitorHeartBeat.Addr)
		}
	}
	return
}
