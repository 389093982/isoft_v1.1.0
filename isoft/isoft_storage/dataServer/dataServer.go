package main

import (
	"./heartbeat"
	"./locate"
	"./objects"
	"./temp"
	"fmt"
	"isoft/isoft_storage/cfg"
	"log"
	"net/http"
	"os"
)

func main() {
	// 启动前初始化参数,参数初始化失败会终止程序
	cfg.InitConfigWithOsArgs(os.Args)

	// 应用启动时对节点本地磁盘上的对象进行定位的,缓存对象定位信息,防止过于频繁的磁盘访问
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
