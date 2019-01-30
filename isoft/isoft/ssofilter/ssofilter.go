package ssofilter

import (
	"encoding/json"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/session"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	// 建立一个全局session mananger对象
	globalSessions *session.Manager
)

func init() {
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

type LoginFilter struct {
	LoginWhiteList *[]string
	LoginUrl string
	Ctx *context.Context
	SsoAddress string
	ErrorFunc func()
}

func (this *LoginFilter) Filter()  {
	// 白名单直接跳过
	if this.checkWhiteList(){
		return
	}
	// 从 cookie 中或者 header 中获取 token
	if this.getTokenString() == "" || !this.checkOrInValidateTokenString() {
		// 校验不通过,则调用回调函数进行登出处理
		this.ErrorFunc()
	}
}

func (this *LoginFilter) checkWhiteList() bool {
	for _,url := range *this.LoginWhiteList{
		return url == this.LoginUrl
	}
	return false
}

// 从 cookie 中或者 header 中获取 token
func (this *LoginFilter) getTokenString() string {
	var tokenString string
	if strings.TrimSpace(this.Ctx.GetCookie("token")) != "" {
		tokenString = this.Ctx.GetCookie("token")
	} else if strings.TrimSpace(this.Ctx.Request.Header.Get("token")) != "" {
		tokenString = this.Ctx.Request.Header.Get("token")
	}
	return tokenString
}

// 验证 token,不通过时注销 token
func (this *LoginFilter) checkOrInValidateTokenString() bool {
	resp, err := http.Get(this.SsoAddress + "/api/sso/user/checkOrInValidateTokenString?tokenString=" + this.getTokenString() + "&operateType=check")
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		jsonStr := string(body)
		var jsonMap map[string]string
		json.Unmarshal([]byte(jsonStr), &jsonMap)
		if jsonMap["status"] == "SUCCESS" {
			this.resetUserName(jsonMap["username"])
			return true
		}
	}
	return false
}

func (this *LoginFilter) resetUserName(username string) {
	if this.Ctx.Input.CruSession == nil {
		// 从未访问过是没有 session 的,需要重新创建
		this.Ctx.Input.CruSession, _ = globalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
		this.Ctx.Input.CruSession.Set("userName", username)
		this.Ctx.Input.CruSession.Set("UserName", username)
	} else {
		// 登录信息认证通过
		this.Ctx.Input.CruSession.Set("userName", username)
		this.Ctx.Input.CruSession.Set("UserName", username)
	}
}
