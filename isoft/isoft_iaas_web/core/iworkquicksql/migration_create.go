package iworkquicksql

import (
	"fmt"
	"strings"
)

type MigrationCreator struct {
	BaseCreator
	tableInfo *TableInfo
}

func (this *MigrationCreator) BuildCreateTableSql() string {
	createTableTemplate := `CREATE TABLE IF NOT EXISTS %s(%s)ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	return fmt.Sprintf(createTableTemplate, this.tableInfo.TableName, this.buildCreateColumunSqls())
}

func (this *MigrationCreator) buildCreateColumunSqls() string {
	columns := make([]string, 0)
	for _, column := range this.tableInfo.TableColumns {
		columns = append(columns, this.buildCreateColumunSql(column))
	}
	sql := strings.Join(columns, "")
	return sql[:len(sql)-1] // 截取最后一个逗号
}

func (this *MigrationCreator) buildCreateColumunSql(column *TableColumn) string {
	appends := make([]string, 0)
	appends = append(appends, column.ColumnName)
	appends = append(appends, this.getColumnTypeWithLength(column))
	appends = append(appends, this.getCommonInfo(column)...)
	return fmt.Sprintf(`%s%s`, strings.Join(appends, " "), ",")
}
