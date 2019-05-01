package iworkanalyzer

import (
	"fmt"
	"github.com/pkg/errors"
	"isoft/isoft/common/stringutil"
	"strings"
)

// 函数执行类
type FuncExecutor struct {
	FuncUUID       string // 函数唯一性 id
	FuncRealName   string
	FuncName       string // 函数实际名称
	FuncLexersName string // 函数 lexers 名称
	FuncRealArgs   []string
	FuncArgs       []string // 函数实际参数列表
	FuncLexersArgs []string // 函数 lexers 参数列表
	FuncLeftIndex  int      // 函数整体在表达式中的左索引位置
	FuncRightIndex int      // 函数整体在表达式中的右索引位置
}

func (c FuncExecutor) String() string {
	return fmt.Sprintf("FuncExecutor-[FuncUUID:%v, FuncLexersName:%v, FuncLexersArgs:%v, FuncLeftIndex:%v, FuncRightIndex:%v]",
		c.FuncUUID, c.FuncLexersName, c.FuncLexersArgs, c.FuncLeftIndex, c.FuncRightIndex)
}

func GetAllFuncExecutor(metas []string, lexers []string) ([]*FuncExecutor, error) {
	lexersExpression := strings.Join(lexers, "")
	executors := make([]*FuncExecutor, 0)
	for {
		executor, err := GetPriorityFuncExecutorFromLexersExpression(lexersExpression)
		if err != nil {
			return nil, err
		}
		if executor != nil {
			executors = append(executors, executor)
			// 在 expression 中将函数用占位符代替
			lexersExpression = lexersExpression[:executor.FuncLeftIndex] + executor.FuncUUID + lexersExpression[executor.FuncRightIndex+1:]
		} else {
			break
		}
	}
	return executors, nil
}

// 获取优先级最高的函数执行体
// 含有 func( 必然有优先函数执行体
func GetPriorityFuncExecutorFromLexersExpression(lexersExpression string) (*FuncExecutor, error) {
	if !strings.Contains(lexersExpression, "func(") && !strings.Contains(lexersExpression, ")") {
		// 非函数类型表达式值
		return nil, nil
	}
	// 获取表达式中所有左括号的索引
	for _, leftBracketIndex := range GetAllLeftBracketIndex(lexersExpression) {
		// 判断表达式 expression 中左括号索引 leftBracketIndex 后面是否有直接右括号
		if bol, rightBracketIndex := CheckHasNearRightBracket(leftBracketIndex, lexersExpression); bol == true {
			// 找到了优先级最高的执行函数
			// 获取参数字符串
			argStr := lexersExpression[leftBracketIndex+1 : rightBracketIndex]
			// 分割得到所有参数
			args := strings.Split(argStr, ",")
			// 获取函数名称
			return &FuncExecutor{
				FuncUUID:       stringutil.RandomUUID(),
				FuncLexersName: "func",
				FuncLexersArgs: args,
				FuncLeftIndex:  leftBracketIndex - len("func"),
				FuncRightIndex: rightBracketIndex,
			}, nil
		}
	}
	return nil, errors.New("invalid func was found...")
}

// 获取表达式中所有左括号的索引
func GetAllLeftBracketIndex(lexersExpression string) []int {
	leftBracketIndexs := make([]int, 0)
	for i := 0; i < len(lexersExpression); i++ {
		if lexersExpression[i] == '(' {
			leftBracketIndexs = append(leftBracketIndexs, i)
		}
	}
	return leftBracketIndexs
}

// 判断表达式 expression 中左括号索引 leftBracketIndex 后面是否有直接右括号
// 返回是否直接跟随右括号,以及右括号的索引位置
func CheckHasNearRightBracket(leftBracketIndex int, lexersExpression string) (bool, int) {
	for i := leftBracketIndex + 1; i < len(lexersExpression); i++ {
		if lexersExpression[i] == '(' {
			return false, -1
		} else if lexersExpression[i] == ')' {
			return true, i
		}
	}
	return false, -1
}
