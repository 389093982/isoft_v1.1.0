<template>
  <div style="background: #FFFFFF;padding: 10px;box-shadow: 2px 2px 1px #888888;">
    <Row>
      <Col v-if="courses.length > 0" span="16">
        <div style="border-bottom: 1px solid #ccc;padding: 15px;border-right: 1px solid #ccc;padding: 15px;" v-for="course in courses">
          <Row>
            <Col span="8">
              <router-link :to="{path:'/ilearning/course_detail',query:{course_id:course.id}}">
                <h6>课程名称：{{course.course_name}}</h6>
                <img v-if="course.small_image" :src="course.small_image" height="120" width="200"/>
                <img v-else src="../../../assets/default.png" height="90px" width="120px"/>
              </router-link>
            </Col>
            <Col span="16">
              <p style="color: #d6241e;">
                浏览量：{{course.watch_number}}
                课程分数：<Rate disabled show-text allow-half v-model="course.score"/> &nbsp;
              </p>
              <p>课程名称：{{course.course_name}}</p>
              <p>作者：{{course.course_author}}</p>
              <p>课程类型：{{course.course_type}}</p>
              <p>课程子类型：{{course.course_sub_type}}</p>
              <p>课程简介：{{course.course_short_desc}}</p>
              <p>课程集数：{{course.course_number}}</p>
              <p>课程更新状态：{{course.course_status}}</p>
            </Col>
          </Row>
        </div>
      </Col>
      <Col v-else span="16">
        未找到相应的搜索结果
      </Col>
      <Col span="8">
        <SiteDashBoard style="border: 1px solid #ccc;padding: 15px;margin-left: 5px;"/>
      </Col>
    </Row>
  </div>
</template>

<script>
  import {SearchCourseList} from "../../../api"
  import SiteDashBoard from "../Site/SiteDashBoard"

  export default {
    name: "CourseSearch",
    components:{SiteDashBoard},
    data(){
      return {
        courses:[],
      }
    },
    methods:{
      refreshCourseSearch:async function (search) {
        const result = await SearchCourseList(search);
        if(result.status=="SUCCESS"){
          this.courses = result.courses;
        }
      }
    },
    mounted:function () {
      // $route为当前router跳转对象里面可以获取name、path、query、params等
      // $router为VueRouter实例,想要导航到不同URL,则使用$router.push方法
      // 返回上一个history也是使用$router.go方法
      const search = this.$route.query.search;
      this.refreshCourseSearch(search);
    }
  }
</script>

<style scoped>
  a{
    color: black;
  }
  a:hover{
    color: red;
  }
</style>
