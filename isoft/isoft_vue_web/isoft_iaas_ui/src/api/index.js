/*
包含n个接口请求函数的模块
函数的返回值: promise对象
 */
import ajax from './ajax'

const BASE_URL = '/api'

// 查询所有的元数据信息
export const FilterPageMetadatas = (current_page,offset) => ajax(BASE_URL+'/ifile/filterPageMetadatas/',{current_page,offset},'POST')


