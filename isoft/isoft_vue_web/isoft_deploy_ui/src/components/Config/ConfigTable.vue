<template>
  <div style="margin-top: 10px;">
    <Table :columns="columns1" :data="configFiles" size="small" height="400"></Table>
    <Page :total="total" show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"/>

    <Modal
      v-model="configUploadModal"
      width="500"
      title="配置包上传"
      :mask-closable="false">
      <div>
        <Upload
          ref="upload"
          multiple
          :on-success="uploadComplete"
          :data="{'configFile_id':uploadConfigFileId}"
          action="/api/config/fileUpload/">
          <Button icon="ios-cloud-upload-outline">配置包上传</Button>
        </Upload>
      </div>
    </Modal>
  </div>
</template>

<script>
  import {ConfigList} from '../../api'
  import {SyncConfigFile} from '../../api'

  export default {
    name: "ConfigTable",
    data () {
      return {
        configUploadModal:false,
        // 上传时附带的额外参数
        uploadConfigFileId:'',
        configFiles:[],
        total:0
      }
    },
    computed:{
      columns1(){
        let columns = [];
        // 通过计算属性动态添加列
        columns.push({
          title: '环境名称',
          key: 'env_name',
          width:100
        });
        columns.push({
          title: 'IP地址',
          key: 'env_ip',
          width:120
        });
        columns.push({
          title: '环境变量',
          key: 'env_property',
          width:200
        });
        columns.push({
          title: '变量值/配置包路径',
          key: 'env_value',
          width:250
        });
        columns.push({
          title: '操作',
          key: 'operate',
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
                    this.configUploadModal = true;
                    this.uploadConfigFileId = this.configFiles[params.index]['id'];
                    this.$refs.upload.fileList = [];
                  }
                }
              }, '配置包上传'),
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
                    this.fileDownload(params.index);
                  }
                }
              }, '配置包下载'),
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
                    this.sync_config_file(params.index);
                  }
                }
              }, '配置包同步'),
            ]);
          }
        });
        return columns;
      }
    },
    methods:{
      async sync_config_file (index){
        // 当前行对应的环境 id
        var configFile_id = this.configFiles[index].id;
        var env_id = this.configFiles[index].env_id;
        const data = await SyncConfigFile(env_id, configFile_id);
        if(data.status=="SUCCESS"){
          // 友好提示
          this.$Notice.success({
            title: '同步操作',
            desc: '同步成功!'
          });
        }else{
          this.$Notice.error({
            title: '同步操作',
            desc: '同步失败!'
          });
        }
      },
      fileDownload(index){
        const configFile_id = this.configFiles[index]['id'];
        window.location='/api/config/fileDownload/?configFile_id=' + configFile_id;
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
      async refreshConfigList(){
        const _this = this;
        const data = await ConfigList(1,10);
        var result = JSON.parse(data);
        var _configFiles=[];
        for(var i=0; i<result.configFiles.length; i++){
          var _configFile = result.configFiles[i];
          _configFile.env_id = _configFile.env_info.id;
          _configFile.env_name = _configFile.env_info.env_name;
          _configFile.env_ip = _configFile.env_info.env_ip;
          _configFiles.push(_configFile);
        }
        _this.configFiles = _configFiles;
        _this.total = result.totalcount;
      }
    },
    mounted:function(){
      this.refreshConfigList();
    },
    watch:{
      "$route": "refreshConfigList"      // 如果路由有变化,会再次执行该方法
    }
  }
</script>

<style scoped>

</style>
