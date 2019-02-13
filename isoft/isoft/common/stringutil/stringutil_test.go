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
