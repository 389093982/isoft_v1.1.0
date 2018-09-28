import Vue from 'vue'
import Router from 'vue-router'
import EnvList from '../components/Env/EnvList.vue'
import ServiceList from '../components/Service/ServiceList.vue'
import MonitorList from '../components/Monitor/MonitorList.vue'
import ConfigList from '../components/Config/ConfigList.vue'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/env/list',
      component: EnvList
    },
    {
      path: '/config/list',
      component: ConfigList
    },
    {
      path: '/service/list',
      component: ServiceList
    },
    {
      path: '/monitor/list',
      component: MonitorList
    },
    {
      path: '/',
      redirect: '/env/list'
    }
  ]
})
