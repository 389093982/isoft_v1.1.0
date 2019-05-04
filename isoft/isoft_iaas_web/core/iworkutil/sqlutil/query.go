package sqlutil

import (
	"database/sql"
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

func Query(sqlstring string, sql_binding []interface{}, dataSourceName string) (
	datacounts int64, rowDatas []map[string]interface{}) {
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
	datacounts, rowDatas = parseRows(rows)
	return
}

func parseRows(rows *sql.Rows) (datacounts int64, rowDatas []map[string]interface{}) {
	// 列名、列值组成的 map,多行数据使用数组存储
	rowDatas = []map[string]interface{}{}
	colNames, _ := rows.Columns()
	for rows.Next() {
		colValues := scanRowData(rows, len(colNames))

		rowData := map[string]interface{}{}
		for index, colValue := range colValues {
			rowData[colNames[index]] = string(colValue)
		}
		rowDatas = append(rowDatas, rowData)
		// 数据量增加 1
		datacounts++
	}
	return
}

func scanRowData(rows *sql.Rows, colSize int) []sql.RawBytes {
	// 存储一行中的每一列值
	colValues := make([]sql.RawBytes, colSize)
	scanArgs := make([]interface{}, len(colValues))
	for i := range colValues {
		scanArgs[i] = &colValues[i]
	}
	rows.Scan(scanArgs...)
	return colValues
}

// 查询sql总数据量
func QuerySelectCount(sqlstring string, sql_binding []interface{}, dataSourceName string) (datacounts int64) {
	db, err := GetConnForMysql("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	stmt, err := db.Prepare(sqlstring)
	if err != nil {
		panic(err)
	}
	row := stmt.QueryRow(sql_binding...)
	err = row.Scan(&datacounts)
	if err != nil {
		panic(err)
	}
	return
}
