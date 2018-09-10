package main

import (
	"./heartbeat"
	"./locate"
	"./objects"
	"./temp"
	"./versions"
	"fmt"
	"isoft/isoft_storage/cfg"
	"log"
	"net/http"
	"os"
)

func main() {
	// 启动前初始化参数,参数初始化失败会终止程序
	cfg.InitConfigWithOsArgs(os.Args)

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
