package file_transfer

import (
	"isoft/isoft/common/fileutil"
	"isoft/isoft_deploy_web/models"
	"path/filepath"
)

// beego 项目部署文件传输
func BeegoDeployFileTransfer(serviceInfo *models.ServiceInfo) []*FileTransfer {
	return DeployFileTransfer(serviceInfo, "beego")
}

// api 项目部署文件传输
func ApiDeployFileTransfer(serviceInfo *models.ServiceInfo) []*FileTransfer {
	return DeployFileTransfer(serviceInfo, "api")
}

func DeployFileTransfer(serviceInfo *models.ServiceInfo, serviceType string) []*FileTransfer {
	FileTransfers := make([]*FileTransfer, 0)
	// .tar.gz 安装包拷贝
	FileTransfer := &FileTransfer{
		LocalFilePath: filepath.Join(SFTP_SRC_DIR, "static", "uploadfile", serviceInfo.ServiceName, serviceInfo.PackageName),
		// 目标机器是 Linux 系统,需要转换为 Linux 路径分隔符
		RemoteDir: fileutil.ChangeToLinuxSeparator(
			filepath.Join(GetRemoteDeployHomePath(serviceInfo.EnvInfo), "upload/"+serviceType+"/packages", serviceInfo.ServiceName)),
	}
	FileTransfers = append(FileTransfers, FileTransfer)
	return FileTransfers
}
