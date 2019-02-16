package funcutil

import "strings"

func IworkStringsToUpper(args ...interface{}) interface{} {
	return strings.ToUpper(args[0].(string))
}

func IworkStringsJoin(args ...interface{}) interface{} {
	sargs := make([]string, 0)
	for _, arg := range args {
		sargs = append(sargs, arg.(string))
	}
	return strings.Join(sargs, "")
}

func IworkStringsJoinWithSep(args ...interface{}) interface{} {
	sargs := make([]string, 0)
	for _, arg := range args {
		sargs = append(sargs, arg.(string))
	}
	return strings.Join(sargs[:len(args)-1], sargs[len(args)-1])
}
