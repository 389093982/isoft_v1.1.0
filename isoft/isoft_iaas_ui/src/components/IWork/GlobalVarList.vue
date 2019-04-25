<template>
  <span>
    <ISimpleBtnTriggerModal ref="triggerModal" btn-size="small" btn-text="全局变量管理" modal-title="新增/编辑全局变量" :modal-width="800">
      GlobalVar管理
    </ISimpleBtnTriggerModal>
  </span>
</template>

<script>
  import ISimpleBtnTriggerModal from "../Common/modal/ISimpleBtnTriggerModal"

  export default {
    name: "GlobalVarList",
    components:{ISimpleBtnTriggerModal},
  }
</script>

<style scoped>

</style>
