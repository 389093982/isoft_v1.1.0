<template>
  <div>
    迁移类型：<a href="javascript:;" @click="createTableMigrate">创建表迁移</a>
    <a href="javascript:;" @click="alterTableMigrate">变更表迁移</a>
    <a href="javascript:;">迁移执行过程查看</a>
  </div>
</template>

<script>
  export default {
    name: "MigrateList",
    data(){
      return {}
    },
    methods:{
      createTableMigrate(){
        alert(1);
      },
      alterTableMigrate(){
        alert(1);
      },
    }
  }
</script>

<style scoped>

</style>
