<template>
  <span>
    <Row v-for="item in paramInputSchemaItems" style="margin-bottom: 10px;">
      <Row>
        <Col span="12">
          {{item.ParamName}}
          <Icon type="ios-book-outline" size="18" style="margin-left: 10px;" @click="showParamDesc(item.ParamDesc)"/>
        </Col>
        <Col span="12">
          <WorkStepParamInputEditDialog :input-label="item.ParamName" :input-text="item.ParamValue"
           @handleSubmit="refreshParamInputSchemaItems"/>
        </Col>
      </Row>
      <Row>
        <Input size="small" v-model.trim="item.ParamValue" readonly type="text" placeholder="small size"/>
      </Row>
    </Row>
  </span>
</template>

<script>
  import WorkStepParamInputEditDialog from "./WorkStepParamInputEditDialog"

  export default {
    name: "WorkStepParamInputEdit",
    components:{WorkStepParamInputEditDialog},
    props:{
      paramInputSchemaItems:{
        type: Array,
        default: () => [],
      },
    },
    methods:{
      // 强制刷新组件
      refreshParamInputSchemaItems:function (label, text) {
        for(var i=0; i<this.paramInputSchemaItems.length; i++){
          var paramInputSchemaItem = this.paramInputSchemaItems[i];
          if(paramInputSchemaItem.ParamName == label){
            paramInputSchemaItem.ParamValue = text;
            this.$set(this.paramInputSchemaItems, i, paramInputSchemaItem);
          }
        }
      },
      showParamDesc:function (paramDesc) {
        this.$Modal.info({
          title: "使用说明",
          content: paramDesc
        });
      }
    },
  }
</script>

<style scoped>

</style>
