<template>
  <div>
    <Divider>我的博客文章</Divider>
    <p v-for="myblog in myblogs" style="margin-bottom: 4px;font-size: 14px;color: #8a8a8a;line-height: 24px;">
      <router-link :to="{path:'/iblog/blog_detail',query:{blog_id:myblog.id}}">
        {{myblog.blog_title}}
      </router-link>
    </p>
  </div>
</template>

<script>
  export default {
    name: "MyBlogList",
    props:["myblogs"],
  }
</script>

<style scoped>

</style>
