<template>
  <div>
    <!-- 对父级CSS选择器加overflow:hidden样式,可以清除父级内使用float产生浮动.优点是可以很少CSS代码即可解决浮动产生 -->
    <ul style="overflow:hidden">
      <li v-for="(configuration,index) in configurations" style="margin:5px;padding:5px;list-style:none;
        float: left;background: rgba(219,167,255,0.34);">
        <a href="javascript:;" style="font-size: 14px;" @click="currentConfiguration=configuration">
          <span>{{configuration.configuration_value}}</span>
        </a>
      </li>
    </ul>
    <ul style="overflow:hidden">
      <li v-for="(sub_configuration,index) in currentConfiguration.sub_configurations"
          style="padding-left: 10px;list-style:none;float: left;">
        <a href="javascript:;" style="font-size: 14px;" @click="submit(sub_configuration.configuration_value)">
          {{sub_configuration.configuration_value}}
        </a>
      </li>
    </ul>
  </div>
</template>

<script>
  import {QueryAllConfigurations} from "../../../api"

  export default {
    name: "HotCourseType",
    data(){
      return {
        configurations:[],
        currentConfiguration:{},
      }
    },
    methods: {
      refreshCourseType: async function () {
        const result = await QueryAllConfigurations("recommand_course_type");
        if (result.status == "SUCCESS") {
          this.configurations = result.configurations;
          this.currentConfiguration = result.configurations[0];
        }
      },
      submit:function (data) {
        alert(data);
      }
    },
    mounted:function () {
      this.refreshCourseType();
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
