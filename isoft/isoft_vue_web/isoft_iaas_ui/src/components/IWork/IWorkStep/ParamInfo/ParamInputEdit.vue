<template>
  <span>
    <Row v-for="(item,index) in paramInputSchemaItems" style="margin-bottom: 10px;">
      <Row>
        <Col span="16">
          {{item.ParamName}}
          <Icon type="ios-book-outline" size="18" style="margin-left: 10px;" @click="showParamDesc(item.ParamDesc)"/>
        </Col>
        <Col span="8" style="text-align: right;">
          <Button type="success" size="small" @click="handleReload(index)">查看/编辑</Button>
        </Col>
      </Row>
      <Row>
        <Input size="small" v-model.trim="item.ParamValue" readonly type="text" placeholder="small size"/>
      </Row>
    </Row>

    <ParamInputEditDialog ref="paramInputEditDialog"
                          @handleSubmit="refreshParamInputSchemaItems" @handleReload="handleReload"/>
  </span>
</template>

<script>
  import ParamInputEditDialog from "./ParamInputEditDialog"

  export default {
    name: "ParamInputEdit",
    components:{ParamInputEditDialog},
    props:{
      paramInputSchemaItems:{
        type: Array,
        default: () => [],
      },
    },
    methods:{
      // 根据 paramIndex 重新加载
      handleReload: function(paramIndex){
        if(paramIndex >=0 && paramIndex <= this.paramInputSchemaItems.length -1){
          let item = this.paramInputSchemaItems[paramIndex];
          this.$refs["paramInputEditDialog"].refreshParamInput(paramIndex, item);
        }
      },
      // 强制刷新组件
      refreshParamInputSchemaItems:function (label, text) {
        for(var i=0; i<this.paramInputSchemaItems.length; i++){
          var paramInputSchemaItem = this.paramInputSchemaItems[i];
          if(paramInputSchemaItem.ParamName == label){
            paramInputSchemaItem.ParamValue = text;
            this.$set(this.paramInputSchemaItems, i, paramInputSchemaItem);
            this.$Message.success('临时参数保存成功!');
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
