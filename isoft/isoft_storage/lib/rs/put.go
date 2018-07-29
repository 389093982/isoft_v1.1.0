package rs

import (
	"fmt"
	"io"
	"isoft/isoft_storage/lib/objectstream"
)

type RSPutStream struct {
	*encoder
}

func NewRSPutStream(dataServers []string, hash string, size int64) (*RSPutStream, error) {
	if len(dataServers) != ALL_SHARDS {
		return nil, fmt.Errorf("dataServers number mismatch")
	}

	perShard := (size + DATA_SHARDS - 1) / DATA_SHARDS
	writers := make([]io.Writer, ALL_SHARDS)
	var e error
	for i := range writers {
		// NewTempPutStream 底层主要是调用数据服务 temp 接口的 post 方法生产临时文件
		writers[i], e = objectstream.NewTempPutStream(dataServers[i],
			fmt.Sprintf("%s.%d", hash, i), perShard)
		if e != nil {
			return nil, e
		}
	}
	enc := NewEncoder(writers)

	return &RSPutStream{enc}, nil
}

func (s *RSPutStream) Commit(success bool) {
	// Flush 方法将数据写入缓存
	s.Flush()
	for i := range s.writers {
		// Commit 方法将临时文件转正
		s.writers[i].(*objectstream.TempPutStream).Commit(success)
	}
}
