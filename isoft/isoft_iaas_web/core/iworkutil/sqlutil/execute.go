package sqlutil

func Execute(sqlstring string, sql_binding []interface{}, dataSourceName string) int64 {
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
	result, err := stmt.Exec(sql_binding...)
	if err != nil {
		panic(err)
	}
	if affected, err := result.RowsAffected(); err == nil {
		return affected
	} else {
		panic(err)
	}
}
