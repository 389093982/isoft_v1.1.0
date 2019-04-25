<template>
  <div class="layout">
    <Layout :style="{minHeight: '100vh'}">
      <Sider :style="{background: '#fff'}">
        <Menu theme="light" width="auto">
          <Submenu name="1">
            <template slot="title">
              <Icon type="ios-navigate"></Icon>
              日志管理
            </template>
            <MenuItem name="1-1"><router-link to="/sso/loginRecord">登录记录</router-link></MenuItem>
          </Submenu>
          <Submenu name="2">
            <template slot="title">
              <Icon type="ios-keypad"></Icon>
              用户管理
            </template>
            <MenuItem name="2-1">用户查询</MenuItem>
            <MenuItem name="2-2">权限管理</MenuItem>
          </Submenu>
          <Submenu name="3">
            <template slot="title">
              <Icon type="ios-analytics"></Icon>
              系统注册
            </template>
            <MenuItem name="3-1"><router-link to="/sso/appRegist">注册管理</router-link></MenuItem>
          </Submenu>
        </Menu>
      </Sider>
      <Layout>
        <Content :style="{padding: '0 16px 16px'}">
          <slot></slot>
        </Content>
      </Layout>
    </Layout>
  </div>
</template>

<script>
  export default {
    name: "LeftMenu"
  }
</script>

<style scoped>
  .menu-item span{
    display: inline-block;
    overflow: hidden;
    width: 69px;
    text-overflow: ellipsis;
    white-space: nowrap;
    vertical-align: bottom;
    transition: width .2s ease .2s;
  }
  .menu-item i{
    transform: translateX(0px);
    transition: font-size .2s ease, transform .2s ease;
    vertical-align: middle;
    font-size: 16px;
  }
  .collapsed-menu span{
    width: 0px;
    transition: width .2s ease;
  }
  .collapsed-menu i{
    transform: translateX(5px);
    transition: font-size .2s ease .2s, transform .2s ease .2s;
    vertical-align: middle;
    font-size: 22px;
  }
</style>
