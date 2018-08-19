package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/utils/pagination"
	"isoft/isoft/common"
	"isoft/isoft_deploy_api/deploy_core/constant"
	"isoft/isoft_deploy_api/deploy_core/executors"
	"isoft/isoft_deploy_api/models"
	"path"
	"strconv"
	"strings"
	"time"
)

type ServiceController struct {
	BaseController
}

func (c *ServiceController) URLMapping() {
	c.Mapping("Edit", c.Edit)
	c.Mapping("List", c.List)
}

// @router /getServiceTrackingLogDetail [post]
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

func (this *ServiceController) EditorMonitor() {
	this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": "保存失败！"}
	url := this.GetString("url")
	monitor := models.ServiceMonitor{
		Url:             url,
		Method:          "get",
		StatusCode:      0,
		CreatedBy:       "AutoInsert",
		CreatedTime:     time.Now(),
		LastUpdatedBy:   "AutoInsert",
		LastUpdatedTime: time.Now(),
	}
	_, err := models.InsertOrUpdateServiceMonitor(&monitor)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	}
	this.ServeJSON()
}

func (this *ServiceController) PostMonitor() {
	condArr := make(map[string]interface{})
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页
	search_text := this.GetString("search_text")

	if search_text != "" {
		condArr["search_text"] = search_text
	}

	serviceMonitors, count, err := models.QueryServiceMonitor(condArr, current_page, offset)
	paginator := pagination.SetPaginator(this.Ctx, offset, count)

	//初始化
	data := make(map[string]interface{}, 1)

	if err == nil {
		data["serviceMonitors"] = serviceMonitors
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

func (this *ServiceController) Monitor() {
	this.Layout = "layout/layout.html"
	this.TplName = "service/monitor.html"
}

// @router /list [post]
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

// @router /edit [post]
func (this *ServiceController) Edit() {
	env_ids := strings.Split(this.GetString("env_ids"), ",")
	service_name := this.GetString("service_name")
	service_type := this.GetString("service_type")
	package_name := this.GetString("package_name")
	run_mode := this.GetString("run_mode")
	service_port, err := this.GetInt64("service_port")

	if strings.TrimSpace(service_name) == "" || strings.TrimSpace(service_type) == "" || err != nil {
		this.RenderJsonErrorWithInvalidParamDetail("不合法的参数")
	}

	if strings.TrimSpace(package_name) == "" {
		package_name = this.preparePackageName(service_name, service_type)
	}

	for _, env_id := range env_ids {
		eid, err := strconv.ParseInt(env_id, 10, 64)
		envInfo, err := models.FilterEnvInfo(map[string]interface{}{"env_id": eid})
		if err != nil {
			this.RenderJsonErrorWithInvalidParamDetail("不正确的 env_id")
		}
		serviceInfo := models.ServiceInfo{
			EnvInfo:         &envInfo,
			ServiceName:     service_name,
			ServiceType:     service_type,
			ServicePort:     service_port,
			PackageName:     package_name,
			RunMode:         run_mode,
			CreatedBy:       "AutoInsert",
			CreatedTime:     time.Now(),
			LastUpdatedBy:   "AutoInsert",
			LastUpdatedTime: time.Now(),
		}
		_, err = models.InsertOrUpdateServiceInfo(&serviceInfo)
		if err != nil {
			this.RenderJsonErrorWithInvalidParamDetail("保存失败！")
		}
	}
	this.RenderJsonSuccessWithResultMap(map[string]interface{}{})
}

func (this *ServiceController) preparePackageName(serviceName, serviceType string) string {
	if strings.EqualFold(serviceType, constant.SERVICE_TYPE_BEEGO) {
		return strings.Join([]string{serviceName, ".tar.gz"}, "") // .tar.gz 格式包
	}
	if strings.EqualFold(serviceType, constant.SERVICE_TYPE_API) {
		return serviceName
	}
	return ""
}

// @router /fileUpload [post]
func (this *ServiceController) FileUpload() {
	_, h, err := this.GetFile("file")
	if err != nil {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": "保存失败！"}
	} else {
		err := this.SaveToFile("file", path.Join("static/uploadfile", h.Filename))
		if err != nil {
			this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": "保存失败！"}
		} else {
			this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
		}
	}
	this.ServeJSON()
}

// @router /runDeployTask [post]
func (this *ServiceController) RunDeployTask() {
	env_id, _ := this.GetInt64("env_id")
	service_id, _ := this.GetInt64("service_id")
	operate_type := this.GetString("operate_type")

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

	// 开启协程执行任务
	tracking_id := executors.RunRemoteShellCommand(&serviceInfo, &envInfo, operate_type)

	this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "tracking_id": tracking_id}
	this.ServeJSON()
}

// @router /queryLastDeployStatus [post]
func (this *ServiceController) QueryLastDeployStatus() {
	this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "finish": false}
	service_id, err := this.GetInt64("service_id")
	if err != nil {
		this.RenderJsonErrorWithInvalidParamDetail("请求参数不合法!")
	}
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
