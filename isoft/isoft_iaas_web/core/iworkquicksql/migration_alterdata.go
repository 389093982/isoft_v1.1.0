package iworkquicksql

import (
	"fmt"
	"strings"
)

func AlterData(tableName, operateType string, tableColunms []*TableColumn) string {
	var sql string
	if operateType == "add" {
		sql = fmt.Sprintf(`INSERT INTO %s(%s) VALUES(%S);`,
			tableName, strings.Join(getColumnNames(tableColunms), ","),
			strings.Join(getColumnValues(tableColunms), ","))
	} else if operateType == "delete" {
		queryColumns := getColumnQuery(tableColunms)
		if len(queryColumns) > 0 {
			sql = fmt.Sprintf(`DELETE FROM %s WHERE %s;`, tableName, strings.Join(queryColumns, " AND "))
		}
	}
	return sql
}

func getColumnValues(tableColunms []*TableColumn) []string {
	columnValues := make([]string, 0)
	for _, tableColunm := range tableColunms {
		columnValues = append(columnValues, tableColunm.Instance)
	}
	return columnValues
}

func getColumnNames(tableColunms []*TableColumn) []string {
	columnNames := make([]string, 0)
	for _, tableColunm := range tableColunms {
		columnNames = append(columnNames, tableColunm.ColumnName)
	}
	return columnNames
}

func getColumnQuery(tableColunms []*TableColumn) []string {
	querys := make([]string, 0)
	for _, tableColunm := range tableColunms {
		if strings.TrimSpace(tableColunm.Instance) != "" {
			querys = append(querys, fmt.Sprintf(`%s = %s`, tableColunm.ColumnName, strings.TrimSpace(tableColunm.Instance)))
		}
	}
	return querys
}
