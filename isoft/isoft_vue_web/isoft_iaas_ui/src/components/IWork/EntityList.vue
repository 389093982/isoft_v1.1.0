<template>
  <span>
    <ISimpleBtnTriggerModal ref="triggerModal" btn-text="实体类管理" modal-title="新增/编辑 work" :modal-width="600">
      Entity管理
    </ISimpleBtnTriggerModal>
  </span>
</template>

<script>
  import ISimpleBtnTriggerModal from "../Common/modal/ISimpleBtnTriggerModal"

  export default {
    name: "EntityList",
    components:{ISimpleBtnTriggerModal},
  }
</script>

<style scoped>

</style>
