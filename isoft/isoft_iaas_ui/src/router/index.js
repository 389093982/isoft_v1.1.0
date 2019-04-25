import Vue from 'vue'
import Router from 'vue-router'

import {getISSORouters} from "./sso"
import {getIWorkRouters} from "./iwork"
import {getILearningRouters} from "./ilearning"
import {getRootRouters} from "./root"

Vue.use(Router);



function getAllRouters() {
  let allRouters = [];
  [].push.apply(allRouters, getIWorkRouters());
  [].push.apply(allRouters, getILearningRouters());
  [].push.apply(allRouters, getISSORouters());
  [].push.apply(allRouters, getRootRouters());
  return allRouters;
}


export default new Router({
  // # 主要用来区分前后台应用, history 模式需要使用 nginx 代理
  // History 模式,去除vue项目中的 #
  // mode: 'history',
  routes: getAllRouters(),
})
