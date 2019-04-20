<template>
  <div style="margin: 10px;">
    <span v-for="tableName in tableNames">
      <QuickTable :table-name="tableName" :table-columns="getTableColumns(tableName)"
        :table-sqls="getTableSqls(tableName)"/>
    </span>
  </div>
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
