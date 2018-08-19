<template>
  <div style="margin-top: 10px;">
    <Table :columns="columns1" :data="serviceInfos" size="small"></Table>
    <Page :total="total" show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"/>
  </div>
</template>

<script>
  import {ServiceList} from '../../api'
  import {RunDeployTask} from '../../api'

  export default {
    name: "ServiceTable",
    data () {
      return {
        columns1: [
          {
            title: '环境ID',
            key: 'env_id',
            width:100
          },
          {
            title: '环境名称',
            key: 'env_name',
            width:100
          },
          {
            title: '服务名称',
            key: 'service_name',
            width:100
          },
          {
            title: '服务类型',
            key: 'service_type',
            width:100
          },
          {
            title: '端口号',
            key: 'service_port',
            width:100
          },
          {
            title: '部署包名',
            key: 'package_name',
            width:100
          },
          {
            title: '运行模式',
            key: 'run_mode',
            width:100
          },
          {
            title: '部署状态',
            key: 'deploy_status',
            width:100
          },
          {
            title: '操作',
            key: 'operate',
            width:550,
            render: (h, params) => {
              return h('div', [
                h('Button', {
                  props: {
                    type: 'info',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px'
                  },
                  on: {
                    click: () => {
                      this.runDeployTask(params.index,"check")
                    }
                  }
                }, '检测'),
                h('Button', {
                  props: {
                    type: 'warning',
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
                }, '安装'),
                h('Button', {
                  props: {
                    type: 'error',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px'
                  },
                  on: {
                    click: () => {
                      this.sync_deploy_home(params.index)
                    }
                  }
                }, '重启'),
                h('Button', {
                  props: {
                    type: 'primary',
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
                }, '详情'),
                h('Button', {
                  props: {
                    type: 'info',
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
                }, '部署'),
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
                }, '启用'),
                h('Button', {
                  props: {
                    type: 'warning',
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
                }, '停用'),
                h('Button', {
                  props: {
                    type: 'error',
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
                }, '服务代理'),
              ]);
            }
          }
        ],
        serviceInfos: [],
        total:0
      }
    },
    methods:{
      refreshServiceList(){
        const _this = this;
        ServiceList(this.$route.query.service_type,1,10).then(function (response) {
          var result = JSON.parse(response);
          var _serviceInfos=[];
          for(var i=0; i<result.serviceInfos.length; i++){
            var _serviceInfo = result.serviceInfos[i];
            _serviceInfo.env_id = _serviceInfo.env_info.id;
            _serviceInfo.env_name = _serviceInfo.env_info.env_name;
            _serviceInfos.push(_serviceInfo);
          }
          _this.serviceInfos = _serviceInfos;
          _this.total = result.totalcount;
        })
      },
      runDeployTask(index,operate_type){
        const _this = this;
        const serviceInfo = this.serviceInfos[index];
        RunDeployTask(serviceInfo.env_id, serviceInfo.service_id, operate_type).then(function (response) {
          alert(response);
        })
      }
    },
    mounted:function(){
      this.refreshServiceList();
    },
    watch:{
      "$route": "refreshServiceList"      // 如果路由有变化,会再次执行该方法
    }
  }
</script>

<style scoped>

</style>
