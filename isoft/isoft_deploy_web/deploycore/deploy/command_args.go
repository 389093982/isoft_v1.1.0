package deploy

import (
	"errors"
	"isoft/isoft_deploy_web/deploycore/deploy/file_transfer"
	"isoft/isoft_deploy_web/models"
	"strconv"
	"strings"
)

type ICommandArgsResolver interface {
	GetCommandArgs(serviceInfo *models.ServiceInfo, operateType string) (*[]string, error)
}

type CommandArgs struct {
	serviceInfo *models.ServiceInfo
}

func (this *CommandArgs) ApiDeployCommandArgs() ([]string, error) {
	var slice []string
	if this.serviceInfo.ServiceName == "" {
		return slice, errors.New("empty param : ServiceName")
	}
	if this.serviceInfo.PackageName == "" {
		return slice, errors.New("empty param : PackageName")
	}
	if this.serviceInfo.RunMode == "" {
		return slice, errors.New("empty param : RunMode")
	}
	slice = append(slice, this.serviceInfo.ServiceName)
	slice = append(slice, strings.Replace(this.serviceInfo.PackageName, ".tar.gz", "", -1))
	slice = append(slice, this.serviceInfo.RunMode)
	return slice, nil
}

func (this *CommandArgs) NginxRestartCommandArgs() ([]string, error) {
	return this.NginxInstallCommandArgs()
}

func (this *CommandArgs) NginxCheckCommandArgs() ([]string, error) {
	return this.NginxInstallCommandArgs()
}

func (this *CommandArgs) NginxInstallCommandArgs() ([]string, error) {
	// 目标机器 deploy_home 路径
	remoteDeployHomePath := file_transfer.GetRemoteDeployHomePath(this.serviceInfo.EnvInfo)
	var slice []string
	slice = append(slice, remoteDeployHomePath)
	slice = append(slice, this.serviceInfo.ServiceName)
	slice = append(slice, "_") // 端口号
	return slice, nil
}

func (this *CommandArgs) BeegoDeployCommandArgs() ([]string, error) {
	var slice []string
	if this.serviceInfo.ServiceName == "" {
		return slice, errors.New("empty param : ServiceName")
	}
	if this.serviceInfo.PackageName == "" {
		return slice, errors.New("empty param : PackageName")
	}
	if this.serviceInfo.RunMode == "" {
		return slice, errors.New("empty param : RunMode")
	}
	if this.serviceInfo.ServicePort < 1 {
		return slice, errors.New("empty param : ServicePort")
	}
	slice = append(slice, this.serviceInfo.ServiceName)
	slice = append(slice, strings.Replace(this.serviceInfo.PackageName, ".tar.gz", "", -1))
	slice = append(slice, this.serviceInfo.RunMode)
	slice = append(slice, strconv.FormatInt(this.serviceInfo.ServicePort, 10))
	return slice, nil
}

func (this *CommandArgs) BeegoCheckCommandArgs() ([]string, error) {
	var slice []string
	if this.serviceInfo.ServiceName == "" {
		return slice, errors.New("empty param : ServiceName")
	}
	if this.serviceInfo.PackageName == "" {
		return slice, errors.New("empty param : PackageName")
	}
	slice = append(slice, this.serviceInfo.ServiceName)
	slice = append(slice, strings.Replace(this.serviceInfo.PackageName, ".tar.gz", "", -1))
	return slice, nil
}

func (this *CommandArgs) MysqlInstallCommandArgs() ([]string, error) {
	// 目标机器 deploy_home 路径
	remoteDeployHomePath := file_transfer.GetRemoteDeployHomePath(this.serviceInfo.EnvInfo)
	var slice []string
	slice = append(slice, remoteDeployHomePath)
	slice = append(slice, this.serviceInfo.ServiceName)
	slice = append(slice, "_")      // 端口号
	slice = append(slice, "123456") // rootPwd
	return slice, nil
}

func (this *CommandArgs) GetCommandArgs(serviceInfo *models.ServiceInfo, operateType, extra_params string) ([]string, error) {
	this.serviceInfo = serviceInfo
	switch operateType {
	case "beego_deploy":
		return this.BeegoDeployCommandArgs()
	case "beego_restart", "beego_shutdown", "beego_startup", "beego_check":
		return this.BeegoCheckCommandArgs()
	case "nginx_install":
		return this.NginxInstallCommandArgs()
	case "nginx_check":
		return this.NginxCheckCommandArgs()
	case "nginx_restart":
		return this.NginxRestartCommandArgs()
	case "mysql_install":
		return this.MysqlInstallCommandArgs()
	case "api_check", "api_deploy", "api_restart", "api_shutdown", "api_startup", "api_undeploy":
		return this.ApiDeployCommandArgs()
	default:
		return this.BeegoCheckCommandArgs()
	}
}
