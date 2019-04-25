<template>
  <div style="padding-top: 20px;">
    <div v-for="history in historys" style="padding: 0 20px 0 20px;">
      <router-link :to="{path:'/ilearning/course_detail',query:{course_id:history.history_link}}">{{history.history_desc}}</router-link>
      <span style="float: right;font-size: 12px;"><Time :time="history.last_updated_time" :interval="1"/></span>
      <Divider />
    </div>
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
      searchFunc:function (data) {
        this.$router.push({ path: '/ilearning/course_search', query: { search: data }});
      },
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
