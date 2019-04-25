import WorkList from "../components/IWork/IWork/WorkList"
import WorkStepList from "../components/IWork/IWorkStep/WorkStepList"
import RunLogList from "../components/IWork/IRunLog/RunLogList"
import WorkHistoryList from "../components/IWork/IHistory/WorkHistoryList"
import RunLogDetail from "../components/IWork/IRunLog/RunLogDetail"
import IWorkLayout from "../components/ILayout/IWorkLayout"
import QuartzList from "../components/IWork/IQuartz/QuartzList"
import ResourceList from "../components/IWork/IResource/ResourceList"
import MigrateList from "../components/IWork/IMigrate/MigrateList"
import EditMigrate from "../components/IWork/IMigrate/EditMigrate"
import QuickSql from "../components/IWork/IAssistant/QuickSql"
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
    {path: 'quickSql',component: QuickSql},
  ]
};

export const getIWorkRouters = function () {
  if (modulesCheck("iwork")){

    return [IWorkRouter];
  }
  return [];
};
