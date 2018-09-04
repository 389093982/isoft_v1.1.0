package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func GetConnection(username, passwd, ip string, port int64, dbname string) (*sql.DB, error) {
	// 新版本 mysql 错误 this user requires mysql native password authentication
	// 在连接mysql的url上加上 ?allowNativePasswords=true 即可解决
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?allowNativePasswords=true", username, passwd, ip, port, dbname)
	fmt.Println(dataSourceName)
	db, err := sql.Open("mysql", dataSourceName)
	return db, err
}
