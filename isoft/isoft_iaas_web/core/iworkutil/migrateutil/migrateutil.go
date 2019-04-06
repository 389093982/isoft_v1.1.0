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
	ForceClean bool
	migrates   []iwork.TableMigrate
}

func (this *MigrateExecutor) ping() (err error) {
	if this.Dsn == "" {
		return errors.New("empty dsn error...")
	}
	// 建立连接
	this.db, err = sql.Open("mysql", this.Dsn)
	return nil
}

func (this *MigrateExecutor) executeForceClean() error {
	if this.ForceClean == false {
		return nil
	}
	dropTables := make([]string, 0)
	for _, migration := range this.migrates {
		if !stringutil.CheckContains(migration.TableName, dropTables) {
			dropTables = append(dropTables, migration.TableName)
		}
	}
	dropTables = append(dropTables, "migrate_version")
	for _, dropTableName := range dropTables {
		if _, err := this.ExecSQL(fmt.Sprintf(`DROP TABLE IF EXISTS %s;`, dropTableName)); err != nil {
			return err
		}
	}
	return nil
}

// 建立迁移文件版本管理表
func (this *MigrateExecutor) initial() (err error) {
	if err = this.executeForceClean(); err != nil {
		return err
	}
	versionTable := `CREATE TABLE IF NOT EXISTS migrate_version (id INT(20) PRIMARY KEY AUTO_INCREMENT, 
	tracking_id CHAR(200), flag CHAR(200), hash CHAR(200),sql_detail TEXT, tracking_detail TEXT, created_time datetime);`
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

func (this *MigrateExecutor) loadAllMigrate() (err error) {
	this.migrates, err = iwork.QueryAllMigrate()
	return
}

func (this *MigrateExecutor) migrate() (err error) {
	for _, migrate := range this.migrates {
		if err = this.migrateOne(migrate); err != nil {
			migrate.ValidateResult = "FAILED"
			iwork.InsertOrUpdateTableMigrate(&migrate)
			return err
		} else {
			migrate.ValidateResult = "SUCCESS"
			iwork.InsertOrUpdateTableMigrate(&migrate)
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

func (this *MigrateExecutor) getMigrate(migrateId int64) *iwork.TableMigrate {
	for _, migrate := range this.migrates {
		if migrate.Id == migrateId {
			return &migrate
		}
	}
	return nil
}

func (this *MigrateExecutor) checkMigrate() error {
	for _, migrate := range this.migrates {
		if migrate.PreMigrateId > 0 {
			if preMigrate := this.getMigrate(migrate.PreMigrateId); preMigrate != nil {
				if migrate.PreMigrateHash != hashutil.CalculateHashWithString(preMigrate.TableInfo) {
					return errors.New(fmt.Sprintf("migrate[id=%d] check pre migrate hash error, please rebuild it...", migrate.Id))
				}
			}
		}
	}
	return nil
}

func (this *MigrateExecutor) migrateOne(migrate iwork.TableMigrate) error {
	if strings.TrimSpace(migrate.TableMigrateSql) != "" {
		// 优先使用用户自定义 sql
		migrate.TableAutoSql = strings.TrimSpace(migrate.TableMigrateSql)
	}
	hash := fmt.Sprintf(`%d-%s`, migrate.Id, hashutil.CalculateHashWithString(migrate.TableAutoSql))
	// 已经执行过则忽略
	if this.checkExecuted(hash) {
		return nil
	}
	// 每次迁移都有可能有多个执行 sql
	executeSqls := strings.Split(migrate.TableAutoSql, ";")
	executeSqls = datatypeutil.FilterSlice(executeSqls, datatypeutil.CheckNotEmpty)
	tx, err := this.db.Begin()
	if err != nil {
		return err
	}
	for index, executeSql := range executeSqls {
		// 防止 executeSql 相同导致的 hash 值相同问题
		detailHash := fmt.Sprintf(`%d-%d-%s`, migrate.Id, index, hashutil.CalculateHashWithString(executeSql))
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
	this.record("true", hash, migrate.TableAutoSql, "")
	return nil
}

func MigrateToDB(dsn string, forceClean bool) (err error) {
	executor := &MigrateExecutor{
		Dsn:        dsn,
		TrackingId: stringutil.RandomUUID(),
		ForceClean: forceClean,
	}
	if err = executor.ping(); err == nil {
		if err = executor.loadAllMigrate(); err != nil {
			return err
		}
		if err = executor.initial(); err != nil {
			return err
		}
		err = executor.migrate()
	}
	if err != nil {
		executor.record("false", "", "", err.Error())
	}
	return
}
