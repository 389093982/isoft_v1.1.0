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
import ILayout from "../components/ILayout/ILayout"
import {getIWorkRouter} from "./iwork";
import {getISSORouter} from "./sso";
import {IQuartzRouter, IResourceRouter} from "./index";
import {modulesCheck} from "../imodules";

const IBlogRouter = {
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

const MonitorRouter ={
  path: '/monitor',
  component: ILayout,
  children: [
    {path: 'filterPageHeartBeat',component: HeartBeat,},
  ]
};

const IFileRouter = {
  path: '/ifile',
  component: ILayout,
  children: [
    {path: 'ifile',component: IFile,},
  ]
};

const ShareListRouter = {
  path: '/share',
  component: ILayout,
  children: [
    {path: 'add',component: ShareAdd,},
    {path: 'list',component: ShareList,},
    {path: 'detail',component: ShareDetail,},
  ]
};

const ILearningRouter = {
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

const CMSRouter = {
  path: '/cms',
  component: ILayout,
  children: [
    {path: 'commonLinkList',component: CommonLinkList},
  ]
};

export const getILearningRouters = function () {
  if (modulesCheck("ilearning")) {
    return [IBlogRouter, IFileRouter, ILearningRouter, ShareListRouter, MonitorRouter, CMSRouter];
  }
  return [];
}

