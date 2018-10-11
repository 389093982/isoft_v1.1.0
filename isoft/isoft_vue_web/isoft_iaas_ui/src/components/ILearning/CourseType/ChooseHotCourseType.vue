<template>
  <span>
    <a href="javascript:;" @click="showCourseType = true"><Icon type="md-book" />&nbsp;选择推荐分类</a>
    <Modal
      v-model="showCourseType"
      title="热门课程分类"
      width="850"
      :footer-hide="true"
      :styles="{top: '20px'}"
      :mask-closable="false">
      <div style="height: 450px;">
        <Scroll height="450">
          <dl v-for="configuration in configurations">
            <Row style="margin: 5px;">
              <Col span="3">
                <dt>
                  <p style="float: right;margin-right: 10px;">{{configuration.configuration_value}}</p>
                </dt>
              </Col>
              <Col span="21">
                <dd v-for="sub_configuration in configuration.sub_configurations">
                  <a href="javascript:;" style="float: left;margin: 3px 0;padding: 0 10px;height: 16px;
              border-left: 1px solid #e0e0e0;line-height: 16px;white-space: nowrap;">{{sub_configuration.configuration_value}}</a>
                </dd>
              </Col>
            </Row>
          </dl>
        </Scroll>
      </div>
    </Modal>
  </span>
</template>

<script>
  import {QueryAllConfigurations} from "../../../api"

  export default {
    name: "ChooseHotCourseType",
    data(){
      return {
        showCourseType:false,
        configurations:[],
      }
    },
    methods:{
      refreshCourseType: async function () {
        const result = await QueryAllConfigurations("recommand_course_type");
        if (result.status == "SUCCESS") {
          this.configurations = result.configurations;
        }
      },
    },
    mounted:function () {
      this.refreshCourseType();
    }
  }
</script>

<style scoped>
  a{
    color: black;
  }
  a:hover{
    color: red;
  }
</style>
