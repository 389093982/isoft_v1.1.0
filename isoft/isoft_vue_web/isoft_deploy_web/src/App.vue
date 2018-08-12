<template>
  <div id="app">
    <Header/>
    <div id="content">
      <div id="left">
        <LeftMenu/>
      </div>
      <div id="right" style="border:1px solid #D1D1D1;min-height: 500px;margin: 10px;">
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
  components: {
    LeftMenu,
    Header,
    Footer
  }
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
        width 15%
      #right
        width 85%
</style>
