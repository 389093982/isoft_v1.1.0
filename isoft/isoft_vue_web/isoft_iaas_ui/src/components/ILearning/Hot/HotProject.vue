<template>
  <div style="margin: 20px;margin-top: 30px;">
    <div style="border-bottom: 2px solid #d9d9d9;">
      <h6 class="hot_project_dd" title="热门开源项目">热门开源项目</h6>
    </div>
    <Row :gutter="50">
      <Col span="8" style="margin-top: 12px;" v-for="hot_project in hot_projects">
        <span style="font-size: 14px;">{{hot_project.link_name}}</span>
        <BeautifulButtonLink msg="点击了解详情" floatstyle="right" :hrefaddr="hot_project.link_addr"/>
      </Col>
    </Row>
  </div>
</template>

<script>
  import BeautifulButtonLink from "../../Common/link/BeautifulButtonLink"
  import {QueryRandomCommonLink} from "../../../api"

  export default {
    name: "HotProject",
    components:{BeautifulButtonLink},
    data(){
      return {
        hot_projects:[],
      }
    },
    methods:{
      async refreshRandomHotProject (){
        const data = await QueryRandomCommonLink("hot_project");
        if(data.status == "SUCCESS"){
          this.hot_projects = data.common_links;
        }
      }
    },
    mounted:function () {
      this.refreshRandomHotProject();
      setInterval(this.refreshRandomHotProject, 30000);
    }
  }
</script>

<style scoped>
  .hot_project_dd{
    width: 200px;
    height: 35px;
    font-size: 18px;
    line-height: 35px;
    text-align: center;
    background: #3b80db;
    color: #fff;
    position: relative;
    font-weight: normal;
    margin: 0;
    padding: 0;
    font-family: "微软雅黑";
  }
</style>
