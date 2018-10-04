import Vue from 'vue'
import Router from 'vue-router'
import IFile from '../components/IFile/IFile.vue'
import IBlog from '../components/IBlog/IBlog.vue'
import BlogList from '../components/IBlog/BlogList.vue'
import CatalogAdd from '../components/IBlog/CatalogAdd.vue'
import BlogAdd from '../components/IBlog/BlogAdd.vue'
import BlogDetail from '../components/IBlog/BlogDetail.vue'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/iblog',
      component: IBlog,
      // 二级路由的配置
      children: [
        {
          path: 'catalog_add',
          component: CatalogAdd
        },
        {
          path: 'blog_add',
          component: BlogAdd
        },
        {
          path: 'blog_list',
          component: BlogList
        },
        {
          path: 'blog_detail',
          component: BlogDetail
        },
      ]
    },
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
