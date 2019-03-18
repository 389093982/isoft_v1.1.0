package iwork

import (
	"github.com/astaxie/beego/utils/pagination"
	"isoft/isoft/common/pageutil"
	"isoft/isoft_iaas_web/core/iworkutil/sftputil"
	"isoft/isoft_iaas_web/core/iworkutil/sqlutil"
	"isoft/isoft_iaas_web/core/iworkutil/sshutil"
	"isoft/isoft_iaas_web/models/iwork"
	"time"
)

func (this *WorkController) AddResource() {
	var resource iwork.Resource
	resource.ResourceName = this.Input().Get("resource_name")
	resource.ResourceType = this.Input().Get("resource_type")
	resource.ResourceUrl = this.Input().Get("resource_url")
	resource.ResourceDsn = this.Input().Get("resource_dsn")
	resource.ResourceUsername = this.Input().Get("resource_username")
	resource.ResourcePassword = this.Input().Get("resource_password")
	resource.CreatedBy = "SYSTEM"
	resource.CreatedTime = time.Now()
	resource.LastUpdatedBy = "SYSTEM"
	resource.LastUpdatedTime = time.Now()
	if _, err := iwork.InsertOrUpdateResource(&resource); err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func (this *WorkController) FilterPageResource() {
	condArr := make(map[string]string)
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页
	if search := this.GetString("search"); search != "" {
		condArr["search"] = search
	}
	resources, count, err := iwork.QueryResource(condArr, current_page, offset)
	paginator := pagination.SetPaginator(this.Ctx, offset, count)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "resources": resources,
			"paginator": pageutil.Paginator(paginator.Page(), paginator.PerPageNums, paginator.Nums())}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func (this *WorkController) DeleteResource() {
	id, _ := this.GetInt64("id")
	err := iwork.DeleteResource(id)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func (this *WorkController) ValidateResource() {
	var err error
	id, _ := this.GetInt64("id")
	resource, err := iwork.QueryResourceById(id)
	if err == nil {
		switch resource.ResourceType {
		case "db":
			db, err1 := sqlutil.GetConnForMysql("mysql", resource.ResourceDsn)
			if err1 == nil {
				defer db.Close()
			}
			err = err1
		case "sftp":
			sshClient, sftpClient, err1 := sftputil.SFTPConnect(resource.ResourceUsername, resource.ResourcePassword, resource.ResourceDsn, 22)
			if err1 == nil {
				defer sshClient.Close()
				defer sftpClient.Close()
			}
			err = err1
		case "ssh":
			err = sshutil.SSHConnectTest(resource.ResourceUsername, resource.ResourcePassword, resource.ResourceDsn, 22)
		}
	}
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}
