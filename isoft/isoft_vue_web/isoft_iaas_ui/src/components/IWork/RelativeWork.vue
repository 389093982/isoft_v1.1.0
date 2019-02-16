<template>
  <div style="margin-top: 20px;">
    <Row :gutter="6">
      <Col span="12">
        <RelativeWorkList ref="parentRelativeWork" title="父级流程清单"/>
      </Col>
      <Col span="12">
        <RelativeWorkList ref="subRelativeWork" title="子级流程清单"/>
      </Col>
    </Row>
  </div>
</template>

<script>
  import {GetRelativeWork} from "../../api"
  import RelativeWorkList from "./RelativeWorkList"

  export default {
    name: "RelativeWork",
    components:{RelativeWorkList},
    methods:{
      refreshRelativeWork:async function (work_id) {
        const result = await GetRelativeWork(work_id);
        if(result.status == "SUCCESS"){
          this.$refs.parentRelativeWork.refreshWorkList(result.parentWorks);
          this.$refs.subRelativeWork.refreshWorkList(result.subwork);
        }
      }
    }
  }
</script>

<style scoped>

</style>
