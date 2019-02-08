package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // _ 的作用,并不需要把整个包都导入进来,仅仅是是希望它执行init()函数而已
	"isoft/isoft/common/apppath"
	"isoft/isoft/common/fileutil"
	"isoft/isoft/common/flyway"
	"isoft/isoft/ssofilter"
	"isoft/isoft_iaas_web/models/cms"
	"isoft/isoft_iaas_web/models/common"
	"isoft/isoft_iaas_web/models/iblog"
	"isoft/isoft_iaas_web/models/ifile"
	"isoft/isoft_iaas_web/models/ilearning"
	"isoft/isoft_iaas_web/models/iquartz"
	"isoft/isoft_iaas_web/models/iresource"
	"isoft/isoft_iaas_web/models/iwork"
	"isoft/isoft_iaas_web/models/monitor"
	"isoft/isoft_iaas_web/models/share"
	"isoft/isoft_iaas_web/models/sso"
	_ "isoft/isoft_iaas_web/routers"
	"isoft/isoft_iaas_web/task"
	"net/url"
	"os"
	"strings"
)

// 数据库连接串
var dsn string

// 数据库同步模式,支持 FLYWAY 和 AUTO
const RunSyncdbMode = "AUTO"

func init() {
	initLog()
	initDB()
}

func initLog() {
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

func initDB() {
	dbhost := beego.AppConfig.String("db.host")
	dbport := beego.AppConfig.String("db.port")
	dbname := beego.AppConfig.String("db.name")
	dbuser := beego.AppConfig.String("db.user")
	dbpass := beego.AppConfig.String("db.pass")
	timezone := beego.AppConfig.String("db.timezone")

	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?allowNativePasswords=true&charset=utf8", dbuser, dbpass, dbhost, dbport, dbname)

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

	if RunSyncdbMode == "FLYWAY" {
		// ilearning 模块
		if strings.Contains(beego.AppConfig.String("open.moudles"), "ilearning") {
			flyway.MigrateToDB(dsn, "./conf/migrations/migrations.sql")
		}
		// sso 模块
		if strings.Contains(beego.AppConfig.String("open.moudles"), "sso") {
			flyway.MigrateToDB(dsn, "./conf/migrations/sso_migrations.sql")
		}
	} else {
		createTable()
	}
}

func registerModel() {
	// ilearning 模块
	if strings.Contains(beego.AppConfig.String("open.moudles"), "ilearning") {
		orm.RegisterModel(new(iblog.Catalog))
		orm.RegisterModel(new(iblog.Blog))

		orm.RegisterModel(new(ilearning.Course))
		orm.RegisterModel(new(ilearning.CourseVideo))
		orm.RegisterModel(new(ilearning.Favorite))
		orm.RegisterModel(new(ilearning.CommentTheme))
		orm.RegisterModel(new(ilearning.CommentReply))
		orm.RegisterModel(new(ilearning.Note))

		orm.RegisterModel(new(ifile.IFile))

		orm.RegisterModel(new(cms.Configuration))
		orm.RegisterModel(new(cms.CommonLink))

		orm.RegisterModel(new(share.Share))

		orm.RegisterModel(new(common.History))

		orm.RegisterModel(new(monitor.HeartBeat2))
		orm.RegisterModel(new(monitor.HeartBeatDetail))
	}
	// sso 模块
	if strings.Contains(beego.AppConfig.String("open.moudles"), "sso") {
		orm.RegisterModel(new(sso.User))
		orm.RegisterModel(new(sso.AppRegister))
		orm.RegisterModel(new(sso.LoginRecord))
		orm.RegisterModel(new(sso.UserToken))
	}

	orm.RegisterModel(new(iquartz.CronMeta))
	orm.RegisterModel(new(iresource.Resource))
	orm.RegisterModel(new(iwork.Work))
	orm.RegisterModel(new(iwork.WorkStep))
	orm.RegisterModel(new(iwork.RunLogRecord))
	orm.RegisterModel(new(iwork.RunLogDetail))
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
	beego.InsertFilter("/api/*", beego.BeforeExec, ssoFilterFunc)

	// 开启定时任务
	task.StartCronTask()
	// 执行 iquartz 组件初始化任务
	task.StartIQuartzInitialTask()

	beego.Run()
}

func ssoFilterFunc(ctx *context.Context) {
	filter := new(ssofilter.LoginFilter)
	filter.LoginWhiteList = &[]string{"/api/sso/user/login", "/api/sso/user/regist", "/api/sso/user/checkOrInValidateTokenString"}
	filter.LoginUrl = ctx.Input.URL()
	filter.Ctx = ctx
	filter.SsoAddress = beego.AppConfig.String("isoft.sso.web.addr")
	filter.ErrorFunc = func() {
		filter.Ctx.ResponseWriter.WriteHeader(401)
	}
	filter.Filter()
}
