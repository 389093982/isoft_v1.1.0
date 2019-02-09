package iworkdata

var datastores = make(map[string]*DataStore, 0)

type DataStore struct {

}

// 注册数据中心
func RegistDataStore(trackingId string) {
	datastores[trackingId] = &DataStore{}
}

// 注销数据中心
func UnRegistDataStore(trackingId string) {
	delete(datastores, trackingId)
}

// 获取数据中心
func GetDataSource(trackingId string) *DataStore {
	return datastores[trackingId]
}