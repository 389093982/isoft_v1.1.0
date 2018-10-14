<template>
  <div style="padding: 20px;background: #fff;">
    <!-- 我的博客分类 topN -->
    <MyCatalogList :mycatalogs="mycatalogs"/>
    <!-- 我的博客文章 topN -->
    <MyBlogList :myblogs="myblogs"/>
  </div>
</template>

<script>
  import MyCatalogList from "./MyCatalogList.vue"
  import MyBlogList from "./MyBlogList.vue"
  import {GetMyCatalogs} from "../../../api/index"
  import {GetMyBlogs} from "../../../api/index"

  export default {
    name: "ShowMyBlog",
    components:{MyCatalogList,MyBlogList},
    data(){
      return {
        // 我的所有文章分类
        mycatalogs:[],
        // 我的所有博客文章
        myblogs:[],
      }
    },
    methods:{
      loadMyCatalogs:async function () {
        const result = await GetMyCatalogs();
        if(result.status=="SUCCESS"){
          this.mycatalogs = result.catalogs;
        }
      },
      loadMyBlogs:async function () {
        const result = await GetMyBlogs();
        if(result.status=="SUCCESS"){
          this.myblogs = result.blogs;
        }
      }
    },
    mounted:function () {
      this.loadMyCatalogs();
      this.loadMyBlogs();
    }
  }
</script>

<style scoped>

</style>
