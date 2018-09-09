package sftputil

import (
	"fmt"
	"testing"
)

func Test_SFTPConnectTest(t *testing.T) {
	// 用来测试的本地文件路径 和 远程机器上的文件夹
	var localFilePath = "D:/zhourui/program/go/goland_workspace/src/isoft/isoft_deploy_web/shell/common/port_check.sh"
	var remoteDir = "/mydata/deploy_home"

	SFTPFileCopy("root", "TU3d26L6GcfWqGn", "193.112.162.61", 22, localFilePath, remoteDir)

	fmt.Println("copy file to remote server finished!")
}

func Test_SFTPDirectoryCopy(t *testing.T) {
	// 用来测试的本地文件路径 和 远程机器上的文件夹
	var localDirectoryPath = "D:/zhourui/program/go/goland_workspace/src/isoft/isoft_deploy_web/shell"
	var remoteDir = "/mydata/deploy_home"

	SFTPDirectoryCopy("root", "TU3d26L6GcfWqGn", "193.112.162.61", 22, localDirectoryPath, remoteDir)

	fmt.Println("copy file to remote server finished!")
}
