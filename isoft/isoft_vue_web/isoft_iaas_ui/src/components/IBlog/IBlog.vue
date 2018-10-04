<template>
  <span>
    <Row>
      <Col span="18">
        <router-view></router-view>
      </Col>
      <Col span="6" style="padding: 10px;">
        <!-- 前往热门博客 -->
        <router-link to="/iblog/blog_list">热门博客</router-link>
        <!-- 我的博客分类 topN -->
        <MyCatalogList/>
        <!-- 我的博客文章 topN -->
        <MyBlogList/>
      </Col>
    </Row>
  </span>
</template>

<script>
  import MyCatalogList from "./MyCatalogList.vue"
  import MyBlogList from "./MyBlogList.vue"
  export default {
    name: "BlogList",
    components:{MyCatalogList,MyBlogList},
  }
</script>

<style scoped>

</style>
