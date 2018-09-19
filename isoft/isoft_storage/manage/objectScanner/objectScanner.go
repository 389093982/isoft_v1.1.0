package main

import (
	"isoft/isoft_storage/apiServer/objects"
	"isoft/isoft_storage/cfg"
	"isoft/isoft_storage/lib"
	"isoft/isoft_storage/lib/utils"
	"log"
	"os"
)

func main() {
	// 启动前初始化参数,参数初始化失败会终止程序
	cfg.InitConfigWithOsArgs(os.Args)
}

func verify(hash string) {
	log.Println("verify", hash)
	// 从元数据服务中获取该散列值对应的对象大小
	size, e := lib.MetaDataProxy{}.SearchHashSize(hash)
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
