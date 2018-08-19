package deploy

import (
	"isoft/isoft/common/fileutil"
	"isoft/isoft_deploy_api/models"
	"path/filepath"
	"strings"
)

const COMMAND_OVER = "COMMAND_OVER"

var (
	ScriptPathMappingMap = make(map[string]string, 0)
)

func init() {
	ScriptPathMappingMap["beego_deploy"] = "shell/beego/beego_deploy.sh"
	ScriptPathMappingMap["beego_restart"] = "shell/beego/beego_restart.sh"
	ScriptPathMappingMap["beego_shutdown"] = "shell/beego/beego_shutdown.sh"
	ScriptPathMappingMap["beego_startup"] = "shell/beego/beego_startup.sh"
	ScriptPathMappingMap["beego_check"] = "shell/beego/beego_check.sh"
	ScriptPathMappingMap["docker_check"] = "shell/docker/docker_check.sh"
	ScriptPathMappingMap["nginx_check"] = "shell/nginx/nginx_check.sh"
	ScriptPathMappingMap["nginx_install"] = "shell/nginx/nginx_install.sh"
	ScriptPathMappingMap["nginx_restart"] = "shell/nginx/nginx_restart.sh"
	ScriptPathMappingMap["mysql_check"] = "shell/mysql/mysql_check.sh"
	ScriptPathMappingMap["mysql_install"] = "shell/mysql/mysql_install.sh"
	ScriptPathMappingMap["mysql_restart"] = "shell/mysql/mysql_restart.sh"
}

// 获取真实的命令脚本执行类型
func getRealCommandType(serviceType, operate_type string) string {
	if _, ok := ScriptPathMappingMap[serviceType+"_"+operate_type]; ok {
		return serviceType + "_" + operate_type
	}
	return operate_type
}

// 根据操作类型获取对应脚本路径
func getScriptFilePathByCommandType(serviceType, operate_type string) string {
	return ScriptPathMappingMap[getRealCommandType(serviceType, operate_type)]
}

// 准备远程执行的 shell 命令
func PrepareCommand(serviceInfo *models.ServiceInfo, operate_type string) (string, error) {
	remoteDeployHome := GetRemoteDeployHomePath(serviceInfo.EnvInfo)
	// 目标机器脚本路径
	scriptPath := remoteDeployHome + "/" + getScriptFilePathByCommandType(serviceInfo.ServiceType, operate_type)
	// 准备 shell 命令相关参数
	args, err := PrepareArgs(serviceInfo, operate_type)
	if err != nil {
		return "", err
	}
	return "cd " + fileutil.ChangeToLinuxSeparator(filepath.Dir(scriptPath)) +
		" && ./" + filepath.Base(scriptPath) + " " + args + " && echo " + COMMAND_OVER, nil
}

// 准备 shell 命令相关参数
func PrepareArgs(serviceInfo *models.ServiceInfo, operate_type string) (string, error) {
	resolver := &CommandArgs{}
	argslices, err := resolver.GetCommandArgs(serviceInfo, getRealCommandType(serviceInfo.ServiceType, operate_type))

	if err != nil {
		return "", err
	}
	if argslices == nil {
		return "", nil
	}
	return strings.Join(*argslices, " "), nil
}
