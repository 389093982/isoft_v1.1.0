package memory

import (
	"fmt"
	"isoft/isoft_iaas_web/models/iwork"
)

var memorycaches = make(map[string]*MemoryCache, 0)

type ItemMemoryCache struct {
	ItemMemoryCacheMap map[string]interface{}
}

type MemoryCache struct {
	TrackingId      string
	MemoryCacheData map[string]interface{}
}

// 向数据中心缓存数据
func (this *MemoryCache) CacheData(paramName string, paramValue interface{}) {
	this.MemoryCacheData[paramName] = paramValue
	iwork.InsertRunLogDetail(this.TrackingId, fmt.Sprintf("[%s]cache memory data: %s:%v", this.TrackingId, paramName, paramValue))
}

// 从 MemoryCache 获取数据
func (this *MemoryCache) GetData(paramName string) interface{} {
	return this.MemoryCacheData[paramName]
}

// 注册 MemoryCache
func RegistMemoryCache(trackingId string) {
	memorycaches[trackingId] = &MemoryCache{
		TrackingId:      trackingId,
		MemoryCacheData: make(map[string]interface{}, 0),
	}
}

// 注销 MemoryCache
func UnRegistMemoryCache(trackingId string) {
	delete(memorycaches, trackingId)
}

// 获取 MemoryCache
func GetMemoryCache(trackingId string) *MemoryCache {
	cache := memorycaches[trackingId]
	if cache != nil {
		return cache
	}
	RegistMemoryCache(trackingId)
	return memorycaches[trackingId]
}
