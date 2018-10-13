<template>
  <div>
    <!-- 对父级CSS选择器加overflow:hidden样式,可以清除父级内使用float产生浮动.优点是可以很少CSS代码即可解决浮动产生 -->
    <ul style="overflow:hidden">
      <li v-for="(configuration,index) in configurations" style="list-style:none;float: left;">
        <a href="javascript:;" style="margin:5px;padding-left:15px;padding-right:15px;font-size: 14px;
            display: inline-block;color: #fff;background: rgb(34,195,0);"
            @click="currentConfiguration=configuration">
          <span>{{configuration.configuration_value}}</span>
        </a>
      </li>
    </ul>
    <ul style="overflow:hidden">
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
