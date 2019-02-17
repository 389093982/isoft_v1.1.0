package miwork

import (
	"github.com/astaxie/beego/orm"
	"isoft/isoft_iaas_web/imodules"
	"isoft/isoft_iaas_web/models/iquartz"
	"isoft/isoft_iaas_web/models/iresource"
	"isoft/isoft_iaas_web/models/iwork"
)

func RegisterModel()  {
	if imodules.CheckModule("iwork"){
		orm.RegisterModel(new(iquartz.CronMeta))
		orm.RegisterModel(new(iresource.Resource))
		orm.RegisterModel(new(iwork.Work))
		orm.RegisterModel(new(iwork.WorkStep))
		orm.RegisterModel(new(iwork.RunLogRecord))
		orm.RegisterModel(new(iwork.RunLogDetail))
		orm.RegisterModel(new(iwork.Entity))
	}
}