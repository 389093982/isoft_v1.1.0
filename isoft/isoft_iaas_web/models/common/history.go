package common

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type History struct {
	Id              int64     `json:"id"`            // history id
	HistoryName     string    `json:"history_name"`  // history 名称
	HistoryValue    string    `json:"history_value"` // history 值
	HistoryDesc     string    `json:"history_desc"`  // history 描述
	HistoryLink     string    `json:"history_link"`  // history 对应的链接地址
	CreatedBy       string    `json:"created_by"`
	CreatedTime     time.Time `json:"created_time"`
	LastUpdatedBy   string    `json:"last_updated_by"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}

func InsertOrUpdateHistory(history *History) (id int64, err error) {
	oldHistory, err := FilterHistory(map[string]interface{}{"created_by": history.CreatedBy, "history_name": history.HistoryName, "history_value": history.HistoryValue})
	if err == nil {
		// 存在则插入
		history.Id = oldHistory.Id
		history.CreatedTime = oldHistory.CreatedTime
		history.LastUpdatedTime = oldHistory.LastUpdatedTime
	}
	o := orm.NewOrm()
	if history.Id > 0 {
		id, err = o.Update(history)
	} else {
		id, err = o.Insert(history)
	}
	return
}

func FilterHistory(condArr map[string]interface{}) (history History, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("history")
	if created_by, ok := condArr["created_by"]; ok {
		qs = qs.Filter("created_by", created_by)
	}
	if history_name, ok := condArr["history_name"]; ok {
		qs = qs.Filter("history_name", history_name)
	}
	if history_value, ok := condArr["history_value"]; ok {
		qs = qs.Filter("history_value", history_value)
	}
	err = qs.One(&history)
	return
}

func FilterHistoryByName(history_name string, page int, offset int, userName string) (history []History, counts int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("history")
	qs = qs.Filter("history_name", history_name).Filter("created_by", userName)
	qs = qs.OrderBy("-last_updated_time")
	counts, _ = qs.Count()
	qs = qs.Limit(offset, (page-1)*offset)
	qs.All(&history)
	return
}
