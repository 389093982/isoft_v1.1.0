<template>
<li>
  <IBeautifulHoverShadow>
    <a :href="item.item_link">
      <div class="cont_main">
        <h3 class="title">{{item.item_label}}</h3>
        <p class="desc">{{item.item_desc}}</p>
      </div>
    </a>
  </IBeautifulHoverShadow>
</li>
</template>

<script>
  import IBeautifulHoverShadow from "./IBeautifulHoverShadow"

  export default {
    name: "IBeautifulHoverBox",
    components:{IBeautifulHoverShadow},
    props:{
      item:{
        type: Object,
        default: function () {
          return {item_label:'云服务器0',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',}
        }
      }
    }
  }
</script>

<style scoped>
  ul li {
    width: 280px;
    float: left;
    margin: 0 10px;
    box-sizing: border-box;
  }

  li a:hover .cont_main h3.title {
    color: #00a4ff;
  }

  li a .cont_main {
    display: table-cell;
    vertical-align: top;
  }

  li a .cont_main h3.title {
    font-size: 18px;
    font-weight: 400;
    color: #333;
    line-height: 32px;
    white-space: nowrap;
    overflow: hidden;
    display: block;
    text-overflow: ellipsis;
  }

  li a .cont_main p.desc {
    font-size: 14px;
    line-height: 1.5;
    color: #666;
    height: 42px;
    margin-top: -2px;
    display: -webkit-box;
    -webkit-box-orient: vertical;
    -webkit-line-clamp: 2;
    overflow: hidden;
  }
</style>
