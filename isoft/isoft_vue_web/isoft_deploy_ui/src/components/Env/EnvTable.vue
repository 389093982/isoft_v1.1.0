<template>
<div style="margin-top: 10px;">
  <Table :columns="columns1" :data="envInfos" stripe border size="small"></Table>
  <Page :total="total" :current="pageNum" :page-size="pageSize" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
        @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>
</div>
</template>

<script>
  import {EnvList} from '../../api'
  import {ConnectTest} from '../../api'
  import {SyncDeployHome} from '../../api'
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
        total:10,     // 总数
        pageNum:1,    // 当前页
        pageSize:10   // 每页数量
      }
    },
    components:{Loading},
    methods:{
      refreshEnvList(){
        const _this = this;
        EnvList(_this.pageNum,_this.pageSize).then(function (response) {
          var result = JSON.parse(response);
          _this.envInfos = result.envInfos;
          _this.pageNum = result.paginator.currpage;
          _this.pageSize = result.paginator.pagesize;
          _this.total = result.paginator.totalcount;
        })
      },
      handlePageSizeChange (value){
        // 设置每页数量
        if(value != 10){
          this.pageSize= value;
          this.refreshEnvList();
        }
      },
      handleChange (value){
        if(value != 1){
          // 设置当前页数
          this.pageNum = value;
          // 重新加载表格数据
          this.refreshEnvList();
        }
      },
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
      this.refreshEnvList();
    }
  }
</script>

<style scoped>
</style>
