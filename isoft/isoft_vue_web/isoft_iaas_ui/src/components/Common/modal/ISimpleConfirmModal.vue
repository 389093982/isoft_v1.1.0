<template>
  <span>
    <Modal
      v-model="showFormModal"
      :width="modalWidth"
      :title="modalTitle"
      :transfer="false"
      :footer-hide="footerHide"
      @on-ok="handleSubmit"
      :mask-closable="false">
      <div>
        <slot></slot>
      </div>
    </Modal>
  </span>
</template>

<script>
  export default {
    name: "ISimpleConfirmModal",
    props: {
      modalTitle: {
        type: String,
        default: "模态框标题"
      },
      modalWidth: {
        type: Number,
        default: 800
      },
      footerHide: {
        type: Boolean,
        default: false,
      }
    },
    data(){
      return {
        showFormModal:false,
      }
    },
    methods:{
      showModal:function () {
        this.showFormModal = true;
      },
      hideModal:function () {
        this.showFormModal = false;
      },
      handleSubmit:function () {
        this.$emit("handleSubmit");
      }
    }
  }
</script>

<style scoped>

</style>
