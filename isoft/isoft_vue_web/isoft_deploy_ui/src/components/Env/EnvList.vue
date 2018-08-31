<template>
<div style="margin: 10px;margin-top: 20px;">
  <EnvAdd/>
  <EnvTable/>
</div>
</template>

<script>

  import EnvTable from "../../components/Env/EnvTable.vue";
  import EnvAdd from "../../components/Env/EnvAdd.vue";

  export default {
    components: {EnvAdd,EnvTable},
    data () {
      return {

      }
    }
  }
</script>

<style scoped>
</style>
