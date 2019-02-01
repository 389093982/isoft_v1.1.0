<template>
  <!-- 按钮触发模态框 -->
  <!-- ref 的作用是为了在其它地方方便的获取到当前子组件 -->
  <ISimpleBtnTriggerModal ref="triggerModal" btn-text="新增" modal-title="新增调度指令" :modal-width="600">
    <!-- 表单信息 -->
    <Form ref="formValidate" :model="formValidate" :rules="ruleValidate" :label-width="100">
      <FormItem label="任务名称" prop="task_name">
        <Input v-model="formValidate.task_name" placeholder="请输入任务名称"></Input>
      </FormItem>
      <FormItem label="任务类型" prop="task_type">
        <Input v-model="formValidate.task_type" placeholder="请输入任务类型"></Input>
      </FormItem>
      <FormItem label="任务ID" prop="task_id">
        <Input v-model="formValidate.task_id" placeholder="请输入任务ID"></Input>
      </FormItem>
      <FormItem label="cron 表达式" prop="cron_str">
        <Input v-model="formValidate.cron_str" placeholder="请输入 cron 表达式"></Input>
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
  import {AddQuartz} from "../../api"

  export default {
    name: "QuartzAdd",
    components:{ISimpleBtnTriggerModal},
    data(){
      return {
        formValidate: {
          task_name: '',
          task_type: '',
          task_id: '',
          cron_str: '',
        },
        ruleValidate: {
          task_name: [
            { required: true, message: '任务名称不能为空', trigger: 'blur' }
          ],
          task_type: [
            { required: true, message: '任务类型不能为空', trigger: 'blur' }
          ],
          task_id: [
            { required: true, message: '任务 ID 不能为空', trigger: 'blur' }
          ],
          cron_str: [
            { required: true, message: 'cron 表达式不能为空', trigger: 'blur' }
          ]
        },
      }
    },
    methods:{
      handleSubmit (name) {
        this.$refs[name].validate(async (valid) => {
          if (valid) {
            const result = await AddQuartz(this.formValidate.task_name,
              this.formValidate.task_type,this.formValidate.task_id,this.formValidate.cron_str);
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
