<template>
  <div>
    <div>
      <h6 class="hot_project_dd" title="热门开源项目">热门开源项目</h6>
    </div>
    <Row :gutter="50">
      <Col span="8" style="margin-top: 20px;" v-for="hot_project in hot_projects">
        {{hot_project.link_name}} <BeautifulButtonLink msg="点击了解详情" floatstyle="right" :hrefaddr="hot_project.link_addr"/>
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
    width: 293px;
    height: 43px;
    font-size: 20px;
    line-height: 43px;
    text-align: center;
    background: #3b80db;
    color: #fff;
    position: relative;
    font-weight: normal;
    margin: 0;
    padding: 0;
    font-family: "微软雅黑";
  }
  .hot_project_dd:after {
    content: "";
    display: block;
    width: 880px;
    height: 2px;
    background: #d9d9d9;
    left: 293px;
    bottom: 0px;
    position: absolute;
  }
</style>
