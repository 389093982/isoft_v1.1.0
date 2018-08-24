package deploy

import (
	"github.com/astaxie/beego"
	"isoft/isoft/common"
	"isoft/isoft/common/fileutil"
	"isoft/isoft_deploy_api/models"
	"path/filepath"
)

var (
	SFTP_SRC_DIR, SFTP_TARGET_DIR string
)

func init() {
	SFTP_SRC_DIR = beego.AppConfig.String("sftp.src.dir")
	SFTP_TARGET_DIR = beego.AppConfig.String("sftp.target.dir.default")
}

// 文件传输器类
type FileTransfer struct {
	// 本地文件路径
	LocalFilePath string
	// 远程虚拟机文件路径
	RemoteDir string
}

func (this *FileTransfer) Transfer(EnvInfo *models.EnvInfo) error {
	sftpClient, err := common.SFTPConnect(EnvInfo.EnvAccount, EnvInfo.EnvPasswd, EnvInfo.EnvIp, 22)
	if err != nil {
		return err
	}
	defer sftpClient.Close()

	if fileutil.IsDir(this.LocalFilePath) {
		common.SFTPClientCopyDirectoryInto(sftpClient, this.LocalFilePath, this.RemoteDir)
	} else {
		common.SFTPClientFileCopy(sftpClient, this.LocalFilePath, this.RemoteDir)
	}
	return nil
}

type IFileTransferCreator interface {
	PrepareFileTransfer() []*FileTransfer
}

type FileTransferCreator struct {
	ServiceInfo *models.ServiceInfo
	OperateType string
}

func (this *FileTransferCreator) PrepareFileTransfer() []*FileTransfer {
	fileTransfers := make([]*FileTransfer, 0)

	// 添加当前 ServiceInfo 所特有的文件传输列表
	switch this.ServiceInfo.ServiceType {
	case "beego":
		creator := BeegoFileTransferCreator{
			ServiceInfo: this.ServiceInfo,
			OperateType: this.OperateType,
		}
		_filetransfers := creator.PrepareFileTransfer()
		fileTransfers = append(fileTransfers, _filetransfers...)
		break
	}

	return fileTransfers
}

type BeegoFileTransferCreator struct {
	ServiceInfo *models.ServiceInfo
	OperateType string
}

// beego 项目部署文件传输
func (this *BeegoFileTransferCreator) BeegoDeployFileTransfer() []*FileTransfer {
	FileTransfers := make([]*FileTransfer, 0)
	// .tar.gz 安装包拷贝
	FileTransfer := &FileTransfer{
		LocalFilePath: filepath.Join(SFTP_SRC_DIR, "static", "uploadfile", this.ServiceInfo.ServiceName, this.ServiceInfo.PackageName),
		// 目标机器是 Linux 系统,需要转换为 Linux 路径分隔符
		RemoteDir: fileutil.ChangeToLinuxSeparator(filepath.Join(GetRemoteDeployHomePath(this.ServiceInfo.EnvInfo), "beego", "packages", this.ServiceInfo.ServiceName)),
	}
	FileTransfers = append(FileTransfers, FileTransfer)

	return FileTransfers
}

func (this *BeegoFileTransferCreator) PrepareFileTransfer() []*FileTransfer {
	switch this.OperateType {
	case "deploy":
		return this.BeegoDeployFileTransfer()
	}
	return nil
}

func GetRemoteDeployHomePath(envInfo *models.EnvInfo) string {
	var deploy_home string
	if envInfo.DpeloyHome != "" {
		deploy_home = envInfo.DpeloyHome
	} else {
		deploy_home = SFTP_TARGET_DIR
	}
	return deploy_home
}

func SyncDeployHome(envInfo *models.EnvInfo) error {
	sftpClient, err := common.SFTPConnect(envInfo.EnvAccount, envInfo.EnvPasswd, envInfo.EnvIp, 22)
	if err != nil {
		return err
	}
	defer sftpClient.Close()

	// 远程机器 deploy_home
	remoteDeployHome := GetRemoteDeployHomePath(envInfo)

	// 拷贝脚本目录
	err = common.SFTPClientCopyDirectoryInto(sftpClient, filepath.Join(SFTP_SRC_DIR, "shell"), remoteDeployHome)
	return err
}
