// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import $ from 'jquery'
import store from './store'

// 工具方法
import {getCookie} from './tools'
import {checkEmpty} from './tools'
import {checkContainsInString} from './tools'

// 引用全局静态数据
import global_ from './components/GlobalData'     //引用文件
Vue.prototype.GLOBAL = global_                    //挂载到Vue实例上面,通过 this.GLOBAL.xxx 访问全局变量

// 使用 iview
import iView from 'iview'
import 'iview/dist/styles/iview.css'
Vue.use(iView);

// 使用 vue-markdown
import mavonEditor from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'
Vue.use(mavonEditor)

Vue.config.productionTip = false

// 登录判断
router.beforeEach((to, from, next) => {
  /* 路由发生变化修改页面title */
  if (to.meta.title) {
    document.title = to.meta.title;
  }else{
    document.title = "iaas统一管理平台";
  }

  // LoadingBar 加载进度条
  iView.LoadingBar.start();

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
});

function checkNotLogin(){
  if(checkContainsInString(window.location.href, "/sso/login") || checkContainsInString(window.location.href, "/sso/regist")){
    return true
  }
  return false
}

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
