package migrateutil

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"isoft/isoft/common/hashutil"
	"isoft/isoft/common/stringutil"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
)

type MigrateExecutor struct {
	Dsn        string // dsn 连接串
	db         *sql.DB
	TrackingId string
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func (this *MigrateExecutor) ping() {
	if this.Dsn == "" {
		panic("empty dsn error...")
	}
	// 建立连接
	db, err := sql.Open("mysql", this.Dsn)
	checkError(err)
	this.db = db
}

// 建立迁移文件版本管理表
func (this *MigrateExecutor) initial() {
	versionTable := `CREATE TABLE IF NOT EXISTS migrate_version 
		(id INT(20) PRIMARY KEY AUTO_INCREMENT, tracking_id CHAR(200), flag CHAR(200), hash CHAR(200),sql_detail TEXT,created_time datetime);`
	this.ExecSQL(versionTable)
}

func (this *MigrateExecutor) ExecSQL(sql string, args ...interface{}) {
	stmt, err := this.db.Prepare(sql)
	checkError(err)
	_, err = stmt.Exec(args...)
	checkError(err)
}

func (this *MigrateExecutor) record(flag, hash, sql string) {
	recordLog := `INSERT INTO migrate_version(tracking_id,flag,hash,sql_detail,created_time) VALUES (?,?,?,?,NOW());`
	this.ExecSQL(recordLog, this.TrackingId, flag, hash, sql)
}

func (this *MigrateExecutor) migrate() {
	migrates, err := iwork.QueryAllMigrate()
	checkError(err)
	for _, migrate := range migrates {
		executeSqls := strings.Split(migrate.TableMigrateSql, ";")
		for _, executeSql := range executeSqls {
			this.ExecSQL(executeSql)
			this.record("true", hashutil.CalculateHashWithString(executeSql), executeSql)
		}
	}
}

func MigrateToDB(dsn string) (err error) {
	executor := &MigrateExecutor{
		Dsn:        dsn,
		TrackingId: stringutil.RandomUUID(),
	}
	defer func() {
		if err1 := recover(); err1 != nil {
			err = err1.(error)
			executor.record("false", "", "")
		}
	}()

	executor.ping()
	executor.initial()
	executor.migrate()
	return nil
}
