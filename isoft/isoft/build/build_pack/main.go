package main

import (
	"isoft/isoft/build/build_pack/pack"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// 设置环境变量 GOOS:目标平台的操作系统
	GOOS := os.Getenv("GOOS")
	os.Setenv("GOOS", "linux")
	defer os.Setenv("GOOS", GOOS)

	gopath := os.Getenv("GOPATH")
	packApps := pack.ReadPackApp(filepath.Join(gopath, "src/isoft/isoft/build/build_pack/pack.xml"))
	err := pack.StartAllPackTask(&packApps, "isoft_blog_web")
	if err != nil {
		log.Println(err)
	}
}
