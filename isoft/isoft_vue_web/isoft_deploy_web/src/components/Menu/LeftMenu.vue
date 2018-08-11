<template>
  <div class="account-l fl" style="margin: 10px;">
    <a class="list-title" href="/user/manage">统一部署管理系统</a>
    <ul id="accordion" class="accordion">
      <li>
        <div class="link">
          <i class="fa fa-leaf"></i>部署环境管理<i class="fa fa-chevron-down"></i>
        </div>
        <ul class="submenu">
          <li><a href="/env/list">环境清单</a></li>
        </ul>
      </li>
      <li>
        <div class="link">
          <i class="fa fa-leaf"></i>基础软件管理<i class="fa fa-chevron-down"></i>
        </div>
        <ul class="submenu">
          <li><a href="#">java运行环境管理</a></li>
          <li><a href="#">python运行环境管理</a></li>
          <li><a href="#">golang运行环境管理</a></li>
          <li><a href="#">mysql数据库管理</a></li>
        </ul>
      </li>
      <li>
        <div class="link">
          <i class="fa fa-leaf"></i>应用服务管理<i class="fa fa-chevron-down"></i>
        </div>
        <ul class="submenu">
          <li><a href="/service/list?service_type=mysql">mysql管理</a></li>
          <li><a href="/service/list?service_type=nginx">nginx管理</a></li>
          <li><a href="/service/list?service_type=beego">beego应用部署</a></li>
          <li><a href="/service/list?service_type=api">api应用部署</a></li>
          <li><a href="/service/monitor">服务监控</a></li>
        </ul>
      </li>
      <li>
        <div class="link">
          <i class="fa fa-leaf"></i>部署教程<i class="fa fa-chevron-down"></i>
        </div>
        <ul class="submenu">
          <li><a href="javascript:">教程</a></li>
        </ul>
      </li>
    </ul>
  </div>
</template>

<script>
  import leftnav from '../../../static/leftnav/js/leftnav.js'
  export default {
      name: "LeftMenu"
  }
</script>

<style scoped  type="text/stylus" rel="stylesheet/stylus">
  @import '../../../static/leftnav/css/leftnav.css';
</style>
