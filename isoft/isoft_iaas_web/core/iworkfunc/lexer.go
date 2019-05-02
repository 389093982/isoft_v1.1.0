package iworkfunc

import (
	"errors"
	"fmt"
	"isoft/isoft/common/stringutil"
	"regexp"
	"strings"
)

// 正则表达式
var regexs = []string{"^[a-zA-Z0-9]+\\(", "^\\)", "^`.*?`", "^[0-9]+", "^\\$[a-zA-Z_0-9]+\\.[a-zA-Z0-9\\-]+", ","}

// 正则表达式对应的词语
var regexLexers = []string{"func(", ")", "S", "N", "V", ","}

func isStringNumberOrVar(s string) bool {
	if _, lexers, err := analysisLexer(s); err == nil && len(lexers) == 1 {
		return true
	}
	return false
}

// 返回 uuid 和 funcCaller
func ParseToFuncCallers(expression string) ([]*FuncCaller, error) {
	callers := make([]*FuncCaller, 0)
	for {
		if strings.TrimSpace(expression) == "" || strings.HasPrefix(expression, "$uuid.") {
			break // 已经被提取完了
		}
		// 对 expression 表达式进行词法分析
		metas, lexers, err := analysisLexer(expression)
		if err != nil {
			return callers, err
		}
		// 提取 func
		caller, err := GetPriorityFuncExecutorFromLexersExpression(strings.Join(lexers, ""))
		if err != nil { // 提取失败
			return callers, err
		}
		if caller == nil { // 未提取到 func
			if !isStringNumberOrVar(expression) {
				return nil, errors.New(fmt.Sprintf(`invalid param for %s`, expression))
			}
			return nil, nil
		}
		uuid := stringutil.RandomUUID()
		// 函数左边部分
		funcLeft := metas[:lexerAt(lexers, caller.FuncLeftIndex)]
		// 函数右边部分
		funcRight := metas[lexerAt(lexers, caller.FuncRightIndex)+1:]
		// 函数部分
		funcArea := metas[lexerAt(lexers, caller.FuncLeftIndex) : lexerAt(lexers, caller.FuncRightIndex)+1]
		// 将 caller 函数替换成 uuid,以便下一轮提取 func 使用
		expression = strings.Join(funcLeft, "") + "$uuid." + uuid + strings.Join(funcRight, "")
		caller.FuncName = strings.Replace(funcArea[0], "(", "", -1) // 去除函数名中的 (
		caller.FuncArgs = funcArea[1 : len(funcArea)-1]
		for _, arg := range caller.FuncArgs {
			if !isStringNumberOrVar(arg) {
				return nil, errors.New(fmt.Sprintf(`invalid param for %s`, arg))
			}
		}
		callers = append(callers, caller)
	}
	return callers, nil
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

func analysisLexer(s string) (metas []string, lexers []string, err error) {
	metas = make([]string, 0)
	lexers = make([]string, 0)
	// 不断地进行词法解析,解析完或者报错
	for {
		s = strings.TrimSpace(s)
		if s == "" {
			// 解析完
			return metas, lexers, nil
		}
		// 标识是否分析到一个词语
		flag := false
		for index, regex := range regexs {
			reg := regexp.MustCompile(regex)
			find := reg.FindString(s)
			if find != "" { // 找到一个词语
				metas = append(metas, find)
				lexers = append(lexers, regexLexers[index])
				s = strings.Replace(s, find, "", 1)
				flag = true
				break
			}
		}
		// 解析报错
		if !flag {
			return metas, lexers, errors.New(fmt.Sprintf("%s is an error lexer data", s))
		}
	}
}
