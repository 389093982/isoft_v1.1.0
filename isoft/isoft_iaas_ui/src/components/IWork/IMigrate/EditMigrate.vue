<template>
  <Scroll height="500">
    <div style="margin: 10px;">
      <Row>
        <Col span="12">
          <Button type="success" size="small" @click="createTableMigrate">创建/变更表迁移</Button>
        </Col>
        <Col span="12">
          <Button type="info" size="small" @click="buildInstanceSql('add')">使用 instance 值插入一条数据</Button>
          <Button type="error" size="small" @click="buildInstanceSql('delete')">使用 instance 值删除数据</Button>
        </Col>
      </Row>
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

      <Table border :columns="columns1" :data="tableColumns" size="small" style="margin-top: 10px;"></Table>
      <Input v-model.trim="table_migrate_sql" placeholder="当自动生成的 sql 不准确时请使用自定义 sql,请输入 table_migrate_sql"
             type="textarea" :rows="10" style="margin-bottom: 10px;margin-top: 10px;"></Input>
      <Button type="success" size="small" @click="handleMigrateSubmit">Submit</Button>
    </div>
  </Scroll>
</template>

<script>
  import ISimpleConfirmModal from "../../Common/modal/ISimpleConfirmModal"
  import {SubmitMigrate} from "../../../api"
  import {GetMigrateInfo} from "../../../api"
  import {BuildInstanceSql} from "../../../api"
  import {validatePatternForString} from "../../../tools"
  import {oneOf} from "../../../tools"
  import {checkEmpty} from "../../../tools"

  export default {
    name: "MigrateList",
    components:{ISimpleConfirmModal},
    data(){
      return {
        tableName:'',
        table_migrate_sql: '',
        tableColumns:[
          {"column_name": "id",
            "column_type": "int",
            "length": "",
            "default":"",
          "primary_key":"Y",
            "auto_increment":"Y",
            "unique":"Y",
            "comment":"主键id"
          }
        ],
        columns1: [
          {
            title: 'column_name',
            key: 'column_name',
            width: 120,
          },
          {
            title: 'column_type',
            key: 'column_type',
            width: 200,
            render: (h, params) => {
              return h('div', [
                h('Select',{
                  props: {
                    value: this.tableColumns[params.index]["column_type"],
                  },
                  on: {
                    'on-change': (event) => {
                      this.tableColumns[params.index]["column_type"] = event;
                    }
                  }
                }, this.GLOBAL.mysql_datatypes.map((item) => {
                  return h('Option',{
                    props:{
                      value: item,
                      label: item,
                    }
                  })
                })),
              ]);
            }
          },
          {
            title: 'length',
            key: 'length',
            width: 100,
            render: (h, params) => {
              return h('div', [
                h('Input',{
                  props: {
                    value: params.row.length,
                  },
                  on:{
                    'on-blur': (event) => {
                      this.tableColumns[params.index]["length"] = event.target.value;
                    }
                  }
                }),
              ]);
            }
          },
          {
            title: 'default',
            key: 'default',
            width: 100,
            render: (h, params) => {
              return h('div', [
                h('Input',{
                  props: {
                    value: params.row.default,
                  },
                  on:{
                    'on-blur': (event) => {
                      this.tableColumns[params.index]["default"] = event.target.value;
                    }
                  }
                }),
              ]);
            }
          },
          {
            title: 'pk',
            key: 'primary_key',
            width: 100,
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
            title: 'increment',
            key: 'auto_increment',
            width: 100,
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
                      if(oneOf(this.tableColumns[params.index]["column_type"], ["int"])){
                        let auto_increment = this.tableColumns[params.index]["auto_increment"];
                        this.tableColumns[params.index]["auto_increment"] = auto_increment == "Y" ? "N" : "Y";
                      }else{
                        this.$Message.error(this.tableColumns[params.index]["column_type"] + "类型不支持自增操作!");
                      }
                    }
                  }
                }),
              ]);
            }
          },
          {
            title: 'unique',
            key: 'unique',
            width: 100,
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
            width: 300,
            render: (h, params) => {
              return h('div', [
                h('Input',{
                  props: {
                    value: params.row.comment,
                  },
                  on:{
                    'on-blur': (event) => {
                      this.tableColumns[params.index]["comment"] = event.target.value;
                    }
                  }
                }),
              ]);
            }
          },
          {
            title: 'instance',
            key: 'instance',
            width: 100,
            render: (h, params) => {
              return h('div', [
                h('Input',{
                  props: {
                    value: params.row.instance,
                  },
                  on:{
                    'on-blur': (event) => {
                      this.tableColumns[params.index]["instance"] = event.target.value;
                    }
                  }
                }),
              ]);
            }
          },
          {
            title: '操作',
            key: 'operate',
            width: 150,
            fixed: 'right',
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
                      this.tableColumns.splice(params.index, 1);
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
             this.$Message.error("格式不正确!");
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
               this.tableColumns.push({"column_name": columnStr, "column_type": "varchar", "length": "200", "default":"",
                 "primary_key":"N", "auto_increment":"N", "unique":"N", "comment":""});
             }
           });
           this.$refs.createTable.hideModal();
          }
        });
      },
      handleMigrateSubmit: async function () {
        if(this.tableColumns.length == 0){
          this.$Message.error("You can't delete all columns with ALTER TABLE; use DROP TABLE instead!");
          return;
        }
        if(!checkEmpty(this.tableName)){
          const result = await SubmitMigrate(this.tableName, this.table_migrate_sql, JSON.stringify(this.tableColumns),
            this.$route.query.id, this.$route.query.operateType);
          if(result.status == "SUCCESS"){
            this.$router.push({ path: '/iwork/migrateList'});
          }else{
            this.$Message.error(result.errorMsg);
          }
        }else{
          this.$Message.error("empty tableName error!");
        }
      },
      refreshMigrateInfo: async function(id){
        const result = await GetMigrateInfo(id);
        if(result.status=="SUCCESS"){
          this.tableName = result.migrate.table_name;
          this.formValidate.tableName = this.tableName;
          this.tableColumns = JSON.parse(result.migrate.table_info).table_columns;
          if(this.$route.query.operateType != undefined && this.$route.query.operateType != null && this.$route.query.operateType == "update"){
            this.table_migrate_sql = result.migrate.table_migrate_sql;
          }

        }
      },
      buildInstanceSql: async function(operateType){
        const result = await BuildInstanceSql(this.tableName, JSON.stringify(this.tableColumns), this.$route.query.id, operateType);
        if(result.status == "SUCCESS"){
          this.table_migrate_sql += result.sql + "\n";
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
