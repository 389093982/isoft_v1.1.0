package iworkdata

import (
	"fmt"
	"isoft/isoft_iaas_web/models/iwork"
)

var datastores = make(map[string]*DataStore, 0)

type DataNodeStore struct {
	NodeOutputDataMap 		map[string]interface{}		// 当前节点的输出参数 map
} 

type DataStore struct {
	TrackingId	 string
	nodeStoreMap map[string]*DataNodeStore
}

// 向数据中心缓存数据
func (this *DataStore) CacheData(nodeName, paramName string, paramValue interface{})  {
	// 为当前 nodeName 绑定 DataNodeStore 数据空间
	if _,ok := this.nodeStoreMap[nodeName];!ok{
		this.nodeStoreMap[nodeName] = &DataNodeStore{
			NodeOutputDataMap:make(map[string]interface{},5),
		}
	}
	// 存数据
	dataNodeStore := this.nodeStoreMap[nodeName]
	dataNodeStore.NodeOutputDataMap[paramName] = paramValue
	iwork.InsertRunLogDetail(this.TrackingId, fmt.Sprintf("cache data for %s:%s",paramName, paramValue))
}

// 从数据中心获取数据
func (this *DataStore) GetData(nodeName, paramName string) interface{} {
	return this.nodeStoreMap[nodeName].NodeOutputDataMap[paramName]
}

// 注册数据中心
func RegistDataStore(trackingId string) {
	datastores[trackingId] = &DataStore{
		TrackingId:trackingId,
		nodeStoreMap:make(map[string]*DataNodeStore,5),
	}
}

// 注销数据中心
func UnRegistDataStore(trackingId string) {
	delete(datastores, trackingId)
}

// 获取数据中心
func GetDataSource(trackingId string) *DataStore {
	return datastores[trackingId]
}
