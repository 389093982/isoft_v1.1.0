package deploy

import (
	"errors"
	"isoft/isoft_deploy_web/deploy_core/deploy/file_transfer"
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

func (this *CommandArgs) BeegoStartupCommandArgs() ([]string, error) {
	return this.BeegoStatusCommandArgs()
}

func (this *CommandArgs) BeegoRestartCommandArgs() ([]string, error) {
	return this.BeegoStatusCommandArgs()
}

func (this *CommandArgs) BeegoShutdownCommandArgs() ([]string, error) {
	return this.BeegoStatusCommandArgs()
}

func (this *CommandArgs) BeegoStatusCommandArgs() ([]string, error) {
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

func (this *CommandArgs) GetCommandArgs(serviceInfo *models.ServiceInfo, operateType, extra_params string) ([]string, error) {
	this.serviceInfo = serviceInfo
	switch operateType {
	case "beego_deploy":
		return this.BeegoDeployCommandArgs()
	case "beego_restart":
		return this.BeegoRestartCommandArgs()
	case "beego_shutdown":
		return this.BeegoShutdownCommandArgs()
	case "beego_startup":
		return this.BeegoStartupCommandArgs()
	case "beego_status":
		return this.BeegoStatusCommandArgs()
	case "nginx_install":
		return this.NginxInstallCommandArgs()
	case "nginx_check":
		return this.NginxCheckCommandArgs()
	case "nginx_restart":
		return this.NginxRestartCommandArgs()
	default:
		return this.BeegoStatusCommandArgs()
	}
}