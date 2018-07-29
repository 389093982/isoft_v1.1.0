package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type EnvInfo struct {
	Id              int64          `json:"id"`
	EnvName         string         `json:"env_name"`
	EnvIp           string         `json:"env_ip"`
	EnvAccount      string         `json:"env_account"`
	EnvPasswd       string         `json:"env_passwd"`
	DpeloyHome      string         `json:"deploy_home"`
	CreatedBy       string         `json:"created_by"`
	CreatedTime     time.Time      `json:"created_time"`
	LastUpdatedBy   string         `json:"last_updated_by"`
	LastUpdatedTime time.Time      `json:"last_updated_time"`
	ServiceInfos    []*ServiceInfo `json:"service_infos" orm:"reverse(many)"` // 设置一对多的反向关系
}

func QueryEnvInfo(condArr map[string]interface{}, page int, offset int) (envInfos []EnvInfo, counts int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("env_info")
	var cond = orm.NewCondition()

	if _, ok := condArr["search"]; ok {
		subCond := orm.NewCondition()
		subCond = cond.And("env_name__contains", condArr["search"]).Or("env_ip__contains", condArr["search"])
		cond = cond.AndCond(subCond)
	}

	qs = qs.SetCond(cond)
	counts, _ = qs.Count()

	qs = qs.Limit(offset, (page-1)*offset)
	qs.All(&envInfos)
	return
}

func InsertOrUpdateEnvInfo(envInfo *EnvInfo) (id int64, err error) {
	// 根据联合条件 env_name 和 env_ip 判断是否存在
	oldEnvInfo, err := FilterEnvInfo(map[string]interface{}{"env_name": envInfo.EnvName, "env_ip": envInfo.EnvIp})
	if err == nil {
		// 存在则插入
		envInfo.Id = oldEnvInfo.Id
		envInfo.CreatedTime = oldEnvInfo.CreatedTime
		envInfo.CreatedBy = oldEnvInfo.CreatedBy
	}
	o := orm.NewOrm()
	if envInfo.Id > 0 {
		id, err = o.Update(envInfo)
	} else {
		id, err = o.Insert(envInfo)
	}
	return
}

// 查询环境信息
func FilterEnvInfo(condArr map[string]interface{}) (envInfo EnvInfo, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("env_info")
	// 根据 env_id 去查询
	if env_id, ok := condArr["env_id"]; ok {
		qs = qs.Filter("id", env_id)
	}
	// 根据环境名称和环境 ip 去查询
	if env_name, ok := condArr["env_name"]; ok {
		qs = qs.Filter("env_name", env_name)
	}
	if env_name, ok := condArr["env_ip"]; ok {
		qs = qs.Filter("env_ip", env_name)
	}
	err = qs.One(&envInfo)
	return
}

func QueryAllEnvInfo() (envInfos []EnvInfo, counts int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("env_info")
	counts, err = qs.All(&envInfos)
	return
}
