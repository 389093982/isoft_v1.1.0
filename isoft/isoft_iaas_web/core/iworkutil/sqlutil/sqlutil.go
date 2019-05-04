package sqlutil

import (
	"fmt"
)

func GetAllTableNames(dataSourceName string) []string {
	_, rowDatas := Query("show tables;", []interface{}{}, dataSourceName)
	tableNames := make([]string, 0)
	for _, rowData := range rowDatas {
		for _, tableName := range rowData {
			tableNames = append(tableNames, tableName.(string))
		}
	}
	return tableNames
}

func GetAllColumnNames(tableName, dataSourceName string) []string {
	return GetMetaDatas(fmt.Sprintf("select * from %s where 1=0", tableName), dataSourceName)
}
