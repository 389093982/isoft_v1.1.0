package models

type LocateMessage struct {
	Id      int64
	Addr    string // 服务地址,机器 ip 和端口
	ShardId int    // 分片 id
}

// 元数据信息
type Metadata struct {
	Name    string
	Version int
	Size    int64
	Hash    string
}
