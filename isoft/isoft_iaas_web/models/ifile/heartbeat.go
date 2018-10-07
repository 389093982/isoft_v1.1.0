package ifile

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type HeartBeat struct {
	Id              int64     `json:"id"`
	Addr            string    `json:"addr"` // 服务器名称
	CreatedBy       string    `json:"created_by"`
	CreatedTime     time.Time `json:"created_time"`
	LastUpdatedBy   string    `json:"last_updated_by"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}

type LocateMessage struct {
	Addr            string    `json:"addr"`     // 服务器名称
	Hash            string    `json:"hash"`     // 对象 hash 值
	ShardId         int       `json:"shard_id"` // 分片 ID
	CreatedBy       string    `json:"created_by"`
	CreatedTime     time.Time `json:"created_time"`
	LastUpdatedBy   string    `json:"last_updated_by"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}

// 查询所有活着的心跳信息
func QueryAllAliveHeartBeat() (heartBeats []HeartBeat, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("heart_beat")
	s, _ := time.ParseDuration("-1s")                                                  // 1s 前
	_, err = qs.Filter("last_updated_time__gt", time.Now().Add(s*10)).All(&heartBeats) // 10s 前
	return
}

// 插入或者更新心跳信息
func InsertOrUpdateHeartBeat(heartBeat *HeartBeat) (id int64, err error) {
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

func FilterHeartBeat(condArr map[string]interface{}) (heartBeat HeartBeat, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("heart_beat")
	if addr, ok := condArr["addr"]; ok {
		qs = qs.Filter("addr", addr)
	}
	err = qs.One(&heartBeat)
	return
}
