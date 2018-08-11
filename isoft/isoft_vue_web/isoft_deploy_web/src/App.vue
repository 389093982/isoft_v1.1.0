<template>
  <div id="app">
    <Header/>
    <div id="content">
      <div id="left">
        <LeftMenu/>
      </div>
      <div id="right">
        <router-link to="/" class="list-group-item menu-item" style="padding: 8px;">home主页</router-link>
        <router-link to="/login" class="list-group-item menu-item" style="padding: 8px;">我的工作</router-link>
        <router-view/>
      </div>
    </div>
    <Footer/>
  </div>
</template>

<script>
import Header from './components/Header/Header.vue'
import Footer from './components/Footer/Footer.vue'
import LeftMenu from './components/Menu/LeftMenu.vue'

export default {
  name: 'App',
  components: {LeftMenu, Header,Footer}
}
</script>

<style lang="stylus" type="text/stylus" rel="stylesheet/stylus">
  #app
    width 100%
    height 100%
    #content
      width 95%
      margin 0 auto
      // css3 盒子模型
      display -webkit-box
      min-height 700px
      #left
        width 20%
      #right
        width 80%
</style>
