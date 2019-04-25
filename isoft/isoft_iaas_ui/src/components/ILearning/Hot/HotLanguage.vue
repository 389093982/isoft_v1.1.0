<template>
  <IBeautifulListLink :pData="pData"/>
</template>

<script>
  import IBeautifulListLink from "../../Common/link/IBeautifulListLink"

  export default {
    name: "HotLanguage",
    components:{IBeautifulListLink},
    data(){
      return {
        pData:{
          title: '热门编程语言推荐',
          pItemData:[
            {
              itemTitle:'Java',
              itemIntroduces:['桌面应用系统开发','分布式系统开发','Web应用系统开发','企业级应用开发','安卓应用开发','大数据技术','服务器程序'],
              itemLink:'https://www.baidu.com'
            },
            {
              itemTitle:'Python',
              itemIntroduces:['WEB开发','桌面软件','网络编程','爬虫开发','云计算开发','人工智能','自动化运维','科学运算'],
              itemLink:'https://www.baidu.com'
            },
            {
              itemTitle:'Golang',
              itemIntroduces:['网络营销规划','线上业务梳理','短、有含义的域名','营销型网站','网络推广计划','线上沟通工具','运营数据分析工具'],
              itemLink:'https://www.baidu.com'
            },
            {
              itemTitle:'C/C++',
              itemIntroduces:['网络营销规划','线上业务梳理','短、有含义的域名','营销型网站','网络推广计划','线上沟通工具','运营数据分析工具'],
              itemLink:'https://www.baidu.com'
            }
          ]
        }
      }
    }
  }
</script>

<style scoped>

</style>
