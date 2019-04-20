package iwork

import (
	"fmt"
	"isoft/isoft_iaas_web/core/iworkutil/sqlutil"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
)

func (this *WorkController) LoadQuickSqlMeta() {
	resource_id, _ := this.GetInt64("resource_id")
	var err error
	// 查询所有的数据库信息
	resource, _ := iwork.QueryResourceById(resource_id)
	tableColumnsMap := make(map[string]interface{}, 0)
	tableSqlMap := make(map[string]interface{}, 0)
	tableNames := sqlutil.GetAllTableNames(resource.ResourceDsn)
	for _, tableName := range tableNames {
		tableColumns := sqlutil.GetAllColumnNames(tableName, resource.ResourceDsn)
		tableColumnsMap[tableName] = tableColumns
		tableSqlMap[tableName] = []string{
			fmt.Sprintf(`select count(*) as count from %s`, tableName),
			fmt.Sprintf(`select count(*) as count from %s where 1 = 0`, tableName),
			strings.Join(tableColumns, ","),
			fmt.Sprintf(`select %s from %s where 1 = 0`, strings.Join(tableColumns, ","), tableName),
		}
	}

	if err == nil {
		this.Data["json"] = &map[string]interface{}{
			"status":          "SUCCESS",
			"tableNames":      tableNames,
			"tableColumnsMap": tableColumnsMap,
			"tableSqlMap":     tableSqlMap,
		}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}
