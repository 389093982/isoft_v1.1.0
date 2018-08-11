import Vue from 'vue'
import Router from 'vue-router'
import Login from '../pages/Login/Login.vue'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/login',
      component: Login
    },
    {
      path: '/',
      redirect: '/login'
    }
  ]
})
