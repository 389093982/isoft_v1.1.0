<template>
  <div style="margin-top: 20px;">
    <Row>
      <Col span="12">父级流程清单</Col>
      <Col span="12">子级流程清单</Col>
    </Row>
  </div>
</template>

<script>
  export default {
    name: "RelativeWork",
    data(){
      return {
        
      }
    },
    methods:{
      refreshRelativeWork:function (work_id) {
        alert(work_id);
      }
    }
  }
</script>

<style scoped>

</style>
