<template>
  <div style="margin: 20px;margin-top: 30px;">
    <div style="border-bottom: 2px solid #d9d9d9;">
      <h6 class="title" :title="title">{{title}}</h6>
    </div>
    <Row :gutter="50">
      <Col span="8" style="margin-top: 12px;" v-for="item in items">
        <span style="font-size: 14px;">{{item.item_label}}</span>
        <IBeautifulButtonLink msg="点击了解详情" floatstyle="right" :hrefaddr="item.item_href"/>
      </Col>
    </Row>
  </div>
</template>

<script>
  import IBeautifulButtonLink from "../../Common/link/IBeautifulButtonLink"

  export default {
    name: "IHotRecommand",
    components:{IBeautifulButtonLink},
    props:{
      title:{
        type: String,
        default: "测试 title"
      },
      items:{
        type: Array,
        default: [
          {
            item_label:"测试 label",
            item_href:"测试 href",
          }
        ],
      },
    },
  }
</script>

<style scoped>
  .title{
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
