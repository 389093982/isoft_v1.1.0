<template>
  <div>
    <p>
      <a href="javascript:;">全部分享</a>
      <a href="javascript:;">最新分享</a>
      <a href="javascript:;">我的分享</a>
      <ShareLinkAdd/>
    </p>
    <span v-for="shareLink in shareLinks">
      {{shareLink.share_type}}
      {{shareLink.author}}
      {{shareLink.link_href}}
      {{shareLink.last_updated_time}}
    </span>
  </div>
</template>

<script>
  import {FilterShareLinkList} from "../../api"
  import ShareLinkAdd from "./ShareLinkAdd.vue"

  export default {
    name: "ShareLinkList",
    components:{ShareLinkAdd},
    data(){
      return {
        shareLinks:[],
      }
    },
    methods:{
      refreshShareLinkList:async function () {
        const result = await FilterShareLinkList(10,1);
        if(result.status == "SUCCESS"){
          this.shareLinks = result.shareLinks;
        }
      }
    },
    mounted(){
      this.refreshShareLinkList();
    }
  }
</script>

<style scoped>

</style>
