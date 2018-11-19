package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // _ 的作用,并不需要把整个包都导入进来,仅仅是是希望它执行init()函数而已
	"isoft/isoft/common/apppath"
	"isoft/isoft/common/fileutil"
	"isoft/isoft_deploy_web/deploycore/prework"
	"isoft/isoft_deploy_web/models"
	_ "isoft/isoft_deploy_web/routers"
	"net/url"
	"os"
)

func init() {
	initLog()
	initDB()
}

func initLog() {
	var logDir string
	if beego.BConfig.RunMode == "dev" || beego.BConfig.RunMode == "local" {
		logDir = "../../../isoft_deploy_web_log"
	} else {
		// 日志文件所在目录
		logDir = fileutil.ChangeToLinuxSeparator(apppath.GetAPPRootPath() + "/isoft_deploy_web_log")
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
		`{"filename":"`+logDir+`/isoft_deploy_web.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)
	// 输出文件名和行号
	logs.EnableFuncCallDepth(true)
	// 异步输出日志
	logs.Async(1e3)
}

func initDB() {
	dbhost := beego.AppConfig.String("db.host")
	dbport := beego.AppConfig.String("db.port")
	dbname := beego.AppConfig.String("db.name")
	dbuser := beego.AppConfig.String("db.user")
	dbpass := beego.AppConfig.String("db.pass")
	timezone := beego.AppConfig.String("db.timezone")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?allowNativePasswords=true&charset=utf8", dbuser, dbpass, dbhost, dbport, dbname)

	if timezone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(timezone)
	}

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dsn)
	orm.SetMaxIdleConns("default", 1000) // SetMaxIdleConns用于设置闲置的连接数
	orm.SetMaxOpenConns("default", 2000) // SetMaxOpenConns用于设置最大打开的连接数,默认值为0表示不限制

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}

	registerModel()

	createTable() // 开启自动建表
}

func registerModel() {
	orm.RegisterModel(new(models.EnvInfo))
	orm.RegisterModel(new(models.ConfigFile))
	orm.RegisterModel(new(models.ServiceInfo))
	orm.RegisterModel(new(models.TrackingTask))
	orm.RegisterModel(new(models.TrackingLog))
}

// 自动建表
func createTable() {
	name := "default"                          // 数据库别名
	force := false                             // 不强制建数据库
	verbose := true                            // 打印建表过程
	err := orm.RunSyncdb(name, force, verbose) // 建表
	if err != nil {
		beego.Error(err)
	}
}

func main() {
	prework.RunPreWork()

	beego.Run()
}
