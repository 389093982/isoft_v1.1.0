// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import $ from 'jquery'
import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootstrap/dist/js/bootstrap.min'
import store from './store'
import {getCookie} from './tools'

import iView from 'iview'
import 'iview/dist/styles/iview.css'

Vue.use(iView);



function checkEmpty(checkStr){
  if(checkStr == null || checkStr == undefined || checkStr == ""){
    return true;
  }
  return false;
}

Vue.config.productionTip = false

// 登录判断
router.beforeEach(async (to, from, next) => {
  // LoadingBar 加载进度条
  iView.LoadingBar.start();

  var userName = getCookie("userName");
  var isLogin = getCookie("isLogin");
  var token = getCookie("token");
  if(checkEmpty(userName) || checkEmpty(isLogin) || checkEmpty(token) || isLogin != "isLogin"){
    // 跳往登录页面
    window.location.href = "/api/auth/redirectToLogin/?redirectUrl=" + window.location.href;
  }else{
    next();
  }
});

router.afterEach(route => {
  // LoadingBar 加载进度条
  iView.LoadingBar.finish();
});

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>',
  store, // 使用上vuex
});
