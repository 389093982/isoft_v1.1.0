<template>
  <div style="margin: 10px;">
    <ISimpleLeftRightRow style="margin-bottom: 10px;margin-right: 10px;">
      <!-- left 插槽部分 -->
      <ResourceAdd slot="left" @handleSuccess="refreshResourceList"/>
      <!-- right 插槽部分 -->
      <ISimpleSearch slot="right" @handleSimpleSearch="handleSearch"/>
    </ISimpleLeftRightRow>

    <Table border :columns="columns1" :data="resources" size="small"></Table>
    <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
          @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>
  </div>
</template>

<script>
  import {formatDate} from "../../../tools/index"
  import {ResourceList} from "../../../api/index"
  import ISimpleLeftRightRow from "../../Common/layout/ISimpleLeftRightRow"
  import ISimpleSearch from "../../Common/search/ISimpleSearch"
  import ResourceAdd from "./ResourceAdd"

  export default {
    name: "ResourceList",
    components:{ISimpleLeftRightRow,ISimpleSearch,ResourceAdd},
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
        resources: [],
        columns1: [
          {
            title: 'resource_name',
            key: 'resource_name',
            width: 100,
          },
          {
            title: 'resource_type',
            key: 'resource_type',
            width: 100,
          },
          {
            title: 'resource_url',
            key: 'resource_url',
            width: 100,
          },
          {
            title: 'resource_dsn',
            key: 'resource_dsn',
            width: 300,
          },
          {
            title: 'resource_username',
            key: 'resource_username',
            width: 120,
          },
          {
            title: 'resource_password',
            key: 'resource_password',
            width: 120,
          },
          {
            title: 'last_updated_time',
            key: 'last_updated_time',
            width: 150,
            render: (h,params)=>{
              return h('div',
                formatDate(new Date(params.row.last_updated_time),'yyyy-MM-dd hh:mm')
              )
            }
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
                  },
                  on: {
                    click: () => {
                    }
                  }
                }, '连接测试'),
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
      refreshResourceList:async function () {
        const result = await ResourceList(this.offset,this.current_page,this.search);
        if(result.status=="SUCCESS"){
          this.resources = result.resources;
          this.total = result.paginator.totalcount;
        }
      },
      handleChange(page){
        this.current_page = page;
        this.refreshResourceList();
      },
      handlePageSizeChange(pageSize){
        this.offset = pageSize;
        this.refreshResourceList();
      },
      handleSearch(data){
        this.offset = 10;
        this.current_page = 1;
        this.search = data;
        this.refreshResourceList();
      }
    },
    mounted: function () {
      this.refreshResourceList();
    },
  }
</script>

<style scoped>

</style>
