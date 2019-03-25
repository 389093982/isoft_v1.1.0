<template>
  <Modal
    v-model="showFormModal"
    width="950"
    title="查看/编辑 workstep 参数"
    :footer-hide="true"
    :transfer="false"
    :mask-closable="false"
    :styles="{top: '20px'}">
    <Row>
      <Col span="8">
        <h3>前置节点输出参数</h3>
        <Scroll height="350">
          <Tree :data="data1" show-checkbox ref="tree1"></Tree>
        </Scroll>
      </Col>
      <Col span="2" style="text-align: center;margin-top: 100px;">
        <Button>
          <Icon type="ios-arrow-forward" @click="appendData"></Icon>
        </Button>
      </Col>
      <Col span="14">
        <h3 style="color: #1600ff;">
          参数名称({{paramIndex}}):{{inputLabel}}
        </h3>
        <QuickFuncList ref="quickFuncList" @chooseFunc="chooseFunc"/>
        <Icon type="md-copy" size="18" style="float: right;" @click="showQuickFunc()"/>
        <Input v-model="inputTextData" type="textarea" :rows="15" placeholder="Enter something..." />
      </Col>
    </Row>
    <Row style="text-align: right;margin-top: 10px;">
      <Button type="success" size="small" @click="handleSubmit(false)">提交</Button>
      <Button type="warning" size="small" @click="handleSubmit(true)">提交并关闭</Button>
      <Button type="success" size="small" @click="showNext(-1)">编辑上一个参数</Button>
      <Button type="warning" size="small" @click="showNext(1)">编辑下一个参数</Button>
    </Row>
  </Modal>
</template>

<script>
  import {LoadPreNodeOutput} from "../../api"
  import ISimpleBtnTriggerModal from "../Common/modal/ISimpleBtnTriggerModal"
  import QuickFuncList from "./QuickFuncList"

  export default {
    name: "WorkStepParamInputEditDialog",
    components:{ISimpleBtnTriggerModal,QuickFuncList},
    data(){
      return {
        showFormModal:false,
        inputLabel:'',
        inputTextData:'',
        paramIndex:1,
        preParamOutputSchemaTreeNodeArr:[],
      }
    },
    methods:{
      handleReload: function(paramIndex){
        this.$emit("handleReload", paramIndex);
      },
      refreshParamInput: function(index, item){
        this.showFormModal = true;
        this.paramIndex = index;
        this.inputLabel = item.ParamName;
        // 文本输入框设置历史值
        this.inputTextData = item.ParamValue;
        this.refreshPreNodeOutput();
      },
      showNext: function(num){
        this.handleReload(this.paramIndex + num);
      },
      chooseFunc: function(funcDemo){
        // 将数据复制到右侧
        this.item.inputTextData = this.item.inputTextData + funcDemo + "\n";
      },
      showQuickFunc: function(){
        this.$refs.quickFuncList.showModal();
      },
      handleSubmit:function (closable) {
        this.$emit("handleSubmit", this.inputLabel, this.inputTextData);
        if(closable){
          this.$refs.triggerModal.hideModal();
        }
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
            this.inputTextData = this.inputTextData + prefix + ";\n";
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
