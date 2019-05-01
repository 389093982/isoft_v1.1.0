package iworkanalyzer

import (
	"fmt"
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
		parse(message)
	}
}

func parse(message string) {
	metas, lexers, err := AnalysisLexer(message)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if err := AnalysisSyntactic(metas, lexers); err != nil {
		fmt.Println(err.Error())
	}
}
