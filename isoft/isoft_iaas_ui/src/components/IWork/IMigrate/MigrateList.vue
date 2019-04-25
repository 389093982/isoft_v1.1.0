<template>
  <div>
    <Row style="margin-bottom: 10px;">
      <Col span="10">
        <Button type="success" size="small" @click="editMigrate(null)" style="margin-bottom: 6px">新建表</Button>
        <Input v-model.trim="filterTableName" placeholder="搜索 tableName" style="width: 200px;" @on-blur="refreshMigrateList"></Input>
      </Col>
      <Col span="14">
        <Select v-model="currentResourceName" style="width:300px">
          <Option v-for="resource in resources" :value="resource.resource_name">
            {{ resource.resource_name }} - {{ resource.resource_dsn }}
          </Option>
        </Select>
        <Button type="success" size="small" @click="executeMigrate(false)" style="margin-bottom: 6px">执行迁移</Button>
        <Button type="success" size="small" @click="executeMigrate(true)" style="margin-bottom: 6px">清理DB并执行迁移</Button>
      </Col>
    </Row>
    <p style="color: red;margin-bottom: 10px;margin-top: 10px;">{{errorMsg}}</p>

    <Table border :columns="columns1" :data="migrates" size="small"></Table>
    <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
          @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>
  </div>
</template>

<script>
  import {FilterPageMigrate} from "../../../api"
  import {ExecuteMigrate} from "../../../api"

  export default {
    name: "MigrateList",
    data(){
      return {
        filterTableName: '',
        errorMsg:'',
        currentResourceName:'',
        resources:[],
        migrates:[],
        // 当前页
        current_page:1,
        // 总页数
        total:1,
        // 每页记录数
        offset:10,
        columns1: [
          {
            title: 'id',
            key: 'id',
            width: 100,
          },
          {
            title: 'table_name',
            key: 'table_name',
            width: 120,
          },
          {
            title: 'table_info',
            key: 'table_info',
            width: 250,
          },
          {
            title: 'migrate_type',
            key: 'migrate_type',
            width: 120,
          },
          {
            title: 'table_auto_sql',
            key: 'table_auto_sql',
            width: 250,
          },
          {
            title: 'table_migrate_sql',
            key: 'table_migrate_sql',
            width: 250,
          },
          {
            title: 'table_info_hash',
            key: 'table_info_hash',
            width: 200,
          },
          {
            title: 'validate_result',
            key: 'validate_result',
            width: 130,
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
                    type: 'success',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px',
                    display: this.migrates[params.index]['is_max_migrate_id'] == true  ? undefined : 'none',
                  },
                  on: {
                    click: () => {
                      this.editMigrate(this.migrates[params.index]['id'], "upgrade");
                    }
                  }
                }, '结构升级'),
                h('Button', {
                  props: {
                    type: 'info',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px',
                    display: this.migrates[params.index]['is_max_migrate_id'] == true ? undefined : 'none',
                  },
                  on: {
                    click: () => {
                      this.editMigrate(this.migrates[params.index]['id'], "dataupgrade");
                    }
                  }
                }, '数据升级'),
                h('Button', {
                  props: {
                    type: 'error',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px',
                    display: this.migrates[params.index]['validate_result'] == "FAILED" ? undefined : 'none',
                  },
                  on: {
                    click: () => {
                      if (this.migrates[params.index]['validate_result'] == "SUCCESS"){
                        this.$Message.error("已执行验证通过的语句不能被更正,请执行升级操作!");
                      }else{
                        if(this.migrates[params.index]['migrate_type'] == "CREATE"){
                          alert("当表存在时 CREATE 语句纠正可能不会重新执行奥!请先执行删除表操作!");
                        }
                        this.editMigrate(this.migrates[params.index]['id'], "update");
                      }
                    }
                  }
                }, '更正'),
              ]);
            }
          }
        ],
      }
    },
    methods:{
      refreshMigrateList: async function(){
        const result = await FilterPageMigrate(this.filterTableName, this.offset, this.current_page);
        this.migrates = result.migrates;
        this.resources = result.resources;
        this.total = result.paginator.totalcount;
      },
      handleChange(page){
        this.current_page = page;
        this.refreshMigrateList();
      },
      handlePageSizeChange(pageSize){
        this.offset = pageSize;
        this.refreshMigrateList();
      },
      editMigrate:function (id, operateType) {
        if(id != undefined && id != null){
          this.$router.push({ path: '/iwork/editMigrate', query: {id: id, "operateType":operateType}});
        }else{
          this.$router.push({ path: '/iwork/editMigrate'});
        }
      },
      executeMigrate: async function (forceClean) {
        this.errorMsg = null;
        const result = await ExecuteMigrate(this.currentResourceName, forceClean);
        if(result.status == "SUCCESS"){
          this.$Message.success("SUCCESS");
          this.refreshMigrateList();
        }else{
          this.errorMsg = result.errorMsg;
        }
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
