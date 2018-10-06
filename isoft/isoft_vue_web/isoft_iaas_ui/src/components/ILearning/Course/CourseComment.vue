<template>
<div>
  <div style="padding: 10px;color: #9c1515;" v-if="topic_theme">
    <span>{{topic_theme.topic_content}}</span>
    <span style="float: right"><Time :time="topic_theme.created_time"/></span>
  </div>
  <div>
    <Row>
      <Col span="14" style="padding-right: 10px;">
        <Input v-model="submit_comment" type="textarea" :rows="8" placeholder="Enter something..." />
        <Button type="success" style="margin-top: 5px;float: right;" @click="submitComment">发表评论</Button>
      </Col>
      <Col span="10" style="border: 1px solid #e9e9e9;font-size:12px;padding: 10px;">
        <p>发表评论需知：</p>
        <p>1、请勿在评论中发表违法违规信息</p>
        <p>2、谢绝人身攻击、地域歧视、刷屏、广告等恶性言论</p>
        <p>3、所有评论均代表玩家本人意见，不代表官方立场</p>
        <p>4、用户发表的评论，经管理员审核后方可显示</p>
        <p>5、如果您有任何疑问，请在此以评论方式留言给我们</p>
      </Col>
    </Row>
  </div>
  <hr>
  <div>
    <!-- 评论列表 -->
    <SubComment v-if="this.course.id" parent_id="0" :topic_id="this.course.id" topic_type="course_topic_type"/>
  </div>
</div>
</template>

<script>
  import {FilterTopicTheme} from "../../../api"
  import {AddTopicReply} from "../../../api"

  import SubComment from "./SubComment.vue"

  export default {
    name: "CourseComment",
    components:{SubComment},
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
        refer_user_name:this.course.last_updated_by,
      }
    },
    methods:{
      // 刷新评论主题
      refreshTopicTheme:async function(){
        // topic_id, topic_type 分别如下参数
        const result = await FilterTopicTheme(this.course.id, "course_topic_type");
        if(result.status=="SUCCESS"){
          this.topic_theme = result.topic_theme;
          // this.refreshTopicReply();
        }
      },
      submitComment: async function () {
        var parent_id = this.parent_id;
        var reply_content = this.submit_comment;
        var topic_id = this.course.id;
        var topic_type = "course_topic_type";
        var refer_user_name = this.refer_user_name;
        const result = await AddTopicReply(parent_id, reply_content, topic_id, topic_type, refer_user_name);
        if(result.status=="SUCCESS"){
          // this.refreshTopicReply();
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
