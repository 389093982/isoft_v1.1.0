/*
包含n个接口请求函数的模块
函数的返回值: promise对象
 */
import ajax from './ajax'

const BASE_URL = '/api'

// 1、编辑环境信息
export const EnvEdit = (env_name,env_ip,env_account,env_passwd,deploy_home) =>
  ajax(BASE_URL+'/env/edit/',{env_name,env_ip,env_account,env_passwd,deploy_home},'POST')

// 分页显示环境信息
export const EnvList = (current_page,offset) =>
  ajax(BASE_URL+'/env/list/',{current_page,offset},'POST')

// 查询所有的环境信息
export const EnvAll = () => ajax(BASE_URL+'/env/all/',{},'POST')

// 分页显示服务信息
export const ServiceList = (service_type,current_page,offset) =>
  ajax(BASE_URL+'/service/list/',{service_type,current_page,offset},'POST')

// 连接测试
export const ConnectTest = (env_id) =>
  ajax(BASE_URL+'/env/connect_test/',{env_id},'POST')

// 同步测试
export const SyncDeployHome = (env_id) =>
  ajax(BASE_URL+'/env/syncDeployHome/',{env_id},'POST')

// 编辑服务接口
export const ServiceEdit = (env_ids,service_name,service_type,package_name,run_mode,service_port,mysql_root_pwd) =>
  ajax(BASE_URL+'/service/edit/',{env_ids,service_name,service_type,package_name,run_mode,service_port,mysql_root_pwd},'POST')

// 运行部署任务
export const RunDeployTask = (env_id,service_id,operate_type, extra_params) =>
  ajax(BASE_URL+'/service/runDeployTask/',{env_id,service_id,operate_type, extra_params},'POST')

// 运行部署任务
export const QueryLastDeployStatus = (service_id) =>
  ajax(BASE_URL+'/service/queryLastDeployStatus/',{service_id},'POST')

// 获取服务运行日志详细信息
export const GetServiceTrackingLogDetail = (service_id) =>
  ajax(BASE_URL+'/service/getServiceTrackingLogDetail/',{service_id},'POST')

export const ConfigEdit = (env_ids, env_property, env_value) =>
  ajax(BASE_URL+'/config/edit/',{env_ids, env_property, env_value},'POST')

// 分页显示配置信息
export const ConfigList = (current_page,page_size) =>
  ajax(BASE_URL+'/config/list/',{current_page,page_size},'POST')

// 同步测试
export const SyncConfigFile = (env_id, configFile_id) =>
  ajax(BASE_URL+'/config/syncConfigFile/',{env_id, configFile_id},'POST')



