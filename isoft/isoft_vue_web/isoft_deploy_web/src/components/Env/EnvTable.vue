<template>
<div style="margin-top: 10px;">
  <Table :columns="columns1" :data="envInfos" stripe border size="small"></Table>
  <Page :total="total" show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"/>
</div>
</template>

<script>
  import {EnvList} from '../../api'
  import {ConnectTest} from '../../api'
  import Loading from '../../components/Common/Loading.vue'

  export default {
    name: "EnvTable",
    data () {
      return {
        columns1: [
          {
            title: '环境名称',
            key: 'env_name',
            width:150
          },
          {
            title: 'ip地址',
            key: 'env_ip',
            width:150
          },
          {
            title: '登录账号',
            key: 'env_account',
            width:150
          },
          {
            title: '密码',
            key: 'env_passwd',
            width:150
          },
          {
            title: '部署空间目录',
            key: 'deploy_home',
            width:200
          },
          {
            title: '操作结果',
            key: 'operate_result',
            width:100,
            render: (h, params) => {
              // 动态渲染子组件
              var _operate_result = params['row']['operate_result'];
              if(_operate_result=='loading'){
                return h(Loading);
              }else{
                return h('div',_operate_result);
              }
            }
          },
          {
            title: '操作',
            key: 'operate',
            width:250,
            render: (h, params) => {
              return h('div', [
                h('Button', {
                  props: {
                    type: 'success',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px'
                  },
                  on: {
                    click: () => {
                      this.connection_test(params.index)
                    }
                  }
                }, '连接测试'),
                h('Button', {
                  props: {
                    type: 'error',
                    size: 'small'
                  },
                  on: {
                    click: () => {
                      this.sync_deploy_home(params.index)
                    }
                  }
                }, '同步deploy_home')
              ]);
            }
          }
        ],
        envInfos: [],
        total:0
      }
    },
    components:{Loading},
    methods:{
      connection_test (index){
        // 当前行对应的环境 id
        var env_id = this.envInfos[index].id;
        const _this = this;
        // 设置转圈效果
        _this.$set(_this.envInfos[index], 'operate_result', 'loading');
        ConnectTest(env_id).then(function (response) {
          if(response.status == 'SUCCESS'){
            _this.$set(_this.envInfos[index], 'operate_result', '连接成功!');
          }else if(response.status == 'ERROR'){
            _this.$set(_this.envInfos[index], 'operate_result', '连接失败!');
          }
        })
      },
      sync_deploy_home (index){
        // 当前行对应的环境 id
        var env_id = this.envInfos[index].id;
        const _this = this;
        // 设置转圈效果
        _this.$set(_this.envInfos[index], 'operate_result', 'loading');
        SyncDeployHome(env_id).then(function (response) {
          if(response.status == 'SUCCESS'){
            _this.$set(_this.envInfos[index], 'operate_result', '同步成功!');
          }else if(response.status == 'ERROR'){
            _this.$set(_this.envInfos[index], 'operate_result', '同步失败!');
          }
        })
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
