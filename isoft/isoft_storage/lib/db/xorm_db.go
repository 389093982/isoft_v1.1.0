package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func init()  {
	engine, _ = xorm.NewEngine("mysql", "root:123456@(106.15.186.139:3306)/isoft_storage?charset=utf8")
}