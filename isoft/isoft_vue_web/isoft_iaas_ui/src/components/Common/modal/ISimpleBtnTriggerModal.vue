<template>
  <span>
    <Button type="success" :size="btnSize" @click="showFormModal = true">{{btnText}}</Button>

    <!-- 设置 transfer 属性为 false 后,保证好两个 Modal 的前后顺序,可以解决顺序问题 -->
    <Modal
      v-model="showFormModal"
      :width="modalWidth"
      :title="modalTitle"
      :footer-hide="true"
      :transfer="false"
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
      btnSize: {
        type: String,
        default: "default"    // large、small、default
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
      // 隐藏 modal
      hideModal:function () {
        this.showFormModal = false;
      }
    }
  }
</script>

<style scoped>

</style>
