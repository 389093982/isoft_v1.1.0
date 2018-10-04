<template>
  <div>
    <Divider>我的博客分类</Divider>
    <p v-for="mycatalog in mycatalogs" style="margin-bottom: 4px;font-size: 14px;color: #8a8a8a;line-height: 24px;">
      {{mycatalog.catalog_name}}
    </p>
  </div>
</template>

<script>
  import {GetMyCatalogs} from "../../api"

  export default {
    name: "MyCatalogList",
    components:{GetMyCatalogs},
    data(){
      return {
        // 我的所有文章分类
        mycatalogs:[],
      }
    },
    mounted:async function () {
      const result = await GetMyCatalogs();
      if(result.status=="SUCCESS"){
        this.mycatalogs = result.catalogs;
      }
    }
  }
</script>

<style scoped>

</style>
