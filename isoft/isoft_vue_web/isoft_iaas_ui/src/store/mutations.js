/*
直接更新state的多个方法的对象
 */
import Vue from 'vue'

import {RECEIVE_HOTCOURSETYPECONFIGURATIONS} from "./mutation-types"

export default {
  [RECEIVE_HOTCOURSETYPECONFIGURATIONS] (state, {hotCourseTypeConfigurations}) {
    state.hotCourseTypeConfigurations = hotCourseTypeConfigurations;
  },
}
