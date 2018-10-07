package sso

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/session"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var (
	isoft_sso_url string
	// 建立一个全局session mananger对象
	globalSessions *session.Manager
	// 登录白名单
	loginWhiteList map[string]string
)

func initWhiteList() {
	loginWhiteList := make(map[string]string, 1)
	loginWhiteList["/common/login"] = "/common/login"
	loginWhiteList["/common/login/"] = "/common/login/"
	loginWhiteList["/common/regist"] = "/common/regist"
	loginWhiteList["/common/regist/"] = "/common/regist/"
}

func init() {
	initWhiteList()
	isoft_sso_url = beego.AppConfig.String("isoft_sso_url")
	// 初始化全局session mananger对象
	sessionConfig := &session.ManagerConfig{
		CookieName:      "gosessionid",
		EnableSetCookie: true,
		Gclifetime:      3600,
		Maxlifetime:     3600,
		Secure:          false,
		CookieLifeTime:  3600,
		ProviderConfig:  "./tmp",
	}
	globalSessions, _ = session.NewManager("memory", sessionConfig)
	go globalSessions.GC()
}

type LoginManager struct {
	ctx *context.Context
}

// 从 cookie 中或者 header 中获取 token
func (this *LoginManager) GetTokenString() string {
	var tokenString string
	if strings.TrimSpace(this.ctx.GetCookie("token")) != "" {
		tokenString = this.ctx.GetCookie("token")
	} else if strings.TrimSpace(this.ctx.Request.Header.Get("token")) != "" {
		tokenString = this.ctx.Request.Header.Get("token")
	}
	return tokenString
}

func (this *LoginManager) IsWhiteUrl() bool {
	if _, ok := loginWhiteList[this.ctx.Input.URL()]; ok {
		return true
	}
	// 鉴权接口也放行
	if strings.HasPrefix(this.ctx.Input.URL(), "/api/auth") {
		return true
	}
	return false
}

func (this *LoginManager) ResponseWithErrorType(errorType string) {
	if errorType == "redirect" {
		// 前去登录
		this.ResponseWithRedirect()
	} else {
		this.ResponseWithStatusCode()
	}
}

func (this *LoginManager) ResponseWithStatusCode() {
	this.ctx.ResponseWriter.WriteHeader(401)
}

func (this *LoginManager) ResponseWithRedirect() {
	// 前去登录
	this.RedirectToLogin("")
}

func (this *LoginManager) ResetUserName(username string) {
	if this.ctx.Input.CruSession == nil {
		// 从未访问过是没有 session 的,需要重新创建
		this.ctx.Input.CruSession, _ = globalSessions.SessionStart(this.ctx.ResponseWriter, this.ctx.Request)
		this.ctx.Input.CruSession.Set("userName", username)
		this.ctx.Input.CruSession.Set("UserName", username)
	} else {
		// 登录信息认证通过
		this.ctx.Input.CruSession.Set("userName", username)
		this.ctx.Input.CruSession.Set("UserName", username)
	}
}

func (this *LoginManager) CheckOrInValidateTokenString() bool {
	resp, err := http.Get(isoft_sso_url + "/user/checkOrInValidateTokenString?tokenString=" + this.GetTokenString() + "&operateType=check")
	defer resp.Body.Close()
	if err == nil {
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			jsonStr := string(body)
			var jsonMap map[string]string
			json.Unmarshal([]byte(jsonStr), &jsonMap)
			if jsonMap["status"] == "SUCCESS" {
				this.ResetUserName(jsonMap["username"])
				return true
			}
		}
	}
	return false
}

func (this *LoginManager) RedirectToLogin(defaultRedirectUrl string) {
	scheme := this.ctx.Input.Scheme()
	if defaultRedirectUrl != "" {
		defaultRedirectUrl = isoft_sso_url + "/user/login?redirectUrl=" + defaultRedirectUrl
	} else {
		if scheme == "https" {
			defaultRedirectUrl = isoft_sso_url + "/user/login?redirectUrl=" + this.ctx.Input.Site() + this.ctx.Input.URI()
		} else {
			defaultRedirectUrl = isoft_sso_url + "/user/login?redirectUrl=" + this.ctx.Input.Site() + ":" + strconv.Itoa(this.ctx.Input.Port()) + this.ctx.Input.URI()
		}
	}
	this.ctx.Redirect(301, defaultRedirectUrl)
	return
}

func (this *LoginManager) RedirectToLogout(defaultRedirectUrl string) {
	username := this.ctx.Input.CruSession.Get("UserName").(string)
	// 使 tokenString 失效
	resp, _ := http.Get(isoft_sso_url + "/user/checkOrInValidateTokenString?username=" + username + "&operateType=invalid")
	defer resp.Body.Close()

	// session 失效
	this.ctx.Input.CruSession.SessionRelease(this.ctx.ResponseWriter)

	// 重新跳往登录页面
	this.RedirectToLogin(defaultRedirectUrl)
	return
}
