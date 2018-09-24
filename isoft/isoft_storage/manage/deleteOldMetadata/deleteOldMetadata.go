package main

import (
	"isoft/isoft_storage/cfg"
	"isoft/isoft_storage/lib"
	"log"
	"os"
)

const MIN_VERSION_COUNT = 5

func main() {
	// 启动前初始化参数,参数初始化失败会终止程序
	cfg.InitConfigWithOsArgs(os.Args)

	// 查询所有版本数量大于等于 MIN_VERSION_COUNT + 1 的对象
	proxy := &lib.MetaDataProxy{}
	versionMap, err := proxy.SearchVersionStatus(MIN_VERSION_COUNT + 1)
	if err != nil {
		log.Println(err)
		return
	}
	// 返回值 key 为对象名, value 为对象现有版本数量、最小版本信息
	for name, versionInfo := range versionMap {
		// 循环遍历每一个 bucket,从该对象当前最小的版本号开始一一删除,直到最后还剩 5 个版本
		for v := 0; v < versionInfo[0]-MIN_VERSION_COUNT; v++ {
			// 根据对象名称,删除对象指定版本
			proxy.DelMetadata(name, v+int(versionInfo[1]))
		}
	}
}
