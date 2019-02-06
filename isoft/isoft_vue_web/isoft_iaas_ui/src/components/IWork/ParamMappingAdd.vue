<template>
  <ISimpleBtnTriggerModal ref="triggerModal" btn-text="新增" btn-size="small" modal-title="新增" :modal-width="500">
    <Input v-model.trim="inputText" placeholder="请输入新增的字段名称"></Input>
    <Row style="text-align: right;">
      <Button type="success" @click="handleSubmit()" style="margin-top: 6px">Submit</Button>
    </Row>
  </ISimpleBtnTriggerModal>
</template>

<script>
  import ISimpleBtnTriggerModal from "../Common/modal/ISimpleBtnTriggerModal"

  export default {
    name: "ParamMappingAdd",
    components:{ISimpleBtnTriggerModal},
    data(){
      return {
        inputText:"",
      }
    },
    methods:{
      handleSubmit:function () {
        this.$emit("handleSubmit", this.inputText);
        this.$refs.triggerModal.hideModal();
      }

    }
  }
</script>

<style scoped>

</style>
