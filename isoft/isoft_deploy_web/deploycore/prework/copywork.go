package prework

import (
	"fmt"
	"github.com/astaxie/beego"
	"isoft/isoft/common/fileutil"
	"os"
	"path/filepath"
)

var (
	SFTP_SRC_DIR, SFTP_LOCAL_DEPLOY_HOME, SFTP_TARGET_DEPLOY_HOME string
)

func init() {
	SFTP_SRC_DIR = beego.AppConfig.String("sftp.src.dir")
	SFTP_LOCAL_DEPLOY_HOME = beego.AppConfig.String("sftp.local.deploy_home.default")
	SFTP_TARGET_DEPLOY_HOME = beego.AppConfig.String("sftp.target.deploy_home.default")
}

// 拷贝项目内部的 shell 文件夹到 deploy_home 中去
func runCopyShellWork()  {
	// 先进行清理
	cleanDeployHomeShell()
	// 再进行拷贝
	copyShellToDeployHome()
}

func copyShellToDeployHome()  {
	err := fileutil.CopyDir(fileutil.ChangeToLinuxSeparator(filepath.Join(SFTP_SRC_DIR,"shell")),
		fileutil.ChangeToLinuxSeparator(filepath.Join(SFTP_LOCAL_DEPLOY_HOME,"shell")))

	if err != nil{
		fmt.Println(fmt.Sprintf("execute copyShellToDeployHome error, %s", err.Error()))
	}else{
		fmt.Println("execute copyShellToDeployHome successful...")
	}
}

func cleanDeployHomeShell()  {
	err := os.RemoveAll(fileutil.ChangeToLinuxSeparator(filepath.Join(SFTP_LOCAL_DEPLOY_HOME,"shell")))
	if err != nil{
		fmt.Println(fmt.Sprintf("execute cleanDeployHomeShell error, %s", err.Error()))
	}else{
		fmt.Println("execute cleanDeployHomeShell successful...")
	}
}
