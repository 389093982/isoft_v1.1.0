<template>
  <div>
    <ul>
      <li v-for="history in historys">
        课程名称：{{history.history_desc}}
        链接地址：{{history.history_link}}
      </li>
    </ul>
  </div>
</template>

<script>
  import {ShowCourseHistory} from "../../../api"

  export default {
    name: "RecentlyViewed",
    data(){
      return {
        current_page:1,
        offset:10,
        total:0,
        historys:[],
      }
    },
    methods:{
      async refreshRecentlyViewed(){
        const data = await ShowCourseHistory(this.offset, this.current_page);
        if(data.status == "SUCCESS"){
          this.historys = data.historys;
          this.total = data.paginator.totalcount;
        }
      },
    },
    mounted:function () {
      this.refreshRecentlyViewed();
    }
  }
</script>

<style scoped>

</style>
