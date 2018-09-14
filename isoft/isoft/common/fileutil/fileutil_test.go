package fileutil

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func Test_GetAllFile(t *testing.T) {
	files, _, err := GetAllFile("D:/zhourui/program/go/goland_workspace/src/isoft/isoft_deploy_web/shell", true)
	if err != nil {
		log.Fatal(err.Error())
	} else {
		for _, filepath := range files {
			fmt.Println(filepath)
		}
	}
}

func Test_CreateDir(t *testing.T)  {
	err := os.Mkdir("D:/build/isoft_storage_log", os.ModePerm)
	if err != nil{
		fmt.Println(err)
	}
}
