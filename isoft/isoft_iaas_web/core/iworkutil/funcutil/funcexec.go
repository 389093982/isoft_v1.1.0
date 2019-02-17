package funcutil

import (
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

func IworkStringsToUpper(args []interface{}) interface{} {
	return strings.ToUpper(args[0].(string))
}

func IworkStringsJoin(args []interface{}) interface{} {
	sargs := make([]string, 0)
	for _, arg := range args {
		sargs = append(sargs, arg.(string))
	}
	return strings.Join(sargs, "")
}

func IworkInt64Add(args []interface{}) interface{} {
	sargs := parseArgsToInt64Arr(args)
	if len(sargs) == 2{
		return sargs[0] + sargs[1]
	}else{
		panic(errors.New("参数个数不足或者参数类型有误！"))
	}
}

func IworkInt64Sub(args []interface{}) interface{} {
	sargs := parseArgsToInt64Arr(args)
	if len(sargs) == 2{
		return sargs[0] - sargs[1]
	}else{
		panic(errors.New("参数个数不足或者参数类型有误！"))
	}
}

func parseArgsToInt64Arr(args []interface{}) []int64 {
	sargs := make([]int64, 0)
	for _, arg := range args {
		if _arg, ok := arg.(int64); ok {
			sargs = append(sargs, _arg)
		} else if _arg, ok := arg.(string); ok {
			if _arg, err := strconv.ParseInt(_arg, 10, 64); err == nil {
				sargs = append(sargs, _arg)
			}
		}
	}
	return sargs
}

func IworkInt64Multi(args []interface{}) interface{} {
	sargs := parseArgsToInt64Arr(args)
	if len(sargs) == 2{
		return sargs[0] * sargs[1]
	}else{
		panic(errors.New("参数有误..."))
	}
}

func IworkStringsJoinWithSep(args []interface{}) interface{} {
	sargs := make([]string, 0)
	for _, arg := range args {
		sargs = append(sargs, arg.(string))
	}
	return strings.Join(sargs[:len(args)-1], sargs[len(args)-1])
}
