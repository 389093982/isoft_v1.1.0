package rs

import (
	"github.com/klauspost/reedsolomon"
	"io"
)

type encoder struct {
	writers []io.Writer         // io.Writer 数组 writers
	enc     reedsolomon.Encoder // RS 编解码的开源库接口(reedsolomon.New(DATA_SHARDS, PARITY_SHARDS) 包含数据片和冗余片数量)
	cache   []byte              // 输入数据缓存的字节数组 cache
}

func NewEncoder(writers []io.Writer) *encoder {
	// 指定 数据片和冗余片数量
	enc, _ := reedsolomon.New(DATA_SHARDS, PARITY_SHARDS)
	return &encoder{writers, enc, nil}
}

func (e *encoder) Write(p []byte) (n int, err error) {
	// 并不是真正的写入,而是将 p 中待写入的数据以块的形式放入缓存
	length := len(p)
	current := 0
	for length != 0 {
		next := BLOCK_SIZE - len(e.cache) // 最大缓存数 - 已有缓存数 = 剩余可缓存数
		if next > length {
			next = length
		}
		e.cache = append(e.cache, p[current:current+next]...)
		if len(e.cache) == BLOCK_SIZE { // 达到最大缓存字节数则立即 Flush
			e.Flush()
		}
		current += next
		length -= next
	}
	return len(p), nil
}

func (e *encoder) Flush() {
	if len(e.cache) == 0 {
		return
	}
	// Split 方法将缓存的数据切成 4 个数据片
	shards, _ := e.enc.Split(e.cache)
	// Encode 方法生成两个校验片(冗余片)
	e.enc.Encode(shards)
	for i := range shards {
		// 循环将分片的数据依次写入 writers 并清空缓存
		e.writers[i].Write(shards[i])
	}
	e.cache = []byte{}
}
