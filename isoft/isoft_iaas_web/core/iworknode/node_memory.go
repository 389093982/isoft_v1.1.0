package iworknode

import (
	"isoft/isoft_iaas_web/core/iworkconst"
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/core/iworkdata/memory"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
)

type MemoryMapCacheNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

// 获取最顶层 trackingId
func getOriginalTrackingId(trackingId string) string {
	if strings.Contains(trackingId, "."){
		return trackingId[:strings.Index(trackingId, ".")]
	}
	return trackingId
}

func getMemoryCache(trackingId string, tmpDataMap map[string]interface{}) *memory.MemoryCache {
	if lifecycle, ok := tmpDataMap[iworkconst.STRING_PREFIX + "lifecycle?"].(string); ok && strings.TrimSpace(lifecycle) != ""{
		// lifecycle 有值表示跨流程共享内存,注册的 trackingId 得使用最顶层 trackingId
		return memory.GetMemoryCache(getOriginalTrackingId(trackingId))
	}else{
		return memory.GetMemoryCache(trackingId)
	}
}

func (this *MemoryMapCacheNode) Execute(trackingId string, skipFunc func(tmpDataMap map[string]interface{}) bool) {
	// 数据中心
	dataStore := datastore.GetDataStore(trackingId)
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, dataStore)
	memoryCache := getMemoryCache(trackingId, tmpDataMap)
	if cachemap_key_get, ok := tmpDataMap[iworkconst.STRING_PREFIX + "cachemap_key_get?"].(string); ok{
		// 往 MemoryCache 中取值
		key := iworkconst.STRING_PREFIX + "cachemap_val_get"
		value := memoryCache.GetData(tmpDataMap[iworkconst.STRING_PREFIX + "cachemap_name"].(string) + "_" + cachemap_key_get)
		dataStore.CacheData(this.WorkStep.WorkStepName, key, value)
	} else if cachemap_key_put, ok := tmpDataMap[iworkconst.STRING_PREFIX + "cachemap_key_put?"].(string); ok{
		// 往 MemoryCache 中放值
		key := tmpDataMap[iworkconst.STRING_PREFIX + "cachemap_name"].(string) + "_" + cachemap_key_put
		value := tmpDataMap[iworkconst.STRING_PREFIX + "cachemap_val_put?"].(string)
		memoryCache.CacheData(key,value)
	}
}

func (this *MemoryMapCacheNode) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	paramMap := map[int][]string{
		1:[]string{iworkconst.STRING_PREFIX + "lifecycle?","内存map存储的生命周期,默认是当前流程,有值的话则表示本次运行时机(即可以跨流程),map不存在会自动创建,运行完后会自动销毁!"},
		2:[]string{iworkconst.STRING_PREFIX + "cachemap_name","存储的map名称"},
		3:[]string{iworkconst.STRING_PREFIX + "cachemap_key_get?","存储的键值"},
		4:[]string{iworkconst.STRING_PREFIX + "cachemap_key_put?","存储的键值"},
		5:[]string{iworkconst.STRING_PREFIX + "cachemap_val_put?","存储的value值"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *MemoryMapCacheNode) GetRuntimeParamInputSchema() *schema.ParamInputSchema {
	return &schema.ParamInputSchema{}
}

func (this *MemoryMapCacheNode) GetDefaultParamOutputSchema() *schema.ParamOutputSchema {
	return schema.BuildParamOutputSchemaWithSlice([]string{iworkconst.STRING_PREFIX + "cachemap_val_get"})
}

func (this *MemoryMapCacheNode) GetRuntimeParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}

func (this *MemoryMapCacheNode) ValidateCustom() {

}