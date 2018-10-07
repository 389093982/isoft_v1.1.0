<template>
  <div style="background: #FFFFFF;padding: 10px;box-shadow: 2px 2px 1px #888888;">
    <router-link to="/ilearning/course_space">我的课程空间</router-link>
    <HotCourseType/>
    <TotalCourseType/>
    <Row style="border: 1px solid #ccc;border-radius: 4px;margin-bottom: 20px;height: 800px;margin: 20px auto;">
      <Col span="18" style="border-right: 1px solid #ccc;height: 800px;display: block;">
        <!-- 热门推荐区域 -->
        <HotRecommend/>
      </Col>
      <Col span="6">
        其它区域
      </Col>
    </Row>
    <p>111</p>
    <p>111</p>
    <p>111</p>
    <p>111</p>
    <p>111</p>
    <p>111</p>
    <p>111</p>
    <p>111</p>
    <p>111</p>
  </div>
</template>

<script>
  import HotCourseType from "./CourseType/HotCourseType.vue"
  import TotalCourseType from "./CourseType/TotalCourseType.vue"
  import HotRecommend from "./Course/HotRecommend.vue"


  export default {
    name: "Index",
    components:{HotCourseType,TotalCourseType,HotRecommend},
  }
</script>

<style scoped>

</style>
