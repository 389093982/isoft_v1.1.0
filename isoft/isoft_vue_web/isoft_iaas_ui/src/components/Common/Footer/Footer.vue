<template>
  <div id="footer">
    <BackTop>
      <div class="top">返回顶端</div>
    </BackTop>

    <Row>
      <Col span="4">
        <dl>
          <dt style="height: 35px;line-height: 35px;font-family: '微软雅黑';font-size: 20px;">关于我们</dt>
          <dd><a href="javascript:;">公司简介</a></dd>
          <dd><a href="javascript:;">站点地图</a></dd>
        </dl>
      </Col>
      <Col span="4">
        <dl>
          <dt style="height: 35px;line-height: 35px;font-family: '微软雅黑';font-size: 20px;">关于我们</dt>
          <dd><a href="javascript:;">公司简介</a></dd>
          <dd><a href="javascript:;">站点地图</a></dd>
        </dl>
      </Col>
      <Col span="8">
        <p>
          <a href="#">关于我们</a> |
          <a href="#">联系我们</a> |
          <a href="#">人才招聘</a> |
          <a href="#">商家入驻</a> |
          <a href="#">广告服务</a> |
          <a href="#">友情链接</a> |
          <a href="#">Isoft论坛</a> |
        </p>
      </Col>
      <Col span="8">
        <p style="font-size: 24px!important;padding: 0;">13590331603</p>
        <p style="line-height: 24px;">24小时客服热线(投诉或意见)</p>
        <p>
          <a href="javascript:;" style="display: block;border-radius: 5px;transition: all .4s;display: inline-block;
          width: 161px;height: 35px;padding: 0;margin-top: 10px;background: #FFF;border: 1px solid #ddd;
          font-size: 14px;line-height: 35px;text-align: center;color: #00991b!important;cursor: pointer;">在线客服</a>
        </p>
        <p>
          <a href="javascript:;" style="display: block;border-radius: 5px;transition: all .4s;display: inline-block;
          width: 161px;height: 35px;padding: 0;margin-top: 10px;background: #FFF;border: 1px solid #ddd;
          font-size: 14px;line-height: 35px;text-align: center;color: #00991b!important;cursor: pointer;">我要举报</a>
        </p>
      </Col>
    </Row>

    <p style="margin: 40px;">Copyright © 2017-2018 Isoft 版权所有</p>
  </div>
</template>

<script>
  export default {
    methods: {

    }
  }
</script>

<style scoped lang="stylus" rel="stylesheet/stylus">
  .top{
    padding: 10px;
    background: rgba(0, 153, 229, .7);
    color: #fff;
    text-align: center;
    border-radius: 2px;
  }

  #footer
    height 80px
    background-color #ffffff
    text-align center
    clear both
    padding 30px
    font-size 12px
    p
      line-height 20px
    a
      color #666
      text-decoration none
    a:hover
      color red
      text-decoration none
</style>
