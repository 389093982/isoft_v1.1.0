<template>
  <div style="background: #FFFFFF;padding: 10px;box-shadow: 2px 2px 1px #888888;">
    <Row style="background: #FFFFFF;">
      <Col span="6" style="padding: 5px;background: #f8f8f9;">
        <Card title="我的课程空间" icon="ios-options" :padding="0" shadow>
          <CellGroup>
            <Cell title="我要开课" to="/ilearning/course_space/newCourse"/>
            <Cell title="我的课程" to="/ilearning/course_space/myCourseList"/>
            <Cell title="最近浏览" to="/ilearning/course_space/recentlyViewed"/>
            <Cell title="我的收藏" to="/components/button"/>
          </CellGroup>
        </Card>
      </Col>
      <Col span="18">
        <div style="padding: 5px;background: #f8f8f9;margin-left: 10px;">
          <div style="padding: 10px;background: #FFFFFF;min-height: 500px;">
            <router-view/>
          </div>
        </div>
      </Col>
    </Row>
  </div>
</template>

<script>
  export default {
    name: "CourseSpace",
  }
</script>

<style scoped>

</style>
