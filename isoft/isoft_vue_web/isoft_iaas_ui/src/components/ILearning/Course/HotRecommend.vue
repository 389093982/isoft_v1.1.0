<template>
  <div>
    <!-- 列表形式显示 -->
    <div v-if="showMode=='list'" style="border: 1px #dbdbdb solid;margin-left: 5px;margin-bottom: 5px;">
      <Card title="热门课程推荐" icon="ios-options" :padding="0" shadow style="width: 100%;">
        <CellGroup>
          <Cell v-for="course in courses" :title="course.course_name" :to="{path:'/ilearning/course_detail',query:{course_id:course.id}}">
            {{course.course_name}}
            <Badge text="查看" slot="extra" />
          </Cell>
        </CellGroup>
      </Card>
    </div>
    <!-- 图标形式显示 -->
    <IBeautifulCard v-else title="热门课程推荐">
      <div slot="content" style="min-height:850px;padding: 20px;">
        <ul class="clear">
          <li v-for="course in courses">
            <router-link :to="{path:'/ilearning/course_detail',query:{course_id:course.id}}">
              <img v-if="course.small_image" :src="course.small_image" height="90px" width="120px"/>
              <img v-else src="../../../assets/default.png" height="90px" width="120px"/>
              <p>{{course.course_name}}</p>
            </router-link>
          </li>
        </ul>
      </div>
    </IBeautifulCard>
  </div>
</template>

<script>
  import IBeautifulCard from "../../Common/card/IBeautifulCard"
  import {GetHotCourseRecommend} from "../../../api"


  export default {
    name: "HotRecommend",
    components:{IBeautifulCard},
    props:{
      // 显示方式,支持 detail 和 list
      showMode: {
        type: [String],
        default: 'detail'
      },
    },
    data(){
      return {
        courses:[],
      }
    },
    methods:{
      refreshHotRecommend:async function(){
        const result = await GetHotCourseRecommend();
        if(result.status == "SUCCESS"){
          this.courses = result.courses;
        }
      }
    },
    mounted:function () {
      this.refreshHotRecommend();
    }
  }
</script>

<style scoped>
  /* 引入公共样式库 */
  @import "../../../../static/css/common.css";

  a{
    color: black;
  }
  li{
    float: left;
    padding: 10px 9px 0;
    width: 140px;
    height: 125px;
    overflow: hidden;
    text-align: center;
    position: relative;
  }
  li:hover{
    background-color: #f4f4f4;
    border: 1px solid #d0cdd2;
  }
  li:hover a{
    color:red;
  }
</style>
