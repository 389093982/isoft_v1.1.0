<template>
    <div :index="index">
      <!-- 表单正文 -->
      <Form ref="formValidate" :model="formValidate" :rules="ruleValidate" :label-width="100">
        <FormItem label="新建账号" prop="create_account">
          <Input v-model="formValidate.create_account" placeholder="请输入新建账号名称"></Input>
        </FormItem>
        <FormItem label="账号密码" prop="create_passwd">
          <Input v-model="formValidate.create_passwd" placeholder="请输入新建账号密码"></Input>
        </FormItem>
        <FormItem label="新建数据库" prop="create_database">
          <Input v-model="formValidate.create_database" placeholder="请输入新建数据库名称"></Input>
        </FormItem>
        <FormItem>
          <Button type="primary" @click="handleSubmit('formValidate')">Submit</Button>
          <Button @click="handleReset('formValidate')" style="margin-left: 8px">Reset</Button>
        </FormItem>
      </Form>
    </div>
</template>

<script>
  export default {
    name: "MysqlInit",
    props:{
      index:{
        type:Number
      }
    },
    data () {
      return {
        formValidate: {
          create_account: '',
          create_passwd: '',
          create_database: ''
        },
        ruleValidate: {
          create_account: [
            { required: true, message: '新建账号名称不能为空', trigger: 'blur' }
          ],
          create_passwd: [
            { required: true, message: '新建账号密码不能为空', trigger: 'blur' }
          ],
          create_database: [
            { required: true, message: '新建数据库名称不能为空', trigger: 'blur' }
          ]
        }
      }
    },
    methods: {
      handleSubmit (name) {
        let data = {
          index:this.index,
          create_account:this.formValidate.create_account,
          create_passwd:this.formValidate.create_passwd,
          create_database:this.formValidate.create_database
        };
        var _this = this;
        this.$refs[name].validate((valid) => {
          if (valid) {
            _this.$emit('handleSubmit', data);
          } else {

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
