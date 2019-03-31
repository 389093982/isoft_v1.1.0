import WorkList from "../components/IWork/WorkList"
import WorkStepList from "../components/IWork/WorkStepList"
import RunLogList from "../components/IWork/RunLogList"
import WorkHistoryList from "../components/IWork/WorkHistoryList"
import RunLogDetail from "../components/IWork/RunLogDetail"
import IWorkLayout from "../components/ILayout/IWorkLayout"
import QuartzList from "../components/IWork/IQuartz/QuartzList"
import ResourceList from "../components/IWork/IResource/ResourceList"
import MigrateList from "../components/IWork/IMigrate/MigrateList"
import EditMigrate from "../components/IWork/IMigrate/EditMigrate"
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
    {path: 'migrateList',component: MigrateList},
    {path: 'editMigrate',component: EditMigrate},
  ]
};

export const getIWorkRouters = function () {
  if (modulesCheck("iwork")){

    return [IWorkRouter];
  }
  return [];
};
