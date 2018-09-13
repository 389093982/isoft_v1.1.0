<template>
  <div>
    <Menu mode="horizontal" :theme="theme1" active-name="1" style="padding-left: 700px;">
      <MenuItem name="1">
        <Icon type="ios-people" />
        用户管理
      </MenuItem>
      <Submenu name="2">
        <template slot="title">
          <Icon type="ios-stats" />
          统计分析
        </template>
        <MenuGroup title="使用">
          <MenuItem name="2-1">新增和启动</MenuItem>
          <MenuItem name="2-2">活跃分析</MenuItem>
          <MenuItem name="2-3">时段分析</MenuItem>
        </MenuGroup>
        <MenuGroup title="留存">
          <MenuItem name="2-4">用户留存</MenuItem>
          <MenuItem name="2-5">流失用户</MenuItem>
        </MenuGroup>
      </Submenu>
      <MenuItem name="3">
        <Icon type="ios-construct" />
        综合设置
      </MenuItem>
      <MenuItem name="4">
        <Icon type="ios-paper" />
        内容管理
      </MenuItem>
    </Menu>
  </div>
</template>

<script>
  export default {
    name: "Header",
    data () {
      return {
        theme1: 'light'
      }
    }
  }
</script>


