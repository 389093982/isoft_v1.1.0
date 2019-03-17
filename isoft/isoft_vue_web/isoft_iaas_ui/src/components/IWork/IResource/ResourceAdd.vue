<template>
  <!-- 按钮触发模态框 -->
  <!-- ref 的作用是为了在其它地方方便的获取到当前子组件 -->
  <ISimpleBtnTriggerModal ref="triggerModal" btn-text="新增" modal-title="新增资源信息" :modal-width="600">
    <!-- 表单信息 -->
    <Form ref="formValidate" :model="formValidate" :rules="ruleValidate" :label-width="140">
      <FormItem label="resource_name" prop="resource_name">
        <Input v-model.trim="formValidate.resource_name" placeholder="请输入 resource_name"></Input>
      </FormItem>
      <FormItem label="resource_type" prop="resource_type">
        <Input v-model.trim="formValidate.resource_type" placeholder="请输入 resource_type"></Input>
      </FormItem>
      <FormItem label="resource_url" prop="resource_url">
        <Input v-model.trim="formValidate.resource_url" placeholder="请输入 resource_url"></Input>
      </FormItem>
      <FormItem label="resource_dsn" prop="resource_dsn">
        <Input v-model.trim="formValidate.resource_dsn" placeholder="请输入 resource_dsn"></Input>
      </FormItem>
      <FormItem label="resource_username" prop="resource_username">
        <Input v-model.trim="formValidate.resource_username" placeholder="请输入 resource_username"></Input>
      </FormItem>
      <FormItem label="resource_password" prop="resource_password">
        <Input v-model.trim="formValidate.resource_password" placeholder="请输入 resource_password"></Input>
      </FormItem>
      <FormItem>
        <Button type="success" @click="handleSubmit('formValidate')" style="margin-right: 6px">Submit</Button>
        <Button type="warning" @click="handleReset('formValidate')" style="margin-right: 6px">Reset</Button>
      </FormItem>
    </Form>
  </ISimpleBtnTriggerModal>
</template>

<script>
  import ISimpleBtnTriggerModal from "../../Common/modal/ISimpleBtnTriggerModal"
  import {AddResource} from "../../../api/index"

  export default {
    name: "ResourceAdd",
    components:{ISimpleBtnTriggerModal},
    data(){
      return {
        formValidate: {
          resource_name: '',
          resource_type: '',
          resource_url: '',
          resource_dsn: '',
          resource_username: '',
          resource_password: '',
          env_name: '',
        },
        ruleValidate: {
          resource_name: [
            { required: true, message: 'resource_name 不能为空!', trigger: 'blur' }
          ],
          resource_type: [
            { required: true, message: 'resource_type 不能为空!', trigger: 'blur' }
          ],
          resource_url: [
            { required: true, message: 'resource_url 不能为空!', trigger: 'blur' }
          ],
          resource_dsn: [
            { required: true, message: 'resource_dsn 不能为空!', trigger: 'blur' }
          ],
          resource_username: [
            { required: true, message: 'resource_username 不能为空!', trigger: 'blur' }
          ],
          resource_password: [
            { required: true, message: 'resource_password 不能为空!', trigger: 'blur' }
          ],
        },
      }
    },
    methods:{
      handleSubmit (name) {
        this.$refs[name].validate(async (valid) => {
          if (valid) {
            const result = await AddResource(this.formValidate.resource_name,
              this.formValidate.resource_type,this.formValidate.resource_url,this.formValidate.resource_dsn,
              this.formValidate.resource_username,this.formValidate.resource_password);
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



