<template>
  <div style="height: 500px;">
    <Button type="primary">Primary</Button>
    <Circle :percent="80" dashboard>
      <span class="demo-circle-inner" style="font-size:24px">80%</span>
    </Circle>
    <Button type="primary">Primary</Button>
  </div>
</template>

<script>
  export default {
    name: "SiteDashBoard"
  }
</script>

<style scoped>

</style>
