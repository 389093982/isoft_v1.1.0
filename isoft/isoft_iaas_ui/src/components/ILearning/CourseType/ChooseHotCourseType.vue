<template>
  <span>
    <a href="javascript:;" @click="showCourseType = true">&nbsp;选择推荐分类</a>
    <Modal
      v-model="showCourseType"
      title="热门课程分类"
      width="850"
      :footer-hide="true"
      :styles="{top: '20px'}"
      :mask-closable="false">
      <div style="height: 450px;">
        <Scroll height="450">
          <HotCourseType @chooseCourseType="chooseCourseType"/>
        </Scroll>
      </div>
    </Modal>
  </span>
</template>

<script>
  import HotCourseType from "./HotCourseType"

  export default {
    name: "ChooseHotCourseType",
    components:{HotCourseType},
    data(){
      return {
        showCourseType:false,
      }
    },
    methods:{
      chooseCourseType:function (course_type, course_sub_type) {
        this.$emit("chooseCourseType", course_type, course_sub_type);
        this.showCourseType = false;
      }
    }
  }
</script>

<style scoped>

</style>
