<template>
  <div>
    <Row style="margin-bottom: 10px;">
      <Col span="12">
        <IFileUpload/>
      </Col>
      <Col span="12">
        <Input v-model="search_name" search enter-button placeholder="搜索对象名称" @on-search="input_search"/>
      </Col>
    </Row>

    <Table :columns="columns1" :data="metadatas" size="small" height="450"></Table>
    <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
          @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>
  </div>
</template>

<script>
  import IFileUpload from "./IFileUpload.vue"
  import {FilterPageMetadatas} from '../../api'

  export default {
    name: "IFile",
    components: {IFileUpload},
    data(){
      return {
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
            title: 'Name',
            key: 'Name',
            width:300
          },
          {
            title: 'Version',
            key: 'Version',
            width:100
          },
          {
            title: 'Size',
            key: 'Size',
            width:100
          },
          {
            title: 'Hash',
            key: 'Hash',
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
      }
    },
    mounted:function(){
      this.refreshMetaDataList();
    },
  }
</script>

<style scoped>

</style>
