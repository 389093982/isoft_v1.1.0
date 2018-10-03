<template>
  <div>
    <Menu mode="horizontal" :theme="theme1" active-name="1" style="padding-left: 650px;">
      <Submenu name="1">
        <template slot="title">
          <Icon type="ios-stats" />
          统计分析
        </template>
        <MenuGroup title="使用">
          <MenuItem name="1-1">新增和启动</MenuItem>
          <MenuItem name="1-2">活跃分析</MenuItem>
          <MenuItem name="1-3">时段分析</MenuItem>
        </MenuGroup>
        <MenuGroup title="留存">
          <MenuItem name="1-4">用户留存</MenuItem>
          <MenuItem name="1-5">流失用户</MenuItem>
        </MenuGroup>
      </Submenu>
      <MenuItem name="2">
        <Icon type="ios-construct" />
        综合设置
      </MenuItem>
      <MenuItem name="3">
        <Icon type="ios-paper" />
        内容管理
      </MenuItem>
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
    </Menu>
  </div>
</template>

<script>
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
        localStorage.removeItem("userName");
        this.loginUserName = "";
        window.location.href = "/api/auth/redirectToLogin/?redirectUrl=" + window.location.href;
      }
    },
    mounted:function(){
      this.loginUserName = localStorage.getItem("userName");
    },
  }
</script>


