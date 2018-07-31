package main

import (
	"../apiServer/objects"
	"flag"
	"fmt"
	"isoft/isoft_storage/cfg"
	"isoft/isoft_storage/lib/es"
	"isoft/isoft_storage/lib/utils"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		sectionSearch := flag.String("sectionSearch", os.Args[1], "sectionSearch")
		cfg.InitConfig(*sectionSearch)
	} else {
		fmt.Println("os args length error...")
	}

	files, _ := filepath.Glob(cfg.GetConfigValue(cfg.STORAGE_ROOT) + "/objects/*")

	for i := range files {
		hash := strings.Split(filepath.Base(files[i]), ".")[0]
		// 检查数据
		verify(hash)
	}
}

func verify(hash string) {
	log.Println("verify", hash)
	// 从元数据服务中获取该散列值对应的对象大小
	size, e := es.SearchHashSize(hash)
	if e != nil {
		log.Println(e)
		return
	}
	// 获取 getStream
	stream, e := objects.GetStream(hash, size)
	if e != nil {
		log.Println(e)
		return
	}
	// 计算 stream 的 hash
	d := utils.CalculateHash(stream)
	if d != hash {
		// hash 值不一致则需要记录 log 输出错误报告
		log.Printf("object hash mismatch, calculated=%s, requested=%s", d, hash)
	}
	stream.Close()
}
