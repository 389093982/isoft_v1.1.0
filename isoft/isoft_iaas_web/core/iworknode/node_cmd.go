package iworknode

import (
	"io/ioutil"
	"isoft/isoft_iaas_web/core/iworkconst"
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/models/iwork"
	"os"
	"os/exec"
	"strings"
)

type RunCmd struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *RunCmd) Execute(trackingId string) {
	// 数据中心
	dataStore := datastore.GetDataStore(trackingId)
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, dataStore)

	if cd := tmpDataMap[iworkconst.STRING_PREFIX+"cd?"].(string); cd != "" {
		if err := os.Chdir(cd); err != nil {
			panic(err)
		}
	}

	command_name := tmpDataMap[iworkconst.STRING_PREFIX+"command_name"].(string)
	command_args := tmpDataMap[iworkconst.STRING_PREFIX+"command_args"].(string)
	args := strings.Split(command_args, " ")
	result := runCommand(command_name, args...)
	dataStore.CacheData(this.WorkStep.WorkStepName, iworkconst.STRING_PREFIX+"command_result", result)
}

func (this *RunCmd) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "cd?", "切换目录"},
		2: {iworkconst.STRING_PREFIX + "command_name", "执行命令"},
		3: {iworkconst.STRING_PREFIX + "command_args", "执行命令参数列表"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *RunCmd) GetRuntimeParamInputSchema() *schema.ParamInputSchema {
	return &schema.ParamInputSchema{}
}

func (this *RunCmd) GetDefaultParamOutputSchema() *schema.ParamOutputSchema {
	return schema.BuildParamOutputSchemaWithSlice([]string{iworkconst.STRING_PREFIX + "command_result"})
}

func (this *RunCmd) GetRuntimeParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}

func (this *RunCmd) ValidateCustom() (checkResult []string) {
	return
}

// 执行系统命令,第一个参数是命令名称,第二个参数是参数列表
func runCommand(name string, arg ...string) string {
	cmd := exec.Command(name, arg...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	defer stdout.Close()
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	opBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		panic(err)
	}
	return string(opBytes)
}
