package sqlutil

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
