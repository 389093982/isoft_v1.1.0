package logger

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"isoft/isoft/common/apppath"
	"isoft/isoft/common/fileutil"
	"os"
)

func ConfigureLogInfo() {
	var logDir string
	if beego.BConfig.RunMode == "dev" || beego.BConfig.RunMode == "local" {
		logDir = "../../../isoft_iaas_web_log"
	} else {
		// 日志文件所在目录
		logDir = fileutil.ChangeToLinuxSeparator(apppath.GetAPPRootPath() + "/isoft_iaas_web_log")
	}
	if ok, _ := fileutil.PathExists(logDir); !ok {
		err := os.MkdirAll(logDir, os.ModePerm)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	// 控制台输出
	logs.SetLogger(logs.AdapterConsole)
	// 多文件输出
	logs.SetLogger(logs.AdapterMultiFile,
		`{"filename":"`+logDir+`/isoft_iaas_web.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)
	// 输出文件名和行号
	logs.EnableFuncCallDepth(true)
	// 异步输出日志
	logs.Async(1e3)
}
