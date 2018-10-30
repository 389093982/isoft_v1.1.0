<template>
  <div style="margin-left: 50px;margin-right: 50px;margin-bottom: 20px;">
    <Row>
      <Col span="2" style="text-align: center;padding-top: 25px;">
        <a href="javascript:;" @click="previous" v-show="showPrevious"><img src="/static/images/arrow_left.png"/></a>
      </Col>
      <Col span="5" v-for="share_type in getCurrentPage">
        <a href="javascript:;" style="color: #999;">
          <div class="item" style="padding:10px; height: 100px;">
            <Row>
              <Col span="6">
                <img :src="share_type.img" :alt="share_type.name"/>
              </Col>
              <Col span="18" style="padding-left: 5px;">
                <p class="share_type_name">{{share_type.name}}</p>
                <p style="font-size: 12px;">{{share_type.detail}}</p>
              </Col>
            </Row>
          </div>
        </a>
      </Col>
      <Col span="2" style="text-align: center;padding-top: 25px;">
        <a href="javascript:;" @click="next"  v-show="showNext"><img src="/static/images/arrow_right.png"/></a>
      </Col>
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
      showPrevious:function () {
        const total_page = Math.ceil(this.hot_share_type.length / 4);
        return this.currentPageNo > 1;
      },
      showNext:function () {
        const total_page = Math.ceil(this.hot_share_type.length / 4);
        return this.currentPageNo < total_page;
      }
    },
    methods:{
      previous: function(){
        if(this.currentPageNo > 1){
          this.currentPageNo = this.currentPageNo - 1;
        }
      },
      next: function(){
        const total_page = Math.ceil(this.hot_share_type.length / 4);
        if(this.currentPageNo < total_page){
          this.currentPageNo = this.currentPageNo + 1;
        }
      },
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
