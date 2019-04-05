package iwork

import (
	"encoding/json"
	"github.com/astaxie/beego/utils/pagination"
	"isoft/isoft/common/pageutil"
	"isoft/isoft_iaas_web/core/iworkquicksql"
	"isoft/isoft_iaas_web/core/iworkutil/migrateutil"
	"isoft/isoft_iaas_web/models/iwork"
	"time"
)

func (this *WorkController) ExecuteMigrate() {
	resource_name := this.GetString("resource_name")
	resource, _ := iwork.QueryResourceByName(resource_name)
	if err := migrateutil.MigrateToDB(resource.ResourceDsn); err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

func (this *WorkController) SubmitMigrate() {
	var err error
	tableName := this.GetString("tableName")
	tableColunmStr := this.GetString("tableColunms")
	operateType := this.GetString("operateType")
	id, _ := this.GetInt64("id")
	tableColunms := make([]*iworkquicksql.TableColumn, 0)
	if err = json.Unmarshal([]byte(tableColunmStr), &tableColunms); err == nil {
		tableInfo := iworkquicksql.TableInfo{
			TableName:    tableName,
			TableColumns: tableColunms,
		}
		// 纠正更新时操作
		if operateType == "update" {
			iwork.DeleteMigrateById(id)
		}
		var migrateSql, migrateType string
		// 有最近一次创建或者修改记录
		if preMigrate, err := iwork.QueryLastMigrate(tableName); err == nil {
			migrateType = "ALTER"
			var preTableInfo iworkquicksql.TableInfo
			json.Unmarshal([]byte(preMigrate.TableInfo), &preTableInfo)
			migrateSql = iworkquicksql.AlterTable(preTableInfo, tableInfo)
		} else {
			migrateType = "CREATE"
			migrateSql = iworkquicksql.CreateTable(tableInfo)
		}
		if tableInfoStr, err1 := json.Marshal(tableInfo); err1 == nil {
			if migrateSql != "" {
				tm := &iwork.TableMigrate{
					TableName:       tableName,
					TableInfo:       string(tableInfoStr),
					TableMigrateSql: migrateSql,
					MigrateType:     migrateType,
					CreatedBy:       "SYSTEM",
					CreatedTime:     time.Now(),
					LastUpdatedBy:   "SYSTEM",
					LastUpdatedTime: time.Now(),
				}
				_, err = iwork.InsertOrUpdateTableMigrate(tm)
			}
		} else {
			err = err1
		}
	}
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

func (this *WorkController) FilterPageMigrate() {
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页
	migrates, count, err := iwork.QueryMigrate(current_page, offset)
	if err == nil {
		resources := iwork.QueryAllResource("db")
		paginator := pagination.SetPaginator(this.Ctx, offset, count)
		this.Data["json"] = &map[string]interface{}{
			"status":    "SUCCESS",
			"migrates":  migrates,
			"paginator": pageutil.Paginator(paginator.Page(), paginator.PerPageNums, paginator.Nums()),
			"resources": resources,
		}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}

func (this *WorkController) GetMigrateInfo() {
	id, _ := this.GetInt64("id")
	migrate, err := iwork.QueryMigrateInfo(id)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "migrate": migrate}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}
