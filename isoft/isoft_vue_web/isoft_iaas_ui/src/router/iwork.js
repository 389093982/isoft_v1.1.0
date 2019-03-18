import WorkList from "../components/IWork/WorkList"
import WorkStepList from "../components/IWork/WorkStepList"
import RunLogList from "../components/IWork/RunLogList"
import WorkHistoryList from "../components/IWork/WorkHistoryList"
import RunLogDetail from "../components/IWork/RunLogDetail"
import IWorkLayout from "../components/ILayout/IWorkLayout"
import QuartzList from "../components/IWork/IQuartz/QuartzList"
import ResourceList from "../components/IWork/IResource/ResourceList"
import {modulesCheck} from "../imodules";

const IWorkRouter = {
  path: '/iwork',
  component: IWorkLayout,
  children: [
    {path: 'quartzList',component: QuartzList},
    {path: 'resourceList',component: ResourceList},
    {path: 'workList',component: WorkList},
    {path: 'workstepList',component: WorkStepList},
    {path: 'runLogList',component: RunLogList},
    {path: 'workHistoryList',component: WorkHistoryList},
    {path: 'runLogDetail',component: RunLogDetail},
  ]
};

export const getIWorkRouters = function () {
  if (modulesCheck("iwork")){

    return [IWorkRouter];
  }
  return [];
};
