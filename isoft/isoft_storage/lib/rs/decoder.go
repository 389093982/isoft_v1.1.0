package rs

import (
	"github.com/klauspost/reedsolomon"
	"io"
)

type decoder struct {
	readers   []io.Reader
	writers   []io.Writer
	enc       reedsolomon.Encoder
	size      int64
	cache     []byte
	cacheSize int
	total     int64 // 表示当前已经读取了多少字节
}

func NewDecoder(readers []io.Reader, writers []io.Writer, size int64) *decoder {
	enc, _ := reedsolomon.New(DATA_SHARDS, PARITY_SHARDS)
	return &decoder{readers, writers, enc, size, nil, 0, 0}
}

func (d *decoder) Read(p []byte) (n int, err error) {
	// cache 中没有更多数据时会调用 getData 方法获取数据
	if d.cacheSize == 0 {
		e := d.getData()
		if e != nil { // 不为 nil 时说明没能够获取足够的数据
			return 0, e
		}
	}
	// 理论上要读取的数据长度,即 p 字节数组长度
	length := len(p)
	// length 超出当前缓存的数据大小
	if d.cacheSize < length {
		// 调整理论长度为实际缓存长度
		length = d.cacheSize
	}
	// 拷贝完成之后缓存数据大小要减少 length 长度,剩余缓存大小
	d.cacheSize -= length
	// 从缓存中读取 length 长度的数据到 p 字节数组中去
	copy(p, d.cache[:length])
	// 剩余缓存数据量
	d.cache = d.cache[length:]
	return length, nil
}

func (d *decoder) getData() error {
	if d.total == d.size {
		return io.EOF
	}
	// 创建 6 个数组 shards
	shards := make([][]byte, ALL_SHARDS)
	// 创建长度为 0 的数组,存放数据丢失的分片 id,用于修复
	repairIds := make([]int, 0)
	for i := range shards {
		if d.readers[i] == nil {
			// reader 为 nil 表示不可读,即数据丢失,需要修复
			repairIds = append(repairIds, i)
		} else {
			// 可读分片,创建最大缓存量的字节数组 8000 个字节
			shards[i] = make([]byte, BLOCK_PER_SHARD)
			// 从 reader 中一次读取 8000 个字节数据保存到字节数组中去
			n, e := io.ReadFull(d.readers[i], shards[i])
			if e != nil && e != io.EOF && e != io.ErrUnexpectedEOF {
				// 读取失败
				shards[i] = nil
			} else if n != BLOCK_PER_SHARD {
				// 读取的数据长度 n 不到 8000 字节,则截取前 n 个
				shards[i] = shards[i][:n]
			}
		}
	}
	// 每个 shard 为空表示分片丢失, Reconstruct 方法尝试将被置为 nil 的 shard 恢复出来,这样每个分片就都有数据了
	e := d.enc.Reconstruct(shards)
	if e != nil {
		return e // 表示不可修复的破坏
	}
	for i := range repairIds {
		id := repairIds[i]
		// 此时 shards[id] 的数据表示需要恢复的分片数据,将其写入相应的 writer
		d.writers[id].Write(shards[id])
	}
	// 遍历 4 个数据分片
	for i := 0; i < DATA_SHARDS; i++ {
		// 当前分片的实际数据量
		shardSize := int64(len(shards[i]))
		if d.total+shardSize > d.size {
			shardSize -= d.total + shardSize - d.size
		}
		// 将当前分片中的数据添加到缓存 cache 中去
		d.cache = append(d.cache, shards[i][:shardSize]...)
		// 将当前分片中的数据大小添加到总缓存大小中去
		d.cacheSize += int(shardSize)
		// 将当前分片中的数据大小计算到当前已经读取的全部数据的大小中去,得到新的已读取数据量
		d.total += shardSize
	}
	return nil
}
