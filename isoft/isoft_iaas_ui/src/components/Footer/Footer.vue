<template>
  <div id="footer">
    <BackTop>
      <div class="top">返回顶端</div>
    </BackTop>

    <Row>
      <FrindLink/>
    </Row>

    <Row>
      <Col span="4" style="padding: 50px;">
        <IVerticalList :showRightBorder="true" title="新手帮助">
          <IVerticalListItem>新手引导</IVerticalListItem>
          <IVerticalListItem>视频发布流程</IVerticalListItem>
          <IVerticalListItem>账号注册</IVerticalListItem>
        </IVerticalList>
      </Col>
      <Col span="4" style="padding: 50px;">
        <IVerticalList :showRightBorder="true" title="安全中心">
          <IVerticalListItem>公司简介</IVerticalListItem>
          <IVerticalListItem>站点地图</IVerticalListItem>
        </IVerticalList>
      </Col>
      <Col span="4" style="padding: 50px;">
        <IVerticalList :showRightBorder="true" title="安全中心">
          <IVerticalListItem>公司简介</IVerticalListItem>
          <IVerticalListItem>站点地图</IVerticalListItem>
        </IVerticalList>
      </Col>
      <Col span="4" style="padding: 50px;">
        <IVerticalList :showRightBorder="true" title="安全中心">
          <IVerticalListItem>关于我们</IVerticalListItem>
          <IVerticalListItem>联系我们</IVerticalListItem>
          <IVerticalListItem>人才招聘</IVerticalListItem>
          <IVerticalListItem>商家入驻</IVerticalListItem>
          <IVerticalListItem>广告服务</IVerticalListItem>
          <IVerticalListItem>友情链接</IVerticalListItem>
          <IVerticalListItem>Isoft论坛</IVerticalListItem>
        </IVerticalList>
      </Col>
      <Col span="8" style="padding: 50px;">
        <IVerticalList title="客服中心">
          <IVerticalListItem>在线客服</IVerticalListItem>
          <IVerticalListItem>我要举报</IVerticalListItem>
          <IVerticalListItem>提出意见</IVerticalListItem>
        </IVerticalList>
      </Col>
    </Row>

    <Row style="text-align: center;">
      <p style="margin: 40px;">Copyright © 2017-2018 Isoft 版权所有 All Rights Reserved.</p>
    </Row>
  </div>
</template>

<script>
  import IVerticalList from "../Common/list/IVerticalList"
  import IVerticalListItem from "../Common/list/IVerticalListItem"
  import FrindLink from "./FrindLink"

  export default {
    name: "Footer",
    components:{IVerticalList,IVerticalListItem, FrindLink},
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
