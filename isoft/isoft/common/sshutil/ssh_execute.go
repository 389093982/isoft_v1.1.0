package sshutil

import (
	"io"
)

func RunSSHShellCommandOnly(sshAccount, sshPwd, sshIp, command string) error {
	return RunSSHShellCommand(sshAccount, sshPwd, sshIp, command, nil, nil)
}

func RunSSHShellCommand(sshAccount, sshPwd, sshIp, command string, stdout, stderr io.Writer) error {
	sshClient, err := SSHConnect(sshAccount, sshPwd, sshIp, 22)
	defer sshClient.Close()
	if err != nil {
		return err
	}
	sshClient.Stdout = stdout
	sshClient.Stderr = stderr
	err = sshClient.Run(command)
	if err != nil {
		return err
	}
	return nil
}
