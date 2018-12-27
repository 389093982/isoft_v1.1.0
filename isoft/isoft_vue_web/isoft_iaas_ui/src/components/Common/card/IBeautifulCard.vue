<template>
  <div class="box_bg">
    <div class="box_area">
      <slot name="header">
        <div class="title">{{title}}</div>
      </slot>
      <slot name="content"></slot>
    </div>
  </div>
</template>

<script>
  export default {
    name: "IBeautifulCard",
    props: {
      title: {
        type: String,
        default: '默认标题'
      },
    },
  }
</script>

<style scoped>
 .box_bg {
   background-color: #eee;
   padding: 6px;
 }
 .box_area {
   background-color: #ffffff;
 }
 .title {
   height: 50px;
   line-height: 50px;
   padding-left: 15px;
   font-size: 18px;
   border-bottom: 1px solid #eee;
   text-align: left;
 }
</style>
