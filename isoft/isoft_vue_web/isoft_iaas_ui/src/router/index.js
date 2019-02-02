import Vue from 'vue'
import Router from 'vue-router'
import IFile from '../components/IFile/IFile'
import IBlog from '../components/IBlog/IBlog'
import BlogList from '../components/IBlog/BlogList'
import CatalogAdd from '../components/IBlog/CatalogAdd'
import BlogAdd from '../components/IBlog/BlogAdd'
import BlogDetail from '../components/IBlog/BlogDetail'
import ILearningIndex from '../components/ILearning/Index'
import CourseSpace from '../components/ILearning/CourseSpace/CourseSpace'
import NewCourse from '../components/ILearning/CourseSpace/NewCourse'
import RecentlyViewed from '../components/ILearning/CourseSpace/RecentlyViewed'
import MyCourseList from '../components/ILearning/CourseSpace/MyCourseList'
import CourseDetail from '../components/ILearning/Course/CourseDetail'
import VideoPay from '../components/ILearning/Course/VideoPay'
import Configuration from '../components/CMS/Configuration'
import CourseSearch from "../components/ILearning/Course/CourseSearch"
import ShareAdd from "../components/Share/ShareAdd"
import ShareList from "../components/Share/ShareList"
import ShareDetail from "../components/Share/ShareDetail"
import HeartBeat from "../components/Monitor/HeartBeat"
import CommonLinkList from "../components/CMS/CommonLinkList"
import Login from "../components/SSO/Login"
import Regist from "../components/SSO/Regist"
import AppRegist from "../components/SSO/AppRegist"
import LoginRecord from "../components/SSO/LoginRecord"
import ISSOLayout from "../components/ILayout/ISSOLayout"
import ILayout from "../components/ILayout/ILayout"
import QuartzList from "../components/IQuartz/QuartzList"
import ResourceList from "../components/IResource/ResourceList"
import WorkList from "../components/IWork/WorkList"
import WorkStepList from "../components/IWork/WorkStepList"

Vue.use(Router);

export const IBlogRouter = {
    path: '/iblog',
    component: ILayout,
    // 二级路由的配置
    children: [
      {
        path: 'blog_index',
        component: IBlog
      },
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

export const MonitorRouter ={
  path: '/monitor',
  component: ILayout,
  children: [
    {path: 'filterPageHeartBeat',component: HeartBeat,},
  ]
};

export const IFileRouter = {
  path: '/ifile',
  component: ILayout,
  children: [
    {path: 'ifile',component: IFile,},
  ]
};

export const ShareListRouter = {
  path: '/share',
  component: ILayout,
  children: [
    {path: 'add',component: ShareAdd,},
    {path: 'list',component: ShareList,},
    {path: 'detail',component: ShareDetail,},
  ]
};


export const ILearningRouter = {
  path: '/ilearning',
  component: ILayout,
  // 二级路由的配置
  children: [
    {
      path: 'index',
      component: ILearningIndex,
    },
    {
      path: 'course_space',
      component: CourseSpace,
      redirect: '/ilearning/course_space/newCourse',
      children: [
        {path: 'newCourse',component: NewCourse,},
        {path: 'myCourseList',component: MyCourseList,},
        {path: 'RecentlyViewed',component: RecentlyViewed,},
      ]
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

export const CMSRouter = {
  path: '/cms',
  component: ILayout,
  children: [
    {path: 'commonLinkList',component: CommonLinkList},
  ]
};

export const IQuartzRouter = {
  path: '/quartz',
  component: ILayout,
  children: [
    {path: 'quartzList',component: QuartzList},
  ]
};

export const IWorkRouter = {
  path: '/iwork',
  component: ILayout,
  children: [
    {path: 'workList',component: WorkList},
    {path: 'workstepList',component: WorkStepList},
  ]
};

export const IResourceRouter = {
  path: '/resource',
  component: ILayout,
  children: [
    {path: 'resourceList',component: ResourceList},
  ]
};

export const ISSOReouter = {
  path: '/sso',
  component: ISSOLayout,
  children: [
    {path: 'login',component: Login},
    {path: 'regist',component: Regist},
    {path: 'appRegist',component: AppRegist},
    {path: 'loginRecord',component: LoginRecord},
  ]
};

export default new Router({
  // History 模式,去除vue项目中的 #
  mode: 'history',
  routes: [
    IWorkRouter,
    IResourceRouter,
    IQuartzRouter,
    ISSOReouter,
    IBlogRouter,
    IFileRouter,
    ILearningRouter,
    ShareListRouter,
    MonitorRouter,
    CMSRouter,
    {
      path: '/',
      redirect: '/ilearning/index'
    }
  ]
})
