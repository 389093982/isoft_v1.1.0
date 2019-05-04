package iworknode

import (
	"fmt"
	"isoft/isoft_iaas_web/core/iworkconst"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworklog"
	"isoft/isoft_iaas_web/core/iworkmodels"
	"isoft/isoft_iaas_web/core/iworkutil/sshutil"
	"isoft/isoft_iaas_web/models/iwork"
	"strconv"
	"strings"
)

type SSHShellLogWriter struct {
	LogType    string
	logwriter  *iworklog.CacheLoggerWriter
	TrackingId string
}

func (this *SSHShellLogWriter) Write(p []byte) (n int, err error) {
	message := string(p)
	messages := strings.Split(message, "\n")
	for _, messageInfo := range messages {
		if strings.TrimSpace(messageInfo) != "" {
			this.logwriter.Write(this.TrackingId, fmt.Sprintf("%s -- %s", this.LogType, strings.TrimSpace(messageInfo)))
		}
	}
	return len(p), nil
}

type SSHShellNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *SSHShellNode) Execute(trackingId string) {
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep)
	sshResource := tmpDataMap[iworkconst.STRING_PREFIX+"ssh_conn"].(iwork.Resource)
	ssh_command := tmpDataMap[iworkconst.STRING_PREFIX+"ssh_command"].(string)

	var timeout int64
	if _timeout, ok := tmpDataMap[iworkconst.NUMBER_PREFIX+"command_timeout?"].(string); ok {
		if _timeout, err := strconv.ParseInt(_timeout, 10, 64); err == nil {
			timeout = _timeout
		}
	}

	stdout := &SSHShellLogWriter{
		LogType:    "INFO",
		logwriter:  this.LogWriter,
		TrackingId: trackingId,
	}
	stderr := &SSHShellLogWriter{
		LogType:    "ERROR",
		logwriter:  this.LogWriter,
		TrackingId: trackingId,
	}

	err := sshutil.RunSSHShellCommand(sshResource.ResourceUsername, sshResource.ResourcePassword,
		sshResource.ResourceDsn, ssh_command, stdout, stderr, timeout)
	if err != nil {
		panic(err)
	}
}

func (this *SSHShellNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "ssh_conn", "ssh连接信息,需要使用 $RESOURCE 全局参数"},
		2: {iworkconst.STRING_PREFIX + "ssh_command", "远程执行的命令,耗时的命令建议使用 nohup xxx > command.log & 格式"},
		3: {iworkconst.NUMBER_PREFIX + "command_timeout?", "执行命令超时时间"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}
