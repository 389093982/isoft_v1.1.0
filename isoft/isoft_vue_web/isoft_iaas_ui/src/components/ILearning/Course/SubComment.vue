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
            <a href="javascript:;"><Icon type="ios-arrow-round-down" />查看子评论({{topic_reply.sub_reply_amount}})</a>
          </Col>
          <Col span="4" style="text-align: right;">
            <a href="javascript:;">回复他/她</a>&nbsp;
            <a href="javascript:;">点赞</a>
          </Col>
        </Row>
      </p>
      <SubComment v-if="topic_reply.sub_reply_amount > 0" :parent_id="topic_reply.id" :topic_id="topic_id" :topic_type="topic_type"/>
    </div>
  </div>
</template>

<script>
  import {FilterTopicReply} from "../../../api"

  export default {
    name: "SubComment",
    // 评论清单
    props:["parent_id","topic_id","topic_type"],
    data(){
      return {
        topic_replys:[],
      }
    },
    methods:{
      // 刷新评论列表
      refreshTopicReply:async function(){
        const result = await FilterTopicReply(this.topic_id, this.topic_type, this.parent_id);
        if(result.status=="SUCCESS"){
          this.topic_replys = result.topic_replys;
        }
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
