<template>
  <IBeautifulCard title="工具盒">
    <div slot="content">
      <ul>
        <li v-for="tool in tools" class="li">
          <i class="icon" :style="{background:'url(' + tool.icon + ') 0 0 no-repeat'}"></i>
          <p class="label">{{tool.label}}</p>
        </li>
      </ul>
      <div style="clear: both;"></div>
    </div>
  </IBeautifulCard>
</template>

<script>
  import IBeautifulCard from "../Common/card/IBeautifulCard"

  export default {
    name: "ToolBox",
    components:{IBeautifulCard},
    data(){
      return {
        tools:[
          {
            "icon":"https://as.zbjimg.com/rake-component/fe-employer/fe-zbj-quickservice-res/0.0.8/images/qs-icons.png",
            "label":"博客中心",
            "href":"/iblog/blog_list"
          },
          {
            "icon":"https://as.zbjimg.com/rake-component/fe-employer/fe-zbj-quickservice-res/0.0.8/images/qs-icons.png",
            "label":"心声社区",
            "href":"https://www.baidu.com"
          },
          {
            "icon":"https://as.zbjimg.com/rake-component/fe-employer/fe-zbj-quickservice-res/0.0.8/images/qs-icons.png",
            "label":"博客天地",
            "href":"https://www.baidu.com"
          },
          {
            "icon":"https://as.zbjimg.com/rake-component/fe-employer/fe-zbj-quickservice-res/0.0.8/images/qs-icons.png",
            "label":"后台管理",
            "href":"https://www.baidu.com"
          },
          {
            "icon":"https://as.zbjimg.com/rake-component/fe-employer/fe-zbj-quickservice-res/0.0.8/images/qs-icons.png",
            "label":"商标查询",
            "href":"https://www.baidu.com"
          },
          {
            "icon":"https://as.zbjimg.com/rake-component/fe-employer/fe-zbj-quickservice-res/0.0.8/images/qs-icons.png",
            "label":"商标查询",
            "href":"https://www.baidu.com"
          }
        ],
      }
    },
  }
</script>

<style scoped>
 .li {
   list-style: none;
   float: left;
   width: 25%;
   height: 75px;
   padding: 10px 0;
   text-align: center;
   border: 1px solid transparent;
   cursor: pointer;
 }
 .icon {
   background-position: 0 0;
   display: inline-block;
   width: 30px;height: 30px;
   vertical-align: top;
 }
 .label {
   margin: 10px 0 0;
   color: #333;
   font-size: 12px;
   line-height: 1;
 }
</style>
