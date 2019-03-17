package sftputil

import (
	"github.com/pkg/sftp"
	"isoft/isoft/common/fileutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func SFTPClientChmodXForShell(sftpClient *sftp.Client, filepath string) {
	if strings.HasSuffix(filepath, ".sh") {
		sftpClient.Chmod(filepath, 700)
	}
}

func sftpClientFileCopy(sftpClient *sftp.Client, localFilePath, remoteDir string) error {
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
		return err
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
func SFTPFileCopy(user, password, host string, port int, localFilePath, remoteDir string) error {
	// 创建 sftp 连接
	sshClient, sftpClient, err := SFTPConnect(user, password, host, port)
	if err != nil {
		return err
	}
	defer sshClient.Close()
	defer sftpClient.Close()
	return sftpClientFileCopy(sftpClient, localFilePath, remoteDir)
}

// localFilePath、remoteDir 本地文件路径和远程机器上的文件夹,相当于移动重命名
func SFTPDirectoryRenameCopy(user, password, host string, port int, localDirectoryPath, remoteDir string) error {
	// 创建 sftp 连接
	sshClient, sftpClient, err := SFTPConnect(user, password, host, port)
	if err != nil {
		return err
	}
	defer sshClient.Close()
	defer sftpClient.Close()
	filepaths, _, err := fileutil.GetAllFile(localDirectoryPath, false)
	if err != nil {
		return err
	}
	for _, fpath := range filepaths {
		if fileutil.IsFile(fpath) {
			err = sftpClientFileCopy(sftpClient, fpath, remoteDir)
			if err != nil {
				return err
			}
		} else {
			err = SFTPDirectoryRenameCopy(user, password, host, port, fpath, remoteDir+"/"+filepath.Base(fpath))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// localFilePath、remoteDir 本地文件路径和远程机器上的文件夹,相当于移动到目标机器子目录
func SFTPDirectoryCopy(user, password, host string, port int, localDirectoryPath, remoteDir string) error {
	// 创建 sftp 连接
	sshClient, sftpClient, err := SFTPConnect(user, password, host, port)
	if err != nil {
		return err
	}
	defer sshClient.Close()
	defer sftpClient.Close()
	return sftpClientCopyDirectoryInto(sftpClient, localDirectoryPath, remoteDir)
}

// localDirectoryPath、remoteDir 本地文件夹路径和远程机器上的文件夹,拷贝本地文件夹到远程机器指定文件夹里面
func sftpClientCopyDirectoryInto(sftpClient *sftp.Client, localDirectoryPath, remoteDir string) error {
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

				err := sftpClientFileCopy(sftpClient, filepath, path.Dir(remoteFilePath))
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
