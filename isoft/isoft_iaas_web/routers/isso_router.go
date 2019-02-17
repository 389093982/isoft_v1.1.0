package routers

import (
	"github.com/astaxie/beego"
	"isoft/isoft_iaas_web/controllers/sso"
	"strings"
)

func loadISSORouter() {
	// sso 模块
	if strings.Contains(beego.AppConfig.String("open.modules"), "sso") {
		loadISSORouterDetail()
	}
}

func loadISSORouterDetail()  {
	beego.Router("/api/sso/user/login", &sso.LoginController{}, "post:PostLogin")
	beego.Router("/api/sso/user/regist", &sso.LoginController{}, "post:PostRegist")
	beego.Router("/api/sso/app/appRegisterList", &sso.AppRegisterController{}, "post:AppRegisterList")
	beego.Router("/api/sso/app/addAppRegister", &sso.AppRegisterController{}, "get,post:AddAppRegister")
	beego.Router("/api/sso/user/loginRecordList", &sso.LoginRecordController{}, "post:LoginRecordList")
	// sso 简单认证模型,每次请求都会在登录系统进行认证,客户端不进行任何认证操作
	beego.Router("/api/sso/user/checkOrInValidateTokenString", &sso.LoginController{}, "get,post:CheckOrInValidateTokenString")
}
