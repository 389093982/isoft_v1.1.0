<template>
  <div style="margin: 10px;">
    <span v-for="tableName in tableNames">
      <p style="color: red;">表名：{{tableName}}</p>
      <span v-for="(tableColumns, _tableName) in tableColumnsMap">
        <span v-if="_tableName == tableName">
          <CheckboxGroup>
            <ul>
              <li v-for="tableColumn in tableColumns" style="list-style: none;">
                <Checkbox :label="tableColumn"></Checkbox>
              </li>
            </ul>
          </CheckboxGroup>
        </span>
      </span>
    </span>
  </div>
</template>

<script>
  import {LoadQuickSqlMeta} from "../../../api"

  export default {
    name: "QuickResource",
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
      }
    },
    methods:{
      loadQuickSqlMeta:async function () {
        const result = await LoadQuickSqlMeta(this.resource.id);
        if(result.status == "SUCCESS"){
          this.tableNames = result.tableNames;
          this.tableColumnsMap = result.tableColumnsMap;
        }
      },
    },
    mounted(){
      this.loadQuickSqlMeta();
    }
  }
</script>

<style scoped>

</style>
