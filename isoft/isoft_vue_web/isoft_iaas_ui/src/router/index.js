import Vue from 'vue'
import Router from 'vue-router'

import ILayout from "../components/ILayout/ILayout"
import QuartzList from "../components/IQuartz/QuartzList"
import ResourceList from "../components/IResource/ResourceList"

import {getISSORouters} from "./sso"
import {getIWorkRouters} from "./iwork"
import {getILearningRouters} from "./ilearning"
import {getRootRouters} from "./root"

Vue.use(Router);

export const IQuartzRouter = {
  path: '/quartz',
  component: ILayout,
  children: [
    {path: 'quartzList',component: QuartzList},
  ]
};



export const IResourceRouter = {
  path: '/resource',
  component: ILayout,
  children: [
    {path: 'resourceList',component: ResourceList},
  ]
};

function getAllRouters() {
  let allRouters = [];
  [].push.apply(allRouters, getIWorkRouters());
  [].push.apply(allRouters, getILearningRouters());
  [].push.apply(allRouters, getISSORouters());
  [].push.apply(allRouters, getRootRouters());
  return allRouters;
}


export default new Router({
  // History 模式,去除vue项目中的 #
  mode: 'history',
  routes: getAllRouters(),
})
