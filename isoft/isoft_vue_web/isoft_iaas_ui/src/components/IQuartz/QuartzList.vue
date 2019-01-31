<template>
  <div style="margin: 10px;">
    <Row type="flex" justify="end" class="code-row-bg" style="margin-bottom: 5px;">
      <Col span="10">
        <Input v-model="search" placeholder="搜索..."/>
      </Col>
      <Col span="2">
        <Button type="success" @click="searchRecord">搜索</Button>
      </Col>
    </Row>

    <Table :columns="columns1" :data="quartzs" size="small"></Table>
    <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
          @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>
  </div>
</template>

<script>
  import {QuartzList} from "../../api"

  export default {
    name: "QuartzList",
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
        quartzs: [],
        columns1: [
          {
            title: 'task_name',
            key: 'task_name',
          },
          {
            title: 'task_type',
            key: 'task_type',
          },
          {
            title: 'task_id',
            key: 'task_id',
          },
          {
            title: 'cron_str',
            key: 'cron_str',
          },
          {
            title: 'created_by',
            key: 'created_by',
          },
          {
            title: 'created_time',
            key: 'created_time',
          },
          {
            title: 'last_updated_by',
            key: 'last_updated_by',
          },
          {
            title: 'last_updated_time',
            key: 'last_updated_time'
          }
        ],
      }
    },
    methods:{
      refreshQuartzList:async function () {
        const result = await QuartzList(this.offset,this.current_page,this.search);
        if(result.status=="SUCCESS"){
          this.quartzs = result.quartzs;
        }
      },
      handleChange(page){
        this.current_page = page;
        this.refreshQuartzList();
      },
      handlePageSizeChange(pageSize){
        this.offset = pageSize;
        this.refreshQuartzList();
      },
      searchRecord(){
        this.offset = 10;
        this.current_page = 1;
        this.refreshQuartzList();
      }
    },
    mounted: function () {
      this.refreshQuartzList();
    },
  }
</script>

<style scoped>

</style>
