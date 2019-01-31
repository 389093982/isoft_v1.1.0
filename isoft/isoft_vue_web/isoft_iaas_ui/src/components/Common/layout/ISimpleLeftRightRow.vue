<template>
  <Row>
    <Col span="12">
      <slot name="left">
        <!-- 设置高度使 div 占用空间-->
        <div style="height: 1px;"/>
      </slot>
    </Col>
    <Col span="12">
      <slot name="right">
        <div style="height: 1px;"/>
      </slot>
    </Col>
  </Row>
</template>

<script>
  export default {
    name: "ISimpleLeftRightRow"
  }
</script>

<style scoped>

</style>
