<template>
  <div id="course_search" data-am-sticky style="background:rgba(255,1,24,0.59);height: 60px;">
    <div style="position: relative;top:10px;float: left;left: 320px;">
      <div id="_search">
        <input class="search" title="请输入搜索文字" placeholder="请输入搜索文字" v-model.trim="search_data">
        <input class="submit" type="submit" title="提交" @click="submitFunc">
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    name: "Search",
    data(){
      return {
        search_data:"",
      }
    },
    methods:{
      submitFunc:function () {
        this.$emit("submitFunc", this.search_data);
      }
    }
  }
</script>

<style scoped>
  .search{
    padding-left: 16px;
    float: left;
    width: 650px;
    background: #FFF;
    vertical-align: middle;
    height: 40px;
    box-shadow: none;
    -webkit-appearance: none;
    border: none;
    outline: 0;
    color: #333;
    font-size: 14px;
  }
  .submit{
    background: url('../../../assets/search.png') 50% 50% no-repeat;
    background-size: 24px;
    position: absolute;
    right: 10px;
    top: 15px;
    margin-top: -7px;
    width: 24px;
    height: 24px;
    border: none;
    color: transparent;
    box-sizing: initial;
    cursor: pointer;
    outline: 0;
  }
</style>
