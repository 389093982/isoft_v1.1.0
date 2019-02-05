<template>
  <ISimpleBtnTriggerModal ref="triggerModal" btn-text="查看/编辑" btn-size="small" modal-title="查看/编辑 workstep 参数" :modal-width="800">
    <Row>
      <Col span="12">前置节点输出参数</Col>
      <Col span="12">{{inputLabel}}
        <Input v-model="inputTextData" type="textarea" :rows="10" placeholder="Enter something..." />
      </Col>
    </Row>
    <Row style="text-align: center;margin-top: 10px;">
      <Button type="success" @click="handleSubmit">Submit</Button>
    </Row>
  </ISimpleBtnTriggerModal>
</template>

<script>
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
      }
    },
    methods:{
      handleSubmit:function () {
        this.$emit("handleSubmit", this.inputLabel, this.inputTextData);
        this.$refs.triggerModal.hideModal();
      }
    }
  }
</script>

<style scoped>

</style>
