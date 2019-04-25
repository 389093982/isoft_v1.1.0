<template>
  <Row type="flex" justify="end" class="code-row-bg" style="margin-bottom: 5px;">
    <Col span="10">
      <Input v-model="search" placeholder="搜索..."/>
    </Col>
    <Col span="2">
      <Button type="success" @click="handleSimpleSearch">搜索</Button>
    </Col>
  </Row>
</template>

<script>
  export default {
    name: "ISimpleSearch",
    data(){
      return {
        // 默认值是空
        search:'',
      }
    },
    methods:{
      handleSimpleSearch:function(){
        this.$emit('handleSimpleSearch', this.search);
      },
    }
  }
</script>

<style scoped>

</style>
