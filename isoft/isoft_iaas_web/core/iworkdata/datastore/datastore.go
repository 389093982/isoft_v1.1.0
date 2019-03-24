package datastore

import (
	"fmt"
	"isoft/isoft_iaas_web/models/iwork"
)

type DataNodeStore struct {
	NodeOutputDataMap map[string]interface{} // 当前节点的输出参数 map
}

type DataStore struct {
	TrackingId       string
	DataNodeStoreMap map[string]*DataNodeStore
}

// 向数据中心缓存数据
func (this *DataStore) CacheData(nodeName, paramName string, paramValue interface{}) {
	this.CacheByteData(nodeName, paramName, paramValue)
	iwork.InsertRunLogDetail(this.TrackingId,
		fmt.Sprintf("<span style='color:#FF99FF;'> [%s] </span>"+
			"<span style='color:#6633FF;'> cache data for $%s.%s: </span>"+
			"<span style='color:#CC0000;'> %v </span>", this.TrackingId, nodeName, paramName, paramValue))
}

// 存储字节数据,不用记录日志
func (this *DataStore) CacheByteData(nodeName, paramName string, paramValue interface{}) {
	// 为当前 nodeName 绑定 DataNodeStore 数据空间
	if _, ok := this.DataNodeStoreMap[nodeName]; !ok {
		this.DataNodeStoreMap[nodeName] = &DataNodeStore{
			NodeOutputDataMap: make(map[string]interface{}, 0),
		}
	}
	// 存数据
	dataNodeStore := this.DataNodeStoreMap[nodeName]
	dataNodeStore.NodeOutputDataMap[paramName] = paramValue
}

// 从数据中心获取数据
func (this *DataStore) GetData(nodeName, paramName string) interface{} {
	store := this.DataNodeStoreMap[nodeName]
	if store == nil {
		return nil
	}
	return this.DataNodeStoreMap[nodeName].NodeOutputDataMap[paramName]
}

// 获取数据中心
func InitDataStore(trackingId string) *DataStore {
	return &DataStore{
		TrackingId:       trackingId,
		DataNodeStoreMap: make(map[string]*DataNodeStore, 0),
	}
}
