<template>
  <div>
    <span style="height: 32px;line-height: 32px;margin-bottom: 5px;color: #000;float: left !important;">课程大类：</span>
    <ul style="overflow:hidden;">
        <li v-for="course_type in course_types"
            style="height: 32px;line-height: 32px;margin: 0 4px 5px;text-align: center;color: #333;float: left;display: inline;">
          <BeautifulLink @onclick="loadSubCourseType(course_type)">{{course_type}}</BeautifulLink>
        </li>
      </ul>
    <span style="height: 32px;line-height: 32px;margin-bottom: 5px;color: #000;float: left !important;">详细分类：</span>
    <ul style="overflow:hidden;">
        <li v-for="sub_course_type in sub_course_types"
            style="height: 32px;line-height: 32px;margin: 0 4px 5px;text-align: center;color: #333;float: left;display: inline;">
          <BeautifulLink @onclick="chooseCourseType(current_course_type, sub_course_type)">{{sub_course_type}}</BeautifulLink>
        </li>
      </ul>
  </div>
</template>

<script>
  import {GetAllCourseType} from "../../../api"
  import {GetAllCourseSubType} from "../../../api"
  import BeautifulLink from "../../Common/link/BeautifulLink.vue"

  export default {
    name: "TotalCourseType",
    components:{BeautifulLink},
    data(){
      return {
        course_types:[],
        current_course_type:'',
        sub_course_types:[],
      }
    },
    methods:{
      refreshCourseType:async function () {
        const result = await GetAllCourseType();
        if(result.status=="SUCCESS"){
          this.course_types = result.course_types;
          this.loadSubCourseType(result.course_types[0]);
        }
      },
      loadSubCourseType:async function(course_type){
        const result = await GetAllCourseSubType(course_type);
        if(result.status=="SUCCESS"){
          this.current_course_type = course_type;
          this.sub_course_types = result.sub_course_types;
        }
      },
      chooseCourseType:function (course_type, course_sub_type) {
        this.$emit("chooseCourseType", course_type, course_sub_type);
      }
    },
    mounted:function () {
      this.refreshCourseType();
    }
  }
</script>

<style scoped>
  a{
    color: #626262;
  }
  a:hover{
    color: red;
  }
</style>
