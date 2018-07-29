package rs

import (
	"fmt"
	"io"
	"isoft/isoft_storage/lib/objectstream"
)

type RSPutStream struct {
	*encoder
}

// 纠删码技术主要是通过纠删码算法将原始的数据进行编码得到冗余,并将数据和冗余一并存储起来,以达到容错的目的.
// 其基本思想是将n块原始的数据元素通过一定的计算,得到m块冗余元素(校验块).对于这n+m块的元素,当其中任意的
// m块元素出错(包括原始数据和冗余数据)时,均可以通过对应的重构算法恢复出原来的n块数据.生成校验的过程被成为
// 编码(encoding),恢复丢失数据块的过程被称为解码(decoding).磁盘利用率为n/(n+m).基于纠删码的方法与多
// 副本方法相比具有冗余度低、磁盘利用率高等优点
func NewRSPutStream(dataServers []string, hash string, size int64) (*RSPutStream, error) {
	if len(dataServers) != ALL_SHARDS { // server 数量与分片数量不匹配
		return nil, fmt.Errorf("dataServers number mismatch")
	}
	// 根据 size 计算出每个分片的大小 perShard
	perShard := (size + DATA_SHARDS - 1) / DATA_SHARDS
	// 创建长度为 6 的 io.Writer 数组,每一个元素用来存放 objectstream.NewTempPutStream,用于上传一个分片对象
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
