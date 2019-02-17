package main

import (
	"github.com/astaxie/beego"
	"isoft/isoft_iaas_web/imodules/misso"
	_ "isoft/isoft_iaas_web/routers"
	"isoft/isoft_iaas_web/startup/db"
	"isoft/isoft_iaas_web/startup/logger"
	"isoft/isoft_iaas_web/task"
)

func init() {
	logger.ConfigureLogInfo()
	db.ConfigureDBInfo()
}

func main() {
	misso.RegisterISSOFilter()

	// 开启定时任务
	task.StartCronTask()
	// 执行 iquartz 组件初始化任务
	task.StartIQuartzInitialTask()

	beego.Run()
}


