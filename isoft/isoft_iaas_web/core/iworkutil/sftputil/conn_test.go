package sftputil

import (
	"fmt"
	"testing"
)

func Test_SFTPConnect(t *testing.T) {
	_, _, err := SFTPConnect("root", "TU3d26L6GcfWqGn", "193.112.162.61", 22)
	fmt.Println(err)
}
