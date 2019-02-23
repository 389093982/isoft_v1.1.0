<template>
  <div style="margin: 10px;">
    <ISimpleLeftRightRow>
      <!-- left 插槽部分 -->
      <QuartzAdd slot="left" @handleSuccess="refreshQuartzList"/>
      <!-- right 插槽部分 -->
      <ISimpleSearch slot="right" @handleSimpleSearch="handleSearch"/>
    </ISimpleLeftRightRow>

    <Table border :columns="columns1" :data="quartzs" size="small"></Table>
    <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
          @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>
  </div>
</template>

<script>
  import {formatDate} from "../../../tools/index"
  import {QuartzList} from "../../../api/index"
  import ISimpleLeftRightRow from "../../Common/layout/ISimpleLeftRightRow"
  import ISimpleSearch from "../../Common/search/ISimpleSearch"
  import QuartzAdd from "./QuartzAdd"

  export default {
    name: "QuartzList",
    components:{ISimpleLeftRightRow,ISimpleSearch,QuartzAdd},
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
            render: (h,params)=>{
              return h('div',
                formatDate(new Date(params.row.created_time),'yyyy-MM-dd hh:mm')
              )
            }
          },
          {
            title: 'last_updated_by',
            key: 'last_updated_by',
          },
          {
            title: 'last_updated_time',
            key: 'last_updated_time',
            render: (h,params)=>{
              return h('div',
                formatDate(new Date(params.row.last_updated_time),'yyyy-MM-dd hh:mm')
              )
            }
          }
        ],
      }
    },
    methods:{
      refreshQuartzList:async function () {
        const result = await QuartzList(this.offset,this.current_page,this.search);
        if(result.status=="SUCCESS"){
          this.quartzs = result.quartzs;
          this.total = result.paginator.totalcount;
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
