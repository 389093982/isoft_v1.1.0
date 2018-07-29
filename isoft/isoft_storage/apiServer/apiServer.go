package main

import (
	"./heartbeat"
	"./locate"
	"./objects"
	"./temp"
	"./versions"
	"flag"
	"fmt"
	"isoft/isoft_storage/cfg"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		sectionSearch := flag.String("sectionSearch", os.Args[1], "sectionSearch")
		cfg.InitConfig(*sectionSearch)
	} else {
		fmt.Println("os args length error...")
	}

	// 使用协程,主要用于接收数据服务节点发送过来的心跳消息
	go heartbeat.ListenHeartbeat()

	// 接口服务提供对外的 REST 接口,接口服务作为 HTTP 客户端向数据服务发送请求
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/temp/", temp.Handler)

	// 接口服务对外提供定位功能,向数据服务节点群发定位消息并接收反馈消息
	http.HandleFunc("/locate/", locate.Handler)

	// 提供对象的列表功能,用于查询所有对象或指定对象的所有版本
	http.HandleFunc("/versions/", versions.Handler)

	fmt.Println(fmt.Sprintf("Start ListenAndServe address %s", cfg.GetConfigValue(cfg.LISTEN_ADDRESS)))

	log.Fatal(http.ListenAndServe(cfg.GetConfigValue(cfg.LISTEN_ADDRESS), nil))
}
