<template>
<div style="margin: 10px;margin-top: 20px;">
  <Button type="success" @click="modal1 = true">新增</Button>
  <Modal
    v-model="modal1"
    title="新增/编辑环境信息"
    :footer-hide="true"
    :mask-closable="false">       <!-- 是否允许点击遮罩层关闭 -->
    <div>
      <!-- 表单正文 -->
      <Form ref="formValidate" :model="formValidate" :rules="ruleValidate" :label-width="80">
        <FormItem label="环境名称" prop="env_name">
          <Input v-model="formValidate.env_name" placeholder="请输入环境名称"></Input>
        </FormItem>
        <FormItem label="ip地址" prop="env_ip">
          <Input v-model="formValidate.env_ip" placeholder="请输入ip地址"></Input>
        </FormItem>
        <FormItem label="登录账号" prop="env_account">
          <Input v-model="formValidate.env_account" placeholder="请输入登录账号"></Input>
        </FormItem>
        <FormItem label="登录密码" prop="env_passwd">
          <Input v-model="formValidate.env_passwd" placeholder="请输入登录密码"></Input>
        </FormItem>
        <FormItem label="部署空间" prop="deploy_home">
          <Input v-model="formValidate.deploy_home" placeholder="请输入部署空间目录"></Input>
        </FormItem>
        <FormItem>
          <Button type="primary" @click="handleSubmit('formValidate')">Submit</Button>
          <Button @click="handleReset('formValidate')" style="margin-left: 8px">Reset</Button>
        </FormItem>
      </Form>
    </div>
  </Modal>

  <EnvTable/>
</div>
</template>

<script>
  import {EnvEdit} from '../../api'
  import EnvTable from "./EnvTable";

  export default {
    components: {EnvTable},
    data () {
      return {
        modal1: false,    // 遮罩层
        formValidate: {
          env_name: '',
          env_ip: '',
          env_account: '',
          env_passwd: '',
          deploy_home: ''
        },
        ruleValidate: {
          env_name: [
            { required: true, message: '环境名称不能为空', trigger: 'blur' }
          ],
          env_ip: [
            { required: true, message: '环境IP不能为空', trigger: 'blur' }
          ],
          env_account: [
            { required: true, message: '登录账号不能为空', trigger: 'blur' }
          ],
          env_passwd: [
            { required: true, message: '登录密码不能为空', trigger: 'blur' }
          ],
          deploy_home: [
            { required: true, message: '部署空间目录不能为空', trigger: 'blur' }
          ]
        }
      }
    },
    methods: {
      // 关闭模态对话框
      closeModalDialog (){
        this.modal1 = false;
      },
      handleSubmit (name) {
        this.$refs[name].validate((valid) => {
          if (valid) {
            // 返回的是个 Promise 对象,需要调用 then 方法
            EnvEdit(this.formValidate.env_name,this.formValidate.env_ip,this.formValidate.env_account,
              this.formValidate.env_passwd,this.formValidate.deploy_home)
              .then(function (response) {
                var result = JSON.parse(response);
                if(result.status == "SUCCESS"){
                }else{
                  alert("保存失败!");
                }
              });
            this.$Message.success('Success!');
            this.closeModalDialog();
          } else {
            this.$Message.error('Fail!');
          }
        })
      },
      handleReset (name) {
        this.$refs[name].resetFields();
      }
    }
  }
</script>

<style scoped>
</style>
