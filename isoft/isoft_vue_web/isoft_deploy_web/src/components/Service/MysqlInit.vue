<template>
    <div :index="index">
      <RadioGroup v-model="mysql_init_group">
        <Radio label="创建数据库"></Radio>
        <Radio label="创建用户"></Radio>
        <Radio label="创建用户并指定数据库"></Radio>
      </RadioGroup>
      <!-- 表单正文 -->
      <Form ref="formValidate" :model="formValidate" :rules="ruleValidate" :label-width="100">
        <FormItem label="新建账号" prop="create_account" v-if="mysql_init_group == '创建用户' || mysql_init_group == '创建用户并指定数据库'">
          <Input v-model="formValidate.create_account" placeholder="请输入新建账号名称"></Input>
        </FormItem>
        <FormItem label="账号密码" prop="create_passwd" v-if="mysql_init_group == '创建用户' || mysql_init_group == '创建用户并指定数据库'">
          <Input v-model="formValidate.create_passwd" placeholder="请输入新建账号密码"></Input>
        </FormItem>
        <FormItem label="新建数据库" prop="create_database" v-if="mysql_init_group == '创建数据库' || mysql_init_group == '创建用户并指定数据库'">
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
        mysql_init_group:'创建数据库',
        formValidate: {
          create_account: '',
          create_passwd: '',
          create_database: '',
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
          create_account:this.formValidate.create_account,
          create_passwd:this.formValidate.create_passwd,
          create_database:this.formValidate.create_database
        };

        var _this = this;
        this.$refs[name].validate((valid) => {
          if (valid) {
            if(_this.mysql_init_group == "创建用户"){
              data.create_database = '';
            }else if(_this.mysql_init_group == "创建数据库"){
              data.create_account = '';
              data.create_passwd = '';
            }
            _this.$emit('handleSubmit', data);
          } else {
            alert(111);
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
