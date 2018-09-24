package objects

import (
	"errors"
	"fmt"
	"isoft/isoft_storage/apiServer/heartbeat"
	"isoft/isoft_storage/lib/rs"
	"time"
)

// putStream(NewRSPutStream) -> NewTempPutStream
func putStream(hash string, size int64) (*rs.RSPutStream, error) {
	startTime := time.Now()
	servers := heartbeat.ChooseRandomDataServers(rs.ALL_SHARDS, nil) // 选择总分片数量个 server
	fmt.Println("ChooseRandomDataServers 3:", time.Now().Sub(startTime))
	if len(servers) != rs.ALL_SHARDS {
		return nil, errors.New("cannot find enough dataServer")
	}

	// NewTempPutStream 底层主要是调用数据服务 temp 接口的 post 方法生产临时文件
	return rs.NewRSPutStream(servers, hash, size)
}
