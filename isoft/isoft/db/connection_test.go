package db

import (
	"fmt"
	"testing"
)

func Test_Connection(t *testing.T) {
	_, err := GetConnection("root", "Admin123456", "193.112.162.61", 4444, "mysql")
	if err != nil {
		fmt.Println(fmt.Sprintf("connection failed, %s"), err.Error())
	} else {
		fmt.Println("connection success...")
	}
}
