<template>
  <div>
    迁移类型：
    <a href="javascript:;" @click="createTableMigrate">创建/变更表迁移</a>
    <ISimpleConfirmModal ref="createTable" modal-title="创建/变更表迁移" :modal-width="800" :footer-hide="true">
      <Form ref="formValidate" :model="formValidate" :rules="ruleValidate" :label-width="140">
        <FormItem label="tableName" prop="tableName">
          <Input v-model.trim="formValidate.tableName" placeholder="请输入 tableName"
                 :readonly="this.tableName != '' && this.tableName != null && this.tableName != undefined"></Input>
        </FormItem>
        <FormItem label="tableColumns" prop="tableColumns">
          <Input v-model.trim="formValidate.tableColumns" placeholder="请输入 tableColumns"></Input>
        </FormItem>
        <FormItem>
          <Button type="success" @click="handleFormSubmit('formValidate')" style="margin-right: 6px">Submit</Button>
        </FormItem>
      </Form>
    </ISimpleConfirmModal>

    <Table border :columns="columns1" :data="tableColumns" size="small"></Table>
    <Button type="success" size="small" @click="handleMigrateSubmit">Submit</Button>
  </div>
</template>

<script>
  import ISimpleConfirmModal from "../../Common/modal/ISimpleConfirmModal"
  import {SubmitMigrate} from "../../../api"
  import {GetMigrateInfo} from "../../../api"
  import {validatePatternForString} from "../../../tools"

  export default {
    name: "MigrateList",
    components:{ISimpleConfirmModal},
    data(){
      return {
        tableName:'',
        tableColumns:[],
        columns1: [
          {
            title: 'column_name',
            key: 'column_name',
          },
          {
            title: 'column_type',
            key: 'column_type',
          },
          {
            title: 'primary_key',
            key: 'primary_key',
            render: (h, params) => {
              return h('div', [
                h('span', params.row.primary_key),
                h('Icon', {
                  props: {
                    type: 'md-create',
                    size: 15,
                  },
                  style: {
                    marginLeft: '30px',
                  },
                  on: {
                    click: () => {
                      let primary_key = this.tableColumns[params.index]["primary_key"];
                      this.tableColumns[params.index]["primary_key"] = primary_key == "Y" ? "N" : "Y";
                    }
                  }
                }),
              ]);
            }
          },
          {
            title: 'auto_increment',
            key: 'auto_increment',
            render: (h, params) => {
              return h('div', [
                h('span', params.row.auto_increment),
                h('Icon', {
                  props: {
                    type: 'md-create',
                    size: 15,
                  },
                  style: {
                    marginLeft: '30px',
                  },
                  on: {
                    click: () => {
                      let auto_increment = this.tableColumns[params.index]["auto_increment"];
                      this.tableColumns[params.index]["auto_increment"] = auto_increment == "Y" ? "N" : "Y";
                    }
                  }
                }),
              ]);
            }
          },
          {
            title: 'unique',
            key: 'unique',
            render: (h, params) => {
              return h('div', [
                h('span', params.row.unique),
                h('Icon', {
                  props: {
                    type: 'md-create',
                    size: 15,
                  },
                  style: {
                    marginLeft: '30px',
                  },
                  on: {
                    click: () => {
                      let unique = this.tableColumns[params.index]["unique"];
                      this.tableColumns[params.index]["unique"] = unique == "Y" ? "N" : "Y";
                    }
                  }
                }),
              ]);
            }
          },
          {
            title: 'comment',
            key: 'comment',
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
                    }
                  }
                }, '删除'),
              ]);
            }
          }
        ],
        formValidate: {
          tableName: '',
          tableColumns: '',
        },
        ruleValidate: {
          tableName: [
            { required: true, message: 'tableName 不能为空!', trigger: 'blur' }
          ],
          tableColumns: [
            { required: true, message: 'tableColumns 不能为空!', trigger: 'blur' }
          ],
        },
      }
    },
    methods:{
      createTableMigrate(){
        this.$refs.createTable.showModal();
      },
      handleFormSubmit(name){
        this.$refs[name].validate((valid) => {
          if (valid) {
           this.tableName = this.formValidate.tableName;
           if(!validatePatternForString(/(^_([a-zA-Z0-9,]_?)*$)|(^[a-zA-Z,](_?[a-zA-Z0-9,])*_?$)/,this.formValidate.tableColumns)){
              alert("格式不正确!");
              return;
           }
           this.formValidate.tableColumns.split(",").forEach(columnStr => {
             let has = false;
             this.tableColumns.forEach(column => {
               // 已经包含
               if(column.column_name == columnStr){
                 has = true;
               }
             });
             if(!has && columnStr != ""){
               this.tableColumns.push({"column_name": columnStr, "column_type": "varchar(200)",
                 "primary_key":"N", "auto_increment":"N", "unique":"N", "comment":""});
             }
           });
           this.$refs.createTable.hideModal();
          }
        });
      },
       handleMigrateSubmit: async function () {
        const result = await SubmitMigrate(this.tableName, JSON.stringify(this.tableColumns));
        if(result.status == "SUCCESS"){
          this.$router.push({ path: '/iwork/migrateList'});
        }else{
          this.$Message.error(result.errorMsg);
        }
      },
      refreshMigrateInfo: async function(id){
        const result = await GetMigrateInfo(id);
        if(result.status=="SUCCESS"){
          this.tableName = result.migrate.table_name;
          this.formValidate.tableName = this.tableName;
          this.tableColumns = JSON.parse(result.migrate.table_info).table_columns;
        }
      }
    },
    mounted(){
      if(this.$route.query.id != undefined && this.$route.query.id != null){
        this.refreshMigrateInfo(this.$route.query.id);
      }
    }
  }
</script>

<style scoped>

</style>
