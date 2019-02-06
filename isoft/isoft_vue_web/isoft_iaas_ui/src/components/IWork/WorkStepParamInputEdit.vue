<template>
  <span>
    <Row v-for="item in paramInputSchemaItems" style="margin-bottom: 10px;">
      <Row>
        <Col span="12">{{item.ParamName}}</Col>
        <Col span="12">
          <WorkStepParamInputEditDialog :input-label="item.ParamName" :input-text="item.ParamValue" @handleSubmit="refreshParamInputSchemaItems"/>
        </Col>
      </Row>
      <Row>
        <Input size="small" v-model="item.ParamValue" readonly type="textarea" :rows="2" placeholder="small size"/>
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
      }
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
      }
    },
  }
</script>

<style scoped>

</style>
