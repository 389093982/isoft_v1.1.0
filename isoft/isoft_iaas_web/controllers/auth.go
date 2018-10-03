package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"isoft/isoft/common/httputil"
)

var (
	isoft_sso_url string
)

func init() {
	isoft_sso_url = beego.AppConfig.String("isoft_sso_url")
}

type AuthController struct {
	beego.Controller
}

func (this *AuthController) RedirectToLogin()  {
	url := fmt.Sprintf("%s/user/login", isoft_sso_url)
	redirectUrl := this.GetString("redirectUrl")
	this.Redirect(url + "?redirectUrl=" + redirectUrl, 301)
}

func (this *AuthController) GetJWTTokenByCode() {
	code := this.GetString("code")
	url := fmt.Sprintf("%s/user/getJWTTokenByCode", isoft_sso_url)
	resultmap, err := httputil.DoPost(url, map[string]interface{}{"code":code})
	if err != nil{
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}else{
		this.Data["json"] = &resultmap
	}
	this.ServeJSON()
	return
}


