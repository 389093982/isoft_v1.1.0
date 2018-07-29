package build_pack

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"isoft/isoft/common/fileutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

type PackApps struct {
	XMLName  xml.Name  `xml:"packapps"` // 指定最外层的标签为 packapps
	PackApps []PackApp `xml:"packapp"`  // 读取packapp标签下的内容
}

type PackApp struct {
	AppName string `xml:"appName"`
	Apppath string `xml:"apppath"`
}

func ReadPackApp(filepath string) (packApps PackApps) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Println(err)
	}
	err = xml.Unmarshal(content, &packApps)
	if err != nil {
		log.Println(err)
	}
	return packApps
}

// 同步所有目录
func StartAllPack(packApps *PackApps, filterAppName string) {
	pas := packApps.PackApps
	for _, packApp := range pas {
		if filterAppName == "" || (filterAppName != "" && filterAppName == packApp.AppName) {
			StartOnePack(packApp.AppName, packApp.Apppath)
		}
	}
}

func StartOnePack(appName, appPath string) {
	gopath := os.Getenv("GOPATH")
	appRootPath := fileutil.ChangeToLinuxSeparator(filepath.Join(gopath, appPath))
	err := os.Chdir(appRootPath)
	if err != nil {
		log.Println(err.Error())
		return
	}

	// 执行系统命令
	// 第一个参数是命令名称
	// 后面参数可以有多个,命令参数
	beeCommandPath := fileutil.ChangeToLinuxSeparator(filepath.Join(gopath, "bin/bee.exe"))
	cmd := exec.Command(beeCommandPath, "pack", "-be", "GOOS=linux")
	// 获取输出对象,可以从该对象中读取输出结果
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Println(err)
	}
	// 保证关闭输出流
	defer stdout.Close()
	// 运行命令
	if err := cmd.Start(); err != nil {
		log.Println(err)
	}
	// 读取输出结果
	opBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(opBytes))

	// 移动文件
	err = os.Rename(fmt.Sprintf("./%s.tar.gz", appName), fmt.Sprintf("D:/%s.tar.gz", appName))
	if err != nil {
		log.Println(err)
	}
}
