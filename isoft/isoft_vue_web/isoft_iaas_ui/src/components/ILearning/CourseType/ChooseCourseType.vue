<template>
  <div>
    <div style="margin-bottom: 10px;">
      <Search @submitFunc="searchFunc"/>
    </div>

    <div style="border: solid 1px #d9d9d9;">
      <div style="background-color: #f6f6f6;height: 38px;border-bottom: solid 1px #d9d9d9;">
        <p style="float: left;line-height: 38px;margin-left: 20px;">课程天地</p>
        <router-link to="/ilearning/course_space" style="float: right;line-height: 38px;
        margin-right: 20px;color: red;">>>进入我的课程空间</router-link>
      </div>
      <div style="padding: 20px;">
        <HotCourseType/>
        <TotalCourseType/>
      </div>
    </div>
  </div>
</template>

<script>
  import HotCourseType from "./HotCourseType.vue"
  import TotalCourseType from "./TotalCourseType.vue"
  import Search from "../../Common/Common/Search.vue"

  export default {
    name: "ChooseCourseType",
    components:{Search,HotCourseType,TotalCourseType},
    methods: {
      searchFunc:function (data) {
        this.$router.push({ name: '/ilearning/course_search', params: { search: data }});
      }
    }
  }
</script>

<style scoped>

</style>
