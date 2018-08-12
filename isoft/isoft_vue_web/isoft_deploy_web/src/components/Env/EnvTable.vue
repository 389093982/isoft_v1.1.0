<template>
<div style="margin-top: 10px;">
  <Table :columns="columns1" :data="envInfos" size="small"></Table>
  <Page :total="total" show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"/>
</div>
</template>

<script>
  import {EnvList} from '../../api'
  export default {
    name: "EnvTable",
    data () {
      return {
        columns1: [
          {
            title: '环境名称',
            key: 'env_name'
          },
          {
            title: 'ip地址',
            key: 'env_ip'
          },
          {
            title: '登录账号',
            key: 'env_account'
          },
          {
            title: '密码',
            key: 'env_passwd'
          },
          {
            title: '部署空间目录',
            key: 'deploy_home'
          },
          {
            title: 'ssh连接',
            key: 'ssh_test'
          },
          {
            title: '操作',
            key: 'operate'
          }
        ],
        envInfos: [
          {
            env_name: '张三',
            env_ip: '127.0.0.1',
            env_account: 'root',
            env_passwd: 'admin@123456!',
            deploy_home: '/mydata/deploy_home'
          }
        ],
        total:0
      }
    },
    mounted:function(){
      const _this = this;
      EnvList(1,10).then(function (response) {
        var result = JSON.parse(response);
        _this.envInfos = result.envInfos;
        _this.total = result.totalcount;
      })
    }
  }
</script>

<style scoped>

</style>
