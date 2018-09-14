package objects

import (
	"fmt"
	"isoft/isoft/common/logutil"
	"isoft/isoft_storage/apiServer/heartbeat"
	"isoft/isoft_storage/lib/rs"
)

// putStream(NewRSPutStream) -> NewTempPutStream
func putStream(hash string, size int64) (*rs.RSPutStream, error) {
	servers := heartbeat.ChooseRandomDataServers(rs.ALL_SHARDS, nil) // 选择总分片数量个 server
	if len(servers) != rs.ALL_SHARDS {
		logutil.Errorln("cannot find enough dataServer:", len(servers))
		return nil, fmt.Errorf("cannot find enough dataServer")
	}

	// NewTempPutStream 底层主要是调用数据服务 temp 接口的 post 方法生产临时文件
	return rs.NewRSPutStream(servers, hash, size)
}
