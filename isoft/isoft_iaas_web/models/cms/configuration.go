package cms

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Configuration struct {
	Id                 int64            `json:"id"`										// 配置项 id
	ParentId		   int64            `json:"parent_id"`								// 父配置项 id,顶级配置为 0
	ConfigurationName  string           `json:"configuration_name"`						// 配置项名称
	ConfigurationValue string           `json:"configuration_value"`					// 配置项值
	SubConfigurations  []*Configuration `json:"sub_configurations" orm:"-"`				// 自配置项列表
	CreatedBy          string           `json:"created_by"`
	CreatedTime        time.Time        `json:"created_time"`
	LastUpdatedBy      string           `json:"last_updated_by"`
	LastUpdatedTime    time.Time        `json:"last_updated_time"`
	Status 			   int 				`json:"status"`									// 状态 -1 表示失效
}

func FilterConfigurations(condArr map[string]string, page int, offset int) (configurations []Configuration, counts int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("configuration")
	var cond = orm.NewCondition()
	if _, ok := condArr["search"]; ok {
		subCond := orm.NewCondition()
		subCond = cond.And("configuration_name__contains", condArr["search"]).Or("configuration_value__contains", condArr["search"])
		cond = cond.AndCond(subCond)
	}
	qs = qs.SetCond(cond)
	counts, _ = qs.Count()
	qs = qs.Limit(offset, (page-1)*offset)
	qs.All(&configurations)
	return
}

func QueryAllConfigurations(configuration_name string, parent_id int64) (configurations []*Configuration, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("configuration").Filter("configuration_name", configuration_name).Filter("parent_id", parent_id).All(&configurations)
	if err == nil && len(configurations) > 0{
		for _,configuration := range configurations{
			sub, err := QueryAllConfigurations(configuration_name, configuration.Id)
			if err == nil && len(sub) > 0{
				configuration.SubConfigurations = sub
			}
		}
	}
	return
}

func AddConfiguration(configuration *Configuration) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(configuration)
	return
}