/*
直接更新state的多个方法的对象
 */
import Vue from 'vue'


export default {
  // 设置当前 work_id 和 work_step_id
  SETCURRENT:function (state, {current_work_id,current_work_step_id}) {
    state.current_work_id=current_work_id;
    state.current_work_step_id=current_work_step_id;
  }
}
