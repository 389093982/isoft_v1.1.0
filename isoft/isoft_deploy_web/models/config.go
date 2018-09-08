package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type ConfigFile struct {
	Id              int64     `json:"id"`
	EnvInfo         *EnvInfo  `json:"env_info" orm:"rel(fk)"`
	EnvProperty     string    `json:"env_property"`
	EnvValue        string    `json:"env_value"`
	CreatedBy       string    `json:"created_by"`
	CreatedTime     time.Time `json:"created_time"`
	LastUpdatedBy   string    `json:"last_updated_by"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}

// 插入或者更新服务
func InsertOrUpdateConfigFile(configFile *ConfigFile) (id int64, err error) {
	// 根据联合条件 env_id 和 service_name 判断是否存在
	oldConfigFile, err := FilterConfigFile(map[string]interface{}{"env_id": configFile.EnvInfo.Id, "env_property": configFile.EnvProperty})
	if err == nil {
		// 存在则插入
		configFile.Id = oldConfigFile.Id
		configFile.CreatedTime = oldConfigFile.CreatedTime
		configFile.CreatedBy = oldConfigFile.CreatedBy
	}
	o := orm.NewOrm()
	if configFile.Id > 0 {
		id, err = o.Update(configFile)
	} else {
		id, err = o.Insert(configFile)
	}
	return
}

func FilterConfigFile(condArr map[string]interface{}) (configFile ConfigFile, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("config_file")
	// 根据 configfile_id 去查询
	if configfile_id, ok := condArr["configfile_id"]; ok {
		qs = qs.Filter("id", configfile_id)
	}
	// 根据 env_property 和环境 id 去查询
	if env_property, ok := condArr["env_property"]; ok {
		qs = qs.Filter("env_property", env_property)
	}
	if env_id, ok := condArr["env_id"]; ok {
		qs = qs.Filter("env_info_id", env_id)
	}
	err = qs.One(&configFile)
	return
}

func QueryConfigFile(condArr map[string]interface{}, page int, offset int) (configFiles []ConfigFile, counts int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("config_file")

	var cond = orm.NewCondition()
	if search, ok := condArr["search_text"]; ok {
		subCond := orm.NewCondition()
		subCond = cond.And("env_property__contains", search).Or("env_property__contains", search)
		cond = cond.AndCond(subCond)
	}

	qs = qs.SetCond(cond)
	counts, _ = qs.Count()

	qs = qs.Limit(offset, (page-1)*offset)

	// 进行关联查询
	qs.RelatedSel().All(&configFiles)
	// 载入关系字段
	for _, configFile := range configFiles {
		_, err = orm.NewOrm().LoadRelated(&configFile, "EnvInfo")
	}
	return
}

func QueryConfigFileById(configfile_id int64) (configFile ConfigFile, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("config_file")
	err = qs.Filter("id", configfile_id).One(&configFile)
	return
}
