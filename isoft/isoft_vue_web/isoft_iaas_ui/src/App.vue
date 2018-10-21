<template>
  <div id="app">
    <Layout :style="{minHeight: '550px'}">
      <Header>
        <Header/>
      </Header>
      <Content :style="{padding: '18px', minHeight: '550px', background: '#f5f6f7', lineHeight: '24px', fontSize: '14px'}">
        <router-view/>
      </Content>
      <Footer>
        <Footer/>
      </Footer>
    </Layout>
  </div>
</template>

<script>
  import {mapActions} from 'vuex'

  import Header from './components/Common/Header/Header.vue'
  import Footer from './components/Common/Footer/Footer.vue'

  export default {
    name: 'App',
    components: {
      Header,
      Footer
    },
    mounted () {
      // this.$store.dispatch('getHotCourseTypeConfigurations')
      this.getHotCourseTypeConfigurations()
    },
    methods: {
      ...mapActions(['getHotCourseTypeConfigurations'])
    },
  }
</script>

<style lang="stylus" type="text/stylus" rel="stylesheet/stylus">

</style>
