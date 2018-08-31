package apppath

import (
	"os"
	"os/exec"
	"path/filepath"
)

// 获取应用根目录
func GetAPPRootPath() string {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return ""
	}
	p, err := filepath.Abs(file)
	if err != nil {
		return ""
	}
	return filepath.Dir(p)
}
