package iworkfunc

import (
	"fmt"
	"github.com/pkg/errors"
	"isoft/isoft/common/stringutil"
	"reflect"
	"strings"
)

// 函数执行类
type FuncExecutor struct {
	FuncUUID       string   // 函数唯一性 id
	FuncName       string   // 函数名称
	FuncArgs       []string // 函数参数列表
	FuncLeftIndex  int      // 函数整体在表达式中的左索引位置
	FuncRightIndex int      // 函数整体在表达式中的右索引位置
}

func GetTrimFuncExecutor(executor *FuncExecutor) {
	executor.FuncName = strings.TrimSpace(executor.FuncName)
	funcArgs := make([]string, 0)
	for _, arg := range executor.FuncArgs {
		funcArgs = append(funcArgs, strings.TrimSpace(arg))
	}
	executor.FuncArgs = funcArgs
}

func GetAllFuncExecutor(expression string) []*FuncExecutor {
	executors := make([]*FuncExecutor, 0)
	for {
		executor := GetFuncExecutorFromExpression(expression)
		if executor != nil {
			executors = append(executors, executor)
			// 在 expression 中将函数用占位符代替
			expression = strings.Replace(expression, expression[executor.FuncLeftIndex:executor.FuncRightIndex+1], executor.FuncUUID, 1)
		} else {
			break
		}
	}
	return executors
}

func GetFuncExecutorFromExpression(expression string) *FuncExecutor {
	// 获取表达式中所有左括号的索引
	for _, leftBracketIndex := range GetAllLeftBracketIndex(expression) {
		// 判断表达式 expression 中左括号索引 leftBracketIndex 后面是否有直接右括号
		if bol, rightBracketIndex := CheckHasNearRightBracket(leftBracketIndex, expression); bol == true {
			// 找到了优先级最高的执行函数
			// 获取参数字符串
			argStr := expression[leftBracketIndex+1 : rightBracketIndex]
			// 分割得到所有参数
			args := strings.Split(argStr, ",")
			// 获取函数名称
			_expression := expression[:leftBracketIndex]
			var funcName string
			if funcNameIndex := strings.LastIndex(strings.ToUpper(_expression), "IWORK"); funcNameIndex >= 0 {
				// 函数名称以 Iwork 开头
				funcName = _expression[funcNameIndex:]
			} else {
				panic(errors.New(fmt.Sprintf("unsupport funcName in %s, must startwith %s", _expression, "Iwork")))
			}
			return &FuncExecutor{
				FuncUUID:       stringutil.RandomUUID(),
				FuncName:       funcName,
				FuncArgs:       args,
				FuncLeftIndex:  leftBracketIndex - len(funcName),
				FuncRightIndex: rightBracketIndex,
			}
		}
	}
	return nil
}

// 获取表达式中所有左括号的索引
func GetAllLeftBracketIndex(expression string) []int {
	leftBracketIndexs := make([]int, 0)
	for i := 0; i < len(expression); i++ {
		if expression[i] == '(' {
			leftBracketIndexs = append(leftBracketIndexs, i)
		}
	}
	return leftBracketIndexs
}

// 判断表达式 expression 中左括号索引 leftBracketIndex 后面是否有直接右括号
// 返回是否直接跟随右括号,以及右括号的索引位置
func CheckHasNearRightBracket(leftBracketIndex int, expression string) (bool, int) {
	flag := true
	for i := leftBracketIndex + 1; i < len(expression); i++ {
		if expression[i] == '(' {
			flag = false
		} else if expression[i] == ')' && flag == true {
			return true, i
		}
	}
	return false, -1
}

// 编码特殊字符, // 对转义字符 \, \; \( \) 等进行编码
func EncodeSpecialForParamVaule(paramVaule string) string {
	paramVaule = strings.Replace(paramVaule, "\\\\n", "__newline__", -1)
	paramVaule = strings.Replace(paramVaule, "\\(", "__leftBracket__", -1)
	paramVaule = strings.Replace(paramVaule, "\\)", "__rightBracket__", -1)
	paramVaule = strings.Replace(paramVaule, "\\,", "__comma__", -1)
	paramVaule = strings.Replace(paramVaule, "\\;", "__semicolon__", -1)
	return paramVaule
}

// 解码特殊字符
func DncodeSpecialForParamVaule(paramVaule string) string {
	paramVaule = strings.Replace(paramVaule, "__newline__", "\n", -1)
	paramVaule = strings.Replace(paramVaule, "__leftBracket__", "(", -1)
	paramVaule = strings.Replace(paramVaule, "__rightBracket__", ")", -1)
	paramVaule = strings.Replace(paramVaule, "__comma__", ",", -1)
	paramVaule = strings.Replace(paramVaule, "__semicolon__", ";", -1)
	return paramVaule
}

func CallFuncExecutor(executor *FuncExecutor, args []interface{}) interface{} {
	defer func() {
		if err := recover(); err != nil {
			panic(errors.New(fmt.Sprintf("execute func error %s, %v, error msg: %s", executor.FuncName, args, err.(error).Error())))
		}
	}()

	proxy := &IWorkFuncProxy{}
	m := reflect.ValueOf(proxy).MethodByName(executor.FuncName)
	rtn := m.Call([]reflect.Value{reflect.ValueOf(args)})
	return rtn[0].Interface()
}
