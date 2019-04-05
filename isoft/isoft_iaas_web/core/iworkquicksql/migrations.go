package iworkquicksql

import (
	"fmt"
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

type BaseCreator struct {
}

func (this *BaseCreator) getColumnTypeWithLength(column *TableColumn) string {
	if strings.TrimSpace(column.Length) != "" {
		return fmt.Sprintf(`%s(%s)`, column.ColumnType, column.Length)
	}
	return column.ColumnType
}

func (this *BaseCreator) getColumnNames(info *TableInfo) []string {
	rs := make([]string, 0)
	for _, column := range info.TableColumns {
		rs = append(rs, column.ColumnName)
	}
	return rs
}

func (this *BaseCreator) getCommonInfo(column *TableColumn) []string {
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

func AlterTable(preTableInfo, tableInfo *TableInfo) string {
	alter := &MigrationAlter{
		preTableInfo: preTableInfo,
		tableInfo:    tableInfo,
	}
	return alter.BuildAlterTableSql()
}

func CreateTable(tableInfo *TableInfo) string {
	creator := &MigrationCreator{
		tableInfo: tableInfo,
	}
	return creator.BuildCreateTableSql()
}
