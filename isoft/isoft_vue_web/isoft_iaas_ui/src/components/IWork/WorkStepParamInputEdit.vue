<template>
  <span>
    <Row v-for="item in paramSchemaItems">
      <Row>
        <Col span="12">{{item.ParamName}}</Col>
        <Col span="12" style="text-align: right;">
          <WorkStepParamInputEditDialog :input-label="item.ParamName" :input-text="item.ParamValue" @handleSubmit="refreshParamSchemaItems"/>
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
      paramSchemaItems:{
        type: Array,
        default: () => [],
      }
    },
    methods:{
      // 强制刷新组件
      refreshParamSchemaItems:function (label, text) {
        for(var i=0; i<this.paramSchemaItems.length; i++){
          var paramSchemaItem = this.paramSchemaItems[i];
          if(paramSchemaItem.ParamName == label){
            paramSchemaItem.ParamValue = text;
            this.$set(this.paramSchemaItems, i, paramSchemaItem);
          }
        }
      }
    },
  }
</script>

<style scoped>

</style>
