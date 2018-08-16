// 通过 mutation 间接更新 state 的多个方法的对象

import {
  EnvAll
} from '../api'

// getEnvAll(context) { context.commit("getEnvAll"); }
// 其实 actions 还可以简写一下,因为函数的参数是一个对象,函数中用的是对象中一个方法,我们可以通过
// 对象的解构赋值直接获取到该方法.
export default {
  // 异步获取环境清单
  getEnvAll: async ({commit, state}) => {
      // 发送异步ajax请求
      const result = await EnvAll();
      const data = JSON.parse(result);
      const envInfos = data.envInfos;
      // 提交一个mutation
      commit("getEnvAll",{envInfos});
  },
}
