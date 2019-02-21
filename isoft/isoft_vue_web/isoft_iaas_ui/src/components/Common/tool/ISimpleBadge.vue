<template>
  <span class="badge" :style="mapStyle"></span>
</template>

<script>
  export default {
    name: "ISimpleBadge",
    props: {
      marginRightStyle: {
        type: Number,
        default: 0,
      },
      backgroundColorStyle: {
        type: String,
        default: 'red',
      },
    },
    data(){
      return {
        mapStyle:{
          marginRight: `${this.marginRightStyle}px`,
          backgroundColor:this.backgroundColorStyle,
        }
      }
    }
  }
</script>

<style scoped>
.badge{
  border-radius: 50%;
  width: 8px;
  height: 8px;
  display: inline-block;
}
</style>
