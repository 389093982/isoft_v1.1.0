<template>
  <span>
    <Row v-for="item in paramDefinitionItems">
      <Row>
        <Col span="12">{{item.ParamName}}</Col>
        <Col span="12" style="text-align: right;">
          <WorkStepInputEditDialog :input-label="item.ParamName" :input-text="item.ParamValue" @handleSubmit="editInputBinding"/>
        </Col>
      </Row>
      <Row>
        <Input size="small" v-model="item.ParamValue" placeholder="small size" />
      </Row>
    </Row>
  </span>
</template>

<script>
  import WorkStepInputEditDialog from "./WorkStepInputEditDialog"

  export default {
    name: "WorkStepInputEdit",
    components:{WorkStepInputEditDialog},
    props:{
      paramDefinitionItems:{
        type: Array,
        default: () => [],
      }
    },
    methods:{
      editInputBinding:function (label, text) {
        for(var i=0; i<this.paramDefinitionItems.length; i++){
          var paramDefinitionItem = this.paramDefinitionItems[i];
          if(paramDefinitionItem.ParamName == label){
            paramDefinitionItem.ParamValue = text;
            this.$set(this.paramDefinitionItems, i, paramDefinitionItem);
          }
        }
      }
    }
  }
</script>

<style scoped>

</style>
