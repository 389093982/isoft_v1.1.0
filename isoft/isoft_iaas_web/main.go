package main

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql" // _ 的作用,并不需要把整个包都导入进来,仅仅是是希望它执行init()函数而已
	"isoft/isoft_iaas_web/startup/db"
	"isoft/isoft_iaas_web/imodules/misso"
	"isoft/isoft_iaas_web/startup/logger"
	_ "isoft/isoft_iaas_web/routers"
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


