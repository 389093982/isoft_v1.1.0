package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/utils/pagination"
	"isoft/isoft/common"
	"isoft/isoft_deploy_api/deploy_core/deploy"
	"isoft/isoft_deploy_api/models"
)

type EnvController struct {
	BaseController
}

func (c *EnvController) URLMapping() {
	c.Mapping("list", c.List)
	c.Mapping("connect_test", c.ConnectonTest)
}

// @Title Get EnvInfo list
// @Description Get EnvInfo list by some info
// @Success 200 {object} models.EnvInfo
// @Param   offset     query   int false       "offset"
// @Param   current_page    query   int false       "current_page"
// @Param   search_text   query   string  false       "search_text"
// @router /list [post]
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

// @Title Get EnvInfo list
// @Description Get EnvInfo list by some info
// @Param   env_id     query   int false       "env_id"
// @router /connect_test [post]
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

// @router /sync_deploy_home [post]
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
