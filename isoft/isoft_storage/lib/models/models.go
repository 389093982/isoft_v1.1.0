package models

type LocateMessage struct {
	Addr string
	Id   int
}

// 元数据信息
type Metadata struct {
	Name    string
	Version int
	Size    int64
	Hash    string
}
