package iworkanalyzer

import (
	"fmt"
	"isoft/isoft/common/stringutil"
	"strings"
	"testing"
)

func Test_GetLexerParse(t *testing.T) {
	messages := []string{
		//"`helloworld`",
		"func1(func1(`helloworld`),func1(123456,123456))",
		//"123456",
		//"$a.b.c.d",
		//"$a.b.c.",
		//"a$a.b.c.d",
		//"`a`$a.b.c.d",
		//"1+2",
		//"`1+2`",
		//"`1+2```",
		//"`1+2````",
	}

	for _, message := range messages {
		if callerids, callers, err := parse(message); err != nil {
			panic(err)
		} else {
			for index, callerid := range callerids {
				fmt.Println(callerid)
				fmt.Println(callers[index].FuncRealName)
				fmt.Println(callers[index].FuncRealArgs)
			}
		}
	}
}

func parse(message string) ([]string, []*FuncExecutor, error) {
	callerids := make([]string, 0)
	callers := make([]*FuncExecutor, 0)
	for {
		if strings.TrimSpace(message) == "" || strings.HasPrefix(message, "$uuid.") {
			break // 已经被提取完了
		}
		metas, lexers, err := AnalysisLexer(message)
		if err != nil {
			return callerids, callers, err
		}
		if caller, err := GetPriorityFuncExecutorFromLexersExpression(strings.Join(lexers, "")); err != nil {
			return callerids, callers, err
		} else {
			uuid := stringutil.RandomUUID()
			// 函数左边部分
			funcLeft := metas[:lexerAt(lexers, caller.FuncLeftIndex)]
			// 函数右边部分
			funcRight := metas[lexerAt(lexers, caller.FuncRightIndex)+1:]
			// 函数部分
			funcArea := metas[lexerAt(lexers, caller.FuncLeftIndex) : lexerAt(lexers, caller.FuncRightIndex)+1]
			// 将 caller 函数替换成 uuid
			message = strings.Join(funcLeft, "") + "$uuid." + uuid + strings.Join(funcRight, "")
			caller.FuncRealName = strings.Replace(funcArea[0], "(", "", -1) // 去除函数名中的 (
			caller.FuncRealArgs = funcArea[1 : len(funcArea)-1]
			callerids = append(callerids, uuid)
			callers = append(callers, caller)
		}
	}
	return callerids, callers, nil
}

// 判断当前索引在整个 lexers 切片中的位置
func lexerAt(lexers []string, index int) int {
	// 统计总长度
	sumIndex := 0
	for _index, lexer := range lexers {
		if index >= sumIndex && index < sumIndex+len(lexer) {
			return _index
		}
		// 总长度增长
		sumIndex += len(lexer)
	}
	return -1
}
