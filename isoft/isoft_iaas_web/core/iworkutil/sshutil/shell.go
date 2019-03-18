package sshutil

import (
	"io"
)

func RunSSHShellCommandOnly(sshAccount, sshPwd, sshIp, command string) error {
	return RunSSHShellCommand(sshAccount, sshPwd, sshIp, command, nil, nil)
}

func RunSSHShellCommand(sshAccount, sshPwd, sshIp, command string, stdout, stderr io.Writer) error {
	client, session, err := SSHConnect(sshAccount, sshPwd, sshIp, 22)
	if err != nil {
		return err
	}
	// 避免 err 导致 session 空指针异常,需要放在 err 判断之后
	defer client.Close()
	defer session.Close()
	session.Stdout = stdout
	session.Stderr = stderr
	err = session.Run(command)
	if err != nil {
		return err
	}
	return nil
}
