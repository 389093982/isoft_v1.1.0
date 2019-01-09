<template>
  <div style="background: #FFFFFF;padding: 10px;box-shadow: 2px 2px 1px #888888;">
    <ChooseCourseType/>

    <Row class="bgcolor" style="margin-bottom: 20px;margin: 10px auto;padding: 3px;">
      <Col span="18">
        <!-- 热门推荐区域 -->
        <HotRecommend/>
      </Col>
      <Col span="6">
        <NewsNotice/>
        <HotIssue/>
        <ToolBox/>
      </Col>
    </Row>

    <IBeautifulListLink/>

    <!-- 热门开源项目 -->
    <HotProject/>
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
  import HotProject from "./Hot/HotProject"
  import ToolBox from "../CMS/ToolBox"
  import IBeautifulListLink from "../Common/link/IBeautifulListLink"

  export default {
    name: "Index",
    components:{ChooseCourseType,HotRecommend,NewsNotice,HotIssue,SiteGuide,HotCommunity,HotProject,ToolBox,IBeautifulListLink},
  }
</script>

<style scoped>
  /* 引入公共样式库 */
  @import "../../../static/css/common.css";

</style>
