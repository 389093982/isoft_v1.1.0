package locate

import (
	"isoft/isoft_storage/lib"
	"isoft/isoft_storage/lib/rs"
)

// 并向数据服务节点群发对象名字的定位消息,并接收反馈消息
func Locate(hash string) (locateInfo map[int]string) {
	proxy := &lib.LocateAndHeartbeatProxy{}
	return proxy.SendAndReceiveLocateInfo(hash)
}

func Exist(hash string) bool {
	// 返回的定位消息数据大于等于 4 个分片数量,则表示存在
	return len(Locate(hash)) >= rs.DATA_SHARDS
}
