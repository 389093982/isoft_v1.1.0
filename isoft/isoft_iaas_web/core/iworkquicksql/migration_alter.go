package iworkquicksql

import (
	"fmt"
	"isoft/isoft/common/stringutil"
	"isoft/isoft_iaas_web/core/iworkutil/datatypeutil"
	"strings"
)

type MigrationAlter struct {
	BaseCreator
	preTableInfo *TableInfo
	tableInfo    *TableInfo
}

func (this *MigrationAlter) BuildAlterTableSql() string {
	migrates := make([]string, 0)
	preColumnNames := this.getColumnNames(this.preTableInfo)
	columnNames := this.getColumnNames(this.tableInfo)
	for _, preColumnName := range preColumnNames {
		if !stringutil.CheckContains(preColumnName, columnNames) {
			migrates = append(migrates, this.deleteField(this.tableInfo.TableName, preColumnName)+",")
		}
	}
	for index, columnName := range columnNames {
		if flag, preindex := stringutil.CheckIndexContains(columnName, preColumnNames); !flag {
			add := this.addField(this.tableInfo.TableName, columnName, this.tableInfo.TableColumns[index])
			migrates = append(migrates, add)
		} else {
			if modify := this.modifyField(this.tableInfo.TableName,
				this.preTableInfo.TableColumns[preindex], this.tableInfo.TableColumns[index]); modify != "" {
				migrates = append(migrates, modify)
			}
		}
	}
	return strings.Join(migrates, "\n")
}

func (this *MigrationAlter) deleteField(tableName, columnName string) string {
	return strings.TrimSpace(fmt.Sprintf(`ALTER TABLE %s DROP COLUMN %s`, tableName, columnName)) + ";"
}

func (this *MigrationAlter) addField(tableName, columnName string, column *TableColumn) string {
	return strings.TrimSpace(fmt.Sprintf(`ALTER TABLE %s ADD %s %s %s`,
		tableName, columnName, this.getColumnTypeWithLength(column), strings.Join(this.getCommonInfo(column), " "))) + ";"
}

func (this *MigrationAlter) modifyField(tableName string, precolumn, column *TableColumn) string {
	modifys := make([]string, 0)
	// unique 设置
	if precolumn.Unique != column.Unique {
		modifys = append(modifys, this.getDropUniqueSql(tableName, column))
		modifys = append(modifys, this.getAddUniqueSql(tableName, column))
	}
	// 类型设置
	if this.getColumnTypeWithLength(precolumn) != this.getColumnTypeWithLength(column) {
		modifys = append(modifys, this.getAlterColumnTypeSql(tableName, column)+";")
	}
	// 注释设置
	if strings.TrimSpace(precolumn.Comment) != strings.TrimSpace(column.Comment) {
		alterColumnTypeSql := this.getAlterColumnTypeSql(tableName, column)
		alterCommentSql := fmt.Sprintf(`%s COMMENT '%s'`, alterColumnTypeSql, strings.TrimSpace(column.Comment))
		// 删除 alter 语句
		modifys = datatypeutil.DeleteSliceItem(modifys, alterColumnTypeSql)
		modifys = append(modifys, alterCommentSql)
	}
	// 设置 default
	if strings.TrimSpace(precolumn.Default) != strings.TrimSpace(column.Default) {
		if strings.TrimSpace(precolumn.Default) != "" {
			dropDefault := fmt.Sprintf(`ALTER TABLE %s ALTER COLUMN %s DROP DEFAULT;`, tableName, column.ColumnName)
			modifys = append(modifys, dropDefault)
		}
		if strings.TrimSpace(column.Default) != "" {
			setDefault := fmt.Sprintf(`ALTER TABLE %s ALTER COLUMN %s SET DEFAULT %v`, tableName, column.ColumnName, column.Default)
			modifys = append(modifys, setDefault)
		}
	}
	// 主键和自增只限制 id 使用,且不可变更
	return strings.Join(modifys, "")
}

func (this *MigrationAlter) getAlterColumnTypeSql(tableName string, column *TableColumn) string {
	return strings.TrimSpace(fmt.Sprintf(`ALTER TABLE %s MODIFY %s %s`,
		tableName, column.ColumnName, this.getColumnTypeWithLength(column)))
}

func (this *MigrationAlter) getDropUniqueSql(tableName string, column *TableColumn) string {
	if column.Unique == "N" {
		return fmt.Sprintf("ALTER TABLE %s DROP INDEX %s;", tableName, column.ColumnName)
	}
	return ""
}

func (this *MigrationAlter) getAddUniqueSql(tableName string, column *TableColumn) string {
	if column.Unique == "Y" {
		return fmt.Sprintf("ALTER TABLE %s ADD UNIQUE(%s);", tableName, column.ColumnName)
	}
	return ""
}
