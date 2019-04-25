import Login from "../components/SSO/Login"
import Regist from "../components/SSO/Regist"
import AppRegist from "../components/SSO/AppRegist"
import LoginRecord from "../components/SSO/LoginRecord"
import ISSOLayout from "../components/ILayout/ISSOLayout"
import {modulesCheck} from "../imodules";

const ISSOReouter = {
  path: '/sso',
  component: ISSOLayout,
  children: [
    {path: 'login',component: Login},
    {path: 'regist',component: Regist},
    {path: 'appRegist',component: AppRegist},
    {path: 'loginRecord',component: LoginRecord},
  ]
};

export const getISSORouters = function () {
  if (modulesCheck("sso")){
    return [ISSOReouter];
  }
  return [];
};
