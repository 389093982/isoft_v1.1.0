package rs

const (
	DATA_SHARDS     = 4                             		// 数据分片
	PARITY_SHARDS   = 2                             		// 冗余分片
	ALL_SHARDS      = DATA_SHARDS + PARITY_SHARDS   		// 总分片数
	BLOCK_PER_SHARD = 1024*20                          	// 每个数据片缓存的上限 BLOCK_PER_SHARD 个字节
	BLOCK_SIZE      = BLOCK_PER_SHARD * DATA_SHARDS 	 	// 4 个数据片最大缓存字节 BLOCK_SIZE 字节
)
