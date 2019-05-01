package iworkanalyzer

import (
	"fmt"
	"testing"
)

func Test_GetLexerParse(t *testing.T) {
	messages := []string{
		//"`helloworld`",
		//"func1(func1(`helloworld`),func1(123456,123456))",
		//"123456",
		//"$a.b.c.d",
		//"$a.b.c.",
		//"a$a.b.c.d",
		//"`a`$a.b.c.d",
		//"1+2",
		//"`1+2`",
		//"`1+2```",
		//"`1+2````",
		"func())",
	}

	for _, message := range messages {
		if callerids, callers, err := ParseToFuncCallers(message); err != nil {
			panic(err)
		} else {
			for index, callerid := range callerids {
				fmt.Println(callerid)
				fmt.Println(callers[index].FuncName)
				fmt.Println(callers[index].FuncArgs)
			}
		}
	}
}
