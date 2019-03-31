package datatypeutil

import (
	"fmt"
	"testing"
)

func Test_GetAllFile(t *testing.T) {
	s := ReverseSlice([]int{1, 2, 3, 4, 5}).([]int)
	fmt.Println(len(s))
	fmt.Println(s)
}
