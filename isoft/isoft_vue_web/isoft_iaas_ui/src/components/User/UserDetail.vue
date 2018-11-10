<template>
  <div v-if="userName" style="border: 1px #dbdbdb solid;margin-left: 5px;margin-bottom: 5px;padding: 15px;">
    <Tag color="primary">{{userName}}</Tag>
    <div style="margin-top: 5px;">
      <Tabs :animated="false">
        <TabPane label="作者课程">
          <ul>
            <li v-for="course in courses" style="list-style: none;">
              <Row>
                <Col span="8">
                  <Tag color="success">
                    <router-link :to="{path:'/ilearning/course_detail',query:{course_id:course.id}}">
                      {{course.course_name}}
                    </router-link>
                  </Tag>
                </Col>
                <Col span="8">
                  <Tag color="warning">{{course.course_type}}</Tag>
                </Col>
                <Col span="8">
                  <Tag color="purple">{{course.course_sub_type}}</Tag>
                </Col>
              </Row>
            </li>
          </ul>
        </TabPane>
        <TabPane label="作者博文">作者博文</TabPane>
        <TabPane label="作者博文">作者博文</TabPane>
      </Tabs>
    </div>
  </div>
</template>

<script>
  import {GetMyCourseList} from "../../api"

  export default {
    name: "UserDetail",
    props:{
      userName: {
        type: String,
        default: ''
      },
    },
    data(){
      return {
        // 当前 userName 的课程列表
        courses:[],
      }
    },
    methods:{
      refreshUserInfo:function () {
        this.refreshCourseList();
      },
      refreshCourseList:async function () {
        const result = await GetMyCourseList(this.userName);
        if(result.status=="SUCCESS"){
          this.courses = result.courses;
        }
      },
    },
    watch:{
      "userName": "refreshUserInfo"      // 如果 userName 有变化,会再次执行该方法
    }
  }
</script>

<style scoped>

</style>
