<template>
  <Menu active-name="1-1" width="auto" :open-names="['1']">

    <Submenu name="1">
      <template slot="title">
        <Icon type="md-barcode"></Icon>
        博客天地
      </template>
      <MenuItem name="1-1">
        <router-link to="/iblog/catalog_add">新增/编辑分类</router-link>
      </MenuItem>
      <MenuItem name="1-2">
        <router-link to="/iblog/blog_add">新增/编辑文章</router-link>
      </MenuItem>
    </Submenu>

    <Submenu name="2">
      <template slot="title">
        <Icon type="ios-navigate"></Icon>
        IFile 文件管理
      </template>
      <MenuItem name="2-1">
        <router-link to="/ifile/ifile">IFile 文件管理</router-link>
      </MenuItem>
      <MenuItem name="2-2">
        <router-link to="/ifile/heartbeat">心跳检测管理</router-link>
      </MenuItem>
      <MenuItem name="2-3">
        <router-link to="/ifile/log">系统日志信息</router-link>
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
