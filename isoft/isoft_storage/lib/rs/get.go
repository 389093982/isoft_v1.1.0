package rs

import (
	"fmt"
	"io"
	"isoft/isoft_storage/lib/objectstream"
)

type RSGetStream struct {
	*decoder
}

func NewRSGetStream(locateInfo map[int]string, dataServers []string, hash string, size int64) (*RSGetStream, error) {
	// 检查总数量是否为 6
	if len(locateInfo)+len(dataServers) != ALL_SHARDS {
		return nil, fmt.Errorf("dataServers number mismatch")
	}
	// 创建长度为 6 的 io.Reader 数组 readers,用于读取 6 个分片的数据
	readers := make([]io.Reader, ALL_SHARDS)
	for i := 0; i < ALL_SHARDS; i++ {
		server := locateInfo[i]
		if server == "" { // 某个分片 id 相对的数据服务节点地址为空,说明该分片丢失,我们需要取一个随机数据服务节点补上
			locateInfo[i] = dataServers[0]
			dataServers = dataServers[1:]
			continue
		}
		// 对象读取流 NewGetStream,用于读取该分片数据
		reader, e := objectstream.NewGetStream(server, fmt.Sprintf("%s.%d", hash, i))
		if e == nil {
			readers[i] = reader
		}
	}

	writers := make([]io.Writer, ALL_SHARDS)
	perShard := (size + DATA_SHARDS - 1) / DATA_SHARDS
	var e error
	for i := range readers {
		if readers[i] == nil { // 表示需要进行恢复
			// 调用 NewTempPutStream 创建相应的临时对象写入流用于恢复分片
			writers[i], e = objectstream.NewTempPutStream(locateInfo[i], fmt.Sprintf("%s.%d", hash, i), perShard)
			if e != nil {
				return nil, e
			}
		}
	}

	// readers、writers 可以同时进行读取和修复功能
	dec := NewDecoder(readers, writers, size)
	return &RSGetStream{dec}, nil
}

func (s *RSGetStream) Close() {
	for i := range s.writers {
		if s.writers[i] != nil {
			// RSGetStream 的 writer 表示丢失的分片修复,临时对象被转正
			s.writers[i].(*objectstream.TempPutStream).Commit(true)
		}
	}
}

// 两个参数分别表示需要跳过的字节数和起跳点
func (s *RSGetStream) Seek(offset int64, whence int) (int64, error) {
	if whence != io.SeekCurrent {
		panic("only support SeekCurrent")
	}
	if offset < 0 {
		panic("only support forward seek")
	}
	for offset != 0 {
		length := int64(BLOCK_SIZE)
		if offset < length {
			length = offset
		}
		buf := make([]byte, length)
		io.ReadFull(s, buf)
		offset -= length
	}
	return offset, nil
}
