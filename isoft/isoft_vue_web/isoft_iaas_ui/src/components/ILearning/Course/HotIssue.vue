<template>
  <div>
    <div style="background-color: #f6f6f6;height: 38px;border-bottom: solid 1px #d9d9d9;">
      <p style="float: left;line-height: 38px;margin-left: 20px;">热门问题</p>
      <a href="javascript:;" style="float: right;line-height: 38px;margin-right: 20px;color: red;">
        <img src="../../../assets/images/common/question.jpg"/>
      </a>
    </div>
    <div style="background: #ffffff;font-size: 14px;padding-bottom: 10px;">
      <ul style="padding-left: 0em;">
        <li style="padding: 10px 15px 0 15px;list-style: none;">
          <span style="float:right;color: #999;">02-21</span>
          <a href="#" target="_blank" style="color: #f55e13!important;">
            Java泛型是什么? <img src="../../../assets/news.gif">
          </a>
        </li>
        <li style="padding: 10px 15px 0 15px;list-style: none;">
          <span style="float:right;color: #999;">02-21</span>
          <a href="#" target="_blank" style="color: #f55e13!important;">
            Java反射是什么? <img src="../../../assets/news.gif">
          </a>
        </li>
        <li style="padding: 10px 15px 0 15px;list-style: none;">
          <span style="float:right;color: #999;">02-21</span>
          <a href="#" target="_blank">
            Python难不难学?
          </a>
        </li>
        <li style="padding: 10px 15px 0 15px;list-style: none;">
          <span style="float:right;color: #999;">02-21</span>
          <a href="#" target="_blank">
            Go语言有发展前途么?
          </a>
        </li>
        <li style="padding: 10px 15px 0 15px;list-style: none;">
          <span style="float:right;color: #999;">02-21</span>
          <a href="#" target="_blank">
            如何快速赚取积分?
          </a>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
  export default {
    name: "HotIssue"
  }
</script>

<style scoped>
 a{
   color: black;
 }
 a:hover{
   color: red;
 }
</style>
