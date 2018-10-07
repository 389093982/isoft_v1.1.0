<template>
  <div>
    热门课程推荐
    <ul>
      <li v-for="(configuration,index) in configurations" style="margin:10px 10px 0 0;list-style:none;float: left;">

        <Poptip trigger="hover" :title="configuration.configuration_value" content="content" placement="bottom">
          <a href="javascript:;" style="font-size: 14px;color: #626262;" @click="submit(configuration.configuration_value)">
            {{configuration.configuration_value}}
          </a>

          <div class="api" slot="content">
            <ul>
              <li v-for="(sub_configuration,index) in configuration.sub_configurations"
                  style="padding-left: 10px;list-style:none;float: left;">
                <a href="javascript:;" style="font-size: 14px;color: #626262;" @click="submit(configuration.configuration_value)">
                  {{sub_configuration.configuration_value}}
                </a>
              </li>
            </ul>
          </div>
        </Poptip>

      </li>
    </ul>
    <div style="clear: both;"></div>
  </div>
</template>

<script>
  import {QueryAllConfigurations} from "../../../api"

  export default {
    name: "HotCourseType",
    data(){
      return {
        configurations:[],
      }
    },
    methods: {
      refreshCourseType: async function () {
        const result = await QueryAllConfigurations("recommand_course_type");
        if (result.status == "SUCCESS") {
          this.configurations = result.configurations;
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

</style>
