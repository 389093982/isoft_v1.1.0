<template>
  <div>
    <div style="margin-bottom: 10px;">
      <ISearch @submitFunc="searchFunc"/>
     </div>

    <div style="border: solid 1px #d9d9d9;">
      <div style="background-color: #f6f6f6;height: 38px;border-bottom: solid 1px #d9d9d9;">
        <p style="float: left;line-height: 38px;margin-left: 20px;">课程天地</p>
        <router-link to="/ilearning/course_space" style="float: right;line-height: 38px;margin-right: 20px;color: red;">
          >>我的课程空间
        </router-link>
      </div>
      <div style="padding: 20px;">
        <div style="border-bottom: 2px solid #edf1f2;">
          <a href="javascript:;" @click="showCourseType=true" style="color: red;">热门课程推荐</a>
          <a href="javascript:;" @click="showCourseType=!showCourseType" style="color: red;float: right;">
            <img src="../../../assets/images/common/more.jpg"/>
          </a>
        </div>
        <div>
          <HotCourseType v-show="showCourseType===true" @chooseCourseType="chooseCourseType"/>
          <TotalCourseType v-show="showCourseType===false" @chooseCourseType="chooseCourseType"/>
        </div>
      </div>

      <div style="text-align: right;padding: 10px;">
        <a href="javascript:;">
          <img src="../../../assets/images/common/free.jpg"/>
        </a>

      </div>
    </div>
  </div>
</template>

<script>
  import HotCourseType from "./HotCourseType"
  import TotalCourseType from "./TotalCourseType"
  import ISearch from "../../Common/search/ISearch"

  export default {
    name: "ChooseCourseType",
    components:{ISearch,HotCourseType,TotalCourseType},
    data(){
      return {
        showCourseType:true,
      }
    },
    methods: {
      searchFunc:function (data) {
        this.$router.push({ path: '/ilearning/course_search', query: { search: data }});
      },
      chooseCourseType:function (course_type, course_sub_type) {
        // params是路由的一部分
        // query是拼接在url后面的参数
        // 由于动态路由也是传递params的,所以在 this.$router.push() 方法中path不能和params一起使用,否则params将无效.需要用name来指定页面
        this.$router.push({ path: '/ilearning/course_search', query: { search: course_sub_type }});
      },
      toggle:function (data) {
        alert(data);
      }
    }
  }
</script>

<style scoped>

</style>
