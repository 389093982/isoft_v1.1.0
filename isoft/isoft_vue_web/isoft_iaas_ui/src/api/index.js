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

// 热门博客分页列表
export const BlogList = (offset,current_page) => ajax(BASE_URL+'/blog/blogList',{offset,current_page},'GET')

// 根据 blog_id 查询 blog 详细信息
export const ShowBlogDetail = (blog_id) => ajax(BASE_URL+'/blog/showBlogDetail',{blog_id},'GET')

// 新建课程
export const NewCourse = (course_name,course_type,course_sub_type,course_short_desc) => ajax(BASE_URL+'/ilearning/newCourse',{course_name,course_type,course_sub_type,course_short_desc},'GET')

// 分页查询我的课程清单
export const GetMyCourseList = () => ajax(BASE_URL+'/ilearning/getMyCourseList',{},'GET')

// 完结视频更新
export const EndUpdate = (course_id) => ajax(BASE_URL+'/ilearning/endUpdate',{course_id},'GET')

// 显示课程详细信息
export const ShowCourseDetail = (course_id) => ajax(BASE_URL+'/ilearning/showCourseDetail',{course_id},'GET')

// 切换收藏点赞
export const ToggleFavorite = (favorite_id, favorite_type) => ajax(BASE_URL+'/ilearning/toggle_favorite',{favorite_id, favorite_type},'GET')

// 查询评论主题
export const FilterCommentTheme = (comment_id, comment_type) => ajax(BASE_URL+'/ilearning/filterCommentTheme',{comment_id, comment_type},'GET')

// 添加评论
export const AddCommentReply = (parent_id, reply_content, comment_id, comment_type, refer_user_name) =>
  ajax(BASE_URL+'/ilearning/addCommentReply',{parent_id, reply_content, comment_id, comment_type, refer_user_name},'GET')

// 获取评论列表
export const FilterCommentReply = (comment_id, comment_type, parent_id) =>
  ajax(BASE_URL+'/ilearning/filterCommentReply',{comment_id, comment_type, parent_id},'GET')

