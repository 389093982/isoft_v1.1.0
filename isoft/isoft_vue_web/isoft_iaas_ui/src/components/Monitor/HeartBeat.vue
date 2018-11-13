<template>
  <div>
    <AddHeartBeat @refreshHeartBeatTable="refreshHeartBeatTable"/>

    <Table :columns="columns1" :data="heartBeats" size="small" height="450"></Table>
    <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
          @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>
  </div>
</template>

<script>
  import {FilterPageHeartBeat} from "../../api"
  import AddHeartBeat from "./AddHeartBeat.vue"

  export default {
    name: "HeartBeat",
    data(){
      return {
        heartBeats:[],
        // 当前页
        current_page:1,
        // 总页数
        total:1,
        // 每页记录数
        offset:10,
        columns1: [
          {
            title: 'id',
            key: 'id'
          },
          {
            title: 'addr',
            key: 'addr'
          },
          {
            title: 'status_code',
            key: 'status_code'
          },
          {
            title: 'last_updated_by',
            key: 'last_updated_by'
          },
          {
            title: 'last_updated_time',
            key: 'last_updated_time'
          },
        ]
      }
    },
    components:{AddHeartBeat},
    methods:{
      refreshHeartBeatTable:async function () {
        const result = await FilterPageHeartBeat(this.offset, this.current_page);
        if(result.status == "SUCCESS"){
          this.total = result.paginator.totalcount;
          this.heartBeats = result.heartBeats;
        }
      },
      handleChange(page){
        this.current_page = page;
        this.refreshHeartBeatTable();
      },
      handlePageSizeChange(pageSize){
        this.offset = pageSize;
        this.refreshHeartBeatTable();
      },
    },
    mounted:function () {
      this.refreshHeartBeatTable();
    }
  }
</script>

<style scoped>

</style>
