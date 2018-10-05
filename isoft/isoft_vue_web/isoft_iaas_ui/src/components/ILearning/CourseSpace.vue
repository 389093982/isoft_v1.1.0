<template>
  <div style="background: #FFFFFF;padding: 10px;">
    <Tabs :animated="false">
      <TabPane label="我要开课">
        <NewCourse/>
      </TabPane>
      <TabPane label="我的课程">
        <MyCourseList/>
      </TabPane>
      <TabPane label="最近浏览">222</TabPane>
      <TabPane label="我的收藏">333</TabPane>
    </Tabs>
  </div>
</template>

<script>
  import MyCourseList from "./CourseSpace/MyCourseList.vue"
  import NewCourse from "./CourseSpace/NewCourse.vue"

  export default {
    name: "CourseSpace",
    components:{MyCourseList,NewCourse}
  }
</script>

<style scoped>

</style>
