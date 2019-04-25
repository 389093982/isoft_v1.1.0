<template>
  <div>
    <Row style="margin-bottom: 10px;">
      <Col span="12">
        <IFileUpload @refreshTable="refreshIFileList" action="/api/ifile/fileUpload/" uploadLabel="上传到文件服务器"/>
      </Col>
      <Col span="12">
        <Input v-model="search_name" search enter-button placeholder="搜索对象名称" @on-search="input_search"/>
      </Col>
    </Row>

    <Table :columns="columns1" :data="ifiles" size="small" height="450"></Table>
    <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
          @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>

    <Modal
      v-model="showImgModel"
      title="显示图片"
      :mask-closable="false">
      <img :src="showImageSrc" alt="smile" style="width: 130px;height: 150px;"/>
    </Modal>

    <Modal
      v-model="playVideoModel"
      title="播放视频"
      :mask-closable="false">
      <video ref="video" width="320" height="240" controls>
        <source type="video/mp4">
        您的浏览器不支持 video 标签。
      </video>
    </Modal>
  </div>
</template>

<script>
  import IFileUpload from "./IFileUpload"
  import {FilterPageIFiles} from '../../api'

  export default {
    name: "IFile",
    components: {IFileUpload},
    data(){
      return {
        showImgModel:false,
        showImageSrc:'',
        playVideoModel:false,
        playVideoSrc:'',
        // 显示对象分片信息对话框
        showShardsModel:false,
        // 搜索的对象名称
        search_name : "",
        // 当前页
        current_page:1,
        // 总页数
        total:1,
        // 每页记录数
        offset:10,
        // 文件元素将清单
        ifiles: [],
        columns1 : [
          {
            title: 'fid',
            key: 'fid',
            width:150
          },
          {
            title: 'file_name',
            key: 'file_name',
            width:200
          },
          {
            title: 'file_size',
            key: 'file_size',
            width:100
          },
          {
            title: 'url',
            key: 'url',
            width:270,
          },
          {
            title: 'created_time',
            key: 'created_time',
            width:200,
          },
          {
            title: '操作',
            key: 'operate',
            width:300,
            render: (h, params) => {
              return h('div', [
                h('Button', {
                  props: {
                    type: 'error',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px',
                  },
                  on: {
                    click: () => {
                      this.fileDownload(params.index);
                    }
                  }
                }, '文件下载'),
                h('Button', {
                  props: {
                    type: 'info',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px',
                  },
                  on: {
                    click: () => {
                      this.showImg(params.index);
                    }
                  }
                }, '图片预览'),
                h('Button', {
                  props: {
                    type: 'warning',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px',
                  },
                  on: {
                    click: () => {
                      this.playVideo(params.index);
                    }
                  }
                }, '视频播放'),
              ]);
            }
          }
        ]
      }
    },
    methods:{
      async refreshIFileList(){
        const data = await FilterPageIFiles(this.search_name, this.offset, this.current_page);
        if(data.status == "SUCCESS"){
          this.ifiles = data.ifiles;
          this.total = data.paginator.totalcount;
        }
      },
      handleChange(page){
        this.current_page = page;
        this.refreshIFileList();
      },
      handlePageSizeChange(pageSize){
        this.offset = pageSize;
        this.refreshIFileList();
      },
      input_search(){
        this.refreshIFileList();
      },
      fileDownload(index){
        const url = this.ifiles[index]['url'];
        window.location=url;
      },
      showImg(index){
        const url = this.ifiles[index]['url'];
        this.showImageSrc = url;
        this.showImgModel = true;
      },
      playVideo(index){
        const url = this.ifiles[index]['url'];
        this.$refs.video.src = url;
        this.playVideoModel = true;
      },
    },
    mounted:function(){
      this.refreshIFileList();
    },
  }
</script>

<style scoped>

</style>
