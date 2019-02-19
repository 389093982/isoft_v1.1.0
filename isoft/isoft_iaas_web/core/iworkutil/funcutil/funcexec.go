package funcutil

import (
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

type IWorkFuncProxy struct {}

func (this *IWorkFuncProxy) IworkStringsContains(args []interface{}) interface{} {
	return strings.Contains(args[0].(string), args[1].(string))
}

func (this *IWorkFuncProxy) IworkStringsHasSuffix(args []interface{}) interface{} {
	return strings.HasSuffix(args[0].(string), args[1].(string))
}

func (this *IWorkFuncProxy) IworkStringsHasPrefix(args []interface{}) interface{} {
	return strings.HasPrefix(args[0].(string), args[1].(string))
}

func (this *IWorkFuncProxy) IworkStringsToLower(args []interface{}) interface{} {
	return strings.ToLower(args[0].(string))
}

func (this *IWorkFuncProxy) IworkStringsToUpper(args []interface{}) interface{} {
	return strings.ToUpper(args[0].(string))
}

func (this *IWorkFuncProxy) IworkStringsJoin(args []interface{}) interface{} {
	sargs := make([]string, 0)
	for _, arg := range args {
		sargs = append(sargs, arg.(string))
	}
	return strings.Join(sargs, "")
}

func (this *IWorkFuncProxy) IworkInt64Add(args []interface{}) interface{} {
	sargs := parseArgsToInt64Arr(args)
	checkArgsAmount(sargs, 2)
	return sargs[0] + sargs[1]
}

func (this *IWorkFuncProxy) IworkInt64Sub(args []interface{}) interface{} {
	sargs := parseArgsToInt64Arr(args)
	checkArgsAmount(sargs, 2)
	return sargs[0] - sargs[1]
}

func (this *IWorkFuncProxy) IworkInt64Gt(args []interface{}) interface{} {
	sargs := parseArgsToInt64Arr(args)
	checkArgsAmount(sargs, 2)
	return sargs[0] > sargs[1]
}

func (this *IWorkFuncProxy) IworkInt64Lt(args []interface{}) interface{} {
	sargs := parseArgsToInt64Arr(args)
	checkArgsAmount(sargs, 2)
	return sargs[0] < sargs[1]
}

func (this *IWorkFuncProxy) IworkInt64Eq(args []interface{}) interface{} {
	sargs := parseArgsToInt64Arr(args)
	checkArgsAmount(sargs, 2)
	return sargs[0] == sargs[1]
}

func (this *IWorkFuncProxy) IworkInt64Multi(args []interface{}) interface{} {
	sargs := parseArgsToInt64Arr(args)
	checkArgsAmount(sargs, 2)
	return sargs[0] * sargs[1]
}

func checkArgsAmount(sargs []int64, amount int)  {
	if len(sargs) < amount{
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

func (this *IWorkFuncProxy) IworkStringsJoinWithSep(args []interface{}) interface{} {
	sargs := make([]string, 0)
	for _, arg := range args {
		sargs = append(sargs, arg.(string))
	}
	return strings.Join(sargs[:len(args)-1], sargs[len(args)-1])
}


