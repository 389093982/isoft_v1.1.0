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
	PrimaryKey    string `json:"primary_key"`
	AutoIncrement string `json:"auto_increment"`
	Comment       string `json:"comment"`
}

func CreateTable(info TableInfo) string {
	create_table := `CREATE TABLE IF NOT EXISTS %s(
%s)ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	return fmt.Sprintf(create_table, info.TableName, CreateColumuns(info))
}

func CreateColumuns(info TableInfo) string {
	columns := make([]string, 0)
	for index, column := range info.TableColumns {
		if index == len(info.TableColumns)-1 {
			columns = append(columns, CreateIndentColumn(column))
		} else {
			columns = append(columns, CreateIndentColumn(column, true))
		}

	}

	return strings.Join(columns, "")
}

func CreateIndentColumn(column *TableColumn, comma ...bool) string {
	indent := `	%s
`
	return fmt.Sprintf(indent, CreateColumn(column, comma...))
}

func CreateColumn(column *TableColumn, comma ...bool) string {
	appends := make([]string, 0)
	appends = append(appends, column.ColumnName)
	appends = append(appends, column.ColumnType)
	if column.PrimaryKey == "Y" {
		appends = append(appends, "PRIMARY KEY")
	}
	if column.AutoIncrement == "Y" {
		appends = append(appends, "AUTO_INCREMENT")
	}
	if strings.TrimSpace(column.Comment) != "" {
		appends = append(appends, fmt.Sprintf(`COMMENT '%s'`, strings.TrimSpace(column.Comment)))
	}
	var commaStr string
	if len(comma) > 0 && comma[0] == true {
		commaStr = ","
	}
	return fmt.Sprintf(`%s%s`, strings.Join(appends, " "), commaStr)
}
