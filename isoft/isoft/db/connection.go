package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func GetConnection(username, passwd, ip string, port int64, dbname string) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, passwd, ip, port, dbname)
	fmt.Println(dataSourceName)
	db, err := sql.Open("mysql", dataSourceName)
	return db, err
}
