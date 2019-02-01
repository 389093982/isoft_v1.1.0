<template>
  <span>
    <Button type="success" @click="showFormModal = true">{{btnText}}</Button>

    <Modal
      v-model="showFormModal"
      :width="modalWidth"
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
    name: "ISimpleBtnTriggerModal",
    props: {
      btnText: {
        type: String,
        default: "按钮标题"
      },
      modalTitle: {
        type: String,
        default: "模态框标题"
      },
      modalWidth: {
        type: Number,
        default: 800
      },
    },
    data(){
      return {
        showFormModal:false,
      }
    },
    methods:{
      // 阴藏 modal
      hideModal:function () {
        this.showFormModal = false;
      }
    }
  }
</script>

<style scoped>

</style>
