package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/utils/pagination"
	"isoft/isoft/common"
	"isoft/isoft_deploy_web/deploy_core/deploy"
	"isoft/isoft_deploy_web/models"
	"time"
)

type EnvController struct {
	BaseController
}

func (this *EnvController) List() {
	this.Layout = "layout/layout.html"
	this.TplName = "env/list.html"
}

func (this *EnvController) PostList() {
	condArr := make(map[string]interface{})
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页
	search_text := this.GetString("search_text")

	if search_text != "" {
		condArr["search_text"] = search_text
	}

	envInfos, count, err := models.QueryEnvInfo(condArr, current_page, offset)
	paginator := pagination.SetPaginator(this.Ctx, offset, count)

	//初始化
	data := make(map[string]interface{}, 1)

	if err == nil {
		data["envInfos"] = envInfos
		data["paginator"] = common.Paginator(paginator.Page(), paginator.PerPageNums, paginator.Nums())
	}
	//序列化
	json_obj, err := json.Marshal(data)
	if err == nil {
		this.Data["json"] = string(json_obj)
	} else {
		fmt.Print(err.Error())
	}
	this.ServeJSON()
}

func (this *EnvController) PostEdit() {
	env_name := this.GetString("env_name")
	env_ip := this.GetString("env_ip")
	env_account := this.GetString("env_account")
	env_passwd := this.GetString("env_passwd")
	deploy_home := this.GetString("deploy_home")

	envInfo := &models.EnvInfo{
		EnvName:         env_name,
		EnvIp:           env_ip,
		EnvAccount:      env_account,
		EnvPasswd:       env_passwd,
		DpeloyHome:      deploy_home,
		CreatedBy:       "AutoInsert",
		CreatedTime:     time.Now(),
		LastUpdatedBy:   "AutoInsert",
		LastUpdatedTime: time.Now(),
	}
	_, err := models.InsertOrUpdateEnvInfo(envInfo)
	if err != nil {
		this.RenderJsonErrorWithInvalidParamDetail("保存失败！")
	}

	serviceInfo := models.ServiceInfo{
		EnvInfo:         envInfo,
		ServiceName:     "docker",
		ServiceType:     "docker",
		CreatedBy:       "AutoInsert",
		CreatedTime:     time.Now(),
		LastUpdatedBy:   "AutoInsert",
		LastUpdatedTime: time.Now(),
	}
	_, err = models.InsertOrUpdateServiceInfo(&serviceInfo)
	if err != nil {
		this.RenderJsonErrorWithInvalidParamDetail("保存失败！")
	}

	this.RenderJsonSuccessWithResultMap(map[string]interface{}{})
}

func (this *EnvController) ConnectonTest() {
	env_id, err := this.GetInt64("env_id")
	if err != nil {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	} else {
		envInfo, err := models.FilterEnvInfo(map[string]interface{}{"env_id": env_id})
		if err != nil {
			this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
		} else {
			err = common.SSHConnectTest(envInfo.EnvAccount, envInfo.EnvPasswd, envInfo.EnvIp, 22)
			if err != nil {
				this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
			} else {
				this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
			}
		}
	}
	this.ServeJSON()
}

func (this *EnvController) SyncDeployHome() {
	env_id, err := this.GetInt64("env_id")
	if err != nil {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
		this.ServeJSON()
	}

	envInfo, err := models.FilterEnvInfo(map[string]interface{}{"env_id": env_id})
	if err != nil {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
		this.ServeJSON()
	}

	err = deploy.SyncDeployHome(&envInfo)
	if err != nil {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
		this.ServeJSON()
	}

	this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	this.ServeJSON()
}
