package db

import (
	"fmt"
	"testing"
)

func Test_Connection(t *testing.T) {
	db, err := GetConnection("root", "123456", "193.112.162.61", 3306, "abc")
	if err != nil {
		fmt.Println(fmt.Sprintf("connection failed, %s"), err.Error())
	} else {
		fmt.Println("connection success...")
	}

	defer db.Close()

	_, err = db.Exec("select * from ddd")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("successful...")
	}
}
