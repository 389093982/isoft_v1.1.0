import {checkContainsInString, checkEmpty, getCookie} from "../tools"
import {modulesCheck} from "./index"

// sso 登陆拦截
export const checkSSOLogin = function (to, from, next) {
  if (modulesCheck("sso")){
    // 登录判断
    var userName = getCookie("userName");
    var isLogin = getCookie("isLogin");
    var token = getCookie("token");
    // 非免登录白名单,并且不含登录标识的需要重新跳往登录页面
    if(!checkNotLogin() && (checkEmpty(userName) || checkEmpty(isLogin) || checkEmpty(token) || isLogin != "isLogin")){
      // 跳往登录页面
      window.location.href = "/sso/login/?redirectUrl=" + window.location.href;
    }else{
      next();
    }
  }else{
    next();
  }
};

function checkNotLogin(){
  if(checkContainsInString(window.location.href, "/sso/login") || checkContainsInString(window.location.href, "/sso/regist")){
    return true;
  }
  return false;
}
