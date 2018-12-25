<template>
  <a @click="onclick" style="display:inline-block;height: inherit;">
    <slot></slot>
  </a>
</template>

<script>
  export default {
    name: "BeautifulLink",
    methods:{
      onclick:function () {
        this.$emit('onclick');
      }
    }
  }
</script>

<style scoped>
  a:hover{
    border-bottom: 2px solid red;
  }
</style>
