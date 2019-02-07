/*
直接更新state的多个方法的对象
 */
import Vue from 'vue'


export default {
  setCurrent:function (state, {current_work_id,current_work_step_id}) {
    state.current_work_id=current_work_id;
    state.current_work_step_id=current_work_step_id;
  }
}
