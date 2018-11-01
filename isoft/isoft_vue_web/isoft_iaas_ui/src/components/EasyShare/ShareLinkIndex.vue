<template>
  <div>
    <router-view/>
  </div>
</template>

<script>
  import HotShareLinkItem from "./HotShareLinkItem.vue"

  export default {
    name: "ShareLinkIndex",
    components:{HotShareLinkItem},
  }
</script>

<style scoped>

</style>
