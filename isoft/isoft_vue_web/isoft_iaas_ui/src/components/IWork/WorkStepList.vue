<template>
  <div style="margin: 10px;">
    <h4 v-if="$route.query.work_name" style="text-align: center;">当前流程为：{{$route.query.work_name}}</h4>

    <ISimpleLeftRightRow style="margin-bottom: 10px;">
      <!-- left 插槽部分 -->
      <Button slot="left" type="success" @click="addWorkStep">新增</Button>
      <Button slot="right" type="success" @click="renderSourceXml" style="float: right;">View Source XML</Button>
    </ISimpleLeftRightRow>

    <WorkStepBaseInfoDialog ref="workStepBaseInfoDialog" @handleSuccess="refreshWorkStepList"/>
    <WorkStepEdit ref="workStepEdit" v-if="$route.query.work_id" :work-id="_workId" @handleSuccess="refreshWorkStepList"/>

    <Table :columns="columns1" :data="worksteps" size="small"></Table>
    <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
          @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>
  </div>
</template>

<script>
  import {formatDate} from "../../tools"
  import {WorkStepList} from "../../api"
  import {DeleteWorkStepById} from "../../api"
  import {ChangeWorkStepOrder} from "../../api"
  import {AddWorkStep} from "../../api"
  import WorkStepEdit from "./WorkStepEdit"
  import ISimpleLeftRightRow from "../Common/layout/ISimpleLeftRightRow"
  import WorkStepBaseInfoDialog from "./WorkStepBaseInfoDialog"

  export default {
    name: "WorkStepList",
    components:{WorkStepEdit,ISimpleLeftRightRow,WorkStepBaseInfoDialog},
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
            render: (h,params)=>{
              return h('div', [
                  h('span', params.row.work_step_id),
                  h('Icon', {
                    props: {
                      type: 'md-arrow-round-up',
                    },
                    style: {
                      marginLeft: '5px',
                    },
                    on: {
                      click: () => {
                        this.changeWorkStepOrder(this.worksteps[params.index]['work_step_id'], "up");
                      }
                    }
                  }),
                  h('Icon', {
                    props: {
                      type: 'md-arrow-round-down',
                    },
                    style: {
                      marginLeft: '5px',
                    },
                    on: {
                      click: () => {
                        this.changeWorkStepOrder(this.worksteps[params.index]['work_step_id'], "down");
                      }
                    }
                  }),
                ]
              )
            }
          },
          {
            title: 'work_step_name',
            key: 'work_step_name',
          },
          {
            title: 'work_step_type',
            key: 'work_step_type',
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
                      this.deleteWorkStepById(this.worksteps[params.index]['id']);
                    }
                  }
                }, '删除'),
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
                      this.$refs.workStepBaseInfoDialog.showWorkStepBaseInfoDialog(this._workId, this.worksteps[params.index]['work_step_id']);
                    }
                  }
                }, '编辑'),
              ]);
            }
          }
        ],
      }
    },
    methods:{
      refreshWorkStepList:async function () {
        const result = await WorkStepList(this.$route.query.work_id, this.offset,this.current_page);
        if(result.status=="SUCCESS"){
          this.worksteps = result.worksteps;
          this.total = result.paginator.totalcount;
          // 子组件同步刷新
          this.$refs.workStepEdit.refreshAllWorkStepsInfo();
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
      deleteWorkStepById:async function(id){
        const result = await DeleteWorkStepById(id);
        if(result.status=="SUCCESS"){
          this.refreshWorkStepList();
        }
      },
      changeWorkStepOrder:async function(work_step_id, type){
        const result = await ChangeWorkStepOrder(this.$route.query.work_id, work_step_id, type);
        if(result.status == "SUCCESS"){
          this.refreshWorkStepList();
        }
      },
      renderSourceXml:function () {
        alert(11111);
      },
      addWorkStep:async function () {
        const result = await AddWorkStep(this.$route.query.work_id);
        if(result.status == "SUCCESS"){
          this.refreshWorkStepList();
        }
      }
    },
    mounted: function () {
      this.refreshWorkStepList();
    },
    computed:{
      _workId:function () {
        return parseInt(this.$route.query.work_id);
      },
    }
  }
</script>

<style scoped>

</style>
