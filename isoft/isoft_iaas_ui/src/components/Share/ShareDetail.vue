<template>
  <div v-if="share">
    <IShowMarkdown v-if="share.content" :content="share.content"/>
  </div>
</template>

<script>
  import {ShowShareDetail} from "../../api"
  import IShowMarkdown from "../Common/markdown/IShowMarkdown"

  export default {
    name: "ShareDetail",
    components:{IShowMarkdown},
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
