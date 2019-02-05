<template>
  <span>
    <Row v-for="item in paramDefinitionItems">
      <Row>
        <Col span="12">{{item.ParamName}}</Col>
        <Col span="12" style="text-align: right;">
          <WorkStepInputEditDialog/>
        </Col>
      </Row>
      <Row>
        <Input v-model="value3" size="small" placeholder="small size" />
      </Row>
    </Row>
  </span>
</template>

<script>
  import WorkStepInputEditDialog from "./WorkStepInputEditDialog"

  export default {
    name: "WorkStepInputEdit",
    components:{WorkStepInputEditDialog},
    props:{
      paramDefinitionItems:{
        type: Array,
        default: () => [],
      }
    }
  }
</script>

<style scoped>

</style>
