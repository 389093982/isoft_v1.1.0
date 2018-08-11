<template>
  <div class="account-l fl" style="margin: 10px;">
    <a class="list-title">{{ title }}</a>
    <ul id="accordion" class="accordion">
      <li v-for="(link,index) in links">
        <div class="link" @click="dropdown($event,index)">
          {{ link.title }}<DownArrow :style="{'float':'right','margin-right':'5%'}" :open="current==index"/>
        </div>
        <ul class="submenu">
          <li v-for="hrefinfo in link.hrefinfos" @click="chooseCurrent($event)">
            <router-link :to="hrefinfo.hrefaddr" style="padding: 8px;">{{ hrefinfo.hrefdesc }}</router-link>
          </li>
        </ul>
      </li>
    </ul>
  </div>
</template>

<script>
  import DownArrow from './DownArrow.vue'

  export default {
    name: "LeftMenu",
    components: {DownArrow},
    data(){
      return {
        current:-1,    // 当前选中的 index
        title:'统一部署管理系统',
        links:[
          {
            title:'部署环境管理',
            hrefinfos:[
              {hrefaddr:'/env/list',hrefdesc:'环境清单'}
            ]
          },
          {
            title:'基础软件管理',
            hrefinfos:[
              {hrefaddr:'/env/list',hrefdesc:'java运行环境管理'},
              {hrefaddr:'/env/list',hrefdesc:'python运行环境管理'},
              {hrefaddr:'/env/list',hrefdesc:'golang运行环境管理'},
              {hrefaddr:'/env/list',hrefdesc:'mysql数据库管理'}
            ]
          },
          {
            title:'应用服务管理',
            hrefinfos:[
              {hrefaddr:'/service/list?service_type=mysql',hrefdesc:'mysql管理'},
              {hrefaddr:'/service/list?service_type=nginx',hrefdesc:'nginx管理'},
              {hrefaddr:'/service/list?service_type=beego',hrefdesc:'beego应用部署'},
              {hrefaddr:'/service/list?service_type=api',hrefdesc:'api应用部署'},
              {hrefaddr:'/service/monitor',hrefdesc:'服务监控'}
            ]
          }
        ]
      }
    },
    methods:{
      // 点击菜单显示下拉
      dropdown(event,index){
        this.current=index;   // 当前菜单的索引
        var el = event.currentTarget;
        $(el).next().slideToggle();
        $(el).parent().toggleClass('open');
        $(el).parent().siblings().find('.submenu').slideUp().parent().removeClass('open');
      },
      // 选中当前设置样式
      chooseCurrent(event){
        var el = event.currentTarget;
        $(el).addClass('current').siblings('li').removeClass('current');
      }
    }
  }
</script>

<style scoped type="text/stylus" rel="stylesheet/stylus">
  @import '../../components/Menu/leftmenu.css'
</style>
