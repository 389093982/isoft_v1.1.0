<template>
  <div style="padding: 10px;margin-top: 50px;">
    <p style="color: #9c1515;">
      {{topic_theme.topic_content}}
    </p>
  </div>
</template>

<script>
  import {FilterTopicTheme} from "../../../api"

  export default {
    name: "CourseComment",
    // 当前评论的课程
    props:["course"],
    data(){
      return {
        topic_theme:{},
      }
    },
    methods:{
      refreshTopicTheme:async function(){
        // topic_id, topic_type 分别如下参数
        const result = await FilterTopicTheme(this.course.id, "course_topic_type");
        if(result.status=="SUCCESS"){
          this.topic_theme = result.topic_theme;
        }
      }
    },
    watch:{
      // 监听 props 修改
      course(curVal,oldVal){
        this.refreshTopicTheme();
      },
    },
    mounted:function () {
      if(this.course && this.course.id){
        // 父组件异步修改子组件 props 值获取了 undefined
        this.refreshTopicTheme();
      }
    }
  }
</script>

<style scoped>

</style>
