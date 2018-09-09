package deploy

import (
	"isoft/isoft/common/fileutil"
	"isoft/isoft_deploy_web/deploy_core"
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
	ScriptPathMappingMap["mysql_install"] = "shell/mysql/mysql_install.sh"
	ScriptPathMappingMap["api_check"] = "shell/api/api_check.sh"
	ScriptPathMappingMap["api_deploy"] = "shell/api/api_deploy.sh"
	ScriptPathMappingMap["api_restart"] = "shell/api/api_restart.sh"
	ScriptPathMappingMap["api_shutdown"] = "shell/api/api_shutdown.sh"
	ScriptPathMappingMap["api_startup"] = "shell/api/api_startup.sh"
	ScriptPathMappingMap["api_undeploy"] = "shell/api/api_undeploy.sh"
	ScriptPathMappingMap["env_writer"] = "shell/common/env_writer.sh"
}

// 准备远程执行的 shell 命令
func PrepareCommand(serviceInfo *models.ServiceInfo, operate_type, extra_params string) (string, error) {
	// 准备 shell 命令相关参数
	args, err := PrepareArgs(serviceInfo, operate_type, extra_params)
	if err != nil {
		return "", err
	}
	// 当前脚本命令
	command := PrepareSimpleCommand(serviceInfo.EnvInfo, deploy_core.GetRealCommandType(serviceInfo.ServiceType, operate_type), args)
	// 获取 next 操作类型对应的脚本命令
	if getNextOperateType(operate_type) != "" {
		nextCommand, err := PrepareCommand(serviceInfo, getNextOperateType(operate_type), extra_params)
		if err == nil {
			command += " && " + nextCommand
		}
	}
	return command + " && echo " + constant.COMMAND_OVER, nil
}

// 准备简单脚本命令
func PrepareSimpleCommand(envInfo *models.EnvInfo, command_type string, args string) string {
	remoteDeployHome := file_transfer.GetRemoteDeployHomePath(envInfo)
	// 目标机器脚本路径, serviceType 和 operate_type 拼接成 command_type
	scriptPath := remoteDeployHome + "/" + ScriptPathMappingMap[command_type]
	// 当前脚本命令
	command := "cd " + fileutil.ChangeToLinuxSeparator(filepath.Dir(scriptPath)) + " && ./" + filepath.Base(scriptPath) + " " + args
	return command
}

// 准备 shell 命令相关参数
func PrepareArgs(serviceInfo *models.ServiceInfo, operate_type, extra_params string) (string, error) {
	resolver := &CommandArgs{}
	argslices, err := resolver.GetCommandArgs(serviceInfo, deploy_core.GetRealCommandType(serviceInfo.ServiceType, operate_type), extra_params)

	if err != nil {
		return "", err
	}
	if argslices == nil {
		return "", nil
	}
	return strings.Join(argslices, " "), nil
}

// 多步骤情况下才有 next 操作类型
func getNextOperateType(operate_type string) string {
	//if operate_type == "mysql_install"{
	//	return "mysql_adjust"
	//}
	return ""
}
