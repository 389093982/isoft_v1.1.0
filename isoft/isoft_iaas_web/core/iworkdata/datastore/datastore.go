package datastore

import (
	"fmt"
	"isoft/isoft_iaas_web/core/iworklog"
	"strings"
)

type DataNodeStore struct {
	NodeOutputDataMap map[string]interface{} // 当前节点的输出参数 map
}

type DataStore struct {
	TrackingId       string
	logwriter        *iworklog.CacheLoggerWriter
	DataNodeStoreMap map[string]*DataNodeStore
}

func (this *DataStore) CacheDatas(nodeName string, paramMap map[string]interface{}) {
	logs := make([]string, 0)
	this.CacheByteData(nodeName, "__output__", paramMap)
	for paramName, paramValue := range paramMap {
		this.CacheByteData(nodeName, paramName, paramValue)
		log := fmt.Sprintf("<span style='color:#FF99FF;'> [%s] </span>"+
			"<span style='color:#6633FF;'> cache data for $%s.%s: </span>"+
			"<span style='color:#CC0000;'> %v </span>", this.TrackingId, nodeName, paramName, paramValue)
		logs = append(logs, log)
	}
	this.logwriter.Write(this.TrackingId, strings.Join(logs, "<br/>"))
}

// 向数据中心缓存数据
func (this *DataStore) CacheData(nodeName, paramName string, paramValue interface{}) {
	this.CacheDatas(nodeName, map[string]interface{}{paramName: paramValue})
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
func InitDataStore(trackingId string, logwriter *iworklog.CacheLoggerWriter) *DataStore {
	return &DataStore{
		TrackingId:       trackingId,
		logwriter:        logwriter,
		DataNodeStoreMap: make(map[string]*DataNodeStore, 0),
	}
}
