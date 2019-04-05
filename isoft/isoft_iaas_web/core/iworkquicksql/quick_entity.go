package iworkquicksql

import (
	"fmt"
	"isoft/isoft/common/stringutil"
	"strings"
)

type TableInfo struct {
	TableName    string         `json:"table_name"`
	TableColumns []*TableColumn `json:"table_columns"`
}

type TableColumn struct {
	ColumnName    string `json:"column_name"`
	ColumnType    string `json:"column_type"`
	Length        string `json:"length"`
	Default       string `json:"default"`
	PrimaryKey    string `json:"primary_key"`
	AutoIncrement string `json:"auto_increment"`
	Unique        string `json:"unique"`
	Comment       string `json:"comment"`
}

func getColumnNames(info TableInfo) []string {
	rs := make([]string, 0)
	for _, column := range info.TableColumns {
		rs = append(rs, column.ColumnName)
	}
	return rs
}

func AlterTable(preTableInfo, tableInfo TableInfo) string {
	migrates := make([]string, 0)
	preColumnNames := getColumnNames(preTableInfo)
	columnNames := getColumnNames(tableInfo)
	for _, preColumnName := range preColumnNames {
		if !stringutil.CheckContains(preColumnName, columnNames) {
			migrates = append(migrates, deleteField(tableInfo.TableName, preColumnName)+",")
		}
	}
	for index, columnName := range columnNames {
		if flag, preindex := stringutil.CheckIndexContains(columnName, preColumnNames); !flag {
			add := addField(tableInfo.TableName, columnName, tableInfo.TableColumns[index])
			migrates = append(migrates, add)
		} else {
			if modify := modifyField(tableInfo.TableName,
				preTableInfo.TableColumns[preindex], tableInfo.TableColumns[index]); modify != "" {
				migrates = append(migrates, modify)
			}
		}
	}
	return strings.Join(migrates, "\n")
}

func deleteField(tableName, columnName string) string {
	return strings.TrimSpace(fmt.Sprintf(`ALTER TABLE %s DROP COLUMN %s`, tableName, columnName)) + ";"
}

func addField(tableName, columnName string, column *TableColumn) string {
	return strings.TrimSpace(fmt.Sprintf(`ALTER TABLE %s ADD %s %s %s`,
		tableName, columnName, getColumnTypeWithLength(column), strings.Join(getCommonInfo(column), " "))) + ";"
}

func modifyField(tableName string, precolumn, column *TableColumn) string {
	modifys := make([]string, 0)
	if precolumn.Unique != column.Unique {
		modifys = append(modifys, getDropUniqueSql(tableName, column))
		modifys = append(modifys, getAddUniqueSql(tableName, column))
	}
	if getColumnTypeWithLength(precolumn) != getColumnTypeWithLength(column) {
		modifyColumnType := strings.TrimSpace(fmt.Sprintf(`ALTER TABLE %s MODIFY %s %s`,
			tableName, column.ColumnName, strings.Join(getCommonInfo(column), " "))) + ";"
		modifys = append(modifys, modifyColumnType)
	}
	if precolumn.PrimaryKey != column.PrimaryKey ||
		precolumn.AutoIncrement != column.AutoIncrement || precolumn.Comment != column.Comment {
		//return strings.TrimSpace(fmt.Sprintf(`ALTER TABLE %s MODIFY %s %s`,
		//	tableName, column.ColumnName, strings.Join(getCommonInfo(column), " "))) + ";"
	}
	return strings.Join(modifys, "")
}

func CreateTable(tableInfo TableInfo) string {
	create_table := `CREATE TABLE IF NOT EXISTS %s(
%s)ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	return fmt.Sprintf(create_table, tableInfo.TableName, createColumuns(tableInfo))
}

func createColumuns(tableInfo TableInfo) string {
	columns := make([]string, 0)
	for index, column := range tableInfo.TableColumns {
		if index == len(tableInfo.TableColumns)-1 {
			columns = append(columns, createIndentColumn(column))
		} else {
			columns = append(columns, createIndentColumn(column, true))
		}

	}

	return strings.Join(columns, "")
}

func createIndentColumn(column *TableColumn, comma ...bool) string {
	indent := `	%s
`
	return fmt.Sprintf(indent, createColumn(column, comma...))
}

func createColumn(column *TableColumn, comma ...bool) string {
	appends := make([]string, 0)
	appends = append(appends, column.ColumnName)
	appends = append(appends, getColumnTypeWithLength(column))
	appends = append(appends, getCommonInfo(column)...)
	var commaStr string
	if len(comma) > 0 && comma[0] == true {
		commaStr = ","
	}
	return fmt.Sprintf(`%s%s`, strings.Join(appends, " "), commaStr)
}

func getCommonInfo(column *TableColumn) []string {
	appends := make([]string, 0)
	if column.PrimaryKey == "Y" {
		appends = append(appends, "PRIMARY KEY")
	}
	if column.AutoIncrement == "Y" {
		appends = append(appends, "AUTO_INCREMENT")
	}
	if column.Unique == "Y" {
		appends = append(appends, "UNIQUE")
	}
	if strings.TrimSpace(column.Comment) != "" {
		appends = append(appends, fmt.Sprintf(`COMMENT '%s'`, strings.TrimSpace(column.Comment)))
	}
	return appends
}

func getDropUniqueSql(tableName string, column *TableColumn) string {
	if column.Unique == "N" {
		return fmt.Sprintf("ALTER TABLE %s DROP INDEX %s;", tableName, column.ColumnName)
	}
	return ""
}

func getAddUniqueSql(tableName string, column *TableColumn) string {
	if column.Unique == "Y" {
		return fmt.Sprintf("ALTER TABLE %s ADD UNIQUE(%s);", tableName, column.ColumnName)
	}
	return ""
}

func getColumnTypeWithLength(column *TableColumn) string {
	if strings.TrimSpace(column.Length) != "" {
		return fmt.Sprintf(`%s(%d)`, column.ColumnType, column.Length)
	}
	return column.ColumnType
}
