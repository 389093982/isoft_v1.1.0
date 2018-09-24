package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"isoft/isoft/common/fileutil"
	"isoft/isoft/common/pageutil"
	"isoft/isoft_deploy_web/deploy_core/constant"
	"isoft/isoft_deploy_web/deploy_core/executors"
	"isoft/isoft_deploy_web/models"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

var (
	SFTP_SRC_DIR string
)

func init() {
	SFTP_SRC_DIR = fileutil.ChangeToLinuxSeparator(beego.AppConfig.String("sftp.src.dir"))
}

type ServiceController struct {
	beego.Controller
}

func (this *ServiceController) GetServiceTrackingLogDetail() {
	service_id, err := this.GetInt64("service_id")
	if err == nil {
		trackingLogs, err := models.QueryLastDeployTrackings(service_id)
		if err == nil {
			this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "trackingLogs": &trackingLogs}
		}
	}
	this.ServeJSON()
}

func (this *ServiceController) List() {
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页

	condArr := make(map[string]interface{})
	if service_type := this.GetString("service_type"); strings.TrimSpace(service_type) != "" {
		condArr["service_type"] = strings.TrimSpace(service_type)
	}
	if search_text := this.GetString("search_text"); strings.TrimSpace(search_text) != "" {
		condArr["search_text"] = strings.TrimSpace(search_text)
	}
	serviceInfos, count, err := models.QueryServiceInfo(condArr, current_page, offset)
	paginator := pagination.SetPaginator(this.Ctx, offset, count)

	//初始化
	data := make(map[string]interface{}, 1)

	if err == nil {
		data["serviceInfos"] = serviceInfos
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

func (this *ServiceController) Edit() {
	env_ids := strings.Split(this.GetString("env_ids"), ",")
	service_name := strings.TrimSpace(this.GetString("service_name"))
	service_type := strings.TrimSpace(this.GetString("service_type"))
	mysql_root_pwd := strings.TrimSpace(this.GetString("mysql_root_pwd"))
	package_name := strings.TrimSpace(this.GetString("package_name"))
	run_mode := strings.TrimSpace(this.GetString("run_mode"))
	service_port, err := this.GetInt64("service_port")

	if service_name == "" || service_type == "" || err != nil {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": "不合法的参数"}
		this.ServeJSON()
	}
	// package_name 为空时动态判断软件包名
	if package_name == "" {
		package_name = this.preparePackageName(service_name, service_type)
	}

	for _, env_id := range env_ids {
		eid, err := strconv.ParseInt(env_id, 10, 64)
		envInfo, err := models.FilterEnvInfo(map[string]interface{}{"env_id": eid})
		if err != nil {
			this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": "不正确的 env_id"}
			this.ServeJSON()
		}
		serviceInfo := models.ServiceInfo{
			EnvInfo:         &envInfo,
			ServiceName:     service_name,
			ServiceType:     service_type,
			ServicePort:     service_port,
			MysqlRootPwd:    mysql_root_pwd,
			PackageName:     package_name,
			RunMode:         run_mode,
			CreatedBy:       "AutoInsert",
			CreatedTime:     time.Now(),
			LastUpdatedBy:   "AutoInsert",
			LastUpdatedTime: time.Now(),
		}
		_, err = models.InsertOrUpdateServiceInfo(&serviceInfo)
		if err != nil {
			this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": "保存失败"}
			this.ServeJSON()
		}
	}
	this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	this.ServeJSON()
}

func (this *ServiceController) preparePackageName(serviceName, serviceType string) string {
	if strings.EqualFold(serviceType, constant.SERVICE_TYPE_BEEGO) {
		return strings.Join([]string{serviceName, ".tar.gz"}, "") // beego应用类型为 .tar.gz 格式包
	}
	if strings.EqualFold(serviceType, constant.SERVICE_TYPE_API) {
		return serviceName
	}
	return ""
}

func (this *ServiceController) FileDownload() {
	service_id, _ := this.GetInt64("service_id")
	serviceInfo, _ := models.QueryServiceInfoById(service_id)
	// 文件下载,第二个参数为下载时自定义的文件名
	this.Ctx.Output.Download(fmt.Sprintf(SFTP_SRC_DIR+"/static/uploadfile/%s/%s",
		serviceInfo.ServiceName, serviceInfo.PackageName), serviceInfo.PackageName)
}

func (this *ServiceController) FileUpload() {
	defer func() {
		if err := recover(); err != nil {
			this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": "保存失败！"}
			this.ServeJSON()
		}
	}()
	var (
		service_id  int64
		serviceInfo models.ServiceInfo
		h           *multipart.FileHeader
		err         error
	)
	if service_id, err = this.GetInt64("service_id"); err != nil {
		panic("invalid param service_id")
	}
	if serviceInfo, err = models.QueryServiceInfoById(service_id); err != nil {
		panic("invalid param service_id, serviceInfo not found")
	}
	if _, h, err = this.GetFile("file"); err != nil {
		panic("invalid file info")
	}
	// 上传文件名和 PackageName 不一致,则校验不通过
	if h.Filename != serviceInfo.PackageName {
		panic("invalid filename, not match with package_name")
	}
	// 根据服务类型创建分级文件夹,一种类型放在同一个目录
	os.MkdirAll("static/uploadfile/"+serviceInfo.ServiceType, os.ModePerm)
	// 安装包上传保存,多实例场景相同的软件包只需要传一次即可,同名的会覆盖
	if err = this.SaveToFile("file", path.Join(SFTP_SRC_DIR+"/static/uploadfile/"+serviceInfo.ServiceType, h.Filename)); err != nil {
		panic("file save err")
	}
	this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	this.ServeJSON()
}

func (this *ServiceController) RunDeployTask() {
	env_id, _ := this.GetInt64("env_id")
	service_id, _ := this.GetInt64("service_id")
	operate_type := this.GetString("operate_type")
	extra_params := this.GetString("extra_params")

	// 获取环境信息
	envInfo, err := models.FilterEnvInfo(map[string]interface{}{"env_id": env_id})
	if err != nil {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
		this.ServeJSON()
	}
	// 获取服务信息
	serviceInfo, err := models.FilterServiceInfo(map[string]interface{}{"service_id": service_id})
	if err != nil {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
		this.ServeJSON()
	}
	// 简单删除操作,只删除 DB 中的 service 信息
	if operate_type == "delete" {
		err := models.DeleteServiceInfo(service_id)
		if err != nil {
			this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
		} else {
			this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
		}
	} else {
		// 开启协程执行任务
		tracking_id := executors.RunCommandTask(&serviceInfo, &envInfo, operate_type, extra_params)
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "tracking_id": tracking_id}
	}
	this.ServeJSON()
}

func (this *ServiceController) QueryLastDeployStatus() {
	// 参数校验
	service_id, err := this.GetInt64("service_id")
	if err != nil {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": "请求参数不合法!"}
		this.ServeJSON()
	}

	this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "finish": false}
	tracking_id, _ := models.QueryLastDeployTrackingId(service_id)
	finish := models.IsFinishTask(tracking_id)
	var trackingStatus string

	trackingLogs, err := models.QueryLastRunTrackingLog(tracking_id)
	if err == nil {
		for _, trackingLog := range trackingLogs {
			if trackingLog.TrackingValue != "" {
				trackingStatus = trackingLog.TrackingValue
				break
			}
		}
	}
	if finish && trackingStatus == "" {
		trackingStatus = "N/A"
	}
	this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "finish": finish, "trackingStatus": trackingStatus}
	this.ServeJSON()
}
