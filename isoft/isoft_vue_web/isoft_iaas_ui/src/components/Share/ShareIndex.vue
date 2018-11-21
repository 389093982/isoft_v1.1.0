<template>
  <div>
    <router-view/>
  </div>
</template>

<script>
  import HotShareItem from "./HotShareItem"

  export default {
    name: "ShareIndex",
    components:{HotShareItem},
  }
</script>

<style scoped>

</style>
