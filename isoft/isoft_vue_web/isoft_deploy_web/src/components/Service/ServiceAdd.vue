<template>
<div>
  <Button type="success" @click="modal1 = true">新增</Button>
  <Modal
    v-model="modal1"
    title="新增/编辑服务信息"
    :footer-hide="true"
    :mask-closable="false">       <!-- 是否允许点击遮罩层关闭 -->
    <div>
      <!-- 表单正文 -->
      <Form ref="formValidate" :model="formValidate" :rules="ruleValidate" :label-width="80">
        <FormItem label="环境名称" prop="env_name">
          <Select v-model="formValidate.env_name" filterable>
            <Option v-for="item in env_names" :value="item.value" :key="item.value">{{ item.label }}</Option>
          </Select>
        </FormItem>
        <FormItem label="服务名称" prop="service_name">
          <Input v-model="formValidate.service_name" placeholder="请输入服务名称"></Input>
        </FormItem>
        <FormItem label="服务类型" prop="service_type">
          <Select v-model="formValidate.service_type" filterable>
            <Option v-for="item in service_types" :value="item.value" :key="item.value">{{ item.label }}</Option>
          </Select>
        </FormItem>
        <FormItem label="端口号" prop="service_port">
          <Input v-model="formValidate.service_port" placeholder="请输入端口号"></Input>
        </FormItem>
        <FormItem>
          <Button type="primary" @click="handleSubmit('formValidate')">Submit</Button>
          <Button @click="handleReset('formValidate')" style="margin-left: 8px">Reset</Button>
        </FormItem>
      </Form>
    </div>
  </Modal>
</div>
</template>

<script>
  import {EnvEdit} from '../../api'
  import {EnvAll} from '../../api'
  import {ServiceEdit} from '../../api'

  export default {
    data () {
      return {
        modal1: false,    // 遮罩层
        formValidate: {
          env_name: '',
          service_name: '',
          service_type: '',
          service_port: ''
        },
        ruleValidate: {
          env_name: [
            { required: true, message: '环境名称不能为空', trigger: 'blur' }
          ],
          service_name: [
            { required: true, message: '服务名称不能为空', trigger: 'blur' }
          ],
          service_type: [
            { required: true, message: '服务类型不能为空', trigger: 'blur' }
          ],
          service_port: [
            { required: true, message: '端口号不能为空', trigger: 'blur' }
          ]
        },
        env_names:[
          {
            value: 'env_1',
            label: 'env_1'
          },
          {
            value: 'env_2',
            label: 'env_2'
          }
        ],
        service_types: [
          {
            value: 'beego',
            label: 'beego'
          },
          {
            value: 'docker',
            label: 'docker'
          },
          {
            value: 'nginx',
            label: 'nginx'
          },
          {
            value: 'mysql',
            label: 'mysql'
          },
          {
            value: 'api',
            label: 'api'
          }
        ]
      }
    },
    methods: {
      // 关闭模态对话框
      closeModalDialog (){
        this.modal1 = false;
      },
      handleSubmit (name) {
        var _this = this;
        this.$refs[name].validate((valid) => {
          if (valid) {
            // 返回的是个 Promise 对象,需要调用 then 方法
            ServiceEdit(this.formValidate.env_name,this.formValidate.service_name,this.formValidate.service_type,
              this.formValidate.service_port)
              .then(function (response) {
                if(response.status == "SUCCESS"){
                  _this.$Message.success('提交成功!');
                  _this.closeModalDialog();
                  _this.$router.go(0);     // 页面刷新,等价于 location.reload()
                }else{
                  _this.$Message.error('提交失败!');
                }
              });
          } else {

          }
        })
      },
      handleReset (name) {
        this.$refs[name].resetFields();
      }
    },
    mounted:function () {
      EnvAll().then(function (response) {
        alert(response);
        if(response.status == "SUCCESS"){
          alert(response);
        }else{
        }
      });
    }
  }
</script>

<style scoped>

</style>
