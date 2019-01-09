<template>
<li>
  <a :href="item.item_link">
    <div class="cont_main">
      <h3 class="title">{{item.item_label}}</h3>
      <p class="desc">{{item.item_desc}}</p>
    </div>
  </a>
</li>
</template>

<script>
  export default {
    name: "IBeautifulHoverBox",
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

  ul li a {
    display: block;
    border: 1px solid #E5E5E5;
    background-color: #fff;
    margin-top: 2px;
    margin-bottom: 10px;
    padding: 13px 20px 15px;
  }

  li a:hover {
    box-shadow: 0 3px 7px 0 rgba(0, 0, 0, .1);
    border-color: rgba(229, 229, 229, .73);
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
