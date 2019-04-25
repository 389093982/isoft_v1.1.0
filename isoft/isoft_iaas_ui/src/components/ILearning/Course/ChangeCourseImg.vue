<template>
  <div>
    <Modal
      v-model="showDialog"
      width="500"
      title="上传/更新课程图标"
      :mask-closable="false">
      <div>
        <img v-if="course.small_image" :src="course.small_image" height="120" width="200"/>
        <img v-else src="../../../assets/default.png" height="120" width="200"/>
        <!-- 属性拼接 -->
        <Upload :action="'/api/ilearning/changeCourseImg?id=' + course.id" :on-success="uploadComplete">
          <Button icon="ios-cloud-upload-outline">Upload files</Button>
        </Upload>
      </div>
    </Modal>

    <a href="javascript:;" style="color: green;" @click="showDialog=true">换张图片</a>
  </div>
</template>

<script>
  export default {
    name: "ChangeCourseImg",
    // 当前需要更改图标的课程
    props:["course"],
    data(){
      return {
        showDialog:false,
      }
    },
    methods:{
      uploadComplete(res, file) {
        if(res.status=="SUCCESS"){
          this.$emit('uploadComplete');
        }else{
          this.$Notice.error({
            title: '文件上传失败',
            desc: '文件 ' + file.name + ' 上传失败!'
          });
        }
      },
    }
  }
</script>

<style scoped>

</style>
