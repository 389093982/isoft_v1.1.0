package sso

import (
	"github.com/astaxie/beego/context"
)

func LoginFilter(ctx *context.Context, errorType string) {
	loginManager := &LoginManager{ctx: ctx}
	// 白名单直接跳过
	if loginManager.IsWhiteUrl() {
		return
	}

	// 从 cookie 中或者 header 中获取 token
	if loginManager.GetTokenString() == "" || !loginManager.CheckOrInValidateTokenString() {
		loginManager.ResponseWithErrorType(errorType)
	}
}

func LoginFilterWithRedirect(ctx *context.Context) {
	LoginFilter(ctx, "redirect")
}

func LoginFilterWithStatusCode(ctx *context.Context) {
	LoginFilter(ctx, "statusCode")
}
