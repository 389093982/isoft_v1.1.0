<template>
  <span>
    <ISimpleBtnTriggerModal ref="triggerModal" btn-text="实体类管理" modal-title="新增/编辑实体类" :modal-width="800">
      <Table :columns="columns1" :data="entities" size="small"></Table>
      <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
            @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>
    </ISimpleBtnTriggerModal>
  </span>
</template>

<script>
  import ISimpleBtnTriggerModal from "../Common/modal/ISimpleBtnTriggerModal"
  import {FilterPageEntity} from "../../api"

  export default {
    name: "EntityList",
    components:{ISimpleBtnTriggerModal},
    data(){
      return {
        // 当前页
        current_page:1,
        // 总页数
        total:1,
        // 每页记录数
        offset:10,
        entities: [],
        columns1: [
          {
            title: 'entity_name',
            key: 'entity_name',
          },
          {
            title: 'entity_field_str',
            key: 'entity_field_str',
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
                    }
                  }
                }, '操作'),
              ]);
            }
          }
        ],
      }
    },
    methods:{
      refreshEntityList: async function () {
        const result = await FilterPageEntity(this.offset, this.current_page);
        if(result.status == "SUCCESS"){
          this.entities = result.entities;
        }
      },
      handleChange(page){
        this.current_page = page;
        this.refreshEntityList();
      },
      handlePageSizeChange(pageSize){
        this.offset = pageSize;
        this.refreshEntityList();
      },
    },
    mounted(){
      this.refreshEntityList();
    }
  }
</script>

<style scoped>

</style>
