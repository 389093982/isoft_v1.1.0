package iwork

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type TableMigrate struct {
	Id              int64     `json:"id"`
	TableName       string    `json:"table_name"`
	MigrateType     string    `json:"migrate_type"`
	TableInfo       string    `json:"table_info"`
	TableInfoHash   string    `json:"table_info_hash"`
	TableMigrateSql string    `json:"table_migrate_sql"`
	TableAutoSql    string    `json:"table_auto_sql"`
	PreMigrateId    int64     `json:"pre_migrate_id"`
	PreMigrateHash  string    `json:"pre_migrate_hash"`
	ValidateResult  string    `json:"validate_result"`
	IsMaxMigrateId  bool      `json:"is_max_migrate_id" orm:"-"`
	CreatedBy       string    `json:"created_by"`
	CreatedTime     time.Time `json:"created_time" orm:"auto_now_add;type(datetime)"`
	LastUpdatedBy   string    `json:"last_updated_by"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}

func InsertOrUpdateTableMigrate(tm *TableMigrate) (id int64, err error) {
	o := orm.NewOrm()
	if tm.Id > 0 {
		id, err = o.Update(tm)
	} else {
		id, err = o.Insert(tm)
	}
	return
}

func QueryAllMigrate() (migrates []TableMigrate, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("table_migrate").OrderBy("id").All(&migrates)
	return
}

func QueryMigrate(filterTableName string, current_page, offset int) (migrates []TableMigrate, counts int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("table_migrate").Filter("table_name__contains", filterTableName)
	counts, _ = qs.Count()
	qs = qs.OrderBy("-id").Limit(offset, (current_page-1)*offset)
	_, err = qs.All(&migrates)
	return
}

func QueryMigrateInfo(id int64) (migrate TableMigrate, err error) {
	o := orm.NewOrm()
	err = o.QueryTable("table_migrate").Filter("id", id).One(&migrate)
	return
}

// 最近一次迁移记录
func QueryLastMigrate(tableName string, id int64, operateType string) (migrate TableMigrate, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("table_migrate").Filter("table_name", tableName)
	if id > 0 { // 非 CREATE 操作
		if operateType == "upgrade" {
			// upgrade 操作前置迁移小于等于当前 id
			qs = qs.Filter("id__lte", id)
		} else if operateType == "update" {
			// update 操作前置迁移小于当前 id
			qs = qs.Filter("id__lt", id)
		}
	}
	err = qs.OrderBy("-last_updated_time").One(&migrate)
	return
}

func QueryMaxMigrationIdForTable(tableName string) (int64, error) {
	var migrate TableMigrate
	o := orm.NewOrm()
	err := o.QueryTable("table_migrate").Filter("table_name", tableName).OrderBy("-id").One(&migrate)
	if err != nil {
		return -1, err
	}
	return migrate.Id, nil
}
