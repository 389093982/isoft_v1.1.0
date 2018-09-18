package executors

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"isoft/isoft/common/sshutil"
	"isoft/isoft_deploy_web/deploy_core/constant"
	"isoft/isoft_deploy_web/deploy_core/deploy"
	"isoft/isoft_deploy_web/models"
	"strings"
)

type WriteSuccessLog struct {
	ServiceInfo         *models.ServiceInfo
	TrackingLogResolver *deploy.TrackingLogResolver
}

func (this *WriteSuccessLog) Write(p []byte) (n int, err error) {
	// 日志多条一行显示时需要去除 COMMAND_OVER 才是最终的日志
	message := strings.Replace(string(p), constant.COMMAND_OVER, "", -1)
	messages := strings.Split(message, "\n")
	for _, messageInfo := range messages {
		if strings.TrimSpace(messageInfo) != "" {
			this.TrackingLogResolver.WriteSuccessLog(strings.TrimSpace(messageInfo))
		}
	}
	// 结束任务
	if strings.Contains(string(p), constant.COMMAND_OVER) {
		this.TrackingLogResolver.EndRecordTask()
	}
	return len(p), nil
}

type WriteErrorLog struct {
	ServiceInfo         *models.ServiceInfo
	TrackingLogResolver *deploy.TrackingLogResolver
}

func (this *WriteErrorLog) Write(p []byte) (n int, err error) {
	this.TrackingLogResolver.WriteErrorLog(string(p))
	return len(p), nil
}

func (this *ExecutorRouter) RunExecuteRemoteScriptTask(operate_type, extra_params string) {
	stdout := &WriteSuccessLog{
		ServiceInfo:         this.ServiceInfo,
		TrackingLogResolver: this.TrackingLogResolver,
	}
	stderr := &WriteErrorLog{
		ServiceInfo:         this.ServiceInfo,
		TrackingLogResolver: this.TrackingLogResolver,
	}

	command, err := deploy.PrepareCommand(this.ServiceInfo, operate_type, extra_params)
	if err != nil {
		logs.Error("prepare command error : %s", err.Error())
		return
	}

	logs.Info("current command is %s", command)
	this.TrackingLogResolver.WriteSuccessLog(fmt.Sprintf("current command is %s", command))
	err = sshutil.RunSSHShellCommand(this.EnvInfo.EnvAccount, this.EnvInfo.EnvPasswd, this.EnvInfo.EnvIp, command, stdout, stderr)
	if err != nil {
		logs.Error("run command error : %s", err.Error())
		return
	}
}
