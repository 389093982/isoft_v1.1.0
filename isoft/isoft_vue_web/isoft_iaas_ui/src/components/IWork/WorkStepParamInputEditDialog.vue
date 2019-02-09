<template>
  <ISimpleBtnTriggerModal ref="triggerModal" btn-text="查看/编辑" btn-size="small" btn-folat="right"
      modal-title="查看/编辑 workstep 参数" :modal-width="800" @btnClick="refreshPreNodeOutput">
    <Row>
      <Col span="11">
        <h3>前置节点输出参数</h3>
        <Tree :data="data1" show-checkbox ref="tree1"></Tree>
      </Col>
      <Col span="2" style="text-align: center;margin-top: 100px;">
        <Button>
          <Icon type="ios-arrow-forward" @click="appendData"></Icon>
        </Button>
      </Col>
      <Col span="11">
        <h3>{{inputLabel}}</h3>
        <Input v-model="inputTextData" type="textarea" :rows="10" placeholder="Enter something..." />
      </Col>
    </Row>
    <Row style="text-align: center;margin-top: 10px;">
      <Button type="success" @click="handleSubmit">Submit</Button>
    </Row>
  </ISimpleBtnTriggerModal>
</template>

<script>
  import {LoadPreNodeOutput} from "../../api"
  import ISimpleBtnTriggerModal from "../Common/modal/ISimpleBtnTriggerModal"

  export default {
    name: "WorkStepParamInputEditDialog",
    components:{ISimpleBtnTriggerModal},
    props: {
      inputLabel: {
        type: String,
        default: "标题",
      },
      inputText: {
        type: String,
        default: "内容",
      },
    },
    data(){
      return {
        inputTextData:this.inputText,
        preParamOutputSchemaTreeNodeArr:[],
      }
    },
    methods:{
      handleSubmit:function () {
        this.$emit("handleSubmit", this.inputLabel, this.inputTextData);
        this.$refs.triggerModal.hideModal();
      },
      refreshPreNodeOutput:async function () {
        const result = await LoadPreNodeOutput(this.$store.state.current_work_id, this.$store.state.current_work_step_id);
        if(result.status == "SUCCESS"){
          this.preParamOutputSchemaTreeNodeArr = result.preParamOutputSchemaTreeNodeArr;
        }
      },
      appendDataWithPrefix:function(prefix, item){
        // 没有子节点
        if(item.children == null){
          if(item.indeterminate == false){
            // 将数据添加到右侧
            this.inputTextData = this.inputTextData + prefix + "\n";
          }
        }else{
          // 有子节点
          let items = item.children;
          for(var i=0; i<items.length; i++){
            let item = items[i];
            this.appendDataWithPrefix(prefix + "." + item.title, item);
          }
        }
      },
      appendData:function () {
        let items = this.$refs.tree1.getCheckedAndIndeterminateNodes();
        for(var i=0; i<items.length; i++){
          let item = items[i];
          // 只统计以 $ 开头的数据
          if(item.title.indexOf("$") != -1){
            this.appendDataWithPrefix(item.title,item);
          }
        }
      }
    },
    computed:{
      data1:function () {
        var appendChildrens = function (paramOutputSchemaTreeNode, node) {       // 父级节点对象、父级节点树元素
          if(paramOutputSchemaTreeNode.NodeChildrens != null && paramOutputSchemaTreeNode.NodeChildrens.length > 0){
            const arr = [];
            for(var i=0; i<paramOutputSchemaTreeNode.NodeChildrens.length; i++) {
              var childParamOutputSchemaTreeNode = paramOutputSchemaTreeNode.NodeChildrens[i];
              var childNode = {title: childParamOutputSchemaTreeNode.NodeName,expand: false,};
              // 递归操作
              appendChildrens(childParamOutputSchemaTreeNode, childNode);
              arr.push(childNode);
            }
            node.children = arr;
          }
        };
        // tree 对应的 arr
        let treeArr = [];
        for(var i=0; i<this.preParamOutputSchemaTreeNodeArr.length; i++){
          let preParamOutputSchemaTreeNode = this.preParamOutputSchemaTreeNodeArr[i];
          const topTreeNode = {
            title: preParamOutputSchemaTreeNode.NodeName,
            expand: false,
          };
          appendChildrens(preParamOutputSchemaTreeNode,topTreeNode);
          treeArr.push(topTreeNode);
        }
        return treeArr;
      }
    }
  }
</script>

<style scoped>

</style>
