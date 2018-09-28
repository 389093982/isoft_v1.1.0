package main

import (
	"fmt"
	"isoft/isoft/business/monitor"
	"isoft/isoft/common/logutil"
	"isoft/isoft_storage/cfg"
	"isoft/isoft_storage/dataServer/heartbeat"
	"isoft/isoft_storage/dataServer/locate"
	"isoft/isoft_storage/dataServer/objects"
	"isoft/isoft_storage/dataServer/temp"
	"net/http"
	"os"
	"strings"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			logutil.Errorln(err)
		}
	}()

	// 启动前初始化参数,参数初始化失败会终止程序
	cfg.InitConfigWithOsArgs(os.Args)

	// 应用启动时对节点本地磁盘上的对象进行定位的,缓存对象定位信息,防止过于频繁的磁盘访问
	locate.CollectObjects()

	// 使用协程,主要用于数据服务节点向所有接口服务节点通知自身的存在,把本服务的监听地址发送出去,发送心跳消息
	go heartbeat.StartHeartbeat()

	// 数据服务提供数据的存储功能
	http.HandleFunc("/objects/", objects.Handler)

	http.HandleFunc("/temp/", temp.Handler)

	// 数据服务提供对象定位功能
	http.HandleFunc("/locate/", locate.Handler)

	LISTEN_ADDRESS := cfg.GetConfigValue(cfg.LISTEN_ADDRESS)

	fmt.Println(fmt.Sprintf("Start ListenAndServe address %s", LISTEN_ADDRESS))

	bind_address := string([]rune(LISTEN_ADDRESS)[strings.Index(LISTEN_ADDRESS, ":"):])

	fmt.Println(fmt.Sprintf("Start bind_address %s", bind_address))

	// 每隔 5 s 发送一次心跳检测信息给监控系统
	go monitor.RecordMonitorHeartBeatLog(cfg.GetConfigValue(cfg.ISOFT_DEPLOY_WEB), LISTEN_ADDRESS)

	if err := http.ListenAndServe(bind_address, nil); err != nil {
		logutil.Errorln(err)
	}
}
