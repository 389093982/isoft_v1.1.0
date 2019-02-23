<template>
  <ISimpleBtnTriggerModal ref="triggerModal" btn-text="项目校验" modal-title="查看校验结果" :modal-width="800">
    <Button type="success" @click="validateAllWork">校验全部</Button>
    <Button type="success" @click="refreshValidateResult">刷新校验结果</Button>

    <div style="margin: 20px;min-height: 300px;">
      <Table border :columns="columns1" :data="details" size="small"></Table>
    </div>
  </ISimpleBtnTriggerModal>
</template>

<script>
  import ISimpleBtnTriggerModal from "../Common/modal/ISimpleBtnTriggerModal"
  import {ValidateAllWork} from "../../api"
  import {LoadValidateResult} from "../../api"

  export default {
    name: "WorkValidate",
    components:{ISimpleBtnTriggerModal},
    data(){
      return {
        validating:false,
        details:[],
        columns1: [
          {
            title: 'work_name',
            key: 'work_name',
          },
          {
            title: 'work_step_name',
            key: 'work_step_name',
          },
          {
            title: 'detail',
            key: 'detail',
          },
        ],
      }
    },
    methods:{
      validateAllWork:async function () {
        if(this.validating == true){
          this.$Message.error("校验中,请稍后！");
        }else{
          this.validating = true;
          const result = await ValidateAllWork();
          if(result.status == "SUCCESS"){
            this.refreshValidateResult();
          }
          this.validating = false;
        }
      },
      refreshValidateResult: async function () {
        const result = await LoadValidateResult();
        if(result.status == "SUCCESS"){
          this.details = result.details;
        }
      }
    }
  }
</script>

<style scoped>

</style>
