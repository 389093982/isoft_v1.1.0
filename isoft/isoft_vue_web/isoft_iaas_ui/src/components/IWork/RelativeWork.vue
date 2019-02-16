<template>
  <div style="margin-top: 20px;">
    <Row>
      <Col span="12">父级流程清单</Col>
      <Col span="12">子级流程清单</Col>
    </Row>
  </div>
</template>

<script>
  import {GetRelativeWork} from "../../api"

  export default {
    name: "RelativeWork",
    data(){
      return {
        
      }
    },
    methods:{
      refreshRelativeWork:async function (work_id) {
        const result = await GetRelativeWork(work_id);
        if(result.status == "SUCCESS"){
          // alert(result);
        }
      }
    }
  }
</script>

<style scoped>

</style>
