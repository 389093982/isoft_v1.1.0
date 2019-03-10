package misso

import (
	"github.com/astaxie/beego/orm"
	"isoft/isoft_iaas_web/imodules"
	"isoft/isoft_iaas_web/models/sso"
)

func RegisterModel() {
	if imodules.CheckModule("sso") {
		orm.RegisterModel(new(sso.User))
		orm.RegisterModel(new(sso.AppRegister))
		orm.RegisterModel(new(sso.LoginRecord))
		orm.RegisterModel(new(sso.UserToken))
	}
}
