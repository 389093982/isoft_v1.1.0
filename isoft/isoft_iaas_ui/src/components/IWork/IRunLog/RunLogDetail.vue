<template>
  <span>
    <p v-for="runLogDetail in runLogDetails">
      <span v-html="runLogDetail.detail"></span>
    </p>
  </span>
</template>

<script>
  import {GetLastRunLogDetail} from "../../../api"

  export default {
    name: "RunLogDetail",
    data(){
      return {
        runLogDetails:[],
      }
    },
    methods:{
      refreshRunLogDetail:async function () {
        const result = await GetLastRunLogDetail(this.$route.query.tracking_id);
        if(result.status=="SUCCESS"){
          this.runLogDetails = result.runLogDetails;
        }
      }
    },
    mounted(){
      this.refreshRunLogDetail();
    }
  }
</script>

<style scoped>

</style>
