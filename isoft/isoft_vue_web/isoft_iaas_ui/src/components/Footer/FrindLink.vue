<template>
  <IHotRecommand title="友情链接" :items="items"/>
</template>

<script>
  import IHotRecommand from "../Common/recommend/IHotRecommand"
  import {QueryRandomCommonLink} from "../../api/index"

  export default {
    name: "FrindLink",
    components:{IHotRecommand},
    data(){
      return {
        frindLinks:[],
      }
    },
    methods:{
      async refreshRandomFrinkLink (){
        const data = await QueryRandomCommonLink("friend_link");
        if(data.status == "SUCCESS"){
          this.frindLinks = data.common_links;
        }
      }
    },
    computed:{
      // 使用计算属性对对象格式进行转换
      items: function () {
        let arr = new Array();
        if(this.fractionLost != null){
          for(var i=0; i<this.frindLinks.length; i++){
            arr.push({"item_label":this.frindLinks[i].link_name,"item_href":this.frindLinks[i].link_addr});
          }
        }
        return arr;
      }
    },
    mounted:function () {
      this.refreshRandomFrinkLink();
      setInterval(this.refreshRandomFrinkLink, 30000);
    }
  }
</script>

<style scoped>

</style>
