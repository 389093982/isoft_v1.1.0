<template>
  <div>
    <IBeautifulCard title="热门问题">
      <div slot="content">
        <ul style="padding-left: 0em;">
          <li v-for="hot_issue in hot_issues">
            <span style="float:right;color: #999;">{{hot_issue.issue_time}}</span>
            <a :href="hot_issue.issue_link" target="_blank" :style="{'color':hot_issue.issue_new == true ? '#f55e13!important':'none'}">
              {{hot_issue.issue_title}} <img v-if="hot_issue.issue_new == true" src="../../../assets/news.gif">
            </a>
          </li>
        </ul>
      </div>
    </IBeautifulCard>
  </div>
</template>

<script>
  import IBeautifulCard from "../../Common/card/IBeautifulCard"

  export default {
    name: "HotIssue",
    components:{IBeautifulCard},
    data(){
      return {
        hot_issues:[
          {
            "issue_title":"Java泛型是什么?",
            "issue_link":"https://www.baidu.com",
            "issue_time":"02-21",
            "issue_new":true
          },
          {
            "issue_title":"Java反射是什么?",
            "issue_link":"https://www.baidu.com",
            "issue_time":"02-21",
            "issue_new":true
          },
          {
            "issue_title":"Python难不难学?",
            "issue_link":"https://www.baidu.com",
            "issue_time":"02-21",
            "issue_new":true
          },
          {
            "issue_title":"Go语言有发展前途么?",
            "issue_link":"https://www.baidu.com",
            "issue_time":"02-21",
            "issue_new":false
          },
          {
            "issue_title":"如何快速赚取积分?",
            "issue_link":"https://www.baidu.com",
            "issue_time":"02-21",
            "issue_new":false
          }
        ]
      }
    }
  }
</script>

<style scoped>
 a{
   color: black;
 }
 a:hover{
   color: red;
 }
 li {
   padding: 10px 15px 0 15px;
   list-style: none;
 }
</style>
