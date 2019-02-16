<template>
  <span>
    <Button :style="{float:btnFolat}" type="success" :size="btnSize" @click="btnClick">{{btnText}}</Button>

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
      btnFolat: {
        type: String,
        default: "none"
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
      },
      // 模拟触发点击事件
      triggerClick:function(){
        this.showFormModal = true;
      },
      btnClick:function () {
        this.showFormModal = true;
        this.$emit("btnClick");
      }
    },
  }
</script>

<style scoped>

</style>
