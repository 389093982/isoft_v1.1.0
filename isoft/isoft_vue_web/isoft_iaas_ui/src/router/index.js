import Vue from 'vue'
import Router from 'vue-router'
import IFile from '../components/IFile/IFile.vue'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/ifile/ifile',
      component: IFile
    },
    {
      path: '/',
      redirect: '/ifile/ifile'
    }
  ]
})
