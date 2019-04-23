<template>
  <span style="margin-top: 10px;">
    选择数据表：
    <Select v-model="choose_table_name" style="width:400px">
      <Option v-for="tableName in tableNames" :value="tableName" :key="tableName">
        {{tableName}}
      </Option>
    </Select>

    <div v-if="choose_table_name == tableName" v-for="tableName in tableNames">
      <QuickTable :table-name="tableName" :table-columns="getTableColumns(tableName)"
        :table-sqls="getTableSqls(tableName)"/>
    </div>
  </span>
</template>

<script>
  import {LoadQuickSqlMeta} from "../../../api"
  import QuickTable from "./QuickTable"

  export default {
    name: "QuickResource",
    components:{QuickTable},
    props:{
      resource:{
        type:Object,
        default:{},
      }
    },
    data(){
      return {
        choose_table_name:'',
        tableNames:[],
        tableColumnsMap:{},
        tableSqlMap:{},
      }
    },
    methods:{
      loadQuickSqlMeta:async function () {
        const result = await LoadQuickSqlMeta(this.resource.id);
        if(result.status == "SUCCESS"){
          this.tableNames = result.tableNames;
          this.tableColumnsMap = result.tableColumnsMap;
          this.tableSqlMap = result.tableSqlMap;
        }
      },
      getTableColumns:function (tableName) {
        return this.tableColumnsMap[tableName];
      },
      getTableSqls:function (tableName) {
        return this.tableSqlMap[tableName];
      },
    },
    mounted(){
      this.loadQuickSqlMeta();
    }
  }
</script>

<style scoped>

</style>
