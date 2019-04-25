<template>
  <div style="background: #FFFFFF;padding: 10px;box-shadow: 2px 2px 1px #888888;">
    <ChooseCourseType/>

    <Row class="bgcolor" style="margin-bottom: 20px;margin: 10px auto;padding: 3px;">
      <Col span="18">
        <!-- 热门推荐区域 -->
        <HotRecommend/>
      </Col>
      <Col span="6">
        <ToolBox/>
        <NewsNotice/>
        <HotIssue/>
      </Col>
    </Row>

    <!-- 热门编程语言 -->
    <HotLanguage/>
    <!-- 热门开源项目 -->
    <HotProject/>
    <SiteGuide/>
    <HotCommunity/>
    <HotUser/>
    <IBeautifulTabLink/>
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
  import HotLanguage from "./Hot/HotLanguage"
  import HotUser from "../User/HotUser"
  import IBeautifulTabLink from "../Common/link/IBeautifulTabLink"

  export default {
    name: "Index",
    components:{ChooseCourseType,HotRecommend,NewsNotice,HotIssue,SiteGuide,HotCommunity,HotProject,ToolBox,HotLanguage,HotUser,IBeautifulTabLink},
  }
</script>

<style scoped>
  /* 引入公共样式库 */
  @import "../../../static/css/common.css";

</style>
