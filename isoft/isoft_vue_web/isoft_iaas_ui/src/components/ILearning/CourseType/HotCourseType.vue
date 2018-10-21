<template>
  <div>
    <div style="float: left;">课程大类：</div>
    <!-- 对父级CSS选择器加overflow:hidden样式,可以清除父级内使用float产生浮动.优点是可以很少CSS代码即可解决浮动产生 -->
    <ul style="overflow:hidden;border-bottom: 2px solid #edf1f2;">
      <li v-for="(configuration,index) in hotCourseTypeConfigurations" style="list-style:none;float: left;">
        <a href="javascript:;" style="margin:5px;padding-left:15px;padding-right:15px;font-size: 14px;
            display: inline-block;color: #fff;background: rgb(34,195,0);"
            @click="currentConfiguration=configuration">
          <span>{{configuration.configuration_value}}</span>
        </a>
      </li>
    </ul>
    <div style="float: left;">详细分类：</div>
    <ul style="overflow:hidden;" v-if="getCurrentConfiguration() && currentConfiguration">
      <li v-for="(sub_configuration,index) in currentConfiguration.sub_configurations" style="list-style:none;float: left;">
        <a href="javascript:;" style="margin:5px;padding-left:15px;padding-right:15px;font-size: 14px;
            display: inline-block;color: #fff;background: rgb(255,98,37);"
            @click="submit(sub_configuration.configuration_value)">
          {{sub_configuration.configuration_value}}
        </a>
      </li>
    </ul>
  </div>
</template>

<script>
  import {mapState} from 'vuex'
  import {QueryAllConfigurations} from "../../../api"

  export default {
    name: "HotCourseType",
    data(){
      return {
        currentConfiguration:undefined,
      }
    },
    computed:{
      ...mapState(['hotCourseTypeConfigurations']),
    },
    methods: {
      submit:function (data) {
        this.$emit("submitFunc", data);
      },
      getCurrentConfiguration:function () {
        if(this.currentConfiguration == undefined){
          this.currentConfiguration = this.hotCourseTypeConfigurations[0];
        }
        return true;
      }
    },
  }
</script>

<style scoped>
  a{
    color: #626262;
  }
  a:hover{
    color: red;
  }
</style>
