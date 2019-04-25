<template>
  <IBeautifulCard title="新闻公告">
    <ul slot="content" style="padding-left: 0em;">
      <li v-for="notice in notices">
        <span style="float:right;color: #999;">{{notice.news_time}}</span>
        <a :href="notice.news_link" target="_blank" :style="{'color':notice.news_new == true ? '#f55e13!important':'none'}">
          {{notice.news_title}} <img v-if="notice.news_new == true" src="../../../assets/news.gif">
        </a>
      </li>
    </ul>
    <span slot="header_right">
      <a href="javascript:;">
        <img src="../../../assets/images/common/more.jpg"/>
      </a>
    </span>
  </IBeautifulCard>
</template>

<script>
  import IBeautifulCard from "../../Common/card/IBeautifulCard"

  export default {
    name: "NewsNotice",
    components:{IBeautifulCard},
    data(){
      return {
        notices:[
          {
            "news_title":"2018年起会员积分调整",
            "news_link":"https://www.baidu.com",
            "news_time":"02-21",
            "news_new":true
          },
          {
            "news_title":"Java泛型是什么?",
            "news_link":"https://www.baidu.com",
            "news_time":"02-21",
            "news_new":false
          },
          {
            "news_title":"2018年起会员积分调整",
            "news_link":"https://www.baidu.com",
            "news_time":"02-21",
            "news_new":true
          },
          {
            "news_title":"Java泛型是什么?",
            "news_link":"https://www.baidu.com",
            "news_time":"02-21",
            "news_new":false
          },
          {
            "news_title":"2018年起会员积分调整",
            "news_link":"https://www.baidu.com",
            "news_time":"02-21",
            "news_new":true
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
    padding: 3px 15px 0 15px;
    list-style: none;
    font-size: 14px;
  }
</style>
