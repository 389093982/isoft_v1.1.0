package iworkquicksql

import (
	"fmt"
	"testing"
)

func Test_CreateTable(t *testing.T) {
	fmt.Println(CreateTable(TableInfo{
		TableName: "helloworld1",
		TableColumns: []*TableColumn{
			&TableColumn{
				ColumnName: "zhangsan1",
				ColumnType: "int",
				PrimaryKey: true,
				Comment:    "zhangsan",
			},
			&TableColumn{
				ColumnName: "zhangsan2",
				ColumnType: "int",
				PrimaryKey: false,
				Comment:    "zhangsan",
			},
			&TableColumn{
				ColumnName: "zhangsan3",
				ColumnType: "int",
				PrimaryKey: false,
				Comment:    "zhangsan",
			},
		},
	}))
}
