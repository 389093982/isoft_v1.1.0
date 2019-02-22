package stringutil

import (
	"fmt"
	"strings"
	"testing"
)

func Test_Rune(t *testing.T) {
	fmt.Print(strings.IndexRune("hello中国.txt", '.'))
}

func Test_GetTypeOfInterface(t *testing.T) {
	fmt.Print(GetTypeOfInterface("demo"))
}

func Test_GetNoRepeatSubStringWithRegexp(t *testing.T)  {
	ss := GetNoRepeatSubStringWithRegexp("a$hello.hello", `\$[a-zA-Z0-9_]+`)
	for _,s := range ss {
		fmt.Println(s)
	}
}