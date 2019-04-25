<template>
  <div>
    <div>
      <span style="height: 32px;line-height: 32px;margin-bottom: 5px;color: #000;float: left !important;">课程大类：</span>
      <!-- 对父级CSS选择器加overflow:hidden样式,可以清除父级内使用float产生浮动.优点是可以很少CSS代码即可解决浮动产生 -->
      <ul style="overflow:hidden;">
        <li v-for="(hot_course_type,index) in hot_course_types"
            style="height: 32px;line-height: 32px;margin: 0 4px 5px;text-align: center;color: #333;float: left;display: inline;">
          <IBeautifulLink @onclick="currentConfiguration=hot_course_type">{{hot_course_type.course_type}}</IBeautifulLink>
        </li>
      </ul>
    </div>
    <div>
      <span style="height: 32px;line-height: 32px;margin-bottom: 5px;color: #000;float: left !important;">详细分类：</span>
      <ul style="overflow:hidden;" v-if="getCurrentConfiguration() && currentConfiguration">
        <li v-for="(sub_course_type,index) in currentConfiguration.sub_course_types"
            style="height: 32px;line-height: 32px;margin: 0 4px 5px;text-align: center;color: #333;float: left;display: inline;">
          <IBeautifulLink @onclick="chooseCourseType(currentConfiguration.course_type, sub_course_type)">
            {{sub_course_type}}
          </IBeautifulLink>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
  import IBeautifulLink from "../../Common/link/IBeautifulLink"

  export default {
    name: "HotCourseType",
    components:{IBeautifulLink},
    data(){
      return {
        hot_course_types: this.GLOBAL.hot_course_types,
        currentConfiguration:undefined,
      }
    },
    methods: {
      chooseCourseType:function (course_type, course_sub_type) {
        this.$emit("chooseCourseType", course_type, course_sub_type);
      },
      getCurrentConfiguration:function () {
        if(this.currentConfiguration == undefined){
          this.currentConfiguration = this.hot_course_types[0];
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
