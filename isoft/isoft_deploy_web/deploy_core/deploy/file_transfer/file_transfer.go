package file_transfer

import (
	"github.com/astaxie/beego"
	"isoft/isoft/common"
	"isoft/isoft/common/fileutil"
	"isoft/isoft_deploy_web/models"
	"path/filepath"
)

var (
	SFTP_SRC_DIR, SFTP_TARGET_DEPLOY_HOME string
)

func init() {
	SFTP_SRC_DIR = beego.AppConfig.String("sftp.src.dir")
	SFTP_TARGET_DEPLOY_HOME = beego.AppConfig.String("sftp.target.deploy_home.default")
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

// 获取远程目标机器 deploy_home 路径
func GetRemoteDeployHomePath(envInfo *models.EnvInfo) string {
	var deploy_home string
	// envInfo 中没有 deploy_home 配置则使用配置文件中的 SFTP_TARGET_DIR
	if envInfo.DpeloyHome != "" {
		deploy_home = envInfo.DpeloyHome
	} else {
		deploy_home = SFTP_TARGET_DEPLOY_HOME
	}
	return deploy_home
}

// 同步本地 deploy_home 到目录机器
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
