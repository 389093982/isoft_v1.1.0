<template>
  <span>
    <ISimpleBtnTriggerModal ref="triggerModal" btn-text="全局变量管理" modal-title="新增/编辑 work" :modal-width="600">
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
