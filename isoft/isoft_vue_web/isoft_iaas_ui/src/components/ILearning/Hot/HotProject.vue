<template>
  <IHotRecommand title="热门项目推荐" :items="items"/>
</template>

<script>
  import IHotRecommand from "../../Common/recommend/IHotRecommand"
  import {QueryRandomCommonLink} from "../../../api"

  export default {
    name: "HotProject",
    components:{IHotRecommand},
    data(){
      return {
        hot_projects:[],
      }
    },
    methods:{
      async refreshRandomHotProject (){
        const data = await QueryRandomCommonLink("hot_project");
        if(data.status == "SUCCESS"){
          this.hot_projects = data.common_links;
        }
      }
    },
    computed:{
      // 使用计算属性对对象格式进行转换
      items: function () {
        let arr = new Array();
        for(var i=0; i<this.hot_projects.length; i++){
          arr.push({"item_label":this.hot_projects[i].link_name,"item_href":this.hot_projects[i].link_addr});
        }
        return arr;
      }
    },
    mounted:function () {
      this.refreshRandomHotProject();
    }
  }
</script>

<style scoped>

</style>
