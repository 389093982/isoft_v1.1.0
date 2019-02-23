<template>
  <span>
    <Table :columns="columns1" :data="runLogRecords" size="small"></Table>
    <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
          @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>
  </span>
</template>

<script>
  import {FilterPageLogRecord} from "../../api"
  import {formatDate} from "../../tools"

  export default {
    name: "RunLogList",
    data(){
      return {
        // 当前页
        current_page:1,
        // 总页数
        total:1,
        // 每页记录数
        offset:10,
        runLogRecords: [],
        columns1: [
          {
            title: 'tracking_id',
            key: 'tracking_id',
            width: 300,
          },
          {
            title: 'work_name',
            key: 'work_name',
          },
          {
            title: 'last_updated_time',
            key: 'last_updated_time',
            render: (h,params)=>{
              return h('div',
                formatDate(new Date(params.row.last_updated_time),'yyyy-MM-dd hh:mm')
              )
            }
          },
          {
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
                    marginRight: '5px',
                  },
                  on: {
                    click: () => {
                      this.$router.push({ path: '/iwork/runLogDetail', query: { tracking_id: this.runLogRecords[params.index]['tracking_id'] }});
                    }
                  }
                }, '查看'),
              ]);
            }
          }
        ],
      }
    },
    methods:{
      refreshRunLogRecordList:async function () {
        const result = await FilterPageLogRecord(this.$route.query.work_id,this.offset,this.current_page);
        if(result.status=="SUCCESS"){
          this.runLogRecords = result.runLogRecords;
          this.total = result.paginator.totalcount;
        }
      },
      handleChange(page){
        this.current_page = page;
        this.refreshRunLogRecordList();
      },
      handlePageSizeChange(pageSize){
        this.offset = pageSize;
        this.refreshRunLogRecordList();
      },
    },
    mounted: function () {
      this.refreshRunLogRecordList();
    },
  }
</script>

<style scoped>

</style>
