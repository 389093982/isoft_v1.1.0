package pack

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
	PackType string `xml:"packType"`
	AppName  string `xml:"appName"`
	Apppath  string `xml:"apppath"`
	PackArgs string `xml:"packArgs"`
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

// 开始打包所有任务
func StartAllPackTask(packApps *PackApps, filterAppName string) (err error) {
	for _, packApp := range packApps.PackApps {
		if filterAppName == "" || (filterAppName == packApp.AppName) {
			if err = ChangeAppRootPath(packApp.Apppath); err != nil {
				return
			}
			if packApp.PackType == "beego" {
				err = StartPackBeegoTask(packApp.AppName, packApp.Apppath)
			} else if packApp.PackType == "common" {
				err = StartPackCommonTask(packApp.AppName, packApp.Apppath, packApp)
			}
		}
	}
	return
}

func StartPackCommonTask(appName, appPath string, packApp PackApp) (err error) {
	// 运行命令
	RunCommand("go", "build", packApp.PackArgs)
	// 移动文件
	err = os.Rename(fmt.Sprintf("./%s", appName), fmt.Sprintf("D:/%s", appName))
	return
}

func StartPackBeegoTask(appName, appPath string) (err error) {
	gopath := os.Getenv("GOPATH")
	// 运行命令
	RunCommand(fileutil.ChangeToLinuxSeparator(filepath.Join(gopath, "bin/bee.exe")), "pack", "-be", "GOOS=linux")
	// 移动文件
	err = os.Rename(fmt.Sprintf("./%s.tar.gz", appName), fmt.Sprintf("D:/%s.tar.gz", appName))
	return
}

func ChangeAppRootPath(appPath string) (err error) {
	gopath := os.Getenv("GOPATH")
	appRootPath := fileutil.ChangeToLinuxSeparator(filepath.Join(gopath, appPath))
	err = os.Chdir(appRootPath)
	return
}

// 执行系统命令,第一个参数是命令名称,第二个参数是参数列表
func RunCommand(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Println(err)
	}
	defer stdout.Close()
	if err := cmd.Start(); err != nil {
		log.Println(err)
	}
	opBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(opBytes))
}
