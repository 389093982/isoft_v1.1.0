package deploy

import (
	"github.com/astaxie/beego/logs"
	"isoft/isoft/common"
	"isoft/isoft_deploy_api/models"
	"strings"
)

var (
	FileTransferHistoryMap = make(map[string]int)
)

type SSHTimerScriptExecutor struct {
	// 环境信息
	EnvInfo             *models.EnvInfo
	ServiceInfo         *models.ServiceInfo
	FileTransfers       []*FileTransfer
	TrackingLogResolver *TrackingLogResolver
}

// 传输文件到目标机器
func (this *SSHTimerScriptExecutor) transfer() {
	for _, transfer := range this.FileTransfers {
		// 传输历史中包含任务则进行舍弃
		if transfer == nil || transfer.localFilePath == "" {
			logs.Error("invalid transfer task!")
			break
		}
		if _, ok := FileTransferHistoryMap[transfer.localFilePath]; !ok {
			// 存入传输历史
			FileTransferHistoryMap[transfer.localFilePath] = 1
			// 开始传输
			transfer.Transfer(this.EnvInfo)
			// 从传输历史中删除
			delete(FileTransferHistoryMap, transfer.localFilePath)
		}
	}
}

type WriteSuccessLog struct {
	ServiceInfo         *models.ServiceInfo
	TrackingLogResolver *TrackingLogResolver
}

func (this *WriteSuccessLog) Write(p []byte) (n int, err error) {
	// 日志多条一行显示时需要去除 COMMAND_OVER 才是最终的日志
	message := strings.Replace(string(p), COMMAND_OVER, "", -1)
	messages := strings.Split(message, "\n")
	for _, messageInfo := range messages {
		if strings.TrimSpace(messageInfo) != "" {
			this.TrackingLogResolver.WriteSuccessLog(strings.TrimSpace(messageInfo))
		}
	}
	// 结束任务
	if strings.Contains(string(p), COMMAND_OVER) {
		this.TrackingLogResolver.EndRecordTask()
	}
	return len(p), nil
}

type WriteErrorLog struct {
	ServiceInfo         *models.ServiceInfo
	TrackingLogResolver *TrackingLogResolver
}

func (this *WriteErrorLog) Write(p []byte) (n int, err error) {
	this.TrackingLogResolver.WriteErrorLog(string(p))
	return len(p), nil
}

func (this *SSHTimerScriptExecutor) RunExecuteRemoteScriptTask(operate_type string) {
	sshClient, err := common.SSHConnect(this.EnvInfo.EnvAccount, this.EnvInfo.EnvPasswd, this.EnvInfo.EnvIp, 22)
	defer sshClient.Close()
	if err != nil {
		logs.Error("ssh connect error : %s", err.Error())
		return
	}

	sshClient.Stdout = &WriteSuccessLog{
		ServiceInfo:         this.ServiceInfo,
		TrackingLogResolver: this.TrackingLogResolver,
	}
	sshClient.Stderr = &WriteErrorLog{
		ServiceInfo:         this.ServiceInfo,
		TrackingLogResolver: this.TrackingLogResolver,
	}

	command, err := PrepareCommand(this.ServiceInfo, operate_type)
	if err != nil {
		logs.Error("prepare command error : %s", err.Error())
	} else {
		logs.Info("current command is %s", command)
		err := sshClient.Run(command)
		if err != nil {
			logs.Error("run command error : %s", err.Error())
		}
	}
}

func (this *SSHTimerScriptExecutor) RunRemoteScriptTask(operate_type string, tracking_id string) {
	// 开启记录任务
	this.TrackingLogResolver = &TrackingLogResolver{
		ServiceInfo: this.ServiceInfo,
	}
	this.TrackingLogResolver.StartRecordNewTask(tracking_id)
	// 传输文件到目标机器
	this.transfer()
	// 执行部署任务
	this.RunExecuteRemoteScriptTask(operate_type)
}
