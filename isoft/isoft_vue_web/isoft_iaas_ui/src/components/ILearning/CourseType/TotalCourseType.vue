<template>
  <div style="margin-top: 25px;">
    <a href="javascript:;" @click="showAll=!showAll" style="color: red;">显示全部课程分类</a>
    <div v-if="showAll == true">
      <ul>
        <li v-for="course_type in course_types" style="margin:10px 10px 0 0;list-style:none;float: left;">
          <Poptip trigger="hover" :title="course_type" content="content" placement="bottom" @on-popper-show="loadSubCourseType(course_type)">
            <a href="javascript:;" style="font-size: 14px;" @click="submit(course_type)">
              {{course_type}}
            </a>

            <div class="api" slot="content">
              <ul>
                <li v-for="sub_course_type in sub_course_types"
                    style="padding-left: 10px;list-style:none;float: left;">
                  <a href="javascript:;" style="font-size: 14px;" @click="submit(sub_course_type)">
                    {{sub_course_type}}
                  </a>
                </li>
              </ul>
            </div>
          </Poptip>

        </li>
      </ul>
    </div>
    <div style="clear: both;"></div>
  </div>
</template>

<script>
  import {GetAllCourseType} from "../../../api"
  import {GetAllCourseSubType} from "../../../api"

  export default {
    name: "TotalCourseType",
    data(){
      return {
        showAll:false,
        course_types:[],
        sub_course_types:[],
      }
    },
    methods:{
      refreshCourseType:async function () {
        const result = await GetAllCourseType();
        if(result.status=="SUCCESS"){
          this.course_types = result.course_types;
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
