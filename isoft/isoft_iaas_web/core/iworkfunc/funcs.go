package iworkfunc

import (
	"github.com/pkg/errors"
	"isoft/isoft/common/stringutil"
	"path/filepath"
	"strconv"
	"strings"
)

type IWorkFuncProxy struct{}

func (this *IWorkFuncProxy) stringsContains(args []interface{}) interface{} {
	return strings.Contains(args[0].(string), args[1].(string))
}

func (this *IWorkFuncProxy) stringsHasSuffix(args []interface{}) interface{} {
	return strings.HasSuffix(args[0].(string), args[1].(string))
}

func (this *IWorkFuncProxy) stringsHasPrefix(args []interface{}) interface{} {
	return strings.HasPrefix(args[0].(string), args[1].(string))
}

func (this *IWorkFuncProxy) stringsToLower(args []interface{}) interface{} {
	return strings.ToLower(args[0].(string))
}

func (this *IWorkFuncProxy) stringsToUpper(args []interface{}) interface{} {
	return strings.ToUpper(args[0].(string))
}

func (this *IWorkFuncProxy) stringsJoin(args []interface{}) interface{} {
	sargs := make([]string, 0)
	for _, arg := range args {
		sargs = append(sargs, arg.(string))
	}
	return strings.Join(sargs, "")
}

func (this *IWorkFuncProxy) int64Add(args []interface{}) interface{} {
	sargs := parseArgsToInt64Arr(args)
	checkArgsAmount(sargs, 2)
	return sargs[0] + sargs[1]
}

func (this *IWorkFuncProxy) int64Sub(args []interface{}) interface{} {
	sargs := parseArgsToInt64Arr(args)
	checkArgsAmount(sargs, 2)
	return sargs[0] - sargs[1]
}

func (this *IWorkFuncProxy) int64Gt(args []interface{}) interface{} {
	sargs := parseArgsToInt64Arr(args)
	checkArgsAmount(sargs, 2)
	return sargs[0] > sargs[1]
}

func (this *IWorkFuncProxy) int64Lt(args []interface{}) interface{} {
	sargs := parseArgsToInt64Arr(args)
	checkArgsAmount(sargs, 2)
	return sargs[0] < sargs[1]
}

func (this *IWorkFuncProxy) int64Eq(args []interface{}) interface{} {
	sargs := parseArgsToInt64Arr(args)
	checkArgsAmount(sargs, 2)
	return sargs[0] == sargs[1]
}

func (this *IWorkFuncProxy) int64Multi(args []interface{}) interface{} {
	sargs := parseArgsToInt64Arr(args)
	checkArgsAmount(sargs, 2)
	return sargs[0] * sargs[1]
}

func checkArgsAmount(sargs []int64, amount int) {
	if len(sargs) < amount {
		panic(errors.New("参数个数不足或者参数类型有误！"))
	}
}

func parseArgsToInt64Arr(args []interface{}) []int64 {
	sargs := make([]int64, 0)
	for _, arg := range args {
		if _arg, ok := arg.(int64); ok {
			sargs = append(sargs, _arg)
		} else if _arg, ok := arg.(int); ok {
			sargs = append(sargs, int64(_arg))
		} else if _arg, ok := arg.(string); ok {
			if _arg, err := strconv.ParseInt(_arg, 10, 64); err == nil {
				sargs = append(sargs, _arg)
			}
		}
	}
	return sargs
}

func (this *IWorkFuncProxy) stringsJoinWithSep(args []interface{}) interface{} {
	sargs := make([]string, 0)
	for _, arg := range args {
		sargs = append(sargs, arg.(string))
	}
	return strings.Join(sargs[:len(args)-1], sargs[len(args)-1])
}

func (this *IWorkFuncProxy) boolOr(args []interface{}) interface{} {
	sargs := make([]bool, 0)
	for _, arg := range args {
		sargs = append(sargs, arg.(bool))
	}
	return sargs[0] || sargs[1]
}

func (this *IWorkFuncProxy) boolAnd(args []interface{}) interface{} {
	sargs := make([]bool, 0)
	for _, arg := range args {
		sargs = append(sargs, arg.(bool))
	}
	return sargs[0] && sargs[1]
}

func (this *IWorkFuncProxy) boolNot(args []interface{}) interface{} {
	sargs := make([]bool, 0)
	for _, arg := range args {
		sargs = append(sargs, arg.(bool))
	}
	return !sargs[0]
}

func (this *IWorkFuncProxy) stringsUUID(args []interface{}) interface{} {
	return stringutil.RandomUUID()
}

func (this *IWorkFuncProxy) stringsCheckEmpty(args []interface{}) interface{} {
	return args[0].(string) == ""
}

func (this *IWorkFuncProxy) checkEmpty(args []interface{}) interface{} {
	return args[0] == nil
}

func (this *IWorkFuncProxy) getDirPath(args []interface{}) string {
	return filepath.Dir(args[0].(string))
}
