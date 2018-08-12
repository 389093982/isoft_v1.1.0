<template>
  <div style="margin-top: 10px;">
    <Table :columns="columns1" :data="serviceInfos" size="small"></Table>
    <Page :total="total" show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"/>
  </div>
</template>

<script>
  import {ServiceList} from '../../api'

  export default {
    name: "ServiceTable",
    data () {
      return {
        columns1: [
          {
            title: '环境ID',
            key: 'env_id'
          },
          {
            title: '环境名称',
            key: 'env_name'
          },
          {
            title: '服务名称',
            key: 'service_name'
          },
          {
            title: '服务类型',
            key: 'service_type'
          },
          {
            title: '端口号',
            key: 'service_port'
          },
          {
            title: '部署包名',
            key: 'package_name'
          },
          {
            title: '运行模式',
            key: 'run_mode'
          },
          {
            title: '部署状态',
            key: 'deploy_status'
          },
          {
            title: '操作',
            key: 'operate'
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
