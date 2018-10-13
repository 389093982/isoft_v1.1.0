<template>
  <span>
    <a href="javascript:;" @click="showCourseType = true"><Icon type="md-book" />&nbsp;选择推荐分类</a>
    <Modal
      v-model="showCourseType"
      title="热门课程分类"
      width="850"
      :footer-hide="true"
      :styles="{top: '20px'}"
      :mask-closable="false">
      <div style="height: 450px;">
        <Scroll height="450">
          <HotCourseType/>
        </Scroll>
      </div>
    </Modal>
  </span>
</template>

<script>
  import HotCourseType from "./HotCourseType.vue"

  export default {
    name: "ChooseHotCourseType",
    components:{HotCourseType},
    data(){
      return {
        showCourseType:false,
      }
    },
  }
</script>

<style scoped>

</style>
