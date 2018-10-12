<template>
  <div>
      <ul style="overflow:hidden">
        <li v-for="course_type in course_types" style="margin:5px;padding:5px;list-style:none;
          float: left;background: rgba(219,167,255,0.34);">
          <a href="javascript:;" style="font-size: 14px;" @click="loadSubCourseType(course_type)">
            <span>{{course_type}}</span>
          </a>
        </li>
      </ul>
      <ul style="overflow:hidden">
        <li v-for="sub_course_type in sub_course_types"
            style="padding-left: 10px;list-style:none;float: left;">
          <a href="javascript:;" style="font-size: 14px;" @click="submit(sub_course_type)">
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
