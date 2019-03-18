package iworknode

import (
	"fmt"
	"isoft/isoft_iaas_web/core/iworkconst"
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/core/iworkdata/param"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworkutil/sshutil"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
)

type SSHShellLogWriter struct {
	LogType    string
	TrackingId string
}

func (this *SSHShellLogWriter) Write(p []byte) (n int, err error) {
	message := string(p)
	messages := strings.Split(message, "\n")
	for _, messageInfo := range messages {
		if strings.TrimSpace(messageInfo) != "" {
			iwork.InsertRunLogDetail(this.TrackingId, fmt.Sprintf("%s -- %s", this.LogType, strings.TrimSpace(messageInfo)))
		}
	}
	return len(p), nil
}

type SSHShellNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *SSHShellNode) Execute(trackingId string) {
	// 数据中心
	dataStore := datastore.GetDataStore(trackingId)
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, dataStore)
	sshResource := param.GetStaticParamValue(iworkconst.STRING_PREFIX+"ssh_conn", this.WorkStep).(iwork.Resource)
	ssh_command := tmpDataMap[iworkconst.STRING_PREFIX+"ssh_command"].(string)
	stdout := &SSHShellLogWriter{
		LogType:    "INFO",
		TrackingId: trackingId,
	}
	stderr := &SSHShellLogWriter{
		LogType:    "ERROR",
		TrackingId: trackingId,
	}
	err := sshutil.RunSSHShellCommand(sshResource.ResourceUsername, sshResource.ResourcePassword,
		sshResource.ResourceDsn, ssh_command, stdout, stderr)
	if err != nil {
		panic(err)
	}
}

func (this *SSHShellNode) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "ssh_conn", "ssh连接信息,需要使用 $RESOURCE 全局参数"},
		2: {iworkconst.STRING_PREFIX + "ssh_command", "远程执行的命令"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *SSHShellNode) GetRuntimeParamInputSchema() *schema.ParamInputSchema {
	return &schema.ParamInputSchema{}
}

func (this *SSHShellNode) GetDefaultParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}

func (this *SSHShellNode) GetRuntimeParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}

func (this *SSHShellNode) ValidateCustom() (checkResult []string) {
	return
}