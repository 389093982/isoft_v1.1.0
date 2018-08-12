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
export const EnvList = (current_page,page_size) =>
  ajax(BASE_URL+'/env/list/',{current_page,page_size},'POST')

// 分页显示服务信息
export const ServiceList = (service_type,current_page,page_size) =>
  ajax(BASE_URL+'/service/list/',{service_type,current_page,page_size},'POST')


