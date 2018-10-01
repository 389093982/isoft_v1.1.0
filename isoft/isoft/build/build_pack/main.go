package main

import (
	"isoft/isoft/build/build_pack/pack"
	"log"
	"os"
)

func main() {
	// 设置环境变量 GOOS:目标平台的操作系统
	GOOS := os.Getenv("GOOS")
	os.Setenv("GOOS", "linux")
	defer os.Setenv("GOOS", GOOS)

	packApps := pack.ReadPackApp("./pack.xml")
	err := pack.StartAllPackTask(&packApps, "dataServer")
	if err != nil {
		log.Println(err)
	}
}
