<template>
  <div class="items items-data">
    <div class="center">
      <p class="main-title">{{pData.title}}</p>
      <div class="flex-row">
        <div class="item" v-for="itemData in pData.pItemData">
          <p class="item-title">{{itemData.itemTitle}}</p>
          <div class="item-block">
            <p class="item-p" v-for="itemIntroduce in itemData.itemIntroduces">{{itemIntroduce}}</p>
          </div>
          <div class="item-btn">
            <a :href="itemData.itemLink" class="item-btn-a" target="_blank">了解更多</a>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    name: "IBeautifulListLink",
    props: {
      // 带有默认值的对象
      pData: {
        type: Object,
        // 对象或数组默认值必须从一个工厂函数获取
        default: function () {
          return {
            title: '网站建设解决方案',
            pItemData:[
              {
                itemTitle:'网络营销解决方案',
                itemIntroduces:['网络营销规划','线上业务梳理','短、有含义的域名','营销型网站','网络推广计划','线上沟通工具','运营数据分析工具'],
                itemLink:'https://www.baidu.com'
              },
              {
                itemTitle:'网络营销解决方案',
                itemIntroduces:['网络营销规划','线上业务梳理','短、有含义的域名','营销型网站','网络推广计划','线上沟通工具','运营数据分析工具'],
                itemLink:'https://www.baidu.com'
              },
              {
                itemTitle:'网络营销解决方案',
                itemIntroduces:['网络营销规划','线上业务梳理','短、有含义的域名','营销型网站','网络推广计划','线上沟通工具','运营数据分析工具'],
                itemLink:'https://www.baidu.com'
              },
              {
                itemTitle:'网络营销解决方案',
                itemIntroduces:['网络营销规划','线上业务梳理','短、有含义的域名','营销型网站','网络推广计划','线上沟通工具','运营数据分析工具'],
                itemLink:'https://www.baidu.com'
              }
            ]
          }
        }
      },
    }
  }
</script>

<style scoped>
  *{padding: 0;margin: 0;}
  body{width: 100%;height: 100%;background: #fff;font-size: 14px;font-family: 'Microsoft Yahei', 'PingFangSC', sans-serif;color: #333;}
  a{text-decoration: none;color: inherit;transition: all .3s ease;}
  .items{padding: 60px 0;overflow: hidden;}
  .center{width: 1200px;margin: 0 auto;position: relative;min-width: 1200px;}
  .main-title{font-size: 30px;text-align: center;margin-bottom: 50px;}
  .items-data .flex-row{justify-content: space-between;display: flex;flex-direction: row;align-items: center;}
  .items-data .item{position: relative;width: 276px;height: 546px;box-shadow: 0 0 3px 2px rgba(135, 135, 135, 0.15);}
  .items-data .item-title{text-align: center;height: 83px;line-height: 83px;background-color: #293949;font-size: 18px;color: #fff;margin-bottom: 40px;}
  .items-data .item-p{text-align: center;width: 62%;margin: 0 auto;height: 40px;line-height: 40px;box-sizing: border-box;border-bottom: 2px dotted #E6E6E6;}
  .items-data .item-p:last-child{border: none;}
  .items-data .item-btn{position: absolute;width: 140px;height: 40px;line-height: 40px;text-align: center;left: 50%;bottom: 50px;transform: translateX(-50%);}
  .items-data .item-btn-a{background-color: #C5C5C5;color: #fff;display: block;}
  .items-data .item:hover .item-title{background-color: #3CA7FC;}
  .items-data .item:hover .item-btn-a{background-color: #3CA7FC;}
</style>
