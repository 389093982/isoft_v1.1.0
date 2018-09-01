package deploy

import (
	"isoft/isoft/common/fileutil"
	"isoft/isoft_deploy_web/deploy_core/constant"
	"isoft/isoft_deploy_web/deploy_core/deploy/file_transfer"
	"isoft/isoft_deploy_web/models"
	"path/filepath"
	"strings"
)

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
}

// 根据操作类型获取对应脚本路径
func getScriptFilePathByCommandType(serviceType, operate_type string) string {
	return ScriptPathMappingMap[GetRealCommandType(serviceType, operate_type)]
}

// 准备远程执行的 shell 命令
func PrepareCommand(serviceInfo *models.ServiceInfo, operate_type, extra_params string) (string, error) {
	remoteDeployHome := file_transfer.GetRemoteDeployHomePath(serviceInfo.EnvInfo)
	// 目标机器脚本路径
	scriptPath := remoteDeployHome + "/" + getScriptFilePathByCommandType(serviceInfo.ServiceType, operate_type)
	// 准备 shell 命令相关参数
	args, err := PrepareArgs(serviceInfo, operate_type, extra_params)
	if err != nil {
		return "", err
	}
	return "cd " + fileutil.ChangeToLinuxSeparator(filepath.Dir(scriptPath)) +
		" && ./" + filepath.Base(scriptPath) + " " + args + " && echo " + constant.COMMAND_OVER, nil
}

// 准备 shell 命令相关参数
func PrepareArgs(serviceInfo *models.ServiceInfo, operate_type, extra_params string) (string, error) {
	resolver := &CommandArgs{}
	argslices, err := resolver.GetCommandArgs(serviceInfo, GetRealCommandType(serviceInfo.ServiceType, operate_type), extra_params)

	if err != nil {
		return "", err
	}
	if argslices == nil {
		return "", nil
	}
	return strings.Join(argslices, " "), nil
}
