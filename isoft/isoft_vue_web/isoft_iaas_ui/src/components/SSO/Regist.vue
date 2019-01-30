<template>
  <div>
    <div id="header">
      <div class="login-link" id="login_link" style="position: absolute;top: 25px;right: 100px;">
        <span style="font-size: 15px;font-weight: inherit;">已有账号,前去<a href="/user/login/">登录</a></span>
      </div>
    </div>
    <div id="nav">
    </div>
    <div id="content" style="width: 100%;">
      <div id="section">
        <div style="margin:80px;margin-left:200px;margin-right: 200px;">
          <div style="text-align: center;padding-left: 50px;">
            <span style="height: 60px;line-height: 60px;font-size: 16px;color: #000;">用户注册</span>
            <span><a href="javascript:;">已有账号,前去登录</a></span>
          </div>
          <Form ref="formValidate" :model="formValidate" :rules="ruleValidate" :label-width="80">
            <FormItem label="用户名" prop="username">
              <Input v-model.trim="formValidate.username" placeholder="请输入用户名"></Input>
            </FormItem>
            <FormItem label="密码" prop="passwd">
              <Input v-model.trim="formValidate.passwd" type="password" placeholder="请输入密码"></Input>
            </FormItem>
            <FormItem label="确认密码" prop="repasswd">
              <Input v-model.trim="formValidate.repasswd" type="password" placeholder="请输入确认密码"></Input>
            </FormItem>
            <FormItem label="用户协议" prop="proxy">
              <CheckboxGroup v-model="formValidate.proxy">
                <Checkbox label="用户协议">
                  <label>阅读并接受</label><a href="#">《Isoft用户协议》</a>
                </Checkbox>
              </CheckboxGroup>
            </FormItem>
            <FormItem>
              <input value="注册" id="submit" @click="handleSubmit('formValidate')">
            </FormItem>
          </Form>
        </div>
      </div>
      <aside id="asideright">
        <div style="margin: 80px;background: #ebfffc;height: 300px;padding:20px;">
          <h3 style="background: url('../../assets/sso/phone.png') left center no-repeat;">
            <span style="padding-left: 30px;">账号特权</span>
          </h3>
          <hr>
          <div style="font-size: 12px;font-family: Tahoma, Helvetica, 'Microsoft Yahei', 微软雅黑, Arial, STHeiti;">
            <p style="line-height: 30px;">初次注册账号送30小时免费学习时间</p>
            <p style="line-height: 30px;">初次注册账号送3000积分</p>
            <p style="line-height: 30px;">初次注册账号送云笔记使用特权</p>
          </div>
        </div>
      </aside>
    </div>

    <LoginFooter/>
  </div>
</template>

<script>
  import LoginFooter from "./LoginFooter"
  import {Regist} from "../../api"
  import {validateUserName} from "../../tools"
  import {validatePasswd} from "../../tools"

  export default {
    name: "Regist",
    components:{LoginFooter},
    data(){
      const _validateUserName = (rule, value, callback) => {
        if (value === '') {
          callback(new Error('用户名不能为空!'));
        } else if (!validateUserName(value)) {
          callback(new Error('6至20位，以字母开头，字母，数字，减号，下划线!'));
        } else {
          callback();
        }
      };
      const _validatePasswd = (rule, value, callback) => {
        if (value === '') {
          callback(new Error('密码不能为空!'));
        } else if (!validatePasswd(value)) {
          callback(new Error('最少6位，至少1个大小写字母，数字和特殊字符!'));
        } else {
          callback();
        }
      };
      // 确认密码校验 validatePassCheck
      const validatePassCheck = (rule, value, callback) => {
        if (value === '') {
          callback(new Error('请输入密码进行确认!'));
        } else if (value !== this.formValidate.passwd) {
          callback(new Error('两次输入密码不一致!'));
        } else {
          callback();
        }
      };
      return {
        formValidate: {
          username: '',
          passwd: '',
          repasswd: '',
          proxy:[],
        },
        ruleValidate: {
          username: [
            { validator: _validateUserName, trigger: 'blur' }
          ],
          passwd: [
            { validator: _validatePasswd, trigger: 'blur' },
          ],
          repasswd: [        // 确认密码校验 validatePassCheck
            { validator: validatePassCheck, trigger: 'blur' }
          ],
          proxy: [
            { required: true, type: 'array', min: 1, message: '用户协议必须同意!', trigger: 'change' },
          ],
        }
      }
    },
    methods:{
      handleSubmit: function (name) {
        this.$refs[name].validate(async (valid) => {
          if (valid) {
            // 校验通过则进行注册
            this.regist();
          } else {
            this.$Message.error('注册信息校验失败!');
          }
        })
      },
      regist:async function () {
        var username = this.formValidate.username;
        var passwd = this.formValidate.passwd;
        const result = await Regist(username,passwd);
        if(result.status=="SUCCESS"){
          this.$Message.success('注册成功!');
        }else{
          if(result.errorMsg == "regist_exist"){
            this.$Message.error("该用户已经被注册!");
          }else if(result.errorMsg == "regist_failed"){
            this.$Message.error("注册失败,请联系管理员获取账号!");
          }
        }
      }
    }
  }
</script>

<style scoped>
  #header {
    background-color: rgb(255, 255, 255);
    text-align:center;
    height:70px;
    padding:5px;
  }
  #nav {
    height: 20px;
    display: block;
    width:100%;
    background: linear-gradient(red, blue);
    opacity:0.1;
  }
  #section {
    width: 60%;
    float:left;
    height: 450px;
  }
  #asideright{
    width:40%;
    float: left;
    height: 450px;
  }
  a:hover {
    color: #E4393C;
    text-decoration: underline;
  }
  .focus:focus {
    background-color: #ffffff;
    border-color: #2c5bff;
  }
  #user_proxy{
    font-size: 12px;
    color:red;
    float: right;
  }
  #submit{
    width: 100%;
    height: 40px;
    display: block;
    line-height: 40px;
    font-size: 16px;
    font-weight: 800;
    cursor: pointer;
    color: #fff;
    background: #3f89ec;
    border: 0;
    text-align: center;
  }
</style>
