package db

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"isoft/isoft/common/flyway"
	"isoft/isoft_iaas_web/imodules"
	"isoft/isoft_iaas_web/imodules/milearning"
	"isoft/isoft_iaas_web/imodules/misso"
	"isoft/isoft_iaas_web/imodules/miwork"
	"isoft/isoft_iaas_web/startup/db/imysql"
	"isoft/isoft_iaas_web/startup/db/isqlite3"
)

// 数据库同步模式,支持 FLYWAY 和 AUTO
const RunSyncdbMode = "AUTO"

func ConfigureDBInfo() {
	if imodules.CheckModule("iwork") {
		//imysql.RegisterDBForMysql()
		isqlite3.RegisterDBForSqlite3()
	}else{
		imysql.RegisterDBForMysql()
	}

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
	registerModel()

	if RunSyncdbMode == "FLYWAY" {
		// ilearning 模块
		if imodules.CheckModule("ilearning") {
			flyway.MigrateToDB(imysql.Dsn, "./conf/migrations/migrations.sql")
		}
		// sso 模块
		if imodules.CheckModule("sso") {
			flyway.MigrateToDB(imysql.Dsn, "./conf/migrations/sso_migrations.sql")
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