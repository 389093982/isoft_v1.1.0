<template>
<div style="margin: 10px;margin-top: 20px;">
  <ConfigAdd/>
  <ConfigTable/>
</div>
</template>

<script>

  import ConfigAdd from "../../components/Config/ConfigAdd.vue";
  import ConfigTable from '../../components/Config/ConfigTable.vue';

  export default {
    components: {ConfigAdd, ConfigTable},
    data () {
      return {

      }
    }
  }
</script>

<style scoped>
</style>
