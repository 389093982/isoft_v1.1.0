<template>
  <div>
    <div v-for="topic_reply in topic_replys" style="margin-bottom:5px;padding: 10px;border: 1px solid #e9e9e9;">
      <p><router-link to="">{{topic_reply.created_by}}</router-link></p>
      <p>
        回复<router-link to="">{{topic_reply.refer_user_name}}</router-link>:{{topic_reply.reply_content}}
        <span style="float: right;"><Time :time="topic_reply.created_time"/></span>
      </p>
      <p>
        <Row>
          <Col span="20">
            <a href="javascript:;" @click="refreshAndShow = !refreshAndShow">
              <Icon type="ios-arrow-round-down" />展开/隐藏子评论({{topic_reply.sub_reply_amount}})
            </a>
          </Col>
          <Col span="4" style="text-align: right;">
            <a href="javascript:;" @click="replyComment(topic_reply.id,topic_type.created_by)">回复他/她</a>&nbsp;
            <a href="javascript:;">点赞</a>
          </Col>
        </Row>
      </p>
      <!-- 递归,子评论区域 -->
      <CommentArea v-if="topic_reply.sub_reply_amount > 0 && refreshAndShow == true"
         :parent_id="topic_reply.id" :topic_id="topic_id" :topic_type="topic_type"/>
    </div>

    <!-- 评论表单 -->
    <Modal
      v-model="showCommentForm"
      width="800"
      title="回复"
      :mask-closable="false">
      <CommentForm v-if="showCommentForm" :parent_id="_parent_id" :topic_id="topic_id" :topic_type="topic_type"
        :refer_user_name="_refer_user_name" @refreshTopicReply="refreshTopicReply"/>
    </Modal>

  </div>
</template>

<script>
  import {FilterTopicReply} from "../../../api/index"
  import CommentForm from "./CommentForm.vue"

  export default {
    name: "CommentArea",
    // 评论清单
    props:["parent_id","topic_id","topic_type"],
    components:{CommentForm},
    data(){
      return {
        topic_replys:[],
        showCommentForm:false,
        // 回复评论,两个参数分别是被评论id,被评论人
        _parent_id:0,
        _refer_user_name:"",
        // 刷新和展开子组件评论
        refreshAndShow:false,
      }
    },
    methods:{
      // 刷新当前父级评论对应的评论列表
      refreshTopicReply:async function(){
        const result = await FilterTopicReply(this.topic_id, this.topic_type, this.parent_id);
        if(result.status=="SUCCESS"){
          this.showCommentForm = false;
          this.topic_replys = result.topic_replys;
          // 同时刷新其子组件对应的评论列表,之前是展开的才需要进行展开
          if(this.refreshAndShow == true){
            this.refreshAndShow = false;
            this.$nextTick(() => {this.refreshAndShow=true});
          }
        }
      },
      // 回复评论,两个参数分别是被评论id,被评论人
      replyComment:function(id,refer_user_name){
        this._parent_id = id;
        this._refer_user_name = refer_user_name;
        this.showCommentForm = true;
      },
    },
    mounted:function () {
      this.refreshTopicReply();
    }
  }
</script>

<style scoped>
  a{
    color:red;
  }
</style>
