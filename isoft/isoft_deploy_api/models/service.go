package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type ServiceInfo struct {
	Id              int64     `json:"id"`
	EnvInfo         *EnvInfo  `json:"env_info" orm:"rel(fk)"`
	ServiceName     string    `json:"service_name"`
	ServiceType     string    `json:"service_type"`
	ServicePort     int64     `json:"service_port"`
	PackageName     string    `json:"package_name"`
	MysqlRootPwd    string    `json:"mysql_root_pwd"`
	RunMode         string    `json:"run_mode"`
	CreatedBy       string    `json:"created_by"`
	CreatedTime     time.Time `json:"created_time"`
	LastUpdatedBy   string    `json:"last_updated_by"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}

type ServiceMonitor struct {
	Id              int64     `json:"id"`
	Url             string    `json:"url"`
	Method          string    `json:"method"`
	StatusCode      int64     `json:"status_code"`
	CreatedBy       string    `json:"created_by"`
	CreatedTime     time.Time `json:"created_time"`
	LastUpdatedBy   string    `json:"last_updated_by"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}

type ServiceMonitorDetail struct {
	Id              int64     `json:"id"`
	Url             string    `json:"url"`
	Method          string    `json:"method"`
	StatusCode      int64     `json:"status_code"`
	CreatedBy       string    `json:"created_by"`
	CreatedTime     time.Time `json:"created_time"`
	LastUpdatedBy   string    `json:"last_updated_by"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}

/*
	载入关系字段示例:
	func (this *Role) GetList() ([]*Role, error) {
		role := make([]*Role, 0)
		_, err := orm.NewOrm().QueryTable(_ROLE_TABLE).RelatedSel().All(&role)

		for _, v := range role {
			_, err = orm.NewOrm().LoadRelated(v, "Permissions")
			_, err = orm.NewOrm().LoadRelated(v, "Users")
		}

		if err != nil {
			return nil, err
		}
		return role, nil
	}
*/
func QueryServiceInfo(condArr map[string]interface{}, page int, offset int) (serviceInfos []ServiceInfo, counts int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("service_info")

	var cond = orm.NewCondition()
	if service_type, ok := condArr["service_type"]; ok {
		cond = cond.And("service_type", service_type)
	}
	if search, ok := condArr["search"]; ok {
		subCond := orm.NewCondition()
		subCond = cond.And("service_name__contains", search).Or("service_type__contains", search)
		cond = cond.AndCond(subCond)
	}

	qs = qs.SetCond(cond)
	counts, _ = qs.Count()

	qs = qs.Limit(offset, (page-1)*offset)

	// 进行关联查询
	qs.RelatedSel().All(&serviceInfos)
	// 载入关系字段
	for _, serviceInfo := range serviceInfos {
		_, err = orm.NewOrm().LoadRelated(&serviceInfo, "EnvInfo")
	}
	return
}

func QueryServiceInfoById(service_id int64) (serviceInfo ServiceInfo, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("service_info")
	err = qs.Filter("id", service_id).One(&serviceInfo)
	return
}

func CheckServiceInfoExists(condArr map[string]interface{}) (exists bool, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("service_info")

	if env_id, ok := condArr["env_id"]; ok {
		qs = qs.Filter("env_info_id", env_id)
	}
	if service_name, ok := condArr["service_name"]; ok {
		qs = qs.Filter("service_name", service_name)
	}
	counts, err := qs.Count()
	if err != nil {
		return false, err
	}
	return counts > 0, nil
}

// 插入或者更新服务
func InsertOrUpdateServiceInfo(serviceInfo *ServiceInfo) (id int64, err error) {
	// 根据联合条件 env_id 和 service_name 判断是否存在
	oldServiceInfo, err := FilterServiceInfo(map[string]interface{}{"env_id": serviceInfo.EnvInfo.Id, "service_name": serviceInfo.ServiceName})
	if err == nil {
		// 存在则插入
		serviceInfo.Id = oldServiceInfo.Id
		serviceInfo.CreatedTime = oldServiceInfo.CreatedTime
		serviceInfo.CreatedBy = oldServiceInfo.CreatedBy
	}
	o := orm.NewOrm()
	if serviceInfo.Id > 0 {
		id, err = o.Update(serviceInfo)
	} else {
		id, err = o.Insert(serviceInfo)
	}
	return
}

// 检查端口号是否被占用
func CheckServicePortExists(env_id int64, service_port int64) (bool, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("service_info")
	count, err := qs.Filter("env_info_id", env_id).Filter("service_port", service_port).Count()
	return count > 0, err
}

// 根据 id 删除服务信息
func DeleteServiceInfo(service_id int64) error {
	o := orm.NewOrm()
	qs := o.QueryTable("service_info")
	_, err := qs.Filter("id", service_id).Delete()
	return err
}

// 根据服务信息
func FilterServiceInfo(condArr map[string]interface{}) (serviceInfo ServiceInfo, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("service_info")
	// 根据 service_id 去查询
	if service_id, ok := condArr["service_id"]; ok {
		qs = qs.Filter("id", service_id)
	}
	// 根据服务名称和环境 id 去查询
	if service_name, ok := condArr["service_name"]; ok {
		qs = qs.Filter("service_name", service_name)
	}
	if env_id, ok := condArr["env_id"]; ok {
		qs = qs.Filter("env_info_id", env_id)
	}
	err = qs.One(&serviceInfo)
	return
}

func QueryAllServiceMonitor() (serviceMonitors []ServiceMonitor, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("service_monitor").All(&serviceMonitors)
	return
}

func QueryServiceMonitor(condArr map[string]interface{}, page int, offset int) (serviceMonitors []ServiceMonitor, counts int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("service_monitor")
	var cond = orm.NewCondition()

	if _, ok := condArr["search"]; ok {
		subCond := orm.NewCondition()
		subCond = cond.And("url__contains", condArr["search"])
		cond = cond.AndCond(subCond)
	}

	qs = qs.SetCond(cond)
	counts, _ = qs.Count()

	qs = qs.Limit(offset, (page-1)*offset)

	// 进行关联查询
	qs.RelatedSel().All(&serviceMonitors)
	return
}

func InsertOrUpdateServiceMonitor(serviceMonitor *ServiceMonitor) (id int64, err error) {
	o := orm.NewOrm()
	if serviceMonitor.Id > 0 {
		id, err = o.Update(serviceMonitor)
	} else {
		id, err = o.Insert(serviceMonitor)
	}
	return
}

func InsertOrUpdateServiceMonitorDetail(serviceMonitorDetail *ServiceMonitorDetail) (id int64, err error) {
	o := orm.NewOrm()
	if serviceMonitorDetail.Id > 0 {
		id, err = o.Update(serviceMonitorDetail)
	} else {
		id, err = o.Insert(serviceMonitorDetail)
	}
	return
}
