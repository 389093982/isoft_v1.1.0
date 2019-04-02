package migrateutil

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"isoft/isoft/common/hashutil"
	"isoft/isoft/common/stringutil"
	"isoft/isoft_iaas_web/models/iwork"
	"strconv"
	"strings"
)

type MigrateExecutor struct {
	Dsn        string // dsn 连接串
	db         *sql.DB
	TrackingId string
}

func checkError(err error, detail ...string) {
	if err != nil {
		if len(detail) > 0 && detail[0] != "" {
			panic(errors.New(fmt.Sprintf("%s : %s", detail[0], err.Error())))
		} else {
			panic(err)
		}
	}
}

func (this *MigrateExecutor) ping() {
	if this.Dsn == "" {
		panic(errors.New("empty dsn error..."))
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

func (this *MigrateExecutor) ExecSQLWithLogger(detail, sql string, args ...interface{}) (rs sql.Result, err error) {
	if detail != "" {
		detail = fmt.Sprintf("[id = %s][sql = %s]", detail, sql)
	}
	stmt, err := this.db.Prepare(sql)
	checkError(err, detail)
	rs, err = stmt.Exec(args...)
	checkError(err, detail)
	return
}

func (this *MigrateExecutor) ExecSQL(sql string, args ...interface{}) (rs sql.Result, err error) {
	return this.ExecSQLWithLogger("", sql, args...)
}

func (this *MigrateExecutor) QueryRowSQL(sql string, args ...interface{}) (row *sql.Row) {
	stmt, err := this.db.Prepare(sql)
	checkError(err)
	row = stmt.QueryRow(args...)
	return
}

func (this *MigrateExecutor) record(flag, hash, sql, tracking_detail string) {
	if this.db != nil {
		recordLog := `INSERT INTO migrate_version(tracking_id,flag,hash,sql_detail,tracking_detail, created_time) VALUES (?,?,?,?,?,NOW());`
		this.ExecSQL(recordLog, this.TrackingId, flag, hash, sql, tracking_detail)
	}
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
	row := this.QueryRowSQL(sql, hash)
	var datacount int64
	if err := row.Scan(&datacount); err == nil && datacount > 0 {
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
			if strings.TrimSpace(executeSql) != "" {
				detailHash := hashutil.CalculateHashWithString(executeSql)
				if !this.checkExecuted(detailHash) {
					this.ExecSQLWithLogger(strconv.FormatInt(migrate.Id, 10), executeSql)
					this.record("true", detailHash, executeSql, "")
				}
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
