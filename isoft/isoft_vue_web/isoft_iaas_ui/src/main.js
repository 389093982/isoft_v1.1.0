// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import $ from 'jquery'
import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootstrap/dist/js/bootstrap.min'
import store from './store'

import iView from 'iview'
import 'iview/dist/styles/iview.css'
import {GetJWTTokenByCode} from "./api"

Vue.use(iView);

Vue.config.productionTip = false

function getQueryVariable(_url, code) {
  if(_url != null && _url != undefined && _url != "" && _url.indexOf("?") > 0){
    _url = _url.split("?")[1];
    var vars = _url.split("&");
    for(var i=0;i<vars.length;i++) {
      var pair = vars[i].split("=");
      if(pair[0] == code){
        var _code = pair[1].replace("#","").replace("/","");
        return _code;
      }
    }
  }
  return "";
}

router.beforeEach(async (to, from, next) => {
  var code = getQueryVariable(window.location.href, "code")
   if(code != null && code != undefined && code != ""){
    const result = await GetJWTTokenByCode(code);
    if(result.status=="SUCCESS"){
      localStorage.setItem("userName", result.userName);
      localStorage.setItem("token", result.token);
      localStorage.setItem("isLogin", result.isLogin);
    }
  }
  if(localStorage.getItem("isLogin") != "isLogin"){
    localStorage.removeItem("userName");
    localStorage.removeItem("token");
    localStorage.removeItem("isLogin");
    window.location.href = "/api/auth/redirectToLogin/?redirectUrl=" + window.location.href;
  }else{
    next();
  }
});

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>',
  store, // 使用上vuex
});
