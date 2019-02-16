package htmlutil

import (
	"fmt"
	"testing"
)

func Test_GetAllHref(t *testing.T) {
	results := GetAllHref("http://www.baidu.com")
	for _, result := range results {
		fmt.Println(result)
	}
}
