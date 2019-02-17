package misso

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"isoft/isoft/ssofilter"
	"isoft/isoft_iaas_web/imodules"
)

func ssoFilterFunc(ctx *context.Context) {
	filter := new(ssofilter.LoginFilter)
	filter.LoginWhiteList = &[]string{"/api/sso/user/login", "/api/sso/user/regist", "/api/sso/user/checkOrInValidateTokenString"}
	filter.LoginUrl = ctx.Input.URL()
	filter.Ctx = ctx
	filter.SsoAddress = beego.AppConfig.String("isoft.sso.web.addr")
	filter.ErrorFunc = func() {
		filter.Ctx.ResponseWriter.WriteHeader(401)
	}
	filter.Filter()
}

func RegisterISSOFilter()  {
	if imodules.CheckModule("sso") {
		beego.InsertFilter("/api/*", beego.BeforeExec, ssoFilterFunc)
	}
}

