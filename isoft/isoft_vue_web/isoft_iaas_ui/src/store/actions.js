// 通过 mutation 间接更新 state 的多个方法的对象
import {
  RECEIVE_HOTCOURSETYPECONFIGURATIONS,
} from './mutation-types'

import {
  QueryAllConfigurations,
} from '../api'

export default {
  // 异步获取热门课程类型配置项
  async getHotCourseTypeConfigurations({commit, state}) {
    // 发送异步ajax请求
    const result = await QueryAllConfigurations("recommand_course_type");
    // 提交一个mutation
    if (result.status == "SUCCESS") {
      const hotCourseTypeConfigurations = result.configurations;
      commit(RECEIVE_HOTCOURSETYPECONFIGURATIONS, {hotCourseTypeConfigurations})
    }
  },
}
