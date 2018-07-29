package main

import (
	"./heartbeat"
	"./locate"
	"./objects"
	"./temp"
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

	locate.CollectObjects()

	// 使用协程,主要用于数据服务节点向所有接口服务节点通知自身的存在,把本服务的监听地址发送出去,发送心跳消息
	go heartbeat.StartHeartbeat()

	// 主要用于接收和处理来自接口服务节点发送过来的定位请求,实际定位对象的存储位置
	go locate.StartLocate()

	// 数据服务提供数据的存储功能
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/temp/", temp.Handler)

	fmt.Println(fmt.Sprintf("Start ListenAndServe address %s", cfg.GetConfigValue(cfg.LISTEN_ADDRESS)))

	log.Fatal(http.ListenAndServe(cfg.GetConfigValue(cfg.LISTEN_ADDRESS), nil))
}
