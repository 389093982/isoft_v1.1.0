<template>
  <div>
    <div v-for="comment_reply in comment_replys" style="margin-bottom:5px;padding: 10px;border: 1px solid #e9e9e9;">
      <p>
        <router-link to="">
          <Avatar size="small" src="https://i.loli.net/2017/08/21/599a521472424.jpg" />
          {{comment_reply.created_by}}
        </router-link>
      </p>
      <p>
        <span v-if="comment_reply.reply_comment_type == 'comment'">
          回复
        </span>
        <span v-else>
          提问
        </span>
        <router-link to="">
        <Avatar size="small" src="https://i.loli.net/2017/08/21/599a521472424.jpg" />
        {{comment_reply.refer_user_name}}
        </router-link>
        :{{comment_reply.reply_content}}
        <span style="float: right;"><Time :time="comment_reply.created_time" :interval="1"/></span>
      </p>
      <p>
        <Row>
          <Col span="20">
            <a href="javascript:;" @click="refreshAndShow = !refreshAndShow">
              <Icon type="ios-arrow-round-down" />展开/隐藏子评论({{comment_reply.sub_reply_amount}})
            </a>
          </Col>
          <Col span="4" style="text-align: right;">
            <a v-if="comment_reply.depth < 4" href="javascript:;" @click="replyComment(comment_reply.id,comment_reply.created_by)">回复他/她</a>&nbsp;
            <a href="javascript:;">点赞</a>
          </Col>
        </Row>
      </p>
      <!-- 递归,子评论区域 -->
      <CommentArea v-if="comment_reply.sub_reply_amount > 0 && refreshAndShow == true"
         :parent_id="comment_reply.id" :comment_id="comment_id" :theme_type="theme_type"/>

    </div>

    <!-- 评论表单 -->
    <Modal
      v-model="showCommentForm"
      width="800"
      title="回复"
      :mask-closable="false">
      <CommentForm v-if="showCommentForm" :parent_id="_parent_id" :comment_id="comment_id" :theme_type="theme_type"
        :refer_user_name="_refer_user_name" @refreshCommentReply="refreshCommentReply"/>
    </Modal>

  </div>
</template>

<script>
  import {FilterCommentReply} from "../../../api/index"
  import CommentForm from "./CommentForm.vue"

  export default {
    name: "CommentArea",
    // 评论清单
    props:["parent_id","comment_id","theme_type"],
    components:{CommentForm},
    data(){
      return {
        comment_replys:[],
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
      refreshCommentReply:async function(reply_comment_type){
        if(reply_comment_type == undefined){
          reply_comment_type = "all";
        }
        const result = await FilterCommentReply(this.comment_id, this.theme_type, this.parent_id, reply_comment_type);
        if(result.status=="SUCCESS"){
          this.showCommentForm = false;
          this.comment_replys = result.comment_replys;
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
      this.refreshCommentReply('all');
    }
  }
</script>

<style scoped>
  a{
    color:red;
  }
</style>
