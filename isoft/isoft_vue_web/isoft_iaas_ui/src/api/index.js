/*
包含n个接口请求函数的模块
函数的返回值: promise对象
 */
import ajax from './ajax'
import store from "../store"

const BASE_URL = '/api'

// 查询所有的元数据信息
export const FilterPageMetadatas = (name,current_page,offset) => ajax(BASE_URL+'/metadata/filterPageMetadatas/',{name,current_page,offset},'POST')

// 分片定位请求
export const LocateShards = (hash) => ajax(BASE_URL+'/ifile/locateShards/',{hash},'POST')

// 单点登录获取 token 地址
export const GetJWTTokenByCode = (code) => ajax(BASE_URL+'/auth/getJWTTokenByCode/',{code},'POST')

