<template>
  <div id="login_area">
    <div id="login_header">用户名密码登录</div>
    <div id="login_content" style="margin: 0 auto;padding:20px;">
      <input class="focus" name="username" placeholder="请输入用户名" type="text" style="width: 100%;height: 40px;" required/>
      <input type="password" style="display:none">
      <input class="focus" name="passwd" placeholder="请输入密码" type="password"
             style="width: 100%;height: 40px;margin-top:20px;" autocomplete="new-password" required/>
      <span id="error_msg" v-if="showError">{{errorMsg}}</span>
      <p>
        <input id="submit" type="submit" value="登录" @click="login">
        <router-link :to="{path:'/sso/user/forgetPwd'}" style="float: right;color: #2e82ff;">忘记密码？</router-link>
      </p>
    </div>
    <div id="login_footer">
      <router-link :to="{path:'/sso/user/friendLogin'}" style="float: left;color: #2e82ff;">友情登录</router-link>
      <router-link :to="{path:'/sso/user/regist'}" style="float: right;color: #2e82ff;">立即注册</router-link>
    </div>
  </div>
</template>

<script>
  import {Login} from "../../api"

  export default {
    name: "LoginForm",
    data(){
      return {
        showError:false,
        errorMsg:"登录失败!",
      }
    },
    methods:{
      login:async function () {
        var username = $("input[name='username']").val();
        var passwd = $("input[name='passwd']").val();
        var result = await Login(username, passwd);
        if(result.loginSuccess == true){
          alert("登录成功!");
        }else{
          this.showError = true;
          this.errorMsg = result.loginStatus;
        }
      }
    }
  }
</script>

<style scoped>
  #login_area{
    width: 320px;
    height: 450px;
    background: #ffffff;
    float: right;
    margin-top: 20px;
    margin-right: 150px;
    position: relative;
  }
  #login_header{
    height: 80px;
    line-height: 80px;
    text-align: center;
    font-size: 16px;
    color: #000;
  }
  #submit{
    width: 100%;
    height: 40px;
    margin-top:30px;
    margin-bottom:20px;
    display: block;
    line-height: 40px;
    font-size: 16px;
    font-weight: 800;
    cursor: pointer;
    color: #fff;
    background: #3f89ec;
    border: 0;
  }
  #login_footer{
    color: #2e82ff;
    position:absolute;
    bottom:0;
    width:100%;
    height: 60px;
    line-height: 60px;
    background: #f0f6ff;
    padding: 0 28px;
  }
  #error_msg{
    font-size: 12px;
    color:red;
    float: right;
    margin-top: 5px;
  }
  a {
    color: #666;
    text-decoration: none;
  }
  a:hover {
    color: #E4393C;
    text-decoration: underline;
  }
  .focus:focus {
    background-color: #ffffff;
    border-color: #2c5bff;
  }
</style>
