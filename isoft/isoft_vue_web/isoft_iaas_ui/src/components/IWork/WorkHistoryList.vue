<template>
  <span>
    <Table :columns="columns1" :data="workHistories" size="small"></Table>
    <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
          @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>
    <Modal
      v-model="modal1"
      title="流程详情"
      :footer-hide="true"
      :mask-closable="false"
      :styles="{top: '20px'}"
      width="1000">
      <Scroll height="400">
        <span v-html="workHistory"></span>
      </Scroll>
    </Modal>
  </span>
</template>

<script>
  import {FilterPageWorkHistory} from "../../api"
  import {formatDate} from "../../tools"

  export default {
    name: "WorkHistoryList",
    data(){
      return {
        modal1: false,
        workHistory: "",
        // 当前页
        current_page:1,
        // 总页数
        total:1,
        // 每页记录数
        offset:10,
        workHistories: [],
        columns1: [
          {
            title: 'id',
            key: 'id',
          },
          {
            title: 'work_name',
            key: 'work_name',
            width: 300,
          },
          {
            title: 'work_desc',
            key: 'work_desc',
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
                      this.modal1 = true;
                      this.workHistory = this.shiftEnter(this.workHistories[params.index]['work_history']);
                    }
                  }
                }, '查看详情'),
              ]);
            }
          }
        ],
      }
    },
    methods:{
      // 转义成 html 符号
      shiftEnter:function(msg){
        return msg
          // 转义嵌套 json
          .replace(/\\\"/g, "&quot;")
          .replace(/\\\'/g, "&#39;")
          .replace(/\\n/g, "\\n\\t\\t\\t")        // 控制嵌套 json 缩进
          .replace(/\\t/g, "&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;")
          .replace(/\\n/g, "<br>")
          // 转义外层 json
          .replace(/\"/g, "&quot;")
          .replace(/\'/g, "&#39;")
          .replace(/\t/g, "&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;")
          .replace(/\n/g, "<br>");
      },
      refreshWorkHistoryList:async function(){
        const result = await FilterPageWorkHistory(this.offset,this.current_page);
        if(result.status=="SUCCESS"){
          this.workHistories = result.workHistories;
          this.total = result.paginator.totalcount;
        }
      },
      handleChange(page){
        this.current_page = page;
        this.refreshWorkHistoryList();
      },
      handlePageSizeChange(pageSize){
        this.offset = pageSize;
        this.refreshWorkHistoryList();
      },
    },
    mounted: function () {
      this.refreshWorkHistoryList();
    },
  }
</script>

<style scoped>

</style>
