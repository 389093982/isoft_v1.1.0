package iworkanalyzer

import (
	"fmt"
	"strings"
)

func RenderFuncCallers(metas, lexers []string) error {
	funcCallers, err := GetAllFuncExecutor(metas, lexers)
	if err != nil {
		return err
	}
	fillFuncCallers(metas, lexers, funcCallers)
	return nil
}

func fillFuncCallers(metas, lexers []string, funcCallers []*FuncExecutor) {
	for _, funcCaller := range funcCallers {
		funcCaller.FuncName = fillFuncCallerLexer(metas, lexers, funcCaller.FuncLexersName)

		fmt.Println(funcCaller)
	}
}

func fillFuncCallerLexer(metas, lexers []string, lexer string) string {
	return ""
}
