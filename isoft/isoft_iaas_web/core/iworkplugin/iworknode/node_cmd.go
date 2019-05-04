package iworknode

import (
	"fmt"
	"isoft/isoft_iaas_web/core/iworkconst"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworklog"
	"isoft/isoft_iaas_web/core/iworkmodels"
	"isoft/isoft_iaas_web/core/iworkutil/cmdutil"
	"isoft/isoft_iaas_web/models/iwork"
	"os"
	"strings"
)

type RunCmdLogWriter struct {
	LogType    string
	TrackingId string
	logwriter  *iworklog.CacheLoggerWriter
}

func (this *RunCmdLogWriter) Write(p []byte) (n int, err error) {
	message := string(p)
	messages := strings.Split(message, "\n")
	for _, messageInfo := range messages {
		if strings.TrimSpace(messageInfo) != "" {
			this.logwriter.Write(this.TrackingId, fmt.Sprintf("%s -- %s", this.LogType, strings.TrimSpace(messageInfo)))
		}
	}
	return len(p), nil
}

type RunCmdNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *RunCmdNode) Execute(trackingId string) {
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep)

	if cd := tmpDataMap[iworkconst.STRING_PREFIX+"cd?"].(string); cd != "" {
		if err := os.Chdir(cd); err != nil {
			panic(err)
		}
	}

	stdout := &RunCmdLogWriter{
		LogType:    "INFO",
		logwriter:  this.LogWriter,
		TrackingId: trackingId,
	}
	stderr := &RunCmdLogWriter{
		LogType:    "ERROR",
		logwriter:  this.LogWriter,
		TrackingId: trackingId,
	}

	command_name := tmpDataMap[iworkconst.STRING_PREFIX+"command_name"].(string)
	command_args := tmpDataMap[iworkconst.STRING_PREFIX+"command_args"].(string)
	args := strings.Split(command_args, " ")
	// 记录当前正在执行的命令
	this.LogWriter.Write(trackingId, fmt.Sprintf("current cmd command is ==> %s %s", command_name, strings.Join(args, " ")))
	err := cmdutil.RunCommand(stdout, stderr, command_name, args...)
	if err != nil {
		panic(err)
	}
}

func (this *RunCmdNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "cd?", "切换目录"},
		2: {iworkconst.STRING_PREFIX + "command_name", "执行命令"},
		3: {iworkconst.STRING_PREFIX + "command_args", "执行命令参数列表"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *RunCmdNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return schema.BuildParamOutputSchemaWithSlice([]string{iworkconst.STRING_PREFIX + "command_result"})
}
