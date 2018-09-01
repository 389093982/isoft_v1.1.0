package file_transfer

import (
	"isoft/isoft/common/fileutil"
	"isoft/isoft_deploy_web/models"
	"path/filepath"
)

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
		RemoteDir: fileutil.ChangeToLinuxSeparator(
			filepath.Join(GetRemoteDeployHomePath(this.ServiceInfo.EnvInfo), "upload/beego/packages", this.ServiceInfo.ServiceName)),
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
