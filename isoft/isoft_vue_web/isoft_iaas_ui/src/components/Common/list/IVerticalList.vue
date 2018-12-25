<template>
  <dl :class="{'showLeftBorder':showLeftBorder,'showRightBorder':showRightBorder}">
    <dt style="font-size: 15px;font-weight: bold;height: 26px;line-height: 26px;">{{title}}</dt>
    <slot></slot>
  </dl>
</template>

<script>
  export default {
    name: "VerticalList",
    props:{
      title:{
        type: [String],
        default: ''
      },
      showLeftBorder: {
        type: [Boolean],
        default: false
      },
      showRightBorder: {
        type: [Boolean],
        default: false
      },
    }
  }
</script>

<style scoped>
  .showLeftBorder{
    border-left: solid 1px #ddd;
  }
  .showRightBorder{
    border-right: solid 1px #ddd;
  }
</style>
