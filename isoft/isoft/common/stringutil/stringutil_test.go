package stringutil

import (
	"fmt"
	"strings"
	"testing"
)

func Test_Rune(t *testing.T) {
	fmt.Print(strings.IndexRune("hello中国.txt",'.'))
}

