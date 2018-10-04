<template>
  <div>
    <Divider>我的博客文章</Divider>
    <p v-for="myblog in myblogs" style="margin-bottom: 4px;font-size: 14px;color: #8a8a8a;line-height: 24px;">
      {{myblog.blog_title}}
    </p>
  </div>
</template>

<script>
  import {GetMyBlogs} from "../../api"
  export default {
    name: "MyBlogList",
    data(){
      return {
        // 我的所有博客文章
        myblogs:[],
      }
    },
    mounted:async function () {
      const result = await GetMyBlogs();
      if(result.status=="SUCCESS"){
        this.myblogs = result.blogs;
      }
    }
  }
</script>

<style scoped>

</style>
