<template>
  <div><a href="javascript:;" class="down-arrow" :class="{ 'down-arrow-up': open}"></a></div>
</template>

<script>
  export default {
    name: "DownArrow",
    props: ['open'],
    data(){
      return {
      }
    }
  }
</script>

<style scoped>
  .down-arrow{
    border:3px solid #666;
    /* [ border-top-width ] || [ border-right-width ] || [ border-bottom-width ] || [ border-left-width ] */
    border-width:0 3px 3px 0;
    display:inline-block;
    height:8px;
    -moz-transform-origin:center center;
    -moz-transform:rotate(45deg);
    -moz-transition:all .5s ease-in .1s;
    -webkit-transform-origin:center center;
    -webkit-transform:rotate(45deg);
    -webkit-transition:all .5s ease-in .1s;
    -o-transform-origin:center center;
    -o-transform:rotate(45deg);
    -o-transition:all .5s ease-in .1s;
    -ms-transform-origin:center center;
    -ms-transform:rotate(45deg);
    -ms-transition:all .5s ease-in .1s;
    transform-origin:center center;
    transform:rotate(45deg);
    transition:all .5s ease-in .1s;
    width:8px;
  }
  /* 箭头向上 180 度旋转 */
  .down-arrow-up{
    -moz-transform:rotate(-105deg);
    -moz-transition:all .5s ease-in .1s;
    -webkit-transform:rotate(315deg);
    -webkit-transition:all .5s ease-in .1s;
    -o-transform:rotate(-135deg);
    -o-transition:all .5s ease-in .1s;
    -ms-transform:rotate(-135deg);
    -ms-transition:all .5s ease-in .1s;
    transform:rotate(-135deg);
    transition:all .5s ease-in .1s;
  }
</style>
