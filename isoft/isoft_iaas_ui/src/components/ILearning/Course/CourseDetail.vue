<template>
  <div style="background: #FFFFFF;padding: 10px;box-shadow: 2px 2px 1px #888888;">
    <Row>
      <!-- 左侧课程详情部分 -->
      <Col span="16">
        <!-- 头部 -->
        <Row>
          <Col span="8">
            <h6>课程名称：{{course.course_name}}</h6>
            <p>
              <img v-if="course.small_image" :src="course.small_image" height="120" width="200"/>
              <img v-else src="../../../assets/default.png" height="120" width="200"/>
            </p>
          </Col>
          <Col span="16">
            <p style="color: #d6241e;">
              浏览量：{{course.watch_number}}
              课程分数：<Rate disabled show-text allow-half v-model="course.score"/> &nbsp;
            </p>
            <p>
              课程名称：{{course.course_name}}
              <router-link to="/ilearning/index" style="margin-left: 30px;">搜索同类资源</router-link>
            </p>
            <p>作者：{{course.course_author}}</p>
            <p>课程类型：{{course.course_type}}</p>
            <p>课程子类型：{{course.course_sub_type}}</p>
            <p>课程简介：{{course.course_short_desc}}</p>
            <p>课程集数：{{course.course_number}}</p>
            <p>课程更新状态：{{course.course_status}}</p>
            <p>
              <a href="javascript:;" v-if="course_collect==true" @click="toggle_favorite(course.id,'course_collect')">取消收藏</a>
              <a href="javascript:;" v-else @click="toggle_favorite(course.id,'course_collect')">加入收藏</a>&nbsp;
              <a href="javascript:;" v-if="course_parise==true" @click="toggle_favorite(course.id,'course_praise')">取消点赞</a>
              <a href="javascript:;" v-else @click="toggle_favorite(course.id,'course_praise')">我要点赞</a>
            </p>
          </Col>
        </Row>
        <hr style="margin-top: 10px;">
        <!-- 视频链接 -->
        <Row style="margin: 10px 0 10px 0">
          <Col span="12" v-for="cVideo in cVideos" style="padding: 5px;">
            <Row>
              <Col span="2">{{cVideo.video_number}}</Col>
              <Col span="18">{{cVideo.video_name}}</Col>
              <Col span="4">
                <router-link :to="{path:'/ilearning/video_play',query:{course_id:course.id,video_id:cVideo.id}}">
                  <Button size="small" type="success">立即播放</Button>
                </router-link>
              </Col>
            </Row>
          </Col>
        </Row>
        <hr>
        <!-- 课程评论 -->
        <CourseComment :course="course" style="margin-top: 50px;"/>
      </Col>

      <Col span="8">
        <UserDetail :userName="course.course_author"/>
        <HotUser/>
        <HotRecommend showMode="list"/>
        <!-- 推荐系统 -->
        <Recommand />
        <CommunicationGroup/>
      </Col>
    </Row>
  </div>
</template>

<script>
  import {ShowCourseDetail} from "../../../api"
  import {ToggleFavorite} from "../../../api"
  import Recommand from "./Recommand"
  import CourseComment from "../Comment/CourseComment"
  import CommunicationGroup from "../Site/CommunicationGroup"
  import HotRecommend from "./HotRecommend"
  import UserDetail from "../../User/UserDetail"
  import HotUser from "../../User/HotUser"

  export default {
    name: "CourseDetail",
    components:{CourseComment,Recommand,CommunicationGroup,HotRecommend,UserDetail,HotUser},
    data(){
      return {
        // 当前课程
        course:{},
        // 视频清单
        cVideos:[],
        // 课程收藏
        course_collect:false,
        // 课程点赞
        course_parise:false,
      }
    },
    methods:{
      refreshCourseDetail:async function(){
        const course_id = this.$route.query.course_id;
        const result = await ShowCourseDetail(course_id);
        if(result.status=="SUCCESS"){
          this.course = result.course;
          this.cVideos = result.cVideos;
          this.course_collect = result.course_collect;
          this.course_parise = result.course_parise;
        }
      },
      toggle_favorite:async function (favorite_id, favorite_type) {
        const result = await ToggleFavorite(favorite_id, favorite_type);
        if(result.status=="SUCCESS"){
          this.refreshCourseDetail();
        }
      }
    },
    mounted:function () {
      this.refreshCourseDetail(this.$route.query.course_id);
    },
    watch:{
      "$route.params": "refreshCourseDetail"      // 如果 $route.params 有变化,会再次执行该方法
    }
  }
</script>

<style scoped>
  a{
    color: red;
  }
</style>
