<template>
  <span>
    <ISimpleBtnTriggerModal ref="triggerModal" btn-text="实体类管理" modal-title="新增/编辑实体类" :modal-width="800" modal-top="10px">
      <!-- 表单信息 -->
      <Form ref="formValidate" :model="formValidate" :rules="ruleValidate" :label-width="100">
        <FormItem label="entity_name" prop="entity_name">
          <Input v-model.trim="formValidate.entity_name" placeholder="请输入 entity_name"></Input>
        </FormItem>
        <FormItem label="entity_field_str" prop="entity_field_str">
          <Input v-model.trim="formValidate.entity_field_str" type="textarea" :rows="4" placeholder="请输入 entity_field_str"></Input>
        </FormItem>
        <FormItem>
          <Button type="success" @click="handleSubmit('formValidate')" style="margin-right: 6px">Submit</Button>
          <Button type="warning" @click="handleReset('formValidate')" style="margin-right: 6px">Reset</Button>
        </FormItem>
      </Form>

      <Table :columns="columns1" :data="entities" size="small"></Table>
      <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
            @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>
    </ISimpleBtnTriggerModal>
  </span>
</template>

<script>
  import ISimpleBtnTriggerModal from "../Common/modal/ISimpleBtnTriggerModal"
  import {FilterPageEntity} from "../../api"
  import {EditEntity} from "../../api"
  import {DeleteEntity} from "../../api"

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
                      this.deleteEntity(this.entities[params.index]['id']);
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
                      this.formValidate.entity_id = this.entities[params.index]['id'];
                      this.formValidate.entity_name = this.entities[params.index]['entity_name'];
                      this.formValidate.entity_field_str = this.entities[params.index]['entity_field_str'];
                    }
                  }
                }, '编辑'),
              ]);
            }
          }
        ],
        formValidate: {
          entity_id:-1,
          entity_name: '',
          entity_field_str: '',
        },
        ruleValidate: {
          entity_name: [
            { required: true, message: 'entity_name 不能为空!', trigger: 'blur' }
          ],
          entity_field_str: [
            { required: true, message: 'entity_field_str 不能为空!', trigger: 'blur' }
          ],
        },
      }
    },
    methods:{
      deleteEntity: async function(entity_id){
        const result = await DeleteEntity(entity_id);
        if(result.status == "SUCCESS"){
          this.refreshEntityList();
        }
      },
      refreshEntityList: async function () {
        const result = await FilterPageEntity(this.offset, this.current_page);
        if(result.status == "SUCCESS"){
          this.entities = result.entities;
          this.worksteps = result.worksteps;
          this.total = result.paginator.totalcount;
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
      handleSubmit (name) {
        this.$refs[name].validate(async (valid) => {
          if (valid) {
            const result = await EditEntity(this.formValidate.entity_id, this.formValidate.entity_name, this.formValidate.entity_field_str);
            if(result.status == "SUCCESS"){
              this.$Message.success('提交成功!');
              // 刷新表格
              this.refreshEntityList();
              // 表单重置,以取消缓存
              this.$refs[name].resetFields();
              this.formValidate.entity_id = -1;
            }else{
              this.$Message.error('提交失败!');
            }
          }
        })
      },
      handleReset (name) {
        this.$refs[name].resetFields();
      },
    },
    mounted(){
      this.refreshEntityList();
    }
  }
</script>

<style scoped>

</style>
