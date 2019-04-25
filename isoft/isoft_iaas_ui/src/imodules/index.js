import {checkContainsInString} from "../tools"
import {checkSSOLogin} from "./sso"

// 前台开启的模块 sso,ilearning,iwork
const open_modules="iwork";

export function modulesCheck(moduleName) {
  return checkContainsInString(open_modules, moduleName);
}

export const CheckSSOLogin = (to, from, next) => checkSSOLogin(to, from, next);

