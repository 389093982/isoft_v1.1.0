<template>
<div>
  <div style="padding: 10px;color: #9c1515;" v-if="topic_theme">
    <span>{{topic_theme.topic_content}}</span>
    <span style="float: right"><Time :time="topic_theme.created_time"/></span>
  </div>
  <div>
    <CommentForm :parent_id="parent_id" :topic_id="course.id" :topic_type="topic_type"
       :refer_user_name="refer_user_name" @refreshTopicReply="refreshTopicReply"/>
  </div>
  <hr>
  <div>
    <!-- 评论列表 -->
    <CommentArea ref="_commentArea" v-if="this.course.id" parent_id="0" :topic_id="this.course.id" :topic_type="topic_type"/>
  </div>
</div>
</template>

<script>
  import {FilterTopicTheme} from "../../../api/index"
  import CommentArea from "./CommentArea.vue"
  import CommentForm from "./CommentForm.vue"

  export default {
    name: "CourseComment",
    components:{CommentForm,CommentArea},
    // 当前评论的课程
    props:["course"],
    data(){
      return {
        topic_theme:null,
        // 父评论 id
        parent_id:0,
        // 提交评论内容
        submit_comment:"",
        // 被评论人
        refer_user_name:"",
        topic_type:"course_topic_type",
      }
    },
    methods:{
      // 重新刷新评论列表
      refreshTopicReply () {
        // 调用子组件的刷新方法
        this.$refs._commentArea.refreshTopicReply();
      },
      // 刷新评论主题
      refreshTopicTheme:async function(){
        // topic_id, topic_type 分别如下参数
        const result = await FilterTopicTheme(this.course.id, this.topic_type);
        if(result.status=="SUCCESS"){
          this.topic_theme = result.topic_theme;
          this.refreshTopicReply();
        }
      },
    },
    watch:{
      // 监听 props 修改
      course(curVal,oldVal){
        this.refer_user_name = curVal.course_author;
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
