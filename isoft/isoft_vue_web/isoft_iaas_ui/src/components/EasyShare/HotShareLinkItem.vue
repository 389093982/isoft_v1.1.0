<template>
  <div style="margin-left: 150px;margin-right: 150px;">
    <Row :gutter="10">
      <Col span="6" style="text-align: center;" v-for="share_type in getCurrentPage">
        <div class="item" style="margin-left:10px; margin-right:10px; height: 100px;">{{share_type}}</div>
      </Col>
    </Row>
    <Row>
      <a href="javascript:;" style="float: right;margin: 10px 30px 10px 0;font-size: 12px;">换一批</a>
    </Row>
  </div>
</template>

<script>
  export default {
    name: "HotShareLinkItem",
    data(){
      return {
        hot_share_type: this.GLOBAL.hot_share_type,
        currentPageNo:1,    // 当前页
      }
    },
    computed:{
      getCurrentPage:function(){
        return this.pagination(this.currentPageNo,4,this.hot_share_type);
      },
    },
    methods:{
      // 对数组进行也分
      pagination: function (pageNo, pageSize, array) {
        let offset = (pageNo - 1) * pageSize;
        return (offset + pageSize >= array.length) ? array.slice(offset, array.length) : array.slice(offset, offset + pageSize);
      }
    }
  }
</script>

<style scoped>
  .item:hover{
    background: #ebebeb;
    border-radius: 10px;
    font-family: 'Helvetica Neue', 'STHeiti', '微软雅黑', 'Microsoft YaHei', Helvetica,Arial,sans-serif;
  }
</style>
