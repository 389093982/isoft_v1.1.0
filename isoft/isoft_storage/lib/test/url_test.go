package test

import (
	"fmt"
	"net/url"
	"testing"
)

func Test_Url(t *testing.T) {
	fmt.Println(url.PathEscape("http://ab/我的/de.txt/ddd"))
}

