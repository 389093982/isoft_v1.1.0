<template>
  <ISimpleBtnTriggerModal ref="triggerModal" btn-text="查看/编辑" btn-size="small" btn-folat="right"
      modal-title="查看/编辑 workstep 参数" :modal-width="800" @btnClick="refreshPreNodeOutput">
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
      }
    },
    methods:{
      handleSubmit:function () {
        this.$emit("handleSubmit", this.inputLabel, this.inputTextData);
        this.$refs.triggerModal.hideModal();
      },
      refreshPreNodeOutput:async function () {
        alert(this.$store.state.current_work_id);
        alert(this.$store.state.current_work_step_id);
        const result = await LoadPreNodeOutput(this.$store.state.current_work_id, this.$store.state.current_work_step_id);
        alert(result);
      }
    },
  }
</script>

<style scoped>

</style>
