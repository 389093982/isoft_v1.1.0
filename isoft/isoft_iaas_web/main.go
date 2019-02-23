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
	task.RegisterCronTask()
	beego.Run()
}


