<template>
  <ISimpleBtnTriggerModal ref="triggerModal" btn-text="项目校验" btn-size="small" modal-title="查看校验结果" :modal-width="800">
    <Button type="success" size="small" @click="validateAllWork">校验全部</Button>
    <Button type="success" size="small" @click="refreshValidateResult">刷新校验结果</Button>

    <div style="margin: 20px;min-height: 300px;">
      <p style="color: green;">last tracking_id: {{tracking_id}}</p>
      <Table border :columns="columns1" :data="details" size="small"></Table>
    </div>
  </ISimpleBtnTriggerModal>
</template>

<script>
  import ISimpleBtnTriggerModal from "../../Common/modal/ISimpleBtnTriggerModal"
  import {ValidateAllWork} from "../../../api/index"
  import {LoadValidateResult} from "../../../api/index"
  import {checkEmpty} from "../../../tools/index"

  export default {
    name: "WorkValidate",
    components:{ISimpleBtnTriggerModal},
    data(){
      return {
        validating:false,
        details:[],
        tracking_id:'',
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
            render: (h,params)=>{
              return h('span',{
                style: {
                  color: checkEmpty(this.details[params.index]['work_step_name']) ? 'green': 'red',
                },
              },this.details[params.index]['detail']
              )
            }
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
          this.tracking_id = result.details[0].tracking_id;
        }
      }
    }
  }
</script>

<style scoped>

</style>
