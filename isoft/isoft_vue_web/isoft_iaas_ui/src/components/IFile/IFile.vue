<template>
  <div>
    <Row style="margin-bottom: 10px;">
      <Col span="12">
        <IFileUpload @refreshTable="refreshMetaDataList" action="/api/ifile/fileUpload/" uploadLabel="上传到文件服务器"/>
      </Col>
      <Col span="12">
        <Input v-model="search_name" search enter-button placeholder="搜索对象名称" @on-search="input_search"/>
      </Col>
    </Row>

    <Table :columns="columns1" :data="metadatas" size="small" height="450"></Table>
    <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
          @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>

    <Modal
      v-model="showShardsModel"
      title="对象分片信息"
      :mask-closable="false">
      <p v-for="(key,value) in shards">
        存储物理机器地址:{{key}}  -  分片id:{{value}}
      </p>
    </Modal>

    <Modal
      v-model="showImgModel"
      title="显示图片"
      :mask-closable="false">
      <img :src="showImageSrc" alt="smile" />
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
  import IFileUpload from "./IFileUpload.vue"
  import {FilterPageMetadatas} from '../../api'
  import {LocateShards} from '../../api'

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
        // 对象分片信息
        shards:[],
        // 搜索的对象名称
        search_name : "",
        // 当前页
        current_page:1,
        // 总页数
        total:1,
        // 每页记录数
        offset:10,
        // 元数据信息
        metadatas: [],
        columns1 : [
          {
            title: 'name',
            key: 'name',
            width:300
          },
          {
            title: 'version',
            key: 'version',
            width:100
          },
          {
            title: 'size',
            key: 'size',
            width:100
          },
          {
            title: 'hash',
            key: 'hash',
            width:380,
          },
          {
            title: 'app_name',
            key: 'app_name',
            width:100,
          },
          {
            title: '操作',
            key: 'operate',
            width:400,
            render: (h, params) => {
              return h('div', [
                h('Button', {
                  props: {
                    type: 'success',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px',
                  },
                  on: {
                    click: () => {
                      // 分片定位
                      this.locateShards(params.index);
                    }
                  }
                }, '分片查询'),
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
      async refreshMetaDataList(){
        const data = await FilterPageMetadatas(this.search_name, this.current_page, this.offset);
        if(data.status == "SUCCESS"){
          this.metadatas = data.metadatas;
          this.total = data.paginator.totalcount;
        }
      },
      handleChange(page){
        this.current_page = page;
        this.refreshMetaDataList();
      },
      handlePageSizeChange(pageSize){
        this.offset = pageSize;
        this.refreshMetaDataList();
      },
      input_search(){
        this.refreshMetaDataList();
      },
      async locateShards(index){
        const hash = this.metadatas[index]['hash'];
        const data = await LocateShards(hash);
        if(data.status=="SUCCESS"){
          this.shards = data.shards;
          this.showShardsModel = true;
        }
      },
      fileDownload(index){
        const name = this.metadatas[index]['name'];
        const version = this.metadatas[index]['version'];
        const app_name = this.metadatas[index]['app_name'];
        window.location='/api/ifile/fileDownload/?name=' + name + "&version=" + version + "&app_name=" + app_name;
      },
      showImg(index){
        const name = this.metadatas[index]['name'];
        const version = this.metadatas[index]['version'];
        const hash = this.metadatas[index]['hash'];
        this.showImageSrc = "http://127.0.0.1:10001/download/" + hash  + ".mp4";
        this.showImgModel = true;
      },
      playVideo(index){
        const name = this.metadatas[index]['name'];
        const version = this.metadatas[index]['version'];
        const hash = this.metadatas[index]['hash'];
        this.$refs.video.src = "http://127.0.0.1:10001/download/" + hash  + ".mp4";
        this.playVideoModel = true;
      },
    },
    mounted:function(){
      this.refreshMetaDataList();
    },
  }
</script>

<style scoped>

</style>
