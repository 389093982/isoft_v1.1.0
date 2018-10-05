<template>
  <span>
    <Modal
      v-model="showDialog"
      width="500"
      title="上传/更新视频"
      :mask-closable="false">
      <div>
        <p>课程名称：{{course.course_name}}</p>

        <div style="margin-top: 20px;margin-bottom: 20px;">
          <Button v-if="course.course_number" v-for="num in course.course_number"
                  type="success" @click="uploadVedioNum = num" style="margin: 5px;">第{{num}}集</Button>
          <Button type="success" @click="uploadVedioNum = course.course_number + 1">新一集{{course.course_number + 1}}</Button>
        </div>

        <Upload :action="'/api/ilearning/uploadVedio?id=' + course.id + '&vedio_number=' + uploadVedioNum" :on-success="uploadComplete">
          <Button icon="ios-cloud-upload-outline">Upload files</Button>
          <span v-if="uploadVedioNum > 0" style="color: #3300ff;">*当前更新第{{uploadVedioNum}}集</span>
          <span v-else style="color: red;">请选择更新集数</span>
        </Upload>
      </div>
    </Modal>

    <a href="javascript:;" style="color:#f55e13;font-family: Arial;font-weight: 700;" @click="showDialog=true">上传/更新视频</a>
  </span>
</template>

<script>
  export default {
    name: "UploadVedio",
    // 当前需要上传视频的课程
    props:["course"],
    data(){
      return {
        showDialog:false,
        // 当前更新视频集数
        uploadVedioNum:-1,
      }
    },
    methods:{
      uploadComplete(res, file) {
        if(res.status=="SUCCESS"){
          this.$Notice.success({
            title: '视频上传',
            desc: "上传成功!"
          });
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
