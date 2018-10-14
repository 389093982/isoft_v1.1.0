<template>
  <div v-if="blog">
    <h3>{{blog.blog_title}}</h3>
    <div style="border-bottom: 1px solid #f4f4f4;margin-top:20px;margin-bottom: 20px;">
      <Row>
        <Col span="18">
          发布于:<Time :time="blog.created_time" style="color:red;"/>&nbsp;
          更新于:<Time :time="blog.last_updated_time" style="color:red;"/>&nbsp;
        </Col>
        <Col span="3">阅读次数 {{blog.views}}</Col>
        <Col span="3">编辑次数 0</Col>
      </Row>
    </div>
    <div class="article" v-html="compiledMarkdown"></div>
  </div>
</template>

<script>
  import {ShowBlogDetail} from "../../api"
  // 导入 showdown
  import showdown from 'showdown';

  export default {
    name: "BlogDetail",
    data(){
      return {
        blog: null,
      }
    },
    methods:{
      refreshBlogDetail:async function () {
        const result = await ShowBlogDetail(this.$route.query.blog_id);
        if(result.status=="SUCCESS"){
          this.blog = result.blog;
        }
      }
    },
    computed:{
      compiledMarkdown () {
        let converter = new showdown.Converter();
        let html = converter.makeHtml(this.blog.content);
        return html;
      }
    },
    mounted:function () {
      this.refreshBlogDetail();
    }
  }
</script>

<style scoped>

</style>
