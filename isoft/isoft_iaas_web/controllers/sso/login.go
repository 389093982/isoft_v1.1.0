package sso

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/pkg/errors"
	"isoft/isoft_iaas_web/models/sso"
	"net/url"
	"strings"
	"time"
)

type LoginController struct {
	beego.Controller
}

var origin_list string

func init() {
	origin_list = beego.AppConfig.String("origin_list")
}

func (this *LoginController) CheckOrInValidateTokenString() {
	this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	tokenString := this.GetString("tokenString")
	username := this.GetString("username")
	operateType := this.GetString("operateType")
	if operateType == "check" {
		username, err := ValidateAndParseJWT(tokenString)
		if err == nil {
			_, err = sso.QueryUserTokenByName(username)
			if err == nil {
				this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "username": username}
			}
		}
	} else {
		// 删除 tokenString,使客户端登录信息失效
		userToken, _ := sso.QueryUserTokenByName(username)
		sso.DeleteUserToken(userToken)
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	}
	this.ServeJSON()
}

func (this *LoginController) PostLogin()  {
	// referer显示来源页面的完整地址,而origin显示来源页面的origin: protocal+host,不包含路径等信息,也就不会包含含有用户信息的敏感内容
	// referer存在于所有请求,而origin只存在于post请求,随便在页面上点击一个链接将不会发送origin
	// 因此origin较referer更安全,多用于防范CSRF攻击
	referer := this.Ctx.Input.Referer()
	origin := this.Ctx.Request.Header.Get("origin")
	username := this.Input().Get("username")
	passwd := this.Input().Get("passwd")
	if IsAdminUser(username) { // 是管理面账号
		loginSuccess, loginStatus, tokenString, _ := AdminUserLogin(referer, origin, username, passwd, this.Ctx.Input.IP())
		this.Data["json"] = &map[string]interface{}{
			"loginSuccess": loginSuccess,
			"loginStatus":loginStatus,
			"tokenString":tokenString,
			"domain":getDomain(GetRedirectUrl(origin)),		 // 管理员登录设置 domain 为 sso 所在站点,不需要指定 redirectUrl
			"adminLogin":"adminLogin",
		}
	} else {
		loginSuccess, loginStatus, tokenString, _ := CommonUserLogin(referer, origin, username, passwd, this.Ctx.Input.IP())
		this.Data["json"] = &map[string]interface{}{
			"loginSuccess": loginSuccess,
			"loginStatus":loginStatus,
			"tokenString":tokenString,
			"redirectUrl": GetRedirectUrl(referer),
			"domain":getDomain(GetRedirectUrl(referer)),
		}
	}
	this.ServeJSON()
}

func getDomain(url string) string {
	if arr := strings.Split(url, "//"); len(arr) > 1{
		return strings.Split(arr[1], "/")[0]
	}
	return ""
}

func IsAdminUser(user_name string) bool {
	if user_name == "admin1" {
		return true
	}
	return false
}

func AdminUserLogin(referer, origin, username, passwd, ip string) (loginSuccess bool, loginStatus, tokenString string, err error){
	if CheckOrigin(origin) { // 非跨站点,不许校验 referer
		user, err := sso.QueryUser(username, passwd)
		if err == nil && &user != nil {
			return SuccessedLogin(username, ip, origin, referer, user)
		} else {
			return ErrorAccountLogin(username, ip, origin, referer)
		}
	} else {
		return ErrorAuthorizedLogin(username, origin, ip, referer)
	}
}

func ErrorAuthorizedLogin(username string, origin string, ip string, referer string) (loginSuccess bool, loginStatus,tokenString string, err error){
	var loginLog sso.LoginRecord
	loginLog.UserName = username
	loginLog.LoginIp = ip
	loginLog.Origin = origin
	loginLog.Referer = GetUnescapeString(referer)
	if !CheckOrigin(origin) {
		loginLog.LoginStatus = "origin_error"
	} else {
		loginLog.LoginStatus = "refer_error"
	}
	loginLog.LoginResult = "FAILED"
	loginLog.CreatedBy = "SYSTEM"
	loginLog.CreatedTime = time.Now()
	loginLog.LastUpdatedBy = "SYSTEM"
	loginLog.LastUpdatedTime = time.Now()
	sso.AddLoginRecord(loginLog)
	return false, loginLog.LoginStatus, "", errors.New(fmt.Sprintf("login error:%s",loginLog.LoginStatus))
}

func ErrorAccountLogin(username string, ip string, origin string, referer string) (loginSuccess bool, loginStatus,tokenString string, err error){
	var loginLog sso.LoginRecord
	loginLog.UserName = username
	loginLog.LoginIp = ip
	loginLog.Origin = origin
	loginLog.Referer = GetUnescapeString(referer)
	loginLog.LoginStatus = "account_error"
	loginLog.LoginResult = "FAILED"
	loginLog.CreatedBy = "SYSTEM"
	loginLog.CreatedTime = time.Now()
	loginLog.LastUpdatedBy = "SYSTEM"
	loginLog.LastUpdatedTime = time.Now()
	sso.AddLoginRecord(loginLog)
	return false, loginLog.LoginStatus, "", errors.New(fmt.Sprintf("login error:%s",loginLog.LoginStatus))
}

func GetRedirectUrl(referer string) string {
	referers := strings.Split(referer, "/sso/login?redirectUrl=")
	if len(referers) == 2{
		return GetUnescapeString(referers[1])
	}
	// 不含 redirectURL 场景
	return GetUnescapeString(referers[0])
}

// 进行编解码
func GetUnescapeString(s string) string {
	if _s, err := url.QueryUnescape(s); err == nil{
		return _s
	}
	return s
}

func CommonUserLogin(referer string, origin string, username string, passwd string, ip string) (loginSuccess bool, loginStatus, tokenString string, err error) {
	referers := strings.Split(referer, "/sso/login?redirectUrl=")
	if CheckOrigin(origin) && len(referers) == 2 && CheckOrigin(referers[0]) && IsValidRedirectUrl(GetRedirectUrl(referer)) {
		user, err := sso.QueryUser(username, passwd)
		if err == nil && &user != nil {
			return SuccessedLogin(username, ip, origin, referer, user)
		} else {
			return ErrorAccountLogin(username, ip, origin, referer)
		}
	} else {
		return ErrorAuthorizedLogin(username, origin, ip, referer)
	}
}

func SuccessedLogin(username string, ip string, origin string, referer string, user sso.User) (loginSuccess bool, loginStatus, tokenString string, err error){
	var loginLog sso.LoginRecord
	loginLog.UserName = username
	loginLog.LoginIp = ip
	loginLog.Origin = origin
	loginLog.Referer = GetUnescapeString(referer)
	loginLog.LoginStatus = "success"
	loginLog.LoginResult = "SUCCESS"
	loginLog.CreatedBy = "SYSTEM"
	loginLog.CreatedTime = time.Now()
	loginLog.LastUpdatedBy = "SYSTEM"
	loginLog.LastUpdatedTime = time.Now()
	sso.AddLoginRecord(loginLog)

	tokenString, err = CreateJWT(username)
	if err == nil {
		var userToken sso.UserToken
		userToken.UserName = username
		userToken.TokenString = tokenString
		userToken.CreatedBy = "SYSTEM"
		userToken.CreatedTime = time.Now()
		userToken.LastUpdatedBy = "SYSTEM"
		userToken.LastUpdatedTime = time.Now()
		sso.SaveUserToken(userToken)
	}
	return true, loginLog.LoginStatus, tokenString, nil
}

func IsValidRedirectUrl(redirectUrl string) bool {
	if redirectUrl != "" && IsHttpProtocol(redirectUrl) {
		// 截取协议名称
		arr := strings.Split(redirectUrl, "//")
		protocol := arr[0]
		// 截取域名
		a1 := arr[1]
		host := strings.Split(a1, "/")[0]
		return CheckRegister(protocol + "//" + host)
	} else {
		return false
	}
}

func IsHttpProtocol(url string) bool {
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		return true
	}
	return false
}

// 判断是否经过注册
func CheckRegister(registUrl string) bool {
	return sso.CheckRegister(registUrl)
}

// 验证 origin 是否合法
func CheckOrigin(origin string) bool {
	origin_slice := strings.Split(origin_list, ",")
	for _, _origin := range origin_slice {
		if origin == _origin {
			return true
		}
	}
	logs.Warn("origin error for %s", origin)
	return false
}

