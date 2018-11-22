<template>
  <div class="layout">
    <Menu mode="horizontal" :theme="theme1" active-name="1">
      <div class="layout-nav">
        <Submenu name="1">
          <template slot="title">
            <Icon type="ios-stats" />
            精品应用
          </template>
          <MenuGroup title="博客天地">
            <MenuItem name="1-1"><router-link to="/iblog/blog_list">热门博文推荐</router-link></MenuItem>
            <MenuItem name="1-2"><router-link to="/iblog/catalog_add">新增/编辑分类</router-link></MenuItem>
            <MenuItem name="1-3"><router-link to="/iblog/blog_add">新增/编辑文章</router-link></MenuItem>
            <MenuItem name="1-4">我的博客空间</MenuItem>
          </MenuGroup>
          <MenuGroup title="在线学习系统">
            <MenuItem name="1-5"><router-link to="/ilearning/index">精品课程</router-link></MenuItem>
            <MenuItem name="1-6"><router-link to="/inote/index">云笔记</router-link></MenuItem>
          </MenuGroup>
          <MenuGroup title="云存储">
            <MenuItem name="1-7"><router-link to="/ifile/ifile">IFile 文件存储</router-link></MenuItem>
          </MenuGroup>
          <MenuGroup title="友情链接">
            <MenuItem name="1-9"><router-link to="/ifile/ifile">心声社区</router-link></MenuItem>
            <MenuItem name="1-10"><router-link to="/share/list">共享空间</router-link></MenuItem>
          </MenuGroup>
        </Submenu>
        <MenuItem name="2">
          <Icon type="ios-construct" />
          综合设置
        </MenuItem>
        <Submenu name="3">
          <template slot="title">
            <Icon type="ios-paper" />
            内容管理
          </template>
          <MenuGroup title="配置项管理">
            <MenuItem name="3-1"><router-link to="/ilearning/configuration">查看配置项</router-link></MenuItem>
            <MenuItem name="3-1"><router-link to="/monitor/filterPageHeartBeat">应用监听</router-link></MenuItem>
            <MenuItem name="3-1"><router-link to="/cms/friendLinkList">友情链接</router-link></MenuItem>
          </MenuGroup>
        </Submenu>
        <Submenu name="4">
          <template slot="title">
            <Icon type="ios-people" />
            <span v-if="loginUserName">{{loginUserName}}</span>
            <span v-else>登录</span>
          </template>
          <MenuGroup title="账号管理">
            <MenuItem name="4-1" @click.native="cancelUser">注销</MenuItem>
            <MenuItem name="4-2" @click.native="cancelUser">切换账号</MenuItem>
          </MenuGroup>
        </Submenu>
      </div>
    </Menu>
  </div>
</template>

<script>
  import {getCookie} from '../../../tools/index'
  import {delCookie} from '../../../tools/index'

  export default {
    name: "Header",
    data () {
      return {
        theme1: 'light',
        loginUserName:'',
      }
    },
    methods:{
      cancelUser() {
        delCookie("userName");
        this.loginUserName = "";
        window.location.href = "/api/auth/redirectToLogin/?redirectUrl=" + window.location.href;
      }
    },
    mounted:function(){
      this.loginUserName = getCookie("userName");
    },
  }
</script>

<style scoped>
  .layout-nav{
    width: 620px;
    margin: 0 auto;
    margin-right: 20px;
  }
</style>

