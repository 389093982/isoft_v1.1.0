import Vue from 'vue'
import Router from 'vue-router'
import IFile from '../components/IFile/IFile.vue'
import IBlog from '../components/IBlog/IBlog.vue'
import BlogList from '../components/IBlog/BlogList.vue'
import CatalogAdd from '../components/IBlog/CatalogAdd.vue'
import BlogAdd from '../components/IBlog/BlogAdd.vue'
import BlogDetail from '../components/IBlog/BlogDetail.vue'
import ILearning from '../components/ILearning/ILearning.vue'
import ILearningIndex from '../components/ILearning/Index.vue'
import CourseSpace from '../components/ILearning/CourseSpace/CourseSpace.vue'
import CourseDetail from '../components/ILearning/Course/CourseDetail.vue'
import VideoPay from '../components/ILearning/Course/VideoPay.vue'
import Configuration from '../components/CMS/Configuration.vue'
import CourseSearch from "../components/ILearning/Course/CourseSearch.vue"

Vue.use(Router);

export const IBlogRouter = {
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
  };

export const IFileRouter = {
    path: '/ifile/ifile',
    component: IFile
  };

export const ILearningRouter = {
    path: '/ilearning',
    component: ILearning,
    // 二级路由的配置
    children: [
      {
        path: 'index',
        component: ILearningIndex,
      },
      {
        path: 'course_space',
        component: CourseSpace,
      },
      {
        path: 'course_detail',
        component: CourseDetail,
      },
      {
        path: 'video_play',
        component: VideoPay,
      },
      {
        path: 'configuration',
        component: Configuration,
      },
      {
        // this.$router.push({ name: 'xxx'});
        // this.$router.push({ path: 'xxx'});
        name:'course_search',
        path: 'course_search',
        component: CourseSearch,
      },
    ]
  };


export default new Router({
  routes: [
    IBlogRouter,
    IFileRouter,
    ILearningRouter,
    {
      path: '/',
      redirect: '/ifile/ifile'
    }
  ]
})
