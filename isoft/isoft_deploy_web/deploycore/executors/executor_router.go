package executors

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"isoft/isoft_deploy_web/deploycore"
	"isoft/isoft_deploy_web/deploycore/deploy"
	"isoft/isoft_deploy_web/deploycore/deploy/file_transfer"
	"isoft/isoft_deploy_web/models"
	"strings"
	"time"
)

type ExecutorRouter struct {
	// 环境信息
	EnvInfo             *models.EnvInfo
	ServiceInfo         *models.ServiceInfo
	FileTransfers       []*file_transfer.FileTransfer
	TrackingLogResolver *deploy.TrackingLogResolver
}

// 传输文件到目标机器
func (this *ExecutorRouter) transfer() {
	for _, transfer := range this.FileTransfers {
		// 传输历史中包含任务则进行舍弃
		if transfer == nil || transfer.LocalFilePath == "" {
			logs.Error("invalid transfer task!")
			break
		}
		// 开始传输
		transfer.Transfer(this.EnvInfo)
		this.TrackingLogResolver.WriteSuccessLog(fmt.Sprintf("copy file %s to %s", transfer.LocalFilePath, transfer.RemoteDir))
	}
}

func (this *ExecutorRouter) RunCommandTask(operate_type string, tracking_id, extra_params string) {
	// 开启记录任务
	this.TrackingLogResolver = &deploy.TrackingLogResolver{
		ServiceInfo: this.ServiceInfo,
	}
	this.TrackingLogResolver.StartRecordNewTask(tracking_id, this.ServiceInfo.ServiceName+"#"+operate_type)
	defer this.TrackingLogResolver.EndRecordTask()

	if len(this.FileTransfers) > 0 {
		this.TrackingLogResolver.WriteSuccessLog("start file transfer...")
		// 传输文件到目标机器
		this.transfer()
		this.TrackingLogResolver.WriteSuccessLog("end file transfer...")
	}

	// 全称的操作类型
	_operate_type := deploycore.GetRealCommandType(this.ServiceInfo.ServiceType, operate_type)

	this.TrackingLogResolver.WriteSuccessLog(fmt.Sprintf("start task at :%v", time.Now()))
	if IsCommonTask(_operate_type) {
		this.RunExecuteCommonTask(_operate_type, extra_params)
	} else {
		// 执行部署任务
		this.RunExecuteRemoteScriptTask(_operate_type, extra_params)
	}
	this.TrackingLogResolver.WriteSuccessLog(fmt.Sprintf("end task at :%v", time.Now()))
}

func IsCommonTask(operate_type string) bool {
	array := [...]string{"mysql_connection_test", "mysql_init"}
	for _, value := range array {
		if strings.EqualFold(value, operate_type) {
			return true
		}
	}
	return false
}
