package iwork

import (
	"encoding/json"
	"github.com/astaxie/beego/utils/pagination"
	"isoft/isoft/common/pageutil"
	"isoft/isoft_iaas_web/core/iworkquicksql"
	"isoft/isoft_iaas_web/models/iwork"
	"time"
)

func (this *WorkController) SubmitMigrate() {
	var err error
	tableName := this.GetString("tableName")
	tableColunmStr := this.GetString("tableColunms")
	tableColunms := make([]*iworkquicksql.TableColumn, 0)
	if err = json.Unmarshal([]byte(tableColunmStr), &tableColunms); err == nil {
		tableInfo := iworkquicksql.TableInfo{
			TableName:    tableName,
			TableColumns: tableColunms,
		}
		if tableInfoStr, err1 := json.Marshal(tableInfo); err1 == nil {
			tm := &iwork.TableMigrate{
				TableName:       tableName,
				TableColumns:    string(tableInfoStr),
				CreatedBy:       "SYSTEM",
				CreatedTime:     time.Now(),
				LastUpdatedBy:   "SYSTEM",
				LastUpdatedTime: time.Now(),
			}
			_, err = iwork.InsertOrUpdateTableMigrate(tm)
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
		paginator := pagination.SetPaginator(this.Ctx, offset, count)
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "migrates": migrates,
			"paginator": pageutil.Paginator(paginator.Page(), paginator.PerPageNums, paginator.Nums())}
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
