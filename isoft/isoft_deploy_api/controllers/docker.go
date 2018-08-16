package controllers

import (
	"fmt"
	"isoft/isoft_deploy_api/deploy_core/executors"
	"isoft/isoft_deploy_api/models"
	"strconv"
)

type DockerController struct {
	BaseController
}

func (this *DockerController) DockerInfoCheck() {
	env_id, err := this.GetInt64("env_id")
	if err != nil {
		this.RenderJsonErrorWithInvalidParamDetail(fmt.Sprintf("无效的环境%s", strconv.FormatInt(env_id, 10)))
	}
	envInfo, err := models.FilterEnvInfo(map[string]interface{}{"env_id": env_id})
	if err != nil {
		this.RenderJsonErrorWithInvalidParamDetail(fmt.Sprintf("无效的环境%s", strconv.FormatInt(env_id, 10)))
	}
	serviceInfo, err := models.FilterServiceInfo(map[string]interface{}{"service_name": "docker", "env_id": env_id})
	if err != nil {
		this.RenderJsonErrorWithInvalidParamDetail(fmt.Sprintf("当前环境 %s 未找到相关的docker服务", strconv.FormatInt(env_id, 10)))
	}

	// 开启协程执行任务
	tracking_id := executors.RunRemoteShellCommand(&serviceInfo, &envInfo, "docker_check")
	this.RenderJsonSuccessWithResultMap(map[string]interface{}{"tracking_id": tracking_id})
}
