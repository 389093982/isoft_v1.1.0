package common

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type History struct {
	Id                 int64            `json:"id"`                         // history id
	HistoryName 	   string			`json:"history_name"`				// 历史名称
	HistoryValue 	   string			`json:"history_value"`				// 历史值
	CreatedBy          string           `json:"created_by"`
	CreatedTime        time.Time        `json:"created_time"`
	LastUpdatedBy      string           `json:"last_updated_by"`
	LastUpdatedTime    time.Time        `json:"last_updated_time"`
}

func AddHistory(history *History) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(history)
	return
}