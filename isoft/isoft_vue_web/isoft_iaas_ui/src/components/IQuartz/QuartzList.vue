<template>
  <div style="margin: 10px;">
    <ISimpleLeftRightRow>
      <ISimpleSearch slot="right" @handleSimpleSearch="handleSearch"/>
    </ISimpleLeftRightRow>

    <Table :columns="columns1" :data="quartzs" size="small"></Table>
    <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
          @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>
  </div>
</template>

<script>
  import {QuartzList} from "../../api"
  import ISimpleLeftRightRow from "../Common/layout/ISimpleLeftRightRow"
  import ISimpleSearch from "../Common/search/ISimpleSearch"

  export default {
    name: "QuartzList",
    components:{ISimpleLeftRightRow,ISimpleSearch},
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
      handleSearch(data){
        this.offset = 10;
        this.current_page = 1;
        this.search = data;
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
