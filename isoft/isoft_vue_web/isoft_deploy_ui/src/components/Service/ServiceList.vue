<template>
<div style="margin: 10px;margin-top: 20px;">
  <ServiceAdd/>
  <ServiceTable/>
</div>
</template>

<script>
  import ServiceTable from '../../components/Service/ServiceTable.vue';
  import ServiceAdd from "../../components/Service/ServiceAdd.vue";
  export default {
    name: "ServiceList",
    components:{ServiceTable,ServiceAdd}
  }
</script>

<style scoped>

</style>
