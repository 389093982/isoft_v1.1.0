import {modulesCheck} from "../imodules";

function getIWorkerRootRouter() {
  return {
    path: '/',
      redirect: '/iwork/workList'
  }
}

function getILearningRootRouter() {
  return {
    path: '/',
    redirect: '/ilearning/index'
  }
}

export const getRootRouters = function () {
  if(modulesCheck("iwork")){
    return [getIWorkerRootRouter()];
  }
  if(modulesCheck("ilearning")){
    return [getILearningRootRouter()];
  }
  return [];
};
