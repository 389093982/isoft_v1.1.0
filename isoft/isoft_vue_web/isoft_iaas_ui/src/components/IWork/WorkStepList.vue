<template>
  <div style="margin: 10px;">
    <h4 v-if="$route.query.work_name" style="text-align: center;">当前流程为：{{$route.query.work_name}}</h4>

    <ISimpleLeftRightRow style="margin-bottom: 10px;">
      <!-- left 插槽部分 -->
      <div slot="left">
        <Row type="flex" justify="start" class="code-row-bg">
          <Col span="5"><Button type="success" @click="addWorkStep('')" style="margin-right: 5px;">新建普通节点</Button></Col>
          <Col span="5"><Button type="error" @click="addWorkStep('empty')" style="margin-right: 5px;">新建空节点</Button></Col>
          <Col span="5"><Button type="warning" style="margin-right: 5px;">Refactor</Button></Col>
        </Row>
      </div>
      <div slot="right">
        <Row type="flex" justify="end" class="code-row-bg">
          <Col span="5"><WorkValidate /></Col>
          <Col span="5"><Button type="success" @click="renderSourceXml">View Source XML</Button></Col>
        </Row>
      </div>
    </ISimpleLeftRightRow>

    <WorkStepBaseInfo ref="workStepBaseInfo" @handleSuccess="refreshWorkStepList"/>
    <WorkStepParamInfo ref="workStepParamInfo" @handleSuccess="refreshWorkStepList"/>

    <Table border :columns="columns1" ref="selection" :data="worksteps" size="small"></Table>

    <!-- 相关流程清单 -->
    <RelativeWork ref="relativeWork"/>
  </div>
</template>

<script>
  import {WorkStepList} from "../../api"
  import {DeleteWorkStepByWorkStepId} from "../../api"
  import {ChangeWorkStepOrder} from "../../api"
  import {AddWorkStep} from "../../api"
  import WorkStepParamInfo from "./WorkStepParamInfo"
  import ISimpleLeftRightRow from "../Common/layout/ISimpleLeftRightRow"
  import WorkStepBaseInfo from "./WorkStepBaseInfo"
  import RelativeWork from "./RelativeWork"
  import {oneOf} from "../../tools"
  import WorkStepColorRender from "./WorkStepColorRender"
  import {checkEmpty} from "../../tools"
  import {EditWorkStepColorInfo} from "../../api"
  import {checkContainsInString} from "../../tools"
  import WorkValidate from "./WorkValidate"

  export default {
    name: "WorkStepList",
    components:{WorkStepParamInfo,ISimpleLeftRightRow,WorkStepBaseInfo,RelativeWork,WorkStepColorRender,WorkValidate},
    data(){
      return {
        default_work_step_types: this.GLOBAL.default_work_step_types,
        worksteps: [],
        columns1: [
          {
            type: 'selection',
            width: 60,
            align: 'center',
          },
          {
            title: 'work_id',
            key: 'work_id',
            width: 100,
          },
          {
            title: 'work_step_id',
            key: 'work_step_id',
            width: 120,
            render: (h,params)=>{
              return h('div', [
                  h('span', params.row.work_step_id),
                  h('Icon', {
                    props: {
                      type: 'md-arrow-round-up',
                      size: 15,
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
                      size: 15,
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
            width: 180,
            render: (h, params) => {
              return h('div', [
                h('span', {
                  style: {
                    display: !checkContainsInString(this.worksteps[params.index]['work_step_name'], "random_")  ? undefined : 'none'
                  },
                }, this.worksteps[params.index]['work_step_name']),
              ]);
            }
          },
          {
            title: 'work_step_type',
            key: 'work_step_type',
            width: 180,
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
            title: 'work_step_color',
            key: 'work_step_color',
            width: 180,
            render: (h, params) => {
              return h('div', this.renderWorkStep(h, this.worksteps[params.index]['work_step_color'],
                this.worksteps[params.index]['work_id'],this.worksteps[params.index]['work_step_id']));
            }
          },
          {
            title: 'work_step_desc',
            key: 'work_step_desc',
            width: 180,
          },
          {
            title: '操作',
            key: 'operate',
            width: 180,
            fixed: 'right',
            render: (h, params) => {
              return h('div', [
                h('Button', {
                  props: {
                    type: 'error',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px',
                    display: !oneOf(this.worksteps[params.index]['work_step_type'], ["work_start","work_end","empty"])  ? undefined : 'none'
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
                    display: !oneOf(this.worksteps[params.index]['work_step_type'], ["empty"])  ? undefined : 'none'
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
                      this.deleteWorkStepByWorkStepId(this.worksteps[params.index]['work_id'], this.worksteps[params.index]['work_step_id']);
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
        const result = await WorkStepList(this.$route.query.work_id);
        if(result.status=="SUCCESS"){
          this.worksteps = result.worksteps;
          // 刷新关联流程信息
          this.$refs.relativeWork.refreshRelativeWork(this.$route.query.work_id);
        }
      },
      deleteWorkStepByWorkStepId:async function(work_id, work_step_id){
        const result = await DeleteWorkStepByWorkStepId(work_id, work_step_id);
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
      addWorkStep:async function (default_work_step_type) {
        let selections = this.$refs.selection.getSelection();
        if(selections.length != 1){
          this.$Message.warning('选中行数不符合要求,请选择一行并在其之后进行添加!');
          return
        }
        const result = await AddWorkStep(this.$route.query.work_id, selections[0].work_step_id, default_work_step_type);
        if(result.status == "SUCCESS"){
          this.refreshWorkStepList();
        }else{
          this.$Message.error('新增失败!');
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
      renderWorkStep:function (h, work_step_color,work_id,work_step_id) {
        var colors = [];
        if(checkEmpty(work_step_color)){
          colors = ["#FFFFFF","#FFFFFF","#FFFFFF","#FFFFFF","#FFFFFF"];
        }else{
          colors = JSON.parse(work_step_color);
        }

        var _this = this;
        var result = [];
        for (var i=0; i<colors.length; i++){
          result.push(h(WorkStepColorRender,
            {
              props:{
                backgroundColorStyle:colors[i],
                marginRightStyle:5,
                backgroundColorIndex:i,
              },
              on:{
                submitColor:async function (index, color) {
                  colors[index] = color;
                  // 更新color信息
                  const result = await EditWorkStepColorInfo(work_id, work_step_id, JSON.stringify(colors));
                  if (result.status == "SUCCESS"){
                    _this.refreshWorkStepList();
                  }
                }
              }
            }
          ));
        }
        return result;
      }
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
