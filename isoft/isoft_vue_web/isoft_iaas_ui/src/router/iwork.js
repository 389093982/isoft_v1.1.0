import WorkList from "../components/IWork/WorkList"
import WorkStepList from "../components/IWork/WorkStepList"
import RunLogList from "../components/IWork/RunLogList"
import RunLogDetail from "../components/IWork/RunLogDetail"
import IWorkLayout from "../components/ILayout/IWorkLayout"
import QuartzList from "../components/IQuartz/QuartzList"
import ResourceList from "../components/IResource/ResourceList"
import {modulesCheck} from "../imodules";

const IQuartzRouter = {
  path: '/quartz',
  component: IWorkLayout,
  children: [
    {path: 'quartzList',component: QuartzList},
  ]
};

const IResourceRouter = {
  path: '/resource',
  component: IWorkLayout,
  children: [
    {path: 'resourceList',component: ResourceList},
  ]
};

const IWorkRouter = {
  path: '/iwork',
  component: IWorkLayout,
  children: [
    {path: 'workList',component: WorkList},
    {path: 'workstepList',component: WorkStepList},
    {path: 'runLogList',component: RunLogList},
    {path: 'runLogDetail',component: RunLogDetail},
  ]
};

export const getIWorkRouters = function () {
  if (modulesCheck("iwork")){

    return [IWorkRouter,IQuartzRouter,IResourceRouter];
  }
  return [];
};
