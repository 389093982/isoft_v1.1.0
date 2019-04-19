<template>
  <div>
    <p v-for="resource in resources">
      {{resource.resource_dsn}}
      <span v-for="(tableNames,_resource_dsn) in tableNamesMap">
        <span v-if="_resource_dsn == resource.resource_dsn">
          <span v-for="tableName in tableNames">
            {{tableName}}
            <span v-for="(tableColumns, _resource_dsn_tableName) in tableColumnsMap">
              <span v-if="_resource_dsn_tableName == _resource_dsn + tableName">
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
        </span>
      </span>
    </p>
  </div>
</template>

<script>
  import {LoadQuickSqlMeta} from "../../../api"

  export default {
    name: "QuickSql",
    data(){
      return {
        resources:[],
        tableNamesMap:{},
        tableColumnsMap:{},
      }
    },
    methods:{
      loadQuickSqlMeta:async function () {
        const result = await LoadQuickSqlMeta();
        if(result.status == "SUCCESS"){
          this.resources = result.resources;
          this.tableNamesMap = result.tableNamesMap;
          this.tableColumnsMap = result.tableColumnsMap;
        }
      }
    },
    mounted:function () {
      this.loadQuickSqlMeta();
    }
  }
</script>

<style scoped>

</style>
