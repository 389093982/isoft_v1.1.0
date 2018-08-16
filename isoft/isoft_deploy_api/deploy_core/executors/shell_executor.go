package executors

import (
	"isoft/isoft/common"
	"isoft/isoft_deploy_api/deploy_core/deploy"
	"isoft/isoft_deploy_api/models"
)

func RunRemoteShellCommand(serviceInfo *models.ServiceInfo, envInfo *models.EnvInfo, operate_type string) (tracking_id string) {
	tracking_id = common.RandomUUID()
	// 开启协程执行任务
	go func() {
		serviceInfo.EnvInfo = envInfo
		// 文件传输生成器
		FileTransferCreator := deploy.FileTransferCreator{
			ServiceInfo: serviceInfo,
			OperateType: operate_type,
		}

		// 任务执行器
		SSHTimerScriptExecutor := deploy.SSHTimerScriptExecutor{
			EnvInfo:       envInfo,
			FileTransfers: FileTransferCreator.PrepareFileTransfer(),
			ServiceInfo:   serviceInfo,
		}
		// 执行任务
		SSHTimerScriptExecutor.RunRemoteScriptTask(operate_type, tracking_id)
	}()

	return
}
