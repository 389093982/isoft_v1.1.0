package sqlutil

func GetMetaDatas(sql, dataSourceName string) []string {
	db, err := GetConnForMysql("mysql", dataSourceName)
	if err != nil{
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query(sql)
	if err != nil{
		panic(err)
	}
	defer rows.Close()
	colNames,err := rows.Columns()
	if err != nil{
		panic(err)
	}
	return colNames
}
