package file_transfer

import (
	"github.com/astaxie/beego"
	"isoft/isoft/common/fileutil"
	"isoft/isoft/common/sftputil"
	"isoft/isoft_deploy_web/deploy_core"
	"isoft/isoft_deploy_web/models"
	"path/filepath"
	"strconv"
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
	sftpClient, err := sftputil.SFTPConnect(EnvInfo.EnvAccount, EnvInfo.EnvPasswd, EnvInfo.EnvIp, 22)
	if err != nil {
		return err
	}
	defer sftpClient.Close()

	if fileutil.IsDir(this.LocalFilePath) {
		sftputil.SFTPClientCopyDirectoryInto(sftpClient, this.LocalFilePath, this.RemoteDir)
	} else {
		sftputil.SFTPClientFileCopy(sftpClient, this.LocalFilePath, this.RemoteDir)
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
	_operate_type := deploy_core.GetRealCommandType(this.ServiceInfo.ServiceType, this.OperateType)
	switch _operate_type {
	case "beego_deploy":
		return BeegoDeployFileTransfer(this.ServiceInfo)
	case "api_deploy":
		return ApiDeployFileTransfer(this.ServiceInfo)
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
	sftpClient, err := sftputil.SFTPConnect(envInfo.EnvAccount, envInfo.EnvPasswd, envInfo.EnvIp, 22)
	if err != nil {
		return err
	}
	defer sftpClient.Close()

	// 远程机器 deploy_home
	remoteDeployHome := GetRemoteDeployHomePath(envInfo)

	// 拷贝脚本目录
	err = sftputil.SFTPClientCopyDirectoryInto(sftpClient, filepath.Join(SFTP_SRC_DIR, "shell"), remoteDeployHome)
	return err
}

// 同步本地 configFile 到目标机器
func SyncConfigFile(envInfo *models.EnvInfo, configFile *models.ConfigFile) error {
	sftpClient, err := sftputil.SFTPConnect(envInfo.EnvAccount, envInfo.EnvPasswd, envInfo.EnvIp, 22)
	if err != nil {
		return err
	}
	defer sftpClient.Close()
	savepath := SFTP_SRC_DIR + "/static/uploadfile/configfile/" + strconv.FormatInt(configFile.Id, 10)
	// 拷贝脚本目录
	err = sftputil.SFTPClientCopyDirectoryInto(sftpClient, savepath, configFile.EnvValue)
	return err
}
