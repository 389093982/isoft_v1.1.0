package main

import (
	"flag"
	"fmt"
	"isoft/isoft_storage/cfg"
	"isoft/isoft_storage/lib/es"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// 删除没有元数据引用的对象数据
func main() {
	if len(os.Args) > 1 {
		sectionSearch := flag.String("sectionSearch", os.Args[1], "sectionSearch")
		cfg.InitConfig(*sectionSearch)
	} else {
		fmt.Println("os args length error...")
	}

	files, _ := filepath.Glob(cfg.GetConfigValue(cfg.STORAGE_ROOT) + "/objects/*")

	for i := range files {
		// 获取对象的 hash 值
		hash := strings.Split(filepath.Base(files[i]), ".")[0]
		// 判断元数据中是否有对象的 hash 值
		hashInMetadata, e := es.HasHash(hash)
		if e != nil {
			log.Println(e)
			return
		}
		if !hashInMetadata {
			del(hash)
		}
	}
}

func del(hash string) {
	log.Println("delete", hash)
	url := "http://" + cfg.GetConfigValue(cfg.LISTEN_ADDRESS) + "/objects/" + hash
	// 根据 hash 值删除对象
	request, _ := http.NewRequest("DELETE", url, nil)
	client := http.Client{}
	client.Do(request)
}
