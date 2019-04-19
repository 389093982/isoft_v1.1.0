package iwork

import (
	"isoft/isoft_iaas_web/core/iworkutil/sqlutil"
	"isoft/isoft_iaas_web/models/iwork"
)

func (this *WorkController) LoadQuickSqlMeta() {
	var err error
	// 查询所有的数据库信息
	resources := iwork.QueryAllResource("db")
	tableNamesMap, tableColumnsMap := make(map[string]interface{}, 0), make(map[string]interface{}, 0)
	for _, resource := range resources {
		tableNames := sqlutil.GetAllTableNames(resource.ResourceDsn)
		tableNamesMap[resource.ResourceDsn] = tableNames
		for _, tableName := range tableNames {
			tableColumns := sqlutil.GetAllColumnNames(tableName, resource.ResourceDsn)
			tableColumnsMap[resource.ResourceDsn+tableName] = tableColumns
		}
	}

	if err == nil {
		this.Data["json"] = &map[string]interface{}{
			"status":          "SUCCESS",
			"resources":       resources,
			"tableNamesMap":   tableNamesMap,
			"tableColumnsMap": tableColumnsMap,
		}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}
