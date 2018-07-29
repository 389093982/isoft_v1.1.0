package rs

const (
	DATA_SHARDS     = 4                           // 数据分片
	PARITY_SHARDS   = 2                           // 冗余分片
	ALL_SHARDS      = DATA_SHARDS + PARITY_SHARDS // 总分片数
	BLOCK_PER_SHARD = 8000
	BLOCK_SIZE      = BLOCK_PER_SHARD * DATA_SHARDS
)
