<template>
  <div>
    <Table border :columns="columns1" :data="migrates" size="small"></Table>
    <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
          @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>
  </div>
</template>

<script>
  import {FilterPageMigrate} from "../../../api"

  export default {
    name: "MigrateList",
    data(){
      return {
        migrates:[],
        // 当前页
        current_page:1,
        // 总页数
        total:1,
        // 每页记录数
        offset:10,
        columns1: [
          {
            title: 'table_name',
            key: 'table_name',
            width: 250,
          },
          {
            title: 'table_columns',
            key: 'table_columns',
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
                      this.$router.push({ path: '/iwork/editMigrate', query: { id: this.migrates[params.index]['id'] }});
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
      refreshMigrateList: async function(){
        const result = await FilterPageMigrate(this.offset, this.current_work_step_id);
        this.migrates = result.migrates;
      },
      handleChange(page){
        this.current_page = page;
        this.refreshMigrateList();
      },
      handlePageSizeChange(pageSize){
        this.offset = pageSize;
        this.refreshMigrateList();
      },
    },
    mounted(){
      this.refreshMigrateList();
    }
  }
</script>

<style scoped>

</style>
