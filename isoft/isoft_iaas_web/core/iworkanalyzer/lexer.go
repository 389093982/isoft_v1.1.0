package iworkanalyzer

import (
	"fmt"
	"github.com/pkg/errors"
	"regexp"
	"strings"
)

func AnalysisLexer(s string) (metas []string, lexers []string, err error) {
	metas = make([]string, 0)
	lexers = make([]string, 0)
	for {
		s = strings.TrimSpace(s)
		if s != "" {
			// 标识是否分析到一个词语
			flag := false
			for index, regex := range regexs {
				reg := regexp.MustCompile(regex)
				find := reg.FindString(s)
				if find != "" {
					metas = append(metas, find)
					lexers = append(lexers, regexLexers[index])
					s = strings.Replace(s, find, "", 1)
					flag = true
					break
				}
			}
			if !flag {
				return metas, lexers, errors.New(fmt.Sprintf("%s is an error lexer data", s))
			}
		} else {
			return metas, lexers, nil
		}
	}
}
