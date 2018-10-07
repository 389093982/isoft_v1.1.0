package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type MonitorHeartBeat struct {
	Id              int64     `json:"id"`
	Addr            string    `json:"addr"` // 服务器名称
	CreatedBy       string    `json:"created_by"`
	CreatedTime     time.Time `json:"created_time"`
	LastUpdatedBy   string    `json:"last_updated_by"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}

// 插入或者更新心跳信息
func InsertOrUpdateMonitorHeartBeat(monitorHeartBeat *MonitorHeartBeat) (id int64, err error) {
	oldMonitorHeartBeat, err := FilterMonitorHeartBeat(map[string]interface{}{"addr": monitorHeartBeat.Addr})
	if err == nil {
		monitorHeartBeat.Id = oldMonitorHeartBeat.Id
		monitorHeartBeat.CreatedTime = oldMonitorHeartBeat.CreatedTime
		monitorHeartBeat.CreatedBy = oldMonitorHeartBeat.CreatedBy
	}
	o := orm.NewOrm()
	if monitorHeartBeat.Id > 0 {
		id, err = o.Update(monitorHeartBeat)
	} else {
		id, err = o.Insert(monitorHeartBeat)
	}
	return
}

func FilterMonitorHeartBeat(condArr map[string]interface{}) (monitorHeartBeat MonitorHeartBeat, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("monitor_heart_beat")
	if addr, ok := condArr["addr"]; ok {
		qs = qs.Filter("addr", addr)
	}
	err = qs.One(&monitorHeartBeat)
	return
}

func FilterPageMonitorHeartBeat(condArr map[string]interface{}, current_page, page_size int) (monitorHeartBeat []MonitorHeartBeat, counts int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("monitor_heart_beat")
	counts, _ = qs.Count()
	_, err = qs.Limit(page_size, (current_page-1)*page_size).All(&monitorHeartBeat)
	return
}
