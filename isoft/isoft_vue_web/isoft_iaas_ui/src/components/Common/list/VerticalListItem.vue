<template>
  <dd><a href="javascript:;" @click="onclick"><slot></slot></a></dd>
</template>

<script>
  export default {
    name: "VerticalListItem",
    methods:{
      onclick:function () {
        this.$emit("onclick");
      }
    }
  }
</script>

<style scoped lang="stylus" rel="stylesheet/stylus">
  a
    font-size: 14px
    color #666
    text-decoration none
  a:hover
    color red
    text-decoration none
</style>
