<template>
  <div id="app">
    <Layout :style="{minHeight: '100vh'}">
      <Header>
        <Header/>
      </Header>
      <Layout>
        <Sider hide-trigger :style="{background: '#fff'}">
          <LeftMenu/>
        </Sider>
        <Content :style="{padding: '24px', minHeight: '550px', background: '#fff'}">
          <router-view/>
        </Content>
      </Layout>
      <Footer>
        <Footer/>
      </Footer>
    </Layout>

  </div>
</template>

<script>
import Header from './components/Common/Header/Header.vue'
import Footer from './components/Common/Footer/Footer.vue'
import LeftMenu from './components/Common/Menu/LeftMenu.vue'

export default {
  name: 'App',
  components: {
    LeftMenu,
    Header,
    Footer
  },

}
</script>

<style lang="stylus" type="text/stylus" rel="stylesheet/stylus">

</style>
