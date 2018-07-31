package main

import (
	"flag"
	"fmt"
	"isoft/isoft_storage/cfg"
	"isoft/isoft_storage/lib/es"
	"log"
	"os"
)

const MIN_VERSION_COUNT = 5

func main() {
	if len(os.Args) > 1 {
		sectionSearch := flag.String("sectionSearch", os.Args[1], "sectionSearch")
		cfg.InitConfig(*sectionSearch)
	} else {
		fmt.Println("os args length error...")
	}

	// 查询所有版本数量大于等于 6 的对象
	buckets, e := es.SearchVersionStatus(MIN_VERSION_COUNT + 1)
	if e != nil {
		log.Println(e)
		return
	}
	for i := range buckets {
		bucket := buckets[i]
		// 循环遍历每一个 bucket,从该对象当前最小的版本号开始一一删除,直到最后还剩 5 个版本
		for v := 0; v < bucket.Doc_count-MIN_VERSION_COUNT; v++ {
			// 根据对象名称,删除对象指定版本
			es.DelMetadata(bucket.Key, v+int(bucket.Min_version.Value))
		}
	}
}
