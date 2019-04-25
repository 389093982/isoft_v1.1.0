<template>
  <div style="width: 88%; height: 88%; margin: 0 auto;box-shadow: 2px 2px 1px #888888;">
    <video ref="video" width="100%" height="100%" controls autoplay="autoplay">
      <source type="video/mp4">
      <source type="video/ogg">
      您的浏览器不支持Video标签。
    </video>
  </div>
</template>

<script>
  import {ShowCourseDetail} from "../../../api"

  export default {
    name: "VideoPay",
    data(){
      return {
      }
    },
    methods:{
      playVideo:async function () {
        const result = await ShowCourseDetail(this.$route.query.course_id);
        if(result.status=="SUCCESS"){
          for(var i=0; i<result.cVideos.length; i++){
            if(result.cVideos[i].id == this.$route.query.video_id){
              this.$refs.video.src = result.cVideos[i].first_play;
            }
          }
        }
      }
    },
    mounted:function () {
      this.playVideo();
    }
  }
</script>

<style scoped>

</style>
