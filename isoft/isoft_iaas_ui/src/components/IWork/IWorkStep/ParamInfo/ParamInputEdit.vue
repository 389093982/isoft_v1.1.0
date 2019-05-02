<template>
  <div>
    <Row v-for="(item,index) in paramInputSchemaItems" style="margin-bottom: 10px;">
      <Row>
        <Col span="16">
          {{item.ParamName}}
          <Icon type="ios-book-outline" size="18" style="margin-left: 3px;" @click="showParamDesc(item.ParamDesc)"/>
        </Col>
        <Col span="8" style="text-align: right;">
          <Button v-if="!item.ParamChoices" type="success" size="small" @click="handleReload(index)">查看/编辑</Button>
        </Col>
      </Row>
      <Row>
        <!-- transfer="true" 表示是否将弹层放置于 body 内,
          在 Tabs、带有 fixed 的 Table 列内使用时,建议添加此属性,它将不受父级样式影响,从而达到更好的效果-->
        <Select v-if="item.ParamChoices" v-model="item.ParamValue" :transfer="true">
          <Option v-for="choice in item.ParamChoices" :value="choice" :key="choice">
            {{choice}}
          </Option>
        </Select>
        <Input v-else size="small" v-model.trim="item.ParamValue" readonly type="text" placeholder="small size"/>
      </Row>
    </Row>

    <ParamInputEditDialog ref="paramInputEditDialog" @handleSubmit="refreshParamInputSchemaItems" @handleReload="handleReload"/>
  </div>
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
      refreshParamInputSchemaItems:function (label, text, pureText) {
        for(var i=0; i<this.paramInputSchemaItems.length; i++){
          var paramInputSchemaItem = this.paramInputSchemaItems[i];
          if(paramInputSchemaItem.ParamName == label){
            paramInputSchemaItem.ParamValue = text;
            paramInputSchemaItem.PureText = pureText;
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
