package fileutil

import (
	"fmt"
	"log"
	"testing"
)

func Test_GetAllFile(t *testing.T) {
	files, err := GetAllFile("D:/zhourui/program/go/goland_workspace/src/isoft/isoft_deploy_web/shell", true)
	if err != nil {
		log.Fatal(err.Error())
	} else {
		for _, filepath := range files {
			fmt.Println(filepath)
		}
	}
}
