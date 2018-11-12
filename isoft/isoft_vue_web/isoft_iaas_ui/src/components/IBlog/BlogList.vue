<template>
 <div>
   热门博客
   <ul>
     <li v-for="searchblog in searchblogs" style="list-style:none;padding: 10px 10px;background: #fff;border-bottom: 1px solid #f4f4f4;">
       <!-- 使用v-bind动态绑定id传递给目标路径 -->
       <router-link :to="{path:'/iblog/blog_detail',query:{blog_id:searchblog.id}}">
         <h5>{{searchblog.blog_title}}</h5>
       </router-link>
       <p style="margin-bottom: 4px;font-size: 14px;color: #8a8a8a;line-height: 24px;">
         {{searchblog.short_desc | filterLimitFunc}}
       </p>
       <p>
         <Row>
           <Col span="18">
             <!-- 作者详情 -->
             <router-link :to="{path:'/iblog/author',query:{author:searchblog.author}}">{{searchblog.author}}</router-link>
             发布于:<Time :time="searchblog.created_time" style="color:red;"/>&nbsp;
             更新于:<Time :time="searchblog.last_updated_time" style="color:red;"/>&nbsp;
             <router-link :to="{path:'/iblog/blog_list',query:{blog_id:searchblog.id}}">
               <span style="color: #3399ea;">{{ searchblog.catalog.catalog_name }}</span>
             </router-link>
           </Col>
           <Col span="3">
             <router-link :to="{path:'/iblog/blog_detail',query:{blog_id:searchblog.id}}">
               <span style="color: red;">{{searchblog.views}}</span>阅读
             </router-link>
           </Col>
           <Col span="3">
             <router-link :to="{path:'/iblog/blog_detail',query:{blog_id:searchblog.id}}">
              <a><span style="color: red;">0</span>条评论</a>
             </router-link>
           </Col>
         </Row>
       </p>
     </li>
   </ul>
 </div>
</template>

<script>
  import {BlogList} from "../../api"

  export default {
    name: "BlogList",
    data(){
      return {
        searchblogs:[],
      }
    },
    mounted:async function () {
      const result = await BlogList(10,1);
      if(result.status=="SUCCESS"){
        this.searchblogs = result.blogs;
      }
    },
    filters:{
      // 内容超长则显示部分
      filterLimitFunc:function (value) {
        if(value && value.length > 100) {
          value= value.substring(0,100) + '...';
        }
        return value;
      },
    }
  }
</script>

<style scoped>
  a{
    color: #3d3d3d;
  }
  a:hover{
    color: red;
  }
</style>
