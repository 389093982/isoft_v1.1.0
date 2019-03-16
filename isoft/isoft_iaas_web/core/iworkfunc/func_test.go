package iworkfunc

import (
	"fmt"
	"os"
	"testing"
)

func Test_GetAllFile(t *testing.T) {
	//expression := "aaaa"
	expression := "func1    (a      ,b,   func2(e,f,a,func3(a,b),b),c,d)"
	for _, executor := range GetAllFuncExecutor(expression) {
		GetTrimFuncExecutor(executor)
		fmt.Println(executor.FuncName)
		for _, arg := range executor.FuncArgs {
			fmt.Println(arg)
		}
	}
}

func Test_CreateDir(t *testing.T) {
	err := os.Mkdir("D:/build/isoft_storage_log", os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}
