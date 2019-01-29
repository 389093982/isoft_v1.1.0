<template>
  <div>
    <LoginHeader/>
    <ImproveAccount/>
    <LoginBg/>
    <LoginFooter/>
  </div>
</template>

<script>
  import LoginHeader from "./LoginHeader"
  import ImproveAccount from "./ImproveAccount"
  import LoginBg from "./LoginBg"
  import LoginFooter from "./LoginFooter"

  export default {
    name: "Login",
    components:{LoginHeader,ImproveAccount,LoginBg,LoginFooter}
  }
</script>

<style scoped>

</style>
