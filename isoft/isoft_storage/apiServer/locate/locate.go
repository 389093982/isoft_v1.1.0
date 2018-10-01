package locate

import (
	"isoft/isoft_storage/apiServer/heartbeat"
	"isoft/isoft_storage/lib"
	"isoft/isoft_storage/lib/rs"
	"isoft/isoft_storage/lib/utils"
	"time"
)

// 并向数据服务节点群发对象名字的定位消息,并接收反馈消息
func Locate(hash string) (locateInfo map[int]string) {
	defer utils.RecordTimeCostForMethod("apiServer locate Locate", time.Now())

	proxy := &lib.LocateAndHeartbeatProxy{}
	return proxy.SendAndReceiveLocateInfo(heartbeat.GetDataServers(),hash, 3)
}

func Exist(hash string) bool {
	defer utils.RecordTimeCostForMethod("apiServer locate Exist", time.Now())

	// 返回的定位消息数据大于等于 4 个分片数量,则表示存在
	return len(Locate(hash)) >= rs.DATA_SHARDS
}
