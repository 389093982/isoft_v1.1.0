<template>
  <div>
    <div>
      <span style="height: 32px;line-height: 32px;margin-bottom: 5px;color: #000;float: left !important;">课程大类：</span>
      <!-- 对父级CSS选择器加overflow:hidden样式,可以清除父级内使用float产生浮动.优点是可以很少CSS代码即可解决浮动产生 -->
      <ul style="overflow:hidden;">
        <li v-for="(configuration,index) in hotCourseTypeConfigurations"
            style="height: 32px;line-height: 32px;margin: 0 4px 5px;text-align: center;color: #333;float: left;display: inline;">
          <a href="javascript:;" style="color: #333;display: block;height: inherit;padding: 0 8px;"
             @click="currentConfiguration=configuration">
            <span>{{configuration.configuration_value}}</span>
          </a>
        </li>
      </ul>
    </div>
    <div>
      <span style="height: 32px;line-height: 32px;margin-bottom: 5px;color: #000;float: left !important;">详细分类：</span>
      <ul style="overflow:hidden;" v-if="getCurrentConfiguration() && currentConfiguration">
        <li v-for="(sub_configuration,index) in currentConfiguration.sub_configurations"
            style="height: 32px;line-height: 32px;margin: 0 4px 5px;text-align: center;color: #333;float: left;display: inline;">
          <a href="javascript:;" style="color: #333;display: block;height: inherit;padding: 0 8px;"
             @click="submit(sub_configuration.configuration_value)">
            {{sub_configuration.configuration_value}}
          </a>
        </li>
      </ul>
    </div>
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
