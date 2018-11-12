<template>
  <div>
    AAA
  </div>
</template>

<script>
  import {ShowShareDetail} from "../../api"

  export default {
    name: "ShareDetail",
    data(){
      return {
        share:null,
      }
    },
    methods:{
      refreshShareDetail:async function () {
        const result = await ShowShareDetail(this.$route.query.share_id);
        if(result.status=="SUCCESS"){
          this.share = result.share;
        }
      }
    },
    mounted:function () {
      this.refreshShareDetail();
    }
  }
</script>

<style scoped>

</style>
