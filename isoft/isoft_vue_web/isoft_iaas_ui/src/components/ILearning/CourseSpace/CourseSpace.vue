<template>
  <div style="background: #FFFFFF;padding: 10px;box-shadow: 2px 2px 1px #888888;">
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
  import MyCourseList from "./MyCourseList.vue"
  import NewCourse from "./NewCourse.vue"

  export default {
    name: "CourseSpace",
    components:{MyCourseList,NewCourse}
  }
</script>

<style scoped>

</style>
