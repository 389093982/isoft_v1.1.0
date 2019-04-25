<template>
  <div>
    <Divider>我的博客分类</Divider>
    <p v-for="mycatalog in mycatalogs" style="margin-bottom: 4px;font-size: 14px;color: #8a8a8a;line-height: 24px;">
      {{mycatalog.catalog_name}}
    </p>
  </div>
</template>

<script>
  export default {
    name: "MyCatalogList",
    props:["mycatalogs"],
  }
</script>

<style scoped>

</style>
