<template>
  <span><a target="_blank" :style="styles" :href="hrefaddr">{{msg}}</a></span>
</template>

<script>
  export default {
    name: "BeautifulButtonLink",
    props: {
      msg: {
        type: String,
        default: '点击了解详情'
      },
      floatstyle: {
        type: String,
        default: 'none'
      },
      marginrightstype: {
        type: Number,
        default: 30
      },
      hrefaddr: {
        type: String,
        default: "javascript:;"
      },
    },
    computed: {
      styles () {
        let style = {};
        style.float = this.floatstyle;
        style.marginRight = this.marginrightstype + 'px';
        return style;
      }
    },
  }
</script>

<style scoped>
  a{
    display: inline-block;
    width: 110px;
    height: 23px;
    line-height: 23px;
    color: #fff;
    background: #ff6225;
    text-align: center;
    text-decoration: none;
    cursor: pointer;
    font-size: 14px;
    font-family: "微软雅黑";
  }
</style>
