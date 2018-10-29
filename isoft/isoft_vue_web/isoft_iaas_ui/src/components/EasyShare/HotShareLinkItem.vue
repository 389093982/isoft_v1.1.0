<template>
  <div style="margin-left: 50px;margin-right: 50px;">
    <Row :gutter="10">
      <Col span="6" v-for="share_type in getCurrentPage">
        <a href="javascript:;" style="color: #999;">
          <div class="item" style="margin:10px; padding:10px; height: 100px;">
            <Row>
              <Col span="6">
                <img src="https://gocn.vip/uploads/nav_menu/12.jpg" alt="每日新闻">
                <!--<img :src="share_type.img" :alt="share_type.name"/>-->
              </Col>
              <Col span="18">
                <p class="share_type_name">{{share_type.name}}</p>
                <p style="font-size: 12px;">{{share_type.detail}}</p>
              </Col>
            </Row>
          </div>
        </a>
      </Col>
    </Row>
    <Row>
      <a href="javascript:;" style="float: right;margin: 0px 30px 10px 0;font-size: 12px;">换一批</a>
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
  *{
    font-family: 'Helvetica Neue', 'STHeiti', '微软雅黑', 'Microsoft YaHei', Helvetica,Arial,sans-serif;
  }
  .item:hover{
    background: #ebebeb;
    border-radius: 10px;
  }
  .share_type_name{
    color:green;
  }
  .share_type_name:hover{
    color:red;
  }
</style>
