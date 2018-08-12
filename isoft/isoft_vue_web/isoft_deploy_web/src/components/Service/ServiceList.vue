<template>
<div style="margin: 10px;margin-top: 20px;">
  <ServiceTable/>
</div>
</template>

<script>
  import ServiceTable from '../../components/Service/ServiceTable.vue'
  export default {
    name: "ServiceList",
    components:{ServiceTable}
  }
</script>

<style scoped>

</style>
