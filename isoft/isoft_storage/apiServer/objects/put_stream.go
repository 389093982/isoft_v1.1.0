package objects

import (
	"../heartbeat"
	"fmt"
	"isoft/isoft_storage/lib/rs"
)

func putStream(hash string, size int64) (*rs.RSPutStream, error) {
	servers := heartbeat.ChooseRandomDataServers(rs.ALL_SHARDS, nil)
	if len(servers) != rs.ALL_SHARDS {
		return nil, fmt.Errorf("cannot find enough dataServer")
	}

	// NewTempPutStream 底层主要是调用数据服务 temp 接口的 post 方法生产临时文件
	return rs.NewRSPutStream(servers, hash, size)
}
