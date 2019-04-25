<template>
<span>
  <slot></slot>
</span>
</template>

<script>
  export default {
    name: "IBeautifulHoverShadow"
  }
</script>

<style scoped>
span{
  display: block;
  border: 1px solid #E5E5E5;
  background-color: #fff;
  margin-top: 2px;
  margin-bottom: 10px;
  padding: 13px 20px 15px;
}

span:hover {
  box-shadow: 0 3px 7px 0 rgba(0, 0, 0, .1);
  border-color: rgba(229, 229, 229, .73);
}
</style>
