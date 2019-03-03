<template>
  <div style="margin-top: 20px;">
    <Tabs :animated="false">
      <TabPane label="父级流程清单">
        <RelativeWorkList ref="parentRelativeWork" title="父级流程清单"/>
      </TabPane>
      <TabPane label="子级流程清单">
        <RelativeWorkList ref="subRelativeWork" title="子级流程清单"/>
      </TabPane>
    </Tabs>
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
          this.$refs.subRelativeWork.refreshWorkList(result.subworks);
        }
      }
    }
  }
</script>

<style scoped>

</style>
