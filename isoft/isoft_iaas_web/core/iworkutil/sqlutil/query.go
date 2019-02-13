package sqlutil

import (
	"database/sql"
	"fmt"
	"reflect"
)

func GetMetaDatas(sql, dataSourceName string) (colNames []string) {
	db, err := GetConnForMysql("mysql", dataSourceName)
	if err != nil {
		return
	}
	defer db.Close()
	rows, err := db.Query(sql)
	if err != nil {
		return
	}
	defer rows.Close()
	colNames, err = rows.Columns()
	if err != nil {
		return
	}
	return colNames
}

func Query(sqlstring string, sql_binding []interface{}, dataSourceName string) (datacounts int64, rowDatas map[string]interface{}) {
	rowDatas = make(map[string]interface{}, 5)
	db, err := GetConnForMysql("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// 使用预编译 sql 防止 sql 注入
	stmt, err := db.Prepare(sqlstring)
	if err != nil {
		panic(err)
	}
	rows, err := stmt.Query(sql_binding...)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	colNames, _ := rows.Columns()
	rows.ColumnTypes()
	for rows.Next() {
		// 存储一行中的每一列值
		colValues := make([]sql.RawBytes, len(colNames))
		scanArgs := make([]interface{}, len(colValues))
		for i := range colValues {
			scanArgs[i] = &colValues[i]
		}
		rows.Scan(scanArgs...)
		for index, colValue := range colValues {
			name := fmt.Sprintf("rows[%d].%s", datacounts, colNames[index])
			// sql.RawBytes 转字符串
			_colValue := reflect.ValueOf(colValue).Interface().(sql.RawBytes)
			rowDatas[name] = string(_colValue)
			// 第一条记录进行简写,去除[0]标识
			if datacounts == 0 {
				_name := fmt.Sprintf("rows.%s", colNames[index])
				rowDatas[_name] = string(_colValue)
			}
		}
		// 数据量增加 1
		datacounts++
	}
	return
}
