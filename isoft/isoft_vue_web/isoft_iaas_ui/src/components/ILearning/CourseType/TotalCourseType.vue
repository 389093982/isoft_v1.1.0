<template>
  <div>
      <ul style="overflow:hidden">
        <li v-for="course_type in course_types" style="list-style:none;float: left;">
          <a href="javascript:;" style="margin:5px;padding-left:15px;padding-right:15px;font-size: 14px;
             display: inline-block;color: #fff;background: rgba(34,195,0,0.55);"
             @click="loadSubCourseType(course_type)">
            <span>{{course_type}}</span>
          </a>
        </li>
      </ul>
      <ul style="overflow:hidden">
        <li v-for="sub_course_type in sub_course_types" style="list-style:none;float: left;">
          <a href="javascript:;" style="margin:5px;padding-left:15px;padding-right:15px;font-size: 14px;
             display: inline-block;color: #fff;background: rgba(255,98,37,0.55);"
             @click="submit(sub_course_type)">
            {{sub_course_type}}
          </a>
        </li>
      </ul>
  </div>
</template>

<script>
  import {GetAllCourseType} from "../../../api"
  import {GetAllCourseSubType} from "../../../api"

  export default {
    name: "TotalCourseType",
    data(){
      return {
        course_types:[],
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
          this.sub_course_types = result.sub_course_types;
        }
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
