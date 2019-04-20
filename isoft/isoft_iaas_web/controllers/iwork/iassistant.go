package iwork

import (
	"isoft/isoft_iaas_web/core/iworkutil/sqlutil"
	"isoft/isoft_iaas_web/models/iwork"
)

func (this *WorkController) LoadQuickSqlMeta() {
	resource_id, _ := this.GetInt64("resource_id")
	var err error
	// 查询所有的数据库信息
	resource, _ := iwork.QueryResourceById(resource_id)
	tableColumnsMap := make(map[string]interface{}, 0)

	tableNames := sqlutil.GetAllTableNames(resource.ResourceDsn)
	for _, tableName := range tableNames {
		tableColumns := sqlutil.GetAllColumnNames(tableName, resource.ResourceDsn)
		tableColumnsMap[tableName] = tableColumns
	}

	if err == nil {
		this.Data["json"] = &map[string]interface{}{
			"status":          "SUCCESS",
			"tableNames":      tableNames,
			"tableColumnsMap": tableColumnsMap,
		}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}
