<template>
  <span>
    <Button type="success" @click="showFormModal = true">{{btnText}}</Button>

    <Modal
      v-model="showFormModal"
      width="850"
      :title="modalTitle"
      :footer-hide="true"
      :mask-closable="false">
      <div>
        <slot></slot>
      </div>
    </Modal>
  </span>
</template>

<script>
  export default {
    name: "IBtnTriggerModal",
    props: {
      btnText: {
        type: String,
        default: "按钮标题"
      },
      modalTitle: {
        type: String,
        default: "模态框标题"
      },
    },
    data(){
      return {
        showFormModal:false,
      }
    }
  }
</script>

<style scoped>

</style>
