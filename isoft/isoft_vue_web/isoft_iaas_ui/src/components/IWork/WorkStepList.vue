<template>
  <div style="margin: 10px;">
    <Table :columns="columns1" :data="worksteps" size="small"></Table>
    <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
          @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>
  </div>
</template>

<script>
  import {formatDate} from "../../tools"
  import {WorkStepList} from "../../api"

  export default {
    name: "WorkStepList",
    data(){
      return {
        // 当前页
        current_page:1,
        // 总页数
        total:1,
        // 每页记录数
        offset:10,
        worksteps: [],
        columns1: [
          {
            title: 'work_id',
            key: 'work_id',
          },
          {
            title: 'work_step_id',
            key: 'work_step_id',
          },
          {
            title: 'work_step_input',
            key: 'work_step_input',
          },
          {
            title: 'work_step_output',
            key: 'work_step_output',
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
      refreshWorkStepList:async function () {
        const result = await WorkStepList(this.offset,this.current_page);
        if(result.status=="SUCCESS"){
          this.worksteps = result.worksteps;
        }
      },
      handleChange(page){
        this.current_page = page;
        this.refreshWorkStepList();
      },
      handlePageSizeChange(pageSize){
        this.offset = pageSize;
        this.refreshWorkStepList();
      },
    },
    mounted: function () {
      this.refreshWorkStepList();
    },
  }
</script>

<style scoped>

</style>
