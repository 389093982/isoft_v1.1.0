package sshutil

import (
	"fmt"
	"testing"
)

func Test_SSHConnectTest(t *testing.T) {
	err := SSHConnectTest("root", "TU3d26L6GcfWqGn", "193.112.162.61", 22)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("sshutil connect test successful...")
	}
}
