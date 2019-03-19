package cmdutil

import (
	"io"
	"os/exec"
)

// 执行系统命令,第一个参数是命令名称,第二个参数是参数列表
func RunCommand(stdout, stderr io.Writer, name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	if err := cmd.Start(); err != nil {
		return err
	}
	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil
}
