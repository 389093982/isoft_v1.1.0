package main

import (
	"fmt"
	"isoft/isoft/business/monitor"
	"isoft/isoft/common/logutil"
	"isoft/isoft_storage/fileserver/cfg"
	"isoft/isoft_storage/fileserver/upload"
	"net/http"
	"os"
	"strings"
)

func main()  {
	// 启动前初始化参数,参数初始化失败会终止程序
	cfg.InitConfigWithOsArgs(os.Args)

	http.HandleFunc("/upload", upload.UploadHandler)

	STORAGE_ROOT := cfg.GetConfigValue(cfg.STORAGE_ROOT)
	http.Handle("/download/", http.StripPrefix("/download/", http.FileServer(http.Dir(STORAGE_ROOT))))

	LISTEN_ADDRESS := cfg.GetConfigValue(cfg.LISTEN_ADDRESS)

	logutil.Infoln(fmt.Sprintf("Start ListenAndServe address %s", LISTEN_ADDRESS))

	bind_address := string([]rune(LISTEN_ADDRESS)[strings.Index(LISTEN_ADDRESS, ":"):])

	logutil.Infoln(fmt.Sprintf("Start bind_address %s", bind_address))
	// 每隔 5 s 发送一次心跳检测信息给监控系统
	go monitor.RecordMonitorHeartBeatLog(cfg.GetConfigValue(cfg.ISOFT_DEPLOY_WEB), LISTEN_ADDRESS)

	if err := http.ListenAndServe(bind_address, nil); err != nil {
		logutil.Errorln(err)
	}
}
