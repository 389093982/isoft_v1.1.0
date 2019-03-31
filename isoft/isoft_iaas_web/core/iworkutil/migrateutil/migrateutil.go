package migrateutil

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"isoft/isoft/common/hashutil"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
)

type MigrateExecutor struct {
	Dsn string // dsn 连接串
	db  *sql.DB
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
	versionTable := `CREATE TABLE IF NOT EXISTS migrate_version (id INT(20) PRIMARY KEY AUTO_INCREMENT,hash CHAR(200),sql_detail TEXT,created_time datetime);`
	this.ExecSQL(versionTable)
}

func (this *MigrateExecutor) ExecSQL(sql string, args ...interface{}) {
	stmt, err := this.db.Prepare(sql)
	checkError(err)
	_, err = stmt.Exec(args...)
	checkError(err)
}

func (this *MigrateExecutor) migrate() {
	migrates, err := iwork.QueryAllMigrate()
	checkError(err)
	for _, migrate := range migrates {
		executeSqls := strings.Split(migrate.TableMigrateSql, ";")
		for _, executeSql := range executeSqls {
			this.ExecSQL(executeSql)
			versionRecord := `INSERT INTO migrate_version(HASH,SQL_DETAIL,CREATED_TIME) VALUES (?,?,NOW());`
			this.ExecSQL(versionRecord, hashutil.CalculateHashWithString(executeSql), executeSql)
		}
	}
}

func MigrateToDB(dsn string) (err error) {
	defer func() {
		if err1 := recover(); err1 != nil {
			err = err1.(error)
		}
	}()
	executor := &MigrateExecutor{
		Dsn: dsn,
	}
	executor.ping()
	executor.initial()
	executor.migrate()
	return nil
}
