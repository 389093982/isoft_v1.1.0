package objects

import (
	"../heartbeat"
	"../locate"
	"fmt"
	"isoft/isoft_storage/lib/rs"
)

func GetStream(hash string, size int64) (*rs.RSGetStream, error) {
	// 根据对象 hash 值获取定位消息
	locateInfo := locate.Locate(hash)
	// 小于 4 个分片数据,数据无法还原
	if len(locateInfo) < rs.DATA_SHARDS {
		return nil, fmt.Errorf("object %s locate fail, result %v", hash, locateInfo)
	}
	dataServers := make([]string, 0)
	// 定位信息数组长度不为 6,表示该对象有部分分片丢失
	if len(locateInfo) != rs.ALL_SHARDS {
		// 随机选取 rs.ALL_SHARDS-len(locateInfo) 个数据服务节点进行修复,排除 locateInfo 所在的已有分片节点
		dataServers = heartbeat.ChooseRandomDataServers(rs.ALL_SHARDS-len(locateInfo), locateInfo)
	}
	return rs.NewRSGetStream(locateInfo, dataServers, hash, size)
}
