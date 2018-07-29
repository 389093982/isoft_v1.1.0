package locate

import (
	"encoding/json"
	"isoft/isoft_storage/cfg"
	"isoft/isoft_storage/lib/rabbitmq"
	"isoft/isoft_storage/lib/rs"
	"isoft/isoft_storage/lib/types"
	"time"
)

// 并向数据服务节点群发对象名字的定位消息,并接收反馈消息
func Locate(hash string) (locateInfo map[int]string) {
	q := rabbitmq.New(cfg.GetConfigValue(cfg.RABBITMQ_SERVER))
	q.Publish("dataServers", hash)
	c := q.Consume()
	go func() {
		time.Sleep(time.Second)
		q.Close()
	}()
	locateInfo = make(map[int]string)
	// 循环获取 6 条定位消息
	for i := 0; i < rs.ALL_SHARDS; i++ {
		msg := <-c
		if len(msg.Body) == 0 {
			return
		}
		var info types.LocateMessage
		json.Unmarshal(msg.Body, &info)
		locateInfo[info.Id] = info.Addr
	}
	return
}

func Exist(hash string) bool {
	// 返回的定位消息数据大于等于 4 个分片数量,则表示存在
	return len(Locate(hash)) >= rs.DATA_SHARDS
}
