<template>
  <div id="content" style="width: 100%;">
    <div id="section">
      <div class="login_bg">
        <LoginForm/>
      </div>
    </div>
  </div>
</template>

<script>
  import LoginForm from "./LoginForm"
  export default {
    name: "LoginBg",
    components:{LoginForm}
  }
</script>

<style scoped>
  #section {
    width: 100%;
    height: 500px;
  }
  .login_bg{
    height:100%;
    width:100%;
    background: url(../../assets/sso/login_bg.jpg) 0px -300px no-repeat;
  }
</style>
