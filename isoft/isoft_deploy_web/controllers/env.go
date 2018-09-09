package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"isoft/isoft/common/pageutil"
	"isoft/isoft/common/sshutil"
	"isoft/isoft_deploy_web/deploy_core/deploy/file_transfer"
	"isoft/isoft_deploy_web/models"
	"time"
)

type EnvController struct {
	beego.Controller
}

func (this *EnvController) All() {
	envInfos, _, err := models.QueryAllEnvInfo()
	data := make(map[string]interface{}, 1)
	if err == nil {
		data["envInfos"] = envInfos
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

func (this *EnvController) List() {
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
		data["paginator"] = pageutil.Paginator(paginator.Page(), paginator.PerPageNums, paginator.Nums())
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

func (this *EnvController) ConnectonTest() {
	env_id, err := this.GetInt64("env_id")
	if err != nil {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	} else {
		envInfo, err := models.FilterEnvInfo(map[string]interface{}{"env_id": env_id})
		if err != nil {
			this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
		} else {
			err = sshutil.SSHConnectTest(envInfo.EnvAccount, envInfo.EnvPasswd, envInfo.EnvIp, 22)
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

	err = file_transfer.SyncDeployHome(&envInfo)
	if err != nil {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
		this.ServeJSON()
	}

	this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	this.ServeJSON()
}

func (this *EnvController) PostEdit() {
	env_name := this.GetString("env_name")
	env_ip := this.GetString("env_ip")
	env_account := this.GetString("env_account")
	env_passwd := this.GetString("env_passwd")

	condArr := make(map[string]interface{})
	condArr["env_ip"] = env_ip
	env_info, err := models.FilterEnvInfo(condArr)
	if env_info.Id > 0 {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": "环境信息已存在！"}
	} else {
		// 插入操作
		envInfo := models.EnvInfo{
			EnvName:         env_name,
			EnvIp:           env_ip,
			EnvAccount:      env_account,
			EnvPasswd:       env_passwd,
			CreatedBy:       "AutoInsert",
			CreatedTime:     time.Now(),
			LastUpdatedBy:   "AutoInsert",
			LastUpdatedTime: time.Now(),
		}
		// 不存在则插入
		_, err = models.InsertOrUpdateEnvInfo(&envInfo)
		if err == nil {
			this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
		} else {
			this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": "保存失败！"}
		}
	}
	this.ServeJSON()
}
