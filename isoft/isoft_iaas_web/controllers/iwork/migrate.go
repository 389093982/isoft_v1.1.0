package iwork

import (
	"encoding/json"
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
