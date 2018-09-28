<template>
  <Menu active-name="1-2" width="auto" :open-names="['1']">
    <Submenu name="1">
      <template slot="title">
        <Icon type="ios-navigate"></Icon>
        部署环境管理
      </template>
      <MenuItem name="1-1">
        <router-link to="/env/list">环境清单</router-link>
      </MenuItem>
    </Submenu>
    <Submenu name="2">
      <template slot="title">
        <Icon type="ios-cube"></Icon>
        配置管理
      </template>
      <MenuItem name="2-1">
        <router-link to="/config/list">配置管理</router-link>
      </MenuItem>
    </Submenu>
    <Submenu name="3">
      <template slot="title">
        <Icon type="ios-analytics"></Icon>
        应用服务管理
      </template>
      <MenuItem name="3-1">
        <router-link to="/service/list?service_type=mysql">mysql管理</router-link>
      </MenuItem>
      <MenuItem name="3-2">
        <router-link to="/service/list?service_type=beego">beego应用部署</router-link>
      </MenuItem>
      <MenuItem name="3-3">
        <router-link to="/service/list?service_type=api">api应用部署</router-link>
      </MenuItem>
    </Submenu>
    <Submenu name="4">
      <template slot="title">
        <Icon type="ios-analytics"></Icon>
        系统监控
      </template>
      <MenuItem name="4-1">
        <router-link to="/monitor/list">系统心跳检测</router-link>
      </MenuItem>
    </Submenu>
  </Menu>
</template>

<script>
  export default {
    name: "LeftMenu"
  }
</script>

<style scoped>
</style>
