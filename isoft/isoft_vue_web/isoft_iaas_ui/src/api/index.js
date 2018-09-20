/*
包含n个接口请求函数的模块
函数的返回值: promise对象
 */
import ajax from './ajax'

const BASE_URL = '/api'

// 查询所有的元数据信息
export const MetaDataList = () => ajax(BASE_URL+'/ifile/metadatalist/',{},'POST')


