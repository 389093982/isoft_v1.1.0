<template>
  <div id="footer">
    <BackTop>
      <div class="top">返回顶端</div>
    </BackTop>



    <Row>
      <Col span="4" style="padding: 50px;">
        <VerticalList :showRightBorder="true" title="新手帮助">
          <VerticalListItem>新手引导</VerticalListItem>
          <VerticalListItem>视频发布流程</VerticalListItem>
          <VerticalListItem>账号注册</VerticalListItem>
        </VerticalList>
      </Col>
      <Col span="4" style="padding: 50px;">
        <VerticalList :showRightBorder="true" title="安全中心">
          <VerticalListItem>公司简介</VerticalListItem>
          <VerticalListItem>站点地图</VerticalListItem>
        </VerticalList>
      </Col>
      <Col span="4" style="padding: 50px;">
        <VerticalList :showRightBorder="true" title="安全中心">
          <VerticalListItem>公司简介</VerticalListItem>
          <VerticalListItem>站点地图</VerticalListItem>
        </VerticalList>
      </Col>
      <Col span="4" style="padding: 50px;">
        <VerticalList :showRightBorder="true" title="安全中心">
          <VerticalListItem>关于我们</VerticalListItem>
          <VerticalListItem>联系我们</VerticalListItem>
          <VerticalListItem>人才招聘</VerticalListItem>
          <VerticalListItem>商家入驻</VerticalListItem>
          <VerticalListItem>广告服务</VerticalListItem>
          <VerticalListItem>友情链接</VerticalListItem>
          <VerticalListItem>Isoft论坛</VerticalListItem>
        </VerticalList>
      </Col>
      <Col span="8" style="padding: 50px;">
        <VerticalList title="客服中心">
          <VerticalListItem>在线客服</VerticalListItem>
          <VerticalListItem>我要举报</VerticalListItem>
          <VerticalListItem>提出意见</VerticalListItem>
        </VerticalList>
      </Col>
    </Row>

    <Row style="text-align: center;">
      <p style="margin: 40px;">Copyright © 2017-2018 Isoft 版权所有 All Rights Reserved.</p>
    </Row>
  </div>
</template>

<script>
  import VerticalList from "../list/VerticalList.vue"
  import VerticalListItem from "../list/VerticalListItem";

  export default {
    name: "Footer",
    components:{VerticalListItem, VerticalList,VerticalListItem},
    methods: {

    }
  }
</script>

<style scoped lang="stylus" rel="stylesheet/stylus">
  .top
    padding: 10px;
    background: rgba(0, 153, 229, .7);
    color: #fff;
    text-align: center;
    border-radius: 2px;
  #footer
    background-color #ffffff
</style>
