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

// 编辑或者新增博客分类
export const CatalogEdit = (catalog_name, catalog_desc) => ajax(BASE_URL+'/catalog/edit',{catalog_name, catalog_desc},'POST')

// 获取我的所有博客分类
export const GetMyCatalogs = () => ajax(BASE_URL+'/catalog/getMyCatalogs',{},'GET')

// 获取我的所有博客文章
export const GetMyBlogs = () => ajax(BASE_URL+'/blog/getMyBlogs',{},'GET')

// 编辑或者新增博客文章
export const BlogEdit = (blog_title, key_words, catalog_id, content) => ajax(BASE_URL+'/blog/edit',{blog_title, key_words, catalog_id, content},'POST')

//
export const BlogList = (offset,current_page) => ajax(BASE_URL+'/blog/blogList',{offset,current_page},'GET')

