package db

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // _ 的作用,并不需要把整个包都导入进来,仅仅是是希望它执行init()函数而已
	"isoft/isoft/common/flyway"
	"isoft/isoft_iaas_web/imodules"
	"isoft/isoft_iaas_web/imodules/milearning"
	"isoft/isoft_iaas_web/imodules/misso"
	"isoft/isoft_iaas_web/imodules/miwork"
	"net/url"
)

// 数据库连接串
var dsn string
// 数据库同步模式,支持 FLYWAY 和 AUTO
const RunSyncdbMode = "AUTO"

func ConfigureDBInfo() {
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
		if imodules.CheckModule("ilearning") {
			flyway.MigrateToDB(dsn, "./conf/migrations/migrations.sql")
		}
		// sso 模块
		if imodules.CheckModule("sso") {
			flyway.MigrateToDB(dsn, "./conf/migrations/sso_migrations.sql")
		}
	} else {
		createTable()
	}
}

func registerModel() {
	milearning.RegisterModel()
	misso.RegisterModel()
	miwork.RegisterModel()
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