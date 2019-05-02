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
      <Col span="7">
        <h3>前置节点输出参数</h3>
        <Scroll height="350">
          <Tree :data="data1" show-checkbox ref="tree1"></Tree>
        </Scroll>
      </Col>
      <Col span="3" style="text-align: center;margin-top: 100px;">
        <Button @click="appendData('parent')" style="margin-top: 10px;"><Icon type="ios-arrow-forward"></Icon>选择父节点</Button>
        <Button @click="appendData('children')" style="margin-top: 10px;"><Icon type="ios-arrow-forward"></Icon>选择子节点</Button>
      </Col>
      <Col span="14">
        <h3 style="color: #1600ff;">参数({{paramIndex}}):{{inputLabel}}</h3>
        <span class="operate_link">
          <ul>
            <li>
              <router-link :to="{ path: '/iwork/quickSql' }" tag="a" target="_blank">
                <Icon type="ios-cube-outline" size="18" style=" float: right;"/>
              </router-link>
            </li>
            <li>
              <QuickFuncList ref="quickFuncList" @chooseFunc="chooseFunc"/>
              <Icon type="md-copy" size="18" style="float: right;" @click="showQuickFunc()"/>
            </li>
            <li>
              <Checkbox v-model="pureText" style="float: right;">纯文本值</Checkbox>
            </li>
            <li>
              <a @click="parseToMultiValue">多值预览</a>
            </li>
          </ul>
        </span>
        <div v-show="showMultiVals" style="margin-top: 20px;">
          <span v-for="(val,index) in multiVals">
            参数 {{index}}
            <Input type="textarea" :value="val" readonly="true"/>
          </span>
        </div>
        <Input v-show="showMultiVals == false" v-model="inputTextData" type="textarea" :rows="15" placeholder="Enter something..." />
      </Col>
    </Row>
    <Row style="text-align: right;margin-top: 10px;">
      <Button type="success" size="small" @click="handleSubmit">Submit</Button>
      <Button type="error" size="small" @click="showNext(-1)">Edit Last</Button>
      <Button type="warning" size="small" @click="showNext(1)">Edit Next</Button>
      <Button type="info" size="small" @click="closeModal">Close</Button>
    </Row>
  </Modal>
</template>

<script>
  import {LoadPreNodeOutput} from "../../../../api"
  import ISimpleBtnTriggerModal from "../../../Common/modal/ISimpleBtnTriggerModal"
  import QuickFuncList from "../../IQuickFunc/QuickFuncList"
  import {ParseToMultiValue} from "../../../../api"

  export default {
    name: "ParamInputEditDialog",
    components:{ISimpleBtnTriggerModal,QuickFuncList},
    data(){
      return {
        showFormModal:false,
        inputLabel:'',
        oldPureText:false,
        pureText:false,
        oldInputTextData:'',
        inputTextData:'',
        showMultiVals:false,  // 默认非多值视图
        multiVals:[],         // 存储多值列表
        paramIndex:1,
        preParamOutputSchemaTreeNodeArr:[],
      }
    },
    methods:{
      parseToMultiValue: async function(){
        const result = await ParseToMultiValue(this.pureText, this.inputTextData);
        if(result.status == "SUCCESS"){
          this.showMultiVals = !this.showMultiVals;
          this.multiVals = result.multiVals;
        }else{
          this.$Message.error('提交失败!' + result.errorMsg);
        }
      },
      handleReload: function(paramIndex){
        this.$emit("handleReload", paramIndex);
      },
      refreshParamInput: function(index, item){
        this.showFormModal = true;
        this.paramIndex = index;
        this.inputLabel = item.ParamName;
        this.pureText = item.PureText;
        // 文本输入框设置历史值
        this.inputTextData = item.ParamValue;
        this.showMultiVals = false;
        this.clearDirty();
        this.refreshPreNodeOutput();
      },
      showNext: function(num){
        var _this = this;
        if(this.checkDrity()){
          this.$Modal.confirm({
            title:"确认",
            content:"是否需要保存上一步操作?",
            onOk: function () {
              _this.handleSubmit();
            },
          });
        }else{
          this.handleReload(this.paramIndex + num);
        }
      },
      closeModal: function(){
        this.showFormModal = false;
      },
      chooseFunc: function(funcDemo){
        // 将数据复制到右侧
        this.inputTextData = this.inputTextData + funcDemo + "\n";
      },
      showQuickFunc: function(){
        this.$refs.quickFuncList.showModal();
      },
      clearDirty: function (){
        this.oldInputTextData = this.inputTextData;
        this.oldPureText = this.pureText;
      },
      checkDrity: function(){
        return this.oldInputTextData != this.inputTextData || this.oldPureText != this.pureText;
      },
      handleSubmit:function () {
        this.$emit("handleSubmit", this.inputLabel, this.inputTextData, this.pureText);
        this.clearDirty();
      },
      refreshPreNodeOutput:async function () {
        const result = await LoadPreNodeOutput(this.$store.state.current_work_id, this.$store.state.current_work_step_id);
        if(result.status == "SUCCESS"){
          this.preParamOutputSchemaTreeNodeArr = result.preParamOutputSchemaTreeNodeArr;
        }
      },
      appendDataWithPrefix:function(prefix, item, chooseType){
        // 没有子节点
        if(item.children == null){
          if(item.indeterminate == false){
            if(item.checked){
              // 将数据添加到右侧
              this.inputTextData = this.inputTextData + prefix + ";\n";
            }
          }
        }else{
          // 有子节点
          let items = item.children;
          if(chooseType == 'parent' && item.indeterminate == false){
            if(item.checked){
              // 将数据添加到右侧
              this.inputTextData = this.inputTextData + prefix + ";\n";
            }
          }else{
            for(var i=0; i<items.length; i++){
              let item = items[i];
              this.appendDataWithPrefix(prefix + "." + item.title, item, chooseType);
            }
          }
        }
      },
      appendData:function (chooseType) {
        let items = this.$refs.tree1.getCheckedAndIndeterminateNodes();
        for(var i=0; i<items.length; i++){
          let item = items[i];
          // 只统计以 $ 开头的数据
          if(item.title.indexOf("$") != -1){
            this.appendDataWithPrefix(item.title,item, chooseType);
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
              var childNode = {title: childParamOutputSchemaTreeNode.NodeName,expand: false};
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
  .operate_link ul li{
   display: inline-block;
   margin-left: 10px;
   float: right;
 }
</style>
