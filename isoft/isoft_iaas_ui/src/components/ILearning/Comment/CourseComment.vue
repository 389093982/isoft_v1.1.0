<template>
<div>
  <div style="padding: 10px;color: #9c1515;" v-if="comment_theme">
    <span>{{comment_theme.comment_content}}</span>
    <span style="float: right"><Time :time="comment_theme.created_time"/></span>
  </div>

  <div>
    <CommentForm :parent_id="parent_id" :comment_id="course.id" :theme_type="theme_type"
       :refer_user_name="refer_user_name" @refreshCommentReply="refreshCommentReply"/>
  </div>

  <div style="margin-top: 30px;">
    <div class="comment_title">
      <span style="text-align: right;">
        <a href="javascript:;" @click="refreshCommentReply('all')">全部</a>
        <Divider type="vertical" />
        <a href="javascript:;" @click="refreshCommentReply('comment')">评论</a>
        <Divider type="vertical" />
        <a href="javascript:;" @click="refreshCommentReply('question')">提问</a>
      </span>
      <span style="float: right;">
        <a href="javascript:;">更多评论话题</a>
      </span>
    </div>
    <!-- 评论列表 -->
    <CommentArea ref="commentArea" v-if="this.course.id" parent_id="0" :comment_id="this.course.id" :theme_type="theme_type"></CommentArea>
  </div>
</div>
</template>

<script>
  import {FilterCommentTheme} from "../../../api/index"
  import CommentArea from "./CommentArea"
  import CommentForm from "./CommentForm"

  export default {
    name: "CourseComment",
    components:{CommentForm,CommentArea},
    // 当前评论的课程
    props:["course"],
    data(){
      return {
        comment_theme:null,
        // 父评论 id
        parent_id:0,
        // 提交评论内容
        submit_comment:"",
        // 被评论人
        refer_user_name:"",
        theme_type:"course_theme_type",
      }
    },
    methods:{
      // 重新刷新评论列表
      refreshCommentReply (reply_comment_type) {
        // 调用子组件的刷新方法
        this.$refs.commentArea.refreshCommentReply(reply_comment_type);
      },
      // 刷新评论主题
      refreshCommentTheme:async function(){
        // comment_id, theme_type 分别如下参数
        const result = await FilterCommentTheme(this.course.id, this.theme_type);
        if(result.status=="SUCCESS"){
          this.comment_theme = result.comment_theme;
          this.refreshCommentReply();
        }
      },
    },
    watch:{
      // 监听 props 修改
      course(curVal,oldVal){
        this.refer_user_name = curVal.course_author;
        this.refreshCommentTheme();
      },
    },
    mounted:function () {
      if(this.course && this.course.id){
        // 父组件异步修改子组件 props 值获取了 undefined
        this.refreshCommentTheme();
      }
    }
  }
</script>

<style scoped lang="stylus" rel="stylesheet/stylus">
  .comment_title
    a
      color: #333
      text-decoration: none
      font-weight: 600
    a:hover
      color:red
</style>
