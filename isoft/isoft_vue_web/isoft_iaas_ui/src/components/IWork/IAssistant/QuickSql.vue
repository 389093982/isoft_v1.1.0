<template>
  <Collapse simple>
    <Panel :name="resource.resource_name"  v-for="resource in resources">
      <span style="color: green;">{{resource.resource_name}} ~ {{resource.resource_dsn}}</span>
      <p slot="content">
        <QuickResource :resource="resource"/>
      </p>
    </Panel>
  </Collapse>
</template>

<script>
  import {GetAllResource} from "../../../api"
  import QuickResource from "./QuickResource"

  export default {
    name: "QuickSql",
    components:{QuickResource},
    data(){
      return {
        resources:[],
      }
    },
    methods:{
      refreshAllResource:async function () {
        const result = await GetAllResource("db");
        if(result.status == "SUCCESS"){
          this.resources = result.resources;
        }
      }
    },
    mounted:function () {
      this.refreshAllResource();
    }
  }
</script>

<style scoped>

</style>
