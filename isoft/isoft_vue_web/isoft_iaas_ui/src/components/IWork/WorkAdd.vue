<template>
  <!-- 按钮触发模态框 -->
  <!-- ref 的作用是为了在其它地方方便的获取到当前子组件 -->
  <ISimpleBtnTriggerModal ref="triggerModal" btn-text="新增" modal-title="新增 work" :modal-width="600">
    <!-- 表单信息 -->
    <Form ref="formValidate" :model="formValidate" :rules="ruleValidate" :label-width="100">
      <FormItem label="work_name" prop="work_name">
        <Input v-model.trim="formValidate.work_name" placeholder="请输入 work_name"></Input>
      </FormItem>
      <FormItem>
        <Button type="success" @click="handleSubmit('formValidate')" style="margin-right: 6px">Submit</Button>
        <Button type="warning" @click="handleReset('formValidate')" style="margin-right: 6px">Reset</Button>
      </FormItem>
    </Form>
  </ISimpleBtnTriggerModal>
</template>

<script>
  import ISimpleBtnTriggerModal from "../Common/modal/ISimpleBtnTriggerModal"
  import {AddWork} from "../../api"

  export default {
    name: "WorkAdd",
    components:{ISimpleBtnTriggerModal},
    data(){
      return {
        formValidate: {
          work_name: '',
        },
        ruleValidate: {
          work_name: [
            { required: true, message: 'work_name 不能为空!', trigger: 'blur' }
          ],
        },
      }
    },
    methods:{
      handleSubmit (name) {
        this.$refs[name].validate(async (valid) => {
          if (valid) {
            const result = await AddWork(this.formValidate.work_name);
            if(result.status == "SUCCESS"){
              this.$Message.success('提交成功!');
              // 调用子组件隐藏 modal (this.refs.xxx.子组件定义的方法())
              this.$refs.triggerModal.hideModal();
              // 通知父组件添加成功
              this.$emit('handleSuccess');
            }else{
              this.$Message.error('提交失败!');
            }
          }
        })
      },
      handleReset (name) {
        this.$refs[name].resetFields();
      },
    }
  }
</script>

<style scoped>

</style>
