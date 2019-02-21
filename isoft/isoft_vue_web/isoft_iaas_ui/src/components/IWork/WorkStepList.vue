<template>
  <div style="margin: 10px;">
    <h4 v-if="$route.query.work_name" style="text-align: center;">当前流程为：{{$route.query.work_name}}</h4>

    <ISimpleLeftRightRow style="margin-bottom: 10px;">
      <!-- left 插槽部分 -->
      <Button slot="left" type="success" @click="addWorkStep">新增</Button>
      <div slot="right" style="text-align: right;">
        <Button type="success" @click="renderSourceXml">View Source XML</Button>
      </div>
    </ISimpleLeftRightRow>

    <WorkStepBaseInfo ref="workStepBaseInfo" @handleSuccess="refreshWorkStepList"/>
    <WorkStepParamInfo ref="workStepParamInfo" @handleSuccess="refreshWorkStepList"/>

    <Table :columns="columns1" :data="worksteps" size="small"></Table>
    <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
          @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>

    <!-- 相关流程清单 -->
    <RelativeWork ref="relativeWork"/>
  </div>
</template>

<script>
  import {WorkStepList} from "../../api"
  import {DeleteWorkStepById} from "../../api"
  import {ChangeWorkStepOrder} from "../../api"
  import {AddWorkStep} from "../../api"
  import WorkStepParamInfo from "./WorkStepParamInfo"
  import ISimpleLeftRightRow from "../Common/layout/ISimpleLeftRightRow"
  import WorkStepBaseInfo from "./WorkStepBaseInfo"
  import RelativeWork from "./RelativeWork"
  import {oneOf} from "../../tools"
  import ISimpleBadge from "../Common/tool/ISimpleBadge"
  import {checkEmpty} from "../../tools"

  export default {
    name: "WorkStepList",
    components:{WorkStepParamInfo,ISimpleLeftRightRow,WorkStepBaseInfo,RelativeWork,ISimpleBadge},
    data(){
      return {
        default_work_step_types: this.GLOBAL.default_work_step_types,
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
            width: 100,
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
            width: 200,
            render: (h, params) => {
              return h('div', [
                h('Icon', {
                  props: {
                    type: this.renderWorkStepTypeIcon(this.worksteps[params.index]['work_step_type']),
                    size: 25,
                  },
                  style: {
                    marginRight: '5px',
                  },
                }),
                h('span', this.worksteps[params.index]['work_step_type']),
              ]);
            }
          },
          {
            title: 'work_step_desc',
            key: 'work_step_desc',
          },
          {
            title: '操作',
            key: 'operate',
            render: (h, params) => {
              return h('div', [
                h('Button', {
                  props: {
                    type: 'error',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px',
                    display: !oneOf(this.worksteps[params.index]['work_step_type'], ["work_start","work_end"])  ? undefined : 'none'
                  },
                  on: {
                    click: () => {
                      this.$refs.workStepBaseInfo.showWorkStepBaseInfo(this.$route.query.work_id, this.worksteps[params.index]['work_step_id']);
                    }
                  }
                }, '编辑'),
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
                      if (this.worksteps[params.index]['work_step_type']){
                        this.$refs.workStepParamInfo.showWorkStepParamInfo(this.$route.query.work_id, this.worksteps[params.index]['work_step_id']);
                      }
                    }
                  }
                }, '参数'),
                h('Button', {
                  props: {
                    type: 'success',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px',
                    display: !oneOf(this.worksteps[params.index]['work_step_type'], ["work_start","work_end"])  ? undefined : 'none'
                  },
                  on: {
                    click: () => {
                      this.deleteWorkStepById(this.worksteps[params.index]['id']);
                    }
                  }
                }, '删除'),
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

          // 刷新关联流程信息
          this.$refs.relativeWork.refreshRelativeWork(this.$route.query.work_id);
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
      },
      renderWorkStepTypeIcon:function (workStepType) {
        for(var i=0; i<this.default_work_step_types.length; i++){
          let default_work_step_type = this.default_work_step_types[i];
          if(default_work_step_type.name == workStepType){
            return default_work_step_type.icon;
          }
        }
      },
    },
    mounted: function () {
      this.refreshWorkStepList();
    },
    watch:{
      // 监听路由是否变化
      '$route' (to, from) {
        this.refreshWorkStepList();
      }
    }
  }
</script>

<style scoped>

</style>
