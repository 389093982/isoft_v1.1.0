package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
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

