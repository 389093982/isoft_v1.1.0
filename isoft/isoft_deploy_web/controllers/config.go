package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"isoft/isoft/common/fileutil"
	"isoft/isoft/common/pageutil"
	"isoft/isoft/common/sshutil"
	"isoft/isoft/common/ziputil"
	"isoft/isoft_deploy_web/deploy_core/deploy"
	"isoft/isoft_deploy_web/deploy_core/deploy/file_transfer"
	"isoft/isoft_deploy_web/models"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type ConfigController struct {
	beego.Controller
}

func (this *ConfigController) Edit() {
	env_ids := strings.Split(this.GetString("env_ids"), ",")
	env_property := strings.TrimSpace(this.GetString("env_property"))
	env_value := strings.TrimSpace(this.GetString("env_value"))
	if env_property == "" || env_value == "" || !fileutil.CheckFilePathValid(env_value) {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": "不合法的参数"}
		this.ServeJSON()
		return
	}
	for _, env_id := range env_ids {
		eid, err := strconv.ParseInt(env_id, 10, 64)
		envInfo, err := models.FilterEnvInfo(map[string]interface{}{"env_id": eid})
		if err != nil {
			this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": "不正确的 env_id"}
			this.ServeJSON()
			return
		}
		configFile := models.ConfigFile{
			EnvInfo:         &envInfo,
			EnvProperty:     env_property,
			EnvValue:        env_value,
			CreatedBy:       "AutoInsert",
			CreatedTime:     time.Now(),
			LastUpdatedBy:   "AutoInsert",
			LastUpdatedTime: time.Now(),
		}
		_, err = models.InsertOrUpdateConfigFile(&configFile)
		if err != nil {
			this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": "保存失败"}
			this.ServeJSON()
			return
		}
	}
	this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	this.ServeJSON()
}

func (this *ConfigController) List() {
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页

	condArr := make(map[string]interface{})
	if search_text := this.GetString("search_text"); strings.TrimSpace(search_text) != "" {
		condArr["search_text"] = strings.TrimSpace(search_text)
	}
	configFiles, count, err := models.QueryConfigFile(condArr, current_page, offset)
	paginator := pagination.SetPaginator(this.Ctx, offset, count)

	//初始化
	data := make(map[string]interface{}, 1)

	if err == nil {
		data["configFiles"] = configFiles
		data["paginator"] = pageutil.Paginator(paginator.Page(), paginator.PerPageNums, paginator.Nums())
	}
	//序列化
	json_obj, err := json.Marshal(data)
	if err == nil {
		this.Data["json"] = string(json_obj)
	} else {
		fmt.Print(err)
	}
	this.ServeJSON()
}

func (this *ConfigController) FileDownload() {
	configFile_id, _ := this.GetInt64("configFile_id")
	savepath := SFTP_SRC_DIR + "/static/uploadfile/configfile/" + strconv.FormatInt(configFile_id, 10)
	configFile, err := models.QueryConfigFileById(configFile_id)
	if err != nil {
		return
	}
	zipPath := SFTP_SRC_DIR + "/static/uploadfile/configfile/" + configFile.EnvProperty + ".zip"
	defer os.Remove(zipPath)
	err = ziputil.CompressZip(savepath, zipPath)
	if err != nil {
		fmt.Println("ziputil.CompressZip() returned %v\n", err)
	}
	// 文件下载,第二个参数为下载时自定义的文件名
	this.Ctx.Output.Download(zipPath, configFile.EnvProperty+".zip")
}

func (this *ConfigController) FileUpload() {
	configFile_id, _ := this.GetInt64("configFile_id")
	_, h, err := this.GetFile("file")
	if err != nil {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": "保存失败！"}
	} else {
		savepath := SFTP_SRC_DIR + "/static/uploadfile/configfile/" + strconv.FormatInt(configFile_id, 10)
		os.MkdirAll(savepath, os.ModePerm)
		// 保存文件
		err := this.SaveToFile("file", path.Join(savepath, h.Filename))
		if err != nil {
			this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": "保存失败！"}
		} else {
			this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
		}
	}
	this.ServeJSON()
}

func (this *ConfigController) SyncConfigFile() {
	defer func() {
		if err := recover(); err != nil {
			this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err}
		} else {
			this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
		}
		this.ServeJSON()
	}()
	env_id, err1 := this.GetInt64("env_id")
	configFile_id, err2 := this.GetInt64("configFile_id")
	if err1 != nil || err2 != nil {
		panic("Invalid param error!")
	}
	envInfo, err := models.FilterEnvInfo(map[string]interface{}{"env_id": env_id})
	if err != nil {
		panic(err.Error())
	}
	configFile, err := models.QueryConfigFileById(configFile_id)
	if err != nil {
		panic(err.Error())
	}
	// 同步配置文件到目标机器指定目录
	err = file_transfer.SyncConfigFile(&envInfo, &configFile)
	if err != nil {
		panic(err.Error())
	}
	command := deploy.PrepareSimpleCommand(&envInfo, "env_writer", configFile.EnvProperty+" "+configFile.EnvValue)
	// 调用脚本设置环境变量
	err = sshutil.RunSSHShellCommandOnly(envInfo.EnvAccount, envInfo.EnvPasswd, envInfo.EnvIp, command)
	if err != nil {
		panic(err.Error())
	}
}
