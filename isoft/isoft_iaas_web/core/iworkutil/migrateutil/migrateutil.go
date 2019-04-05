package migrateutil

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"isoft/isoft/common/hashutil"
	"isoft/isoft/common/stringutil"
	"isoft/isoft_iaas_web/core/iworkutil/datatypeutil"
	"isoft/isoft_iaas_web/models/iwork"
	"strconv"
	"strings"
)

type MigrateExecutor struct {
	Dsn        string // dsn 连接串
	db         *sql.DB
	TrackingId string
}

func (this *MigrateExecutor) ping() (err error) {
	if this.Dsn == "" {
		return errors.New("empty dsn error...")
	}
	// 建立连接
	this.db, err = sql.Open("mysql", this.Dsn)
	return nil
}

// 建立迁移文件版本管理表
func (this *MigrateExecutor) initial() (err error) {
	versionTable := `CREATE TABLE IF NOT EXISTS migrate_version 
		(id INT(20) PRIMARY KEY AUTO_INCREMENT, tracking_id CHAR(200), flag CHAR(200), hash CHAR(200),sql_detail TEXT, tracking_detail TEXT, created_time datetime);`
	_, err = this.ExecSQL(versionTable)
	return
}

func (this *MigrateExecutor) ExecSQL(sql string, args ...interface{}) (rs sql.Result, err error) {
	stmt, err := this.db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	rs, err = stmt.Exec(args...)
	return
}

func (this *MigrateExecutor) QueryRowSQL(sql string, args ...interface{}) (row *sql.Row, err error) {
	if stmt, err := this.db.Prepare(sql); err == nil {
		row = stmt.QueryRow(args...)
	}
	return
}

func (this *MigrateExecutor) record(flag, hash, sql, tracking_detail string) error {
	if this.db != nil {
		recordLog := `INSERT INTO migrate_version(tracking_id,flag,hash,sql_detail,tracking_detail, created_time) VALUES (?,?,?,?,?,NOW());`
		_, err := this.ExecSQL(recordLog, this.TrackingId, flag, hash, sql, tracking_detail)
		return err
	}
	return nil
}

func (this *MigrateExecutor) migrate() (err error) {
	migrates, err := iwork.QueryAllMigrate()
	if err == nil {
		for _, migrate := range migrates {
			if err = this.migrateOne(migrate); err != nil {
				return err
			}
		}
	}
	return
}

func (this *MigrateExecutor) checkExecuted(hash string) bool {
	sql := `SELECT COUNT(*) FROM migrate_version WHERE hash = ?`
	if row, err := this.QueryRowSQL(sql, hash); err == nil {
		var datacount int64
		if err := row.Scan(&datacount); err == nil && datacount > 0 {
			return true
		}
	}
	return false
}

func (this *MigrateExecutor) migrateOne(migrate iwork.TableMigrate) error {
	hash := hashutil.CalculateHashWithString(migrate.TableMigrateSql)
	// 已经执行过则忽略
	if this.checkExecuted(hash) {
		return nil
	}
	// 每次迁移都有可能有多个执行 sql
	executeSqls := strings.Split(migrate.TableMigrateSql, ";")
	executeSqls = datatypeutil.FilterSlice(executeSqls, datatypeutil.CheckNotEmpty)
	tx, err := this.db.Begin()
	if err != nil {
		return err
	}
	for _, executeSql := range executeSqls {
		detailHash := hashutil.CalculateHashWithString(executeSql)
		if this.checkExecuted(detailHash) {
			break
		}
		if _, err := this.ExecSQL(executeSql); err == nil {
			this.record("true", detailHash, executeSql, "")
		} else {
			tx.Rollback()
			errorMsg := fmt.Sprintf("[%s] - [%s] : %s", strconv.FormatInt(migrate.Id, 10), executeSql, err.Error())
			return errors.New(errorMsg)
		}
	}
	tx.Commit()
	// 计算hash 值
	this.record("true", hash, migrate.TableMigrateSql, "")
	return nil
}

func MigrateToDB(dsn string) (err error) {
	executor := &MigrateExecutor{
		Dsn:        dsn,
		TrackingId: stringutil.RandomUUID(),
	}
	if err = executor.ping(); err == nil {
		if err = executor.initial(); err == nil {
			err = executor.migrate()
		}
	}
	if err != nil {
		executor.record("false", "", "", err.Error())
	}
	return
}