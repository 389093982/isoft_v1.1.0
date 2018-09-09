package sftputil

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"isoft/isoft/common/fileutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func SFTPConnect(user, password, host string, port int) (*sftp.Client, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		sshClient    *ssh.Client
		sftpClient   *sftp.Client
		err          error
	)
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User:            user,
		Auth:            auth,
		Timeout:         30 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	addr = fmt.Sprintf("%s:%d", host, port)

	if sshClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		return nil, err
	}

	return sftpClient, nil
}

func SFTPClientChmodXForShell(sftpClient *sftp.Client, filepath string) {
	if strings.HasSuffix(filepath, ".sh") {
		sftpClient.Chmod(filepath, 700)
	}
}

func SFTPClientFileCopy(sftpClient *sftp.Client, localFilePath, remoteDir string) error {
	if ok, err := fileutil.PathExists(localFilePath); ok == false {
		return err
	}

	srcFile, err := os.Open(localFilePath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	var remoteFileName = filepath.Base(localFilePath)
	err = sftpClient.MkdirAll(remoteDir)
	if err != nil {
		return err
	}

	dstFilePath := path.Join(remoteDir, remoteFileName)
	dstFile, err := sftpClient.Create(dstFilePath)

	SFTPClientChmodXForShell(sftpClient, dstFilePath)

	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()

	buf := make([]byte, 1024)
	for {
		n, _ := srcFile.Read(buf)
		if n == 0 {
			break
		}
		// buf[:n]
		dstFile.Write(buf[:n])
	}

	return nil
}

// localFilePath、remoteDir 本地文件路径和远程机器上的文件夹
func SFTPFileCopy(user, password, host string, port int, localFilePath, remoteDir string) {
	var (
		err        error
		sftpClient *sftp.Client
	)
	// 创建 sftp 连接
	sftpClient, err = SFTPConnect(user, password, host, port)
	if err != nil {
		log.Fatal(err)
	}
	defer sftpClient.Close()

	SFTPClientFileCopy(sftpClient, localFilePath, remoteDir)
}

// localFilePath、remoteDir 本地文件路径和远程机器上的文件夹
func SFTPDirectoryCopy(user, password, host string, port int, localDirectoryPath, remoteDir string) {
	var (
		err        error
		sftpClient *sftp.Client
	)
	// 创建 sftp 连接
	sftpClient, err = SFTPConnect(user, password, host, port)
	if err != nil {
		log.Fatal(err)
	}
	defer sftpClient.Close()

	SFTPClientCopyDirectoryInto(sftpClient, localDirectoryPath, remoteDir)
}

// localDirectoryPath、remoteDir 本地文件夹路径和远程机器上的文件夹,拷贝本地文件夹到远程机器指定文件夹里面
func SFTPClientCopyDirectoryInto(sftpClient *sftp.Client, localDirectoryPath, remoteDir string) error {
	filepaths, _, err := fileutil.GetAllFile(localDirectoryPath, true)
	if err != nil {
		return err
	}

	localDirectoryPath = fileutil.ChangeToLinuxSeparator(localDirectoryPath)
	targetDirectoryPath := fileutil.ChangeToLinuxSeparator(remoteDir + "/" + path.Base(fileutil.ChangeToLinuxSeparator(localDirectoryPath)))

	for _, filepath := range filepaths {
		if ok, _ := fileutil.PathExists(filepath); ok {
			if !fileutil.IsDir(filepath) {
				localFilePath := fileutil.ChangeToLinuxSeparator(filepath)

				// 目标机器对应的文件路径
				remoteFilePath := strings.Replace(localFilePath, localDirectoryPath, targetDirectoryPath, -1)

				err := SFTPClientFileCopy(sftpClient, filepath, path.Dir(remoteFilePath))
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
