<template>
  <div style="margin-top: 10px;">
    <Row style="margin-bottom: 5px;">
      <Col span="12"><ServiceAdd style="float: left;"/></Col>
      <Col span="12"><Button type="info" @click="hiddenColumnFunc()" style="float: right;">隐藏列</Button></Col>
    </Row>

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

    <!--
        一般来讲,获取DOM元素,需document.querySelector(".input1")获取这个dom节点,然后在获取input1的值.
        但是用ref绑定之后,我们就不需要在获取dom节点了,直接在上面的input上绑定input1,然后$refs里面调用就行.
        然后在javascript里面这样调用:this.$refs.input1 这样就可以减少获取dom节点的消耗了
     -->
    <Modal
      v-model="packageUploadModal"
      width="500"
      title="新增/编辑服务信息"
      :mask-closable="false">
      <div>
          <Upload
            ref="upload"
            multiple
            :on-success="uploadComplete"
            :data="{'service_id':packageUploadServiceId}"
            action="/api/service/fileUpload/">
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
        <MysqlInit @handleSubmit="mysqlInit_handleSubmit"/>
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
  import ServiceAdd from "../../components/Service/ServiceAdd.vue"
  import MysqlInit from "./MysqlInit"

  export default {
    name: "ServiceTable",
    components: {MysqlInit,ServiceAdd},
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
        serviceInfos: [],
        total:0,
        // 需要隐藏的列
        hiddenColumn:[]
      }
    },
    computed:{
      columns1(){
        let columns = [];
        // 通过计算属性动态添加列
        if($.inArray("env_name", this.hiddenColumn) < 0){
          columns.push({
            title: '环境名称',
            key: 'env_name',
            width:100
          });
        };
        if($.inArray("env_ip", this.hiddenColumn) < 0){
          columns.push({
            title: 'IP地址',
            key: 'env_ip',
            width:120
          });
        };
        if($.inArray("service_name", this.hiddenColumn) < 0){
          columns.push({
            title: '服务名称',
            key: 'service_name',
            width:150
          });
        };
        if($.inArray("service_type", this.hiddenColumn) < 0){
          columns.push({
            title: '服务类型',
            key: 'service_type',
            width:100
          });
        };
        if($.inArray("service_port", this.hiddenColumn) < 0){
          columns.push({
            title: '端口号',
            key: 'service_port',
            width:80
          });
        };
        if($.inArray("package_name", this.hiddenColumn) < 0 && $.inArray(this.$route.query.service_type, ["beego","api"])>=0){
          columns.push({
            title: '部署包名',
            key: 'package_name',
            width:200
          });
        }
        if($.inArray("run_mode", this.hiddenColumn) < 0 && $.inArray(this.$route.query.service_type, ["beego","api"])>=0){
          columns.push({
            title: '运行模式',
            key: 'run_mode',
            width:150
          });
        }
        if($.inArray("deploy_status", this.hiddenColumn) < 0){
          columns.push({
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
          });
        }
        columns.push({
          title: '操作',
          key: 'operate',
          width:650,
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
                    this.packageUploadServiceId = this.serviceInfos[params.index]['id'];
                    this.$refs.upload.fileList = [];
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
                  type: 'warning',
                  size: 'small'
                },
                style: {
                  marginRight: '5px',
                  // 控制按钮是否显示
                  display: $.inArray(this.$route.query.service_type, ["nginx","beego","api"])>=0  ? undefined : 'none'
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
                  display: $.inArray(this.$route.query.service_type, ["mysql"])>=0  ? undefined : 'none'
                },
                on: {
                  click: () => {
                    this.runDeployTask(params.index,"install", null)
                  }
                }
              }, '安装'),
              h('Button', {
                props: {
                  type: 'success',
                  size: 'small'
                },
                style: {
                  marginRight: '5px',
                  // 控制按钮是否显示
                  display: $.inArray(this.$route.query.service_type, ["other"])>=0  ? undefined : 'none'
                },
                on: {
                  click: () => {
                    this.runDeployTask(params.index,"uninstall", null)
                  }
                }
              }, '卸载'),
              h('Button', {
                props: {
                  type: 'error',
                  size: 'small'
                },
                style: {
                  marginRight: '5px',
                  // 控制按钮是否显示
                  display: $.inArray(this.$route.query.service_type, ["mysql"])>=0  ? undefined : 'none'
                },
                on: {
                  click: () => {
                    const _this = this;
                    _this.runDeployTask(params.index,"delete", null, function (data) {
                      if(data.status=="SUCCESS"){
                        // 删除当前元素
                        _this.serviceInfos = _this.serviceInfos.filter(t => t != _this.serviceInfos[params.index]);
                        // 友好提示
                        _this.$Notice.success({
                          title: '删除操作',
                          desc: '删除成功!'
                        });
                      }else{
                        _this.$Notice.error({
                          title: '删除操作',
                          desc: '删除失败!'
                        });
                      }
                    })
                  }
                }
              }, '删除'),
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
                  display: $.inArray(this.$route.query.service_type, ["nginx","beego","api"])>=0  ? undefined : 'none'
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
        });
        return columns;
      }
    },
    methods:{
      hiddenColumnFunc(){
        if(this.hiddenColumn == false){
          this.hiddenColumn = ["service_type","service_port","package_name","run_mode"];
        }else{
          this.hiddenColumn = [];
        }
      },
      mysqlInit_handleSubmit(data){
        // 关闭模态对话框
        this.mysqlInitModal = false;
        this.runDeployTask(this.mysqlInit_index, "mysql_init", JSON.stringify(data));
      },
      fileDownload(index){
        const service_id = this.serviceInfos[index]['id'];
        window.location='/api/service/fileDownload/?service_id=' + service_id;
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
          _serviceInfo.env_ip = _serviceInfo.env_info.env_ip;
          // 动态添加属性
          _serviceInfo['deploy_status'] = '';
          _serviceInfos.push(_serviceInfo);
        }
        _this.serviceInfos = _serviceInfos;
        _this.total = result.totalcount;
      },
      async runDeployTask(index, operate_type, extra_params, callback){
        const serviceInfo = this.serviceInfos[index];
        const data = await RunDeployTask(serviceInfo.env_id, serviceInfo.id, operate_type, extra_params);
        if(data.status=="SUCCESS"){
          // 有 tracking_id 才渲染转圈效果
          if(data.tracking_id != "" && data.tracking_id != null && data.tracking_id != undefined){
            // 设置转圈效果
            this.$set(this.serviceInfos[index], 'deploy_status', 'loading');
            // 获取 vue 实例
            const _this = this;
            var interval = setInterval(function () {
              // 渲染最后一次部署状态
              _this.renderLastDeployStatus(index,serviceInfo.id,interval);
            }, 2000);
          }
          if(callback != null){
            // 任务回调函数
            callback(data);
          }
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
