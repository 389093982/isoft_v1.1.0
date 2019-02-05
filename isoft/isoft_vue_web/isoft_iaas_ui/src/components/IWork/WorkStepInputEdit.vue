<template>
  <span>
    <Row v-for="item in paramDefinitionItems">
      <Col span="4">{{item.ParamName}}</Col>
      <Col span="16"><Input v-model="value3" size="small" placeholder="small size" /></Col>
      <Col span="4">edit</Col>
    </Row>
  </span>
</template>

<script>
  export default {
    name: "WorkStepInputEdit",
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
