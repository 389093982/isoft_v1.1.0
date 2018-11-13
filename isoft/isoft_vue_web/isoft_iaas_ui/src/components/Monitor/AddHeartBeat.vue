<template>
  <span>
    <Button type="success" @click="showAddHeartBeat = true">新增</Button>
    <Modal
      v-model="showAddHeartBeat"
      width="500"
      title="新增应用监控"
      :footer-hide="true"
      :mask-closable="false">
      <!-- 表单正文 -->
      <Form ref="formValidate" :model="formValidate" :rules="ruleValidate" :label-width="100">
        <FormItem label="应用地址" prop="addr">
          <Input v-model="formValidate.addr" placeholder="请输入应用地址"></Input>
        </FormItem>
        <FormItem>
          <Button type="primary" @click="handleSubmit('formValidate')">Submit</Button>
          <Button @click="handleReset('formValidate')" style="margin-left: 8px">Reset</Button>
        </FormItem>
      </Form>
    </Modal>
  </span>
</template>

<script>
  import {RegisterHeartBeat} from "../../api"

  export default {
    name: "AddHeartBeat",
    data(){
      return {
        showAddHeartBeat:false,
        formValidate: {
          addr: '',
        },
        ruleValidate: {
          addr: [
            { required: true, message: '应用地址不能为空', trigger: 'blur' }
          ]
        }
      }
    },
    methods: {
      handleSubmit (name) {
        var _this = this;
        this.$refs[name].validate(async (valid) => {
          if (valid) {
            const result = await RegisterHeartBeat(_this.formValidate.addr);
            if(result.status == "SUCCESS"){
              this.showAddHeartBeat = false;
              this.$emit('refreshHeartBeatTable');
            }else{
              _this.$Message.error('提交失败!');
            }
          } else {
            _this.$Message.error('验证失败!');
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
