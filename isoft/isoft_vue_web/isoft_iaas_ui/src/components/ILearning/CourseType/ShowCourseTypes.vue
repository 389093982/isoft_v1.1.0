<template>
  <span>
    <Icon type="md-expand" style="color: red;"/><a href="javascript:;" @click="toggleShow" style="color:black;">查看更多课程分类</a>&nbsp;
    <Icon type="ios-home" style="color: red;"/><router-link to="/ilearning/course_space" style="color: black;">我的课程空间</router-link>

    <div v-if="showCourseType==true" style="border:2px solid #b3d5a8;min-height: 150px;padding:10px;">ASAAA</div>
    <ChooseCourseType/>
  </span>
</template>

<script>
  import ChooseCourseType from "./ChooseCourseType.vue"

  export default {
    name: "ShowCourseTypes",
    components:{ChooseCourseType},
    data(){
      return {
        showCourseType:false,
      }
    },
    methods:{
      toggleShow(){
        this.showCourseType = !this.showCourseType;
      }
    }
  }
</script>

<style scoped>

</style>
