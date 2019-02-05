<template>
  <span>
    <Row v-for="item in paramDefinitionItems">
      <Row>
        <Col span="12">{{item.ParamName}}</Col>
        <Col span="12" style="text-align: right;">
          <WorkStepParamInputEditDialog :input-label="item.ParamName" :input-text="item.ParamValue" @handleSubmit="refreshParamDefinitionItems"/>
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
      paramDefinitionItems:{
        type: Array,
        default: () => [],
      }
    },
    methods:{
      // 强制刷新组件
      refreshParamDefinitionItems:function (label, text) {
        for(var i=0; i<this.paramDefinitionItems.length; i++){
          var paramDefinitionItem = this.paramDefinitionItems[i];
          if(paramDefinitionItem.ParamName == label){
            paramDefinitionItem.ParamValue = text;
            this.$set(this.paramDefinitionItems, i, paramDefinitionItem);
          }
        }
      }
    },
  }
</script>

<style scoped>

</style>
