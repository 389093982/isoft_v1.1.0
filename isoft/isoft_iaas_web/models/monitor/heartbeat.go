package monitor

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type HeartBeat2 struct {
	Id              int64     `json:"id"`
	Addr            string    `json:"addr"` 				// 请求地址
	StatusCode 		int	  	  `json:"status_code"`			// 请求返回的状态码
	CreatedBy       string    `json:"created_by"`
	CreatedTime     time.Time `json:"created_time"`
	LastUpdatedBy   string    `json:"last_updated_by"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}

type HeartBeatDetail struct {
	Id              int64     `json:"id"`
	Addr            string    `json:"addr"` 				// 请求地址
	StatusCode 		int	  `json:"status_code"`			// 请求返回的状态码
	CreatedBy       string    `json:"created_by"`
	CreatedTime     time.Time `json:"created_time"`
	LastUpdatedBy   string    `json:"last_updated_by"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}

func InsertHeartBeatDetail(heartBeatDetail *HeartBeatDetail) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(heartBeatDetail)
	return
}

// 插入或者更新心跳信息
func InsertOrUpdateHeartBeat(heartBeat *HeartBeat2) (id int64, err error) {
	oldHeartBeat, err := FilterHeartBeat(map[string]interface{}{"addr": heartBeat.Addr})
	if err == nil {
		heartBeat.Id = oldHeartBeat.Id
		heartBeat.CreatedTime = oldHeartBeat.CreatedTime
		heartBeat.CreatedBy = oldHeartBeat.CreatedBy
	}
	o := orm.NewOrm()
	if heartBeat.Id > 0 {
		id, err = o.Update(heartBeat)
	} else {
		id, err = o.Insert(heartBeat)
	}
	return
}

func FilterHeartBeat(condArr map[string]interface{}) (heartBeat HeartBeat2, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("heart_beat2")
	if addr, ok := condArr["addr"]; ok {
		qs = qs.Filter("addr", addr)
	}
	err = qs.One(&heartBeat)
	return
}

func FilterPageHeartBeat(condArr map[string]interface{}, current_page, page_size int) (heartBeat []HeartBeat2, counts int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("heart_beat2")
	counts, _ = qs.Count()
	_, err = qs.Limit(page_size, (current_page-1)*page_size).All(&heartBeat)
	return
}

func GetAllHeartBeat() (heartBeat []HeartBeat2, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("heart_beat2")
	_, err = qs.All(&heartBeat)
	return
}