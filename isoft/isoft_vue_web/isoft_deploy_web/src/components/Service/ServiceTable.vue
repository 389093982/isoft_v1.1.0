<template>
  <div style="margin-top: 10px;">
    <Table :columns="columns1" :data="serviceInfos" size="small" height="400"></Table>
    <Page :total="total" show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"/>

    <Modal
      v-model="showServiceTrackingLogDetailFlag"
      width="800"
      title="最近一次日志详情"
      :mask-closable="false">
        <Scroll>
          <p v-for="trackingLog in trackingLogs">
            {{trackingLog.tracking_detail}}
          </p>
        </Scroll>
    </Modal>

    <Modal
      v-model="packageUploadModal"
      width="500"
      title="新增/编辑服务信息"
      :mask-closable="false">
      <div>
          <Upload
            :on-success="uploadComplete"
            :data="{'service_id':packageUploadServiceId}"
            action="/api/v1/service/fileUpload/">
            <Button icon="ios-cloud-upload-outline">软件上传</Button>
          </Upload>
      </div>
    </Modal>

    <Modal
      v-model="mysqlInitModal"
      width="500"
      title="新建账号和数据库"
      :footer-hide="true"
      :mask-closable="false">
      <div>
        <MysqlInit @handleSubmit="mysqlInit_handleSubmit" :index="mysqlInit_index"/>
      </div>
    </Modal>
  </div>
</template>

<script>
  import {ServiceList} from '../../api'
  import {RunDeployTask} from '../../api'
  import {QueryLastDeployStatus} from '../../api'
  import {GetServiceTrackingLogDetail} from '../../api'
  import {FileDownload} from '../../api'
  import Loading from '../../components/Common/Loading.vue'
  import MysqlInit from "./MysqlInit";

  export default {
    name: "ServiceTable",
    components: {MysqlInit},
    data () {
      return {
        // 日志详情标志
        showServiceTrackingLogDetailFlag: false,
        // 日志详情内容
        trackingLogs:[],
        // 软件包上传 modal
        packageUploadModal: false,
        // 上传时附带的额外参数
        packageUploadServiceId:'',
        // 新建账号和数据库 modal
        mysqlInitModal: false,
        // 新建账号和数据库 modal 时所选择的 index
        mysqlInit_index: -1,
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
            width:200
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
            width:200
          },
          {
            title: '运行模式',
            key: 'run_mode',
            width:100
          },
          {
            title: '操作结果',
            key: 'deploy_status',
            width:100,
            render: (h, params) => {
              // 动态渲染子组件
              var operate_result = params['row']['deploy_status'];
              if(operate_result=='loading'){
                return h(Loading);
              }else{
                return h('div',operate_result);
              }
            }
          },
          {
            title: '操作',
            key: 'operate',
            width:550,
            render: (h, params) => {
              return h('div', [
                h('Button', {
                  props: {
                    type: 'success',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px',
                    // 控制按钮是否显示
                    display: $.inArray(this.$route.query.service_type, ["beego","api"])>=0  ? undefined : 'none'
                  },
                  on: {
                    click: () => {
                      this.packageUploadModal = true;
                      this.packageUploadServiceId = this.serviceInfos[params.index]['id']
                    }
                  }
                }, '软件包上传'),
                h('Button', {
                  props: {
                    type: 'error',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px',
                    // 控制按钮是否显示
                    display: $.inArray(this.$route.query.service_type, ["beego","api"])>=0  ? undefined : 'none'
                  },
                  on: {
                    click: () => {
                      this.fileDownload(params.index);
                    }
                  }
                }, '软件包下载'),
                h('Button', {
                  props: {
                    type: 'info',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px',
                    // 控制按钮是否显示
                    display: $.inArray(this.$route.query.service_type, ["mysql","nginx","beego","api"])>=0  ? undefined : 'none'
                  },
                  on: {
                    click: () => {
                      this.runDeployTask(params.index,"check", null)
                    }
                  }
                }, '检测'),
                h('Button', {
                  props: {
                    type: 'warning',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px',
                    // 控制按钮是否显示
                    display: $.inArray(this.$route.query.service_type, ["mysql","nginx"])>=0  ? undefined : 'none'
                  },
                  on: {
                    click: () => {
                      this.runDeployTask(params.index,"install", null)
                    }
                  }
                }, '安装'),
                h('Button', {
                  props: {
                    type: 'info',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px',
                    // 控制按钮是否显示
                    display: $.inArray(this.$route.query.service_type, ["beego","api"])>=0  ? undefined : 'none'
                  },
                  on: {
                    click: () => {
                      this.runDeployTask(params.index,"deploy", null)
                    }
                  }
                }, '部署'),
                h('Button', {
                  props: {
                    type: 'error',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px',
                    // 控制按钮是否显示
                    display: $.inArray(this.$route.query.service_type, ["mysql","nginx","beego","api"])>=0  ? undefined : 'none'
                  },
                  on: {
                    click: () => {
                      this.runDeployTask(params.index,"restart", null)
                    }
                  }
                }, '重启'),
                h('Button', {
                  props: {
                    type: 'success',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px',
                    // 控制按钮是否显示
                    display: $.inArray(this.$route.query.service_type, ["beego","api"])>=0  ? undefined : 'none'
                  },
                  on: {
                    click: () => {
                      this.runDeployTask(params.index,"startup", null)
                    }
                  }
                }, '启用'),
                h('Button', {
                  props: {
                    type: 'warning',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px',
                    // 控制按钮是否显示
                    display: $.inArray(this.$route.query.service_type, ["beego","api"])>=0  ? undefined : 'none'
                  },
                  on: {
                    click: () => {
                      this.runDeployTask(params.index,"shutdown", null)
                    }
                  }
                }, '停用'),
                h('Button', {
                  props: {
                    type: 'error',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px',
                    // 控制按钮是否显示
                    display: $.inArray(this.$route.query.service_type, ["nginx"])>=0  ? undefined : 'none'
                  },
                  on: {
                    click: () => {
                    }
                  }
                }, '服务代理'),
                h('Button', {
                  props: {
                    type: 'info',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px',
                    // 控制按钮是否显示
                    display: $.inArray(this.$route.query.service_type, ["mysql"])>=0  ? undefined : 'none'
                  },
                  on: {
                    click: () => {
                      this.runDeployTask(params.index,"connection_test", null)
                    }
                  }
                }, '连接测试'),
                h('Button', {
                  props: {
                    type: 'success',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px',
                    // 控制按钮是否显示
                    display: $.inArray(this.$route.query.service_type, ["mysql"])>=0  ? undefined : 'none'
                  },
                  on: {
                    click: () => {
                      this.mysqlInit_index = params.index;
                      this.mysqlInitModal = true;
                    }
                  }
                }, '新建账号和DB'),
                h('Button', {
                  props: {
                    type: 'default',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px',
                    // 控制按钮是否显示
                    display: $.inArray(this.$route.query.service_type, ["mysql","nginx","beego","api"])>=0  ? undefined : 'none'
                  },
                  on: {
                    click: () => {
                      this.getServiceTrackingLogDetail(params.index);
                    }
                  }
                }, '详情'),
              ]);
            }
          }
        ],
        serviceInfos: [],
        total:0
      }
    },
    methods:{
      mysqlInit_handleSubmit(data){
        // 关闭模态对话框
        this.mysqlInitModal = false;
        this.runDeployTask(this.mysqlInit_index, "mysql_init", JSON.stringify(data));
      },
      fileDownload(index){
        const service_id = this.serviceInfos[index]['id'];
        window.location='/api/v1/service/fileDownload/?service_id=' + service_id;
      },
      uploadComplete(res, file) {
        if(res.status=="SUCCESS"){
          this.$Notice.success({
            title: '文件上传成功',
            desc: '文件 ' + file.name + ' 上传成功!'
          });
        }else{
          this.$Notice.error({
            title: '文件上传失败',
            desc: '文件 ' + file.name + ' 上传失败!'
          });
        }
      },
      async getServiceTrackingLogDetail(index){
        const data = await GetServiceTrackingLogDetail(this.serviceInfos[index]['id']);
        if(data.status=="SUCCESS"){
          this.trackingLogs = data.trackingLogs;
          this.showServiceTrackingLogDetailFlag = true;
        }
      },
      async renderLastDeployStatus(index,service_id,interval){
        // 设置转圈效果
        this.$set(this.serviceInfos[index], 'deploy_status', 'loading');
        // 异步调用接口
        const data = await QueryLastDeployStatus(service_id);
        if(data.status == 'SUCCESS' && data.finish == true){
          this.$set(this.serviceInfos[index], 'deploy_status', data.trackingStatus);
          clearInterval(interval);
        }
      },
      async refreshServiceList(){
        const _this = this;
        const data = await ServiceList(this.$route.query.service_type,1,10);
        var result = JSON.parse(data);
        var _serviceInfos=[];
        for(var i=0; i<result.serviceInfos.length; i++){
          var _serviceInfo = result.serviceInfos[i];
          _serviceInfo.env_id = _serviceInfo.env_info.id;
          _serviceInfo.env_name = _serviceInfo.env_info.env_name;
          // 动态添加属性
          _serviceInfo['deploy_status'] = '';
          _serviceInfos.push(_serviceInfo);
        }
        _this.serviceInfos = _serviceInfos;
        _this.total = result.totalcount;
      },
      async runDeployTask(index,operate_type, extra_params){
        const serviceInfo = this.serviceInfos[index];
        const data = await RunDeployTask(serviceInfo.env_id, serviceInfo.id, operate_type, extra_params);
        if(data.status=="SUCCESS"){
          // 设置转圈效果
          this.$set(this.serviceInfos[index], 'deploy_status', 'loading');
          // 获取 vue 实例
          const _this = this;
          var interval = setInterval(function () {
            // 渲染最后一次部署状态
            _this.renderLastDeployStatus(index,serviceInfo.id,interval);
          }, 2000);
        }
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
