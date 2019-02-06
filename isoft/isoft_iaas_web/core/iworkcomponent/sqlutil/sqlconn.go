package sqlutil

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //导入mysql驱动包
	"isoft/isoft_iaas_web/core/iworkcomponent"
)


func GetConnForMysql(driverName,dataSourceName string) (db *sql.DB, err error) {
	db, err = sql.Open(driverName, dataSourceName)
	if err != nil{
		panic(err)
	}
	if err = db.Ping(); err != nil{
		panic(err)
	}
	return db, err
}

func GetDataSourceName(db_conn string) string {
	if !iworkcomponent.IsDynamicParam(db_conn){
		return db_conn
	}else{
		return ""
	}
}