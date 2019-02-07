// 通过 mutation 间接更新 state 的多个方法的对象
export default {
  commitSetCurrent:function ({commit}, {current_work_id,current_work_step_id}) {
    commit('SETCURRENT',{current_work_id,current_work_step_id});
  }
}
