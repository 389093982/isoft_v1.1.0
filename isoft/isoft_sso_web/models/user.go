package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id              int       `pk json:"id"`
	UserName        string    `json:"user_name"`
	PassWd          string    `json:"pass_wd"`
	CreatedBy       string    `json:"created_by"`        // 创建人
	CreatedTime     time.Time `json:"created_time"`      // 创建时间
	LastUpdatedBy   string    `json:"last_updated_by"`   // 修改人
	LastUpdatedTime time.Time `json:"last_updated_time"` // 修改时间
}

type UserToken struct {
	Id              int       `pk json:"id"`
	UserName        string    `json:"user_name"`
	// 一次性 code,客户端无法根据 cookie 获取 token_string,则使用 redirectUrl 中的请求参数 code 获取 token_string
	Code			string    `json:"code"`
	TokenString     string    `json:"token_string"`
	CreatedBy       string    `json:"created_by"`        // 创建人
	CreatedTime     time.Time `json:"created_time"`      // 创建时间
	LastUpdatedBy   string    `json:"last_updated_by"`   // 修改人
	LastUpdatedTime time.Time `json:"last_updated_time"` // 修改时间
}

func SaveUser(user User) error {
	o := orm.NewOrm()
	count, _ := o.QueryTable("user").Filter("user_name", user.UserName).Count()
	if count > 0 {
		return errors.New("用户已注册!")
	} else {
		_, err := o.Insert(&user)
		return err
	}
	return nil
}

func QueryUser(username, passwd string) (user User, err error) {
	o := orm.NewOrm()
	err = o.QueryTable("user").Filter("user_name", username).Filter("pass_wd", passwd).One(&user)
	return
}

func SaveUserToken(userToken UserToken) error {
	o := orm.NewOrm()
	var uToken UserToken
	err := o.QueryTable("user_token").Filter("user_name", userToken.UserName).One(&uToken)
	if err == nil {
		uToken.TokenString = userToken.TokenString
		uToken.Code = userToken.Code
		o.Update(&uToken, "token_string", "code")
	} else {
		_, err := o.Insert(&userToken)
		return err
	}
	return nil
}

func DeleteUserToken(userToken UserToken) error {
	o := orm.NewOrm()
	_, err := o.Delete(&userToken)
	return err
}

func QueryUserTokenByName(username string) (userToken UserToken, err error) {
	o := orm.NewOrm()
	err = o.QueryTable("user_token").Filter("user_name", username).OrderBy("-created_time").One(&userToken)
	return
}

func QueryUserTokenByCode(code string) (userToken UserToken, err error) {
	o := orm.NewOrm()
	err = o.QueryTable("user_token").Filter("code", code).OrderBy("-created_time").One(&userToken)
	return
}
