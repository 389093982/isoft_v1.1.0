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
		(id INT(20) PRIMARY KEY AUTO_INCREMENT, tracking_id CHAR(200), flag CHAR(200), hash CHAR(200),sql_detail TEXT, tracking_detail TEXT, created_time datetime);`
	this.ExecSQL(versionTable)
}

func (this *MigrateExecutor) ExecSQL(sql string, args ...interface{}) (rs sql.Result, err error) {
	stmt, err := this.db.Prepare(sql)
	checkError(err)
	rs, err = stmt.Exec(args...)
	checkError(err)
	return
}

func (this *MigrateExecutor) record(flag, hash, sql, tracking_detail string) {
	recordLog := `INSERT INTO migrate_version(tracking_id,flag,hash,sql_detail,tracking_detail, created_time) VALUES (?,?,?,?,?,NOW());`
	this.ExecSQL(recordLog, this.TrackingId, flag, hash, sql, tracking_detail)
}

func (this *MigrateExecutor) migrate() {
	migrates, err := iwork.QueryAllMigrate()
	checkError(err)
	for _, migrate := range migrates {
		this.migrateOne(migrate)
	}
}

func (this *MigrateExecutor) checkExecuted(hash string) bool {
	sql := `SELECT COUNT(*) FROM migrate_version WHERE hash = ?`
	rs, _ := this.ExecSQL(sql, hash)
	if count, err := rs.RowsAffected(); err == nil && count > 0 {
		return true
	}
	return false
}

func (this *MigrateExecutor) migrateOne(migrate iwork.TableMigrate) {
	hash := hashutil.CalculateHashWithString(migrate.TableMigrateSql)
	if !this.checkExecuted(hash) {
		// 每次迁移都有可能有多个执行 sql
		executeSqls := strings.Split(migrate.TableMigrateSql, ";")
		for _, executeSql := range executeSqls {
			detailHash := hashutil.CalculateHashWithString(executeSql)
			if !this.checkExecuted(detailHash) {
				this.ExecSQL(executeSql)
				this.record("true", detailHash, executeSql, "")
			}
		}
		// 计算hash 值
		this.record("true", hash, migrate.TableMigrateSql, "")
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
			executor.record("false", "", "", err.Error())
		}
	}()

	executor.ping()
	executor.initial()
	executor.migrate()
	return nil
}
