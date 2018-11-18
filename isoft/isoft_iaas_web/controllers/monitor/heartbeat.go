package monitor

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"isoft/isoft/common/pageutil"
	"isoft/isoft_iaas_web/models/monitor"
	"time"
)

type HeartBeatController struct {
	beego.Controller
}

func (this *HeartBeatController) RegisterHeartBeat() {
	addr := this.GetString("addr")
	heartBeat := monitor.HeartBeat2{
		Addr:            addr,
		CreatedBy:       "AutoInsert",
		CreatedTime:     time.Now(),
		LastUpdatedBy:   "AutoInsert",
		LastUpdatedTime: time.Now(),
	}
	_, err := monitor.InsertOrUpdateHeartBeat(&heartBeat)
	if err != nil {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	}
	this.ServeJSON()
}

func (this *HeartBeatController) FilterPageHeartBeat() {
	page_size, _ := this.GetInt("page_size", 10)      // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页

	condArr := make(map[string]interface{})
	heartBeats, count, err := monitor.FilterPageHeartBeat(condArr, current_page, page_size)
	paginator := pagination.SetPaginator(this.Ctx, page_size, count)
	//初始化
	dataMap := make(map[string]interface{}, 1)
	if err == nil {
		dataMap["status"] = "SUCCESS"
		dataMap["heartBeats"] = heartBeats
		dataMap["paginator"] = pageutil.Paginator(paginator.Page(), paginator.PerPageNums, paginator.Nums())
	} else {
		dataMap["status"] = "ERROR"
		dataMap["errorMsg"] = err.Error()
	}
	this.Data["json"] = &dataMap
	this.ServeJSON()
}
