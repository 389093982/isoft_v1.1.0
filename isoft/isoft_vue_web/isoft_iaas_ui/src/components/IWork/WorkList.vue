<template>
  <div style="margin: 10px;">
    <Row type="flex" justify="center">
      <Col span="3"><EntityList/></Col>
      <Col span="3"><GlobalVarList/></Col>
      <Col span="3"><IWorkDL/></Col>
      <Col span="3"><WorkValidate/></Col>
    </Row>

    <ISimpleLeftRightRow style="margin-bottom: 10px;margin-right: 10px;">
      <!-- left 插槽部分 -->
      <WorkEdit ref="workEdit" slot="left" @handleSuccess="refreshWorkList"/>
      <!-- right 插槽部分 -->
      <ISimpleSearch slot="right" @handleSimpleSearch="handleSearch"/>
    </ISimpleLeftRightRow>

    <Table border :columns="columns1" :data="works" size="small"></Table>
    <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
          @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>
  </div>
</template>

<script>
  import {WorkList} from "../../api"
  import {DeleteWorkById} from "../../api"
  import {RunWork} from "../../api"
  import {SaveHistory} from "../../api"
  import ISimpleLeftRightRow from "../Common/layout/ISimpleLeftRightRow"
  import ISimpleSearch from "../Common/search/ISimpleSearch"
  import WorkEdit from "./WorkEdit"
  import EntityList from "./EntityList"
  import GlobalVarList from "./GlobalVarList"
  import IWorkDL from "./IWorkDL"
  import WorkValidate from "./WorkValidate"

  export default {
    name: "WorkList",
    components:{ISimpleLeftRightRow,ISimpleSearch,WorkEdit,EntityList,GlobalVarList,IWorkDL,WorkValidate},
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
        works: [],
        columns1: [
          {
            title: 'work_name',
            key: 'work_name',
            width: 250,
          },
          {
            title: 'work_desc',
            key: 'work_desc',
            width: 250,
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
                  },
                  on: {
                    click: () => {
                      this.$refs.workEdit.triggerWorkEdit(this.works[params.index]);
                    }
                  }
                }, '编辑'),
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
                      this.deleteWorkById(this.works[params.index]['id']);
                    }
                  }
                }, '删除'),
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
                      this.saveHistory(this.works[params.index]['id']);
                    }
                  }
                }, '保存为历史版本'),
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
                      this.editWork(this.works[params.index]['id'], this.works[params.index]['work_name']);
                    }
                  }
                }, '编辑步骤'),
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
                      this.runWork(this.works[params.index]['id']);
                    }
                  }
                }, '运行流程'),
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
                      this.$router.push({ path: '/iwork/runLogList', query: { work_id: this.works[params.index]['id'] }});
                    }
                  }
                }, '运行日志'),
              ]);
            }
          }
        ],
      }
    },
    methods:{
      refreshWorkList:async function () {
        const result = await WorkList(this.offset,this.current_page,this.search);
        if(result.status=="SUCCESS"){
          this.works = result.works;
          this.total = result.paginator.totalcount;
        }
      },
      handleChange(page){
        this.current_page = page;
        this.refreshWorkList();
      },
      handlePageSizeChange(pageSize){
        this.offset = pageSize;
        this.refreshWorkList();
      },
      handleSearch(data){
        this.offset = 10;
        this.current_page = 1;
        this.search = data;
        this.refreshWorkList();
      },
      deleteWorkById:async function(id){
        const result = await DeleteWorkById(id);
        if(result.status=="SUCCESS"){
          this.refreshWorkList();
        }
      },
      editWork:function (id, work_name) {
        this.$router.push({ path: '/iwork/workstepList', query: { work_id: id, work_name: work_name }});
      },
      runWork:async function (work_id) {
        const result = await RunWork(work_id);
        if(result.status == "SUCCESS"){
          this.$Message.success("运行任务已触发!");
        }
      },
      saveHistory:async function (work_id) {
        const result = await SaveHistory(work_id);
        if(result.status == "SUCCESS"){
          this.$Message.success("保存成功!");
        }
      }
    },
    mounted: function () {
      this.refreshWorkList();
    },
  }
</script>

<style scoped>

</style>
