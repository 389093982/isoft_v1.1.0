<template>
  <span>
    <Row>
      <Col span="18">
        <router-view></router-view>
      </Col>
      <Col span="6" style="padding: 10px;">
        <!-- 前往热门博客 -->
        <router-link to="/iblog/blog_list">热门博客</router-link>
        <ShowMyBlog/>
      </Col>
    </Row>
  </span>
</template>

<script>
  import ShowMyBlog from "./MyBlog/ShowMyBlog"

  export default {
    name: "IBlog",
    components:{ShowMyBlog},
  }
</script>

<style scoped>

</style>
