<template>
  <div style="margin: 10px;">
    <h4 v-if="$route.query.work_name" style="text-align: center;margin-bottom: 10px;">当前流程为：{{$route.query.work_name}}</h4>

    <ISimpleLeftRightRow style="margin-bottom: 10px;margin-right: 10px;">
      <!-- left 插槽部分 -->
      <div slot="left">
        <Row type="flex" justify="start" class="code-row-bg">
          <Col span="6"><Button type="error" @click="addWorkStep('empty')" style="margin-right: 5px;">新建空节点</Button></Col>
          <Col span="6"><Button type="warning" @click="showRefactorModal">Refactor</Button></Col>
          <Col span="6"><Button type="info" @click="batchChangeIndent('left')">向左缩进</Button></Col>
          <Col span="6"><Button type="success" @click="batchChangeIndent('right')">向右缩进</Button></Col>

          <ISimpleConfirmModal ref="refactor_modal" modal-title="重构为子流程" :modal-width="500" @handleSubmit="refactor">
            <Input v-model.trim="refactor_worksub_name" placeholder="请输入重构的子流程名称"></Input>
          </ISimpleConfirmModal>
        </Row>
      </div>
      <div slot="right">
        <Row type="flex" justify="end" class="code-row-bg">
          <Col span="6"><WorkValidate /></Col>
          <Col span="6"><Button type="success" @click="renderSourceXml">View Source XML</Button></Col>
        </Row>
      </div>
    </ISimpleLeftRightRow>

    <WorkStepBaseInfo ref="workStepBaseInfo" @handleSuccess="refreshWorkStepList"/>
    <WorkStepParamInfo ref="workStepParamInfo" @handleSuccess="refreshWorkStepList"/>

    <Table border :columns="columns1" ref="selection" :data="worksteps" size="small"></Table>

    <!-- 相关流程清单 -->
    <RelativeWork id="relativeWork" ref="relativeWork"/>
  </div>
</template>

<script>
  import {WorkStepList} from "../../api"
  import {DeleteWorkStepByWorkStepId} from "../../api"
  import {ChangeWorkStepOrder} from "../../api"
  import {RefactorWorkStepInfo} from "../../api"
  import {BatchChangeIndent} from "../../api"
  import {AddWorkStep} from "../../api"
  import WorkStepParamInfo from "./WorkStepParamInfo"
  import ISimpleLeftRightRow from "../Common/layout/ISimpleLeftRightRow"
  import WorkStepBaseInfo from "./WorkStepBaseInfo"
  import RelativeWork from "./RelativeWork"
  import {oneOf} from "../../tools"
  import {checkContainsInString} from "../../tools"
  import WorkValidate from "./WorkValidate"
  import ISimpleConfirmModal from "../Common/modal/ISimpleConfirmModal"
  import {getRepeatStr} from "../../tools"
  import {GetRelativeWork} from "../../api"

  export default {
    name: "WorkStepList",
    components:{WorkStepParamInfo,ISimpleLeftRightRow,WorkStepBaseInfo,RelativeWork,WorkValidate,ISimpleConfirmModal},
    data(){
      return {
        showRelativeWorkFlag:false,
        refactor_worksub_name:'',
        default_work_step_types: this.GLOBAL.default_work_step_types,
        worksteps: [],
        columns1: [
          {
            type: 'selection',
            width: 60,
            align: 'center',
          },
          {
            title: '步骤编号',
            key: 'work_step_id',
            width: 100,
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
                      display: !oneOf(this.worksteps[params.index]['work_step_type'], ["work_start","work_end"]) ? undefined : 'none',
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
                      display: !oneOf(this.worksteps[params.index]['work_step_type'], ["work_start","work_end"]) ? undefined : 'none',
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
            width: 250,
            render: (h, params) => {
              return h('div', [
                h('span', {
                  style: {
                    // work_step_name 根据缩进级别进行缩进,不同级别使用不同颜色
                    display: !checkContainsInString(this.worksteps[params.index]['work_step_name'], "random_")  ? undefined : 'none',
                    color: ['green','blue','grey','red'][params.row.work_step_indent],
                  },
                }, getRepeatStr('\xa0\xa0\xa0\xa0\xa0', params.row.work_step_indent) + this.worksteps[params.index]['work_step_name']),
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
                h('Badge', {    // 延迟执行函数显示效果
                  props: {
                    status: "error",
                  },
                  style: {
                    marginLeft: '5px',
                    display: oneOf(this.worksteps[params.index]['is_defer'], ["true"])  ? undefined : 'none',
                  },
                }),
              ]);
            }
          },
          {
            title: 'work_step_desc',
            key: 'work_step_desc',
            width: 250,
          },
          {
            title: '操作',
            key: 'operate',
            width: 220,
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
                      this.deleteWorkStepByWorkStepId(this.worksteps[params.index]['work_id'], this.worksteps[params.index]['work_step_id']);
                    }
                  }
                }, '删除'),
                h('Button', {
                  props: {
                    type: 'warning',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px',
                    display: oneOf(this.worksteps[params.index]['work_step_type'], ["work_sub"]) ? undefined : 'none'
                  },
                  on: {
                    click: () => {
                      this.showWorkSubDetail(this.worksteps[params.index]);
                    }
                  }
                }, '详情'),
                h('Button', {
                  props: {
                    type: 'warning',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px',
                    display: oneOf(this.worksteps[params.index]['work_step_type'], ["work_start"]) ? undefined : 'none'
                  },
                  on: {
                    click: () => {
                      this.goAnchor('#relativeWork');
                    }
                  }
                }, '关联流程'),
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
        }else{
          this.$Message.error(result.errorMsg);
        }
      },
      renderSourceXml:function () {
        alert(11111);
      },
      addWorkStep:async function () {
        let selections = this.$refs.selection.getSelection();
        if(selections.length != 1){
          this.$Message.warning('选中行数不符合要求,请选择一行并在其之后进行添加!');
          return
        }
        const result = await AddWorkStep(this.$route.query.work_id, selections[0].work_step_id);
        if(result.status == "SUCCESS"){
          this.refreshWorkStepList();
        }else{
          this.$Message.error(result.errorMsg);
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
      showRefactorModal:function (){
        this.$refs.refactor_modal.showModal();
      },
      getSelectionArr:function(){
        let selectionArr = [];
        let selections = this.$refs.selection.getSelection();
        for(var i=0; i<selections.length; i++){
          selectionArr.push(selections[i].work_step_id);
        }
        return selectionArr;
      },
      refactor: async function () {
        let selections = this.getSelectionArr();
        const result = await RefactorWorkStepInfo(this.$route.query.work_id, this.refactor_worksub_name, JSON.stringify(selections));
        if(result.status == "SUCCESS"){
          this.refreshWorkStepList();
        }else{
          this.$Message.error(result.errorMsg);
        }
      },
      batchChangeIndent:async function (mod) {
        let selections = this.getSelectionArr();
        const result = await BatchChangeIndent(this.$route.query.work_id, mod, JSON.stringify(selections));
        if(result.status == "SUCCESS"){
          this.refreshWorkStepList();
        }else{
          this.$Message.error(result.errorMsg);
        }
      },
      showWorkSubDetail:async function (currentworkStep) {
        const result = await GetRelativeWork(currentworkStep['work_id']);
        const work_subs = result.subworks.filter(subWork => subWork.id === currentworkStep['work_sub_id']);
        if(work_subs.length > 0){
          this.$router.push({ path: '/iwork/workstepList',
            query: { work_id: work_subs[0].id, work_name: work_subs[0].work_name }});
        }
      },
      // 前往锚点方法
      goAnchor: function(selector) {
        var anchor = this.$el.querySelector(selector);
        document.documentElement.scrollTop = anchor.offsetTop;
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
