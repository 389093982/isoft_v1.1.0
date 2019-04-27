package iworknode

import (
	"isoft/isoft_iaas_web/core/iworkconst"
	"isoft/isoft_iaas_web/core/iworkdata/memory"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworkmodels"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
)

type MemoryMapCacheNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

// 获取最顶层 trackingId
func getOriginalTrackingId(trackingId string) string {
	if strings.Contains(trackingId, ".") {
		return trackingId[:strings.Index(trackingId, ".")]
	}
	return trackingId
}

func getMemoryCache(trackingId string, tmpDataMap map[string]interface{}) *memory.MemoryCache {
	if lifecycle, ok := tmpDataMap[iworkconst.STRING_PREFIX+"lifecycle?"].(string); ok && strings.TrimSpace(lifecycle) != "" {
		// lifecycle 有值表示跨流程共享内存,注册的 trackingId 得使用最顶层 trackingId
		return memory.GetMemoryCache(getOriginalTrackingId(trackingId))
	} else {
		return memory.GetMemoryCache(trackingId)
	}
}

func (this *MemoryMapCacheNode) Execute(trackingId string) {
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, this.DataStore)
	memoryCache := getMemoryCache(trackingId, tmpDataMap)
	if cachemap_key_get, ok := tmpDataMap[iworkconst.STRING_PREFIX+"cachemap_key_get?"].(string); ok {
		// 往 MemoryCache 中取值
		key := iworkconst.STRING_PREFIX + "cachemap_val_get"
		value := memoryCache.GetData(tmpDataMap[iworkconst.STRING_PREFIX+"cachemap_name"].(string) + "_" + cachemap_key_get)
		this.DataStore.CacheDatas(this.WorkStep.WorkStepName, map[string]interface{}{key: value})
	} else if cachemap_key_put, ok := tmpDataMap[iworkconst.STRING_PREFIX+"cachemap_key_put?"].(string); ok {
		// 往 MemoryCache 中放值
		key := tmpDataMap[iworkconst.STRING_PREFIX+"cachemap_name"].(string) + "_" + cachemap_key_put
		value := tmpDataMap[iworkconst.STRING_PREFIX+"cachemap_val_put?"].(string)
		memoryCache.CacheData(key, value)
	}
}

func (this *MemoryMapCacheNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: []string{iworkconst.STRING_PREFIX + "lifecycle?", "内存map存储的生命周期,默认是当前流程,有值的话则表示本次运行时机(即可以跨流程),map不存在会自动创建,运行完后会自动销毁!"},
		2: []string{iworkconst.STRING_PREFIX + "cachemap_name", "存储的map名称"},
		3: []string{iworkconst.STRING_PREFIX + "cachemap_key_get?", "存储的键值"},
		4: []string{iworkconst.STRING_PREFIX + "cachemap_key_put?", "存储的键值"},
		5: []string{iworkconst.STRING_PREFIX + "cachemap_val_put?", "存储的value值"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *MemoryMapCacheNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return schema.BuildParamOutputSchemaWithSlice([]string{iworkconst.STRING_PREFIX + "cachemap_val_get"})
}
