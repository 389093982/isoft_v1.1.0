<template>
 <div>
   热门博客
   <p v-for="searchblog in searchblogs">{{searchblog.blog_title}}</p>
 </div>
</template>

<script>
  import {BlogList} from "../../api"

  export default {
    name: "BlogList",
    data(){
      return {
        searchblogs:[],
      }
    },
    mounted:async function () {
      const result = await BlogList(10,1);
      if(result.status=="SUCCESS"){
        this.searchblogs = result.blogs;
      }
    }
  }
</script>

<style scoped>

</style>
