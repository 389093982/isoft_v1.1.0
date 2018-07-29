package main

import (
	"isoft/isoft/build/build_pack"
	"os"
	"path/filepath"
)

func main() {
	gopath := os.Getenv("GOPATH")
	packApps := build_pack.ReadPackApp(filepath.Join(gopath, "src/isoft/isoft/build/build_pack/pack.xml"))
	build_pack.StartAllPack(&packApps, "")
}
