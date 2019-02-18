<template>
  <ISimpleBtnTriggerModal ref="triggerModal" btn-text="管理IWorkDL" modal-title="IWork服务定义文件管理" :modal-width="800">
    <Button type="success" @click="buildIWorkDL">生成IWorkDL文件</Button>
    <Button type="success">导入IWorkDL文件</Button>
    <p>11111111111111111111111111111111</p>
    <p>11111111111111111111111111111111</p>
    <p>11111111111111111111111111111111</p>
    <p>11111111111111111111111111111111</p>
    <p>11111111111111111111111111111111</p>
    <p>11111111111111111111111111111111</p>
    <p>11111111111111111111111111111111</p>
    <p>11111111111111111111111111111111</p>
  </ISimpleBtnTriggerModal>
</template>

<script>
  import ISimpleBtnTriggerModal from "../Common/modal/ISimpleBtnTriggerModal"

  export default {
    name: "IWorkDL",
    components:{ISimpleBtnTriggerModal},
    methods:{
      buildIWorkDL(){
        alert(111);
      }
    }
  }
</script>

<style scoped>

</style>
