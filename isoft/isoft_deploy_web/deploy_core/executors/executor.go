package executors

import (
	"isoft/isoft/common/stringutil"
	"isoft/isoft_deploy_web/deploy_core/deploy/file_transfer"
	"isoft/isoft_deploy_web/models"
)

func RunCommandTask(serviceInfo *models.ServiceInfo, envInfo *models.EnvInfo, operate_type, extra_params string) (tracking_id string) {
	tracking_id = stringutil.RandomUUID()
	// 开启协程执行任务
	go func() {
		serviceInfo.EnvInfo = envInfo
		// 文件传输生成器
		FileTransferCreator := file_transfer.FileTransferCreator{
			ServiceInfo: serviceInfo,
			OperateType: operate_type,
		}

		// 任务执行器
		ExecutorRouter := ExecutorRouter{
			EnvInfo:       envInfo,
			FileTransfers: FileTransferCreator.PrepareFileTransfer(),
			ServiceInfo:   serviceInfo,
		}
		// 执行任务
		ExecutorRouter.RunCommandTask(operate_type, tracking_id, extra_params)
	}()

	return
}
