<template>
  <LeftMenu>
    <div style="margin: 10px;">
      <Row type="flex" justify="end" class="code-row-bg" style="margin-bottom: 5px;">
        <Col span="10">
          <Input v-model="search" placeholder="搜索..."/>
        </Col>
        <Col span="2">
          <Button type="success" @click="searchRecord">搜索</Button>
        </Col>
      </Row>

      <Table :columns="columns1" :data="loginRecords" size="small"></Table>
      <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
            @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>
    </div>
  </LeftMenu>
</template>

<script>
  import {LoginRecordList} from "../../api"
  import LeftMenu from "./LeftMenu";

  export default {
    name: "LoginRecord",
    components: {LeftMenu},
    data(){
      return {
        // 当前页
        current_page:1,
        // 总页数
        total:1,
        // 每页记录数
        offset:10,
        // 搜索条件
        search:"",
        loginRecords: [],
        columns1: [
          {
            title: 'origin',
            key: 'origin',
            width:200
          },
          {
            title: 'referer',
            key: 'referer',
            width:400
          },
          {
            title: '登录ip',
            key: 'login_ip',
            width:80
          },
          {
            title: '登录用户',
            key: 'user_name',
            width:80
          },
          {
            title: '登录状态',
            key: 'login_status',
            width:80
          },
          {
            title: '登录时间',
            key: 'created_time',
          },
        ],
      }
    },
    methods:{
      refreshLoginRecordList: async function(){
        const result = await LoginRecordList(this.offset, this.current_page, this.search);
        if(result.status=="SUCCESS"){
          this.loginRecords = result.loginRecords;
          this.total = result.paginator.totalcount;
        }
      },
      handleChange(page){
        this.current_page = page;
        this.refreshLoginRecordList();
      },
      handlePageSizeChange(pageSize){
        this.offset = pageSize;
        this.refreshLoginRecordList();
      },
      searchRecord(){
        this.offset = 10;
        this.current_page = 1;
        this.refreshLoginRecordList();
      }
    },
    mounted: function () {
      this.refreshLoginRecordList();
    },
  }
</script>

<style scoped>

</style>
