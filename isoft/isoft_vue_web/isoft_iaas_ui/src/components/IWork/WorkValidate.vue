<template>
  <ISimpleBtnTriggerModal ref="triggerModal" btn-text="项目校验" modal-title="查看校验结果" :modal-width="800">
    <Button type="success" @click="validateAllWork">校验全部</Button>
    <Button type="success">查看校验结果</Button>
  </ISimpleBtnTriggerModal>
</template>

<script>
  import ISimpleBtnTriggerModal from "../Common/modal/ISimpleBtnTriggerModal"
  import {ValidateAllWork} from "../../api"

  export default {
    name: "WorkValidate",
    components:{ISimpleBtnTriggerModal},
    methods:{
      validateAllWork:async function () {
        const result = await ValidateAllWork();
      }
    }
  }
</script>

<style scoped>

</style>
