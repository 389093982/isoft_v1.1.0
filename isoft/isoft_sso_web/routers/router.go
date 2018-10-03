package routers

import (
	"github.com/astaxie/beego"
	"isoft/isoft_sso_web/controllers"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	// 这段代码放在router.go文件的init()的开头
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	beego.Router("/user/regist", &controllers.UserController{}, "get,post:Regist")
	beego.Router("/user/login", &controllers.UserController{}, "get,post:Login")
	beego.Router("/user/getJWTTokenByCode", &controllers.UserController{}, "get,post:GetJWTTokenByCode")

	beego.Router("/userlogin/loginRecordList", &controllers.LoginRecordController{}, "get,post:LoginRecordList")

	beego.Router("/appregister/appRegisterList", &controllers.AppRegisterController{}, "get,post:AppRegisterList")
	beego.Router("/appregister/addAppRegister", &controllers.AppRegisterController{}, "get,post:AddAppRegister")

	// sso 简单认证模型,每次请求都会在登录系统进行认证,客户端不进行任何认证操作
	beego.Router("/user/checkOrInValidateTokenString", &controllers.UserController{}, "get,post:CheckOrInValidateTokenString")
}
