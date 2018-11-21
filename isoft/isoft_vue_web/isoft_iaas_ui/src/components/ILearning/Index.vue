<template>
  <div style="background: #FFFFFF;padding: 10px;box-shadow: 2px 2px 1px #888888;">
    <ChooseCourseType/>
    <Row style="border: 1px solid #ccc;border-radius: 4px;margin-bottom: 20px;height: 800px;margin: 20px auto;">
      <Col span="18" style="border-right: 1px solid #ccc;height: 800px;display: block;">
        <!-- 热门推荐区域 -->
        <HotRecommend/>
      </Col>
      <Col span="6">
        <NewsNotice/>
        <HotIssue/>
      </Col>
    </Row>
    <SiteGuide/>
    <HotCommunity/>
  </div>
</template>

<script>
  import ChooseCourseType from "./CourseType/ChooseCourseType"
  import HotRecommend from "./Course/HotRecommend"
  import NewsNotice from "./Course/NewsNotice"
  import HotIssue from "./Course/HotIssue"
  import SiteGuide from "./Site/SiteGuide"
  import HotCommunity from "./Community/HotCommunity"

  export default {
    name: "Index",
    components:{ChooseCourseType,HotRecommend,NewsNotice,HotIssue,SiteGuide,HotCommunity},
  }
</script>

<style scoped>

</style>
