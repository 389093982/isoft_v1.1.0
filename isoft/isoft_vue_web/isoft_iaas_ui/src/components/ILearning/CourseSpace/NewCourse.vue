<template>
  <div>
    <Row>
      <Col span="16">
        <Form ref="formValidate" :model="formValidate" :rules="ruleValidate" :label-width="100">
          <FormItem label="课程名称" prop="course_name">
            <Input v-model="formValidate.course_name" placeholder="Enter course name..."/>
          </FormItem>
          <FormItem label="课程类型" prop="course_type">
            <Row>
              <Col span="19"><Input v-model="formValidate.course_type" placeholder="Enter course type..."></Input></Col>
              <Col span="5"><ChooseHotCourseType @chooseCourseType="chooseCourseType"/></Col>
            </Row>
          </FormItem>
          <FormItem label="课程子类型" prop="course_sub_type">
            <Input v-model="formValidate.course_sub_type" placeholder="Enter course sub type..."></Input>
          </FormItem>
          <FormItem label="课程描述" prop="course_short_desc">
            <Input v-model="formValidate.course_short_desc" placeholder="Enter course short desc..."></Input>
          </FormItem>
          <FormItem>
            <Button type="primary" @click="handleSubmit('formValidate')">Submit</Button>
            <Button style="margin-left: 8px" @click="handleReset('formValidate')">Cancel</Button>
          </FormItem>
        </Form>
      </Col>
      <Col span="8">
        课程发布规范说明：AAAA
      </Col>
    </Row>
  </div>
</template>

<script>
  import {NewCourse} from "../../../api"
  import ChooseHotCourseType from "../CourseType/ChooseHotCourseType.vue"

  export default {
    name: "NewCourse",
    components:{ChooseHotCourseType},
    data(){
      return {
        formValidate: {
          course_name: '',
          course_type: '',
          course_sub_type: "",
          course_short_desc:"",
        },
        ruleValidate: {
          course_name: [
            { required: true, message: '课程名称不能为空', trigger: 'blur' }
          ],
          course_type: [
            { required: true, message: '课程类型不能为空', trigger: 'blur' }
          ],
          course_sub_type: [
            { required: true, message: '课程子类型不能为空', trigger: 'blur' }
          ],
          course_short_desc: [
            { required: true, message: '课程描述不能为空', trigger: 'blur' }
          ],
        },
      }
    },
    methods:{
      handleSubmit (name) {
        var _this = this;
        this.$refs[name].validate(async (valid) => {
          if (valid) {
            const result = await NewCourse(_this.formValidate.course_name,
              _this.formValidate.course_type, _this.formValidate.course_sub_type, _this.formValidate.course_short_desc);
            if(result.status == "SUCCESS"){
              _this.$Message.success('提交成功!');
              _this.$router.go(0);     // 页面刷新,等价于 location.reload()
            }else{
              _this.$Message.error('提交失败!');
            }
          } else {
            _this.$Message.error('验证失败!');
          }
        })
      },
      handleReset (name) {
        this.$refs[name].resetFields();
      },
      chooseCourseType:function (course_type, course_sub_type) {
        this.formValidate.course_type = course_type;
        this.formValidate.course_sub_type = course_sub_type;
      }
    }
  }
</script>

<style scoped>

</style>
