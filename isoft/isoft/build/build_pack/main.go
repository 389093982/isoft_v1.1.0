package main

import (
	"isoft/isoft/build/build_pack/pack"
	"log"
	"os"
)

func main() {
	// 设置环境变量 GOOS:目标平台的操作系统
	// 交叉编译依赖下面几个环境变量:
	// $GOARCH 目标平台（编译后的目标平台）的处理器架构（386、amd64、arm）
	// $GOOS 目标平台（编译后的目标平台）的操作系统（darwin、freebsd、linux、windows）
	GOOS := os.Getenv("GOOS")
	os.Setenv("GOOS", "linux")
	defer os.Setenv("GOOS", GOOS)

	packApps := pack.ReadPackApp("./pack.xml")
	err := pack.StartAllPackTask(&packApps, "")
	if err != nil {
		log.Println(err)
	}
}
