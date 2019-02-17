import WorkList from "../components/IWork/WorkList"
import WorkStepList from "../components/IWork/WorkStepList"
import RunLogList from "../components/IWork/RunLogList"
import RunLogDetail from "../components/IWork/RunLogDetail"
import ILayout from "../components/ILayout/ILayout"
import {modulesCheck} from "../imodules";

const IWorkRouter = {
  path: '/iwork',
  component: ILayout,
  children: [
    {path: 'workList',component: WorkList},
    {path: 'workstepList',component: WorkStepList},
    {path: 'runLogList',component: RunLogList},
    {path: 'runLogDetail',component: RunLogDetail},
  ]
};

export const getIWorkRouter = function () {
  if (modulesCheck("iwork")){
    return IWorkRouter;
  }
  return {}
};
