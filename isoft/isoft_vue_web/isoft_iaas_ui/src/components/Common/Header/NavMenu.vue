<template>
  <div>
    <div class="nav" id="mainNav">
      <ul class="list">
        <li>
          <a href="#" class="nav_menu_hover">精品应用</a>
          <div class="nav_menu_content jpyy">
            <div class="nav_content">
              <div class="nav_li">
                <div class="nav_li_l">
                  博客天地
                </div>
                <div class="nav_li_r">
                  <router-link to="/iblog/blog_list">热门博文推荐</router-link>┊
                  <router-link to="/iblog/catalog_add">新增/编辑分类</router-link>┊
                  <router-link to="/iblog/blog_add">新增/编辑文章</router-link>┊
                  我的博客空间<br />
                </div>
              </div>
              <div class="nav_li">
                <div class="nav_li_l">
                  在线学习系统
                </div>
                <div class="nav_li_r">
                  <router-link to="/ilearning/index">精品课程</router-link>┊
                  <router-link to="/inote/index">云笔记</router-link><br />
                </div>
              </div>
              <div class="nav_li">
                <div class="nav_li_l">
                  云存储
                </div>
                <div class="nav_li_r">
                  <router-link to="/ifile/ifile">IFile 文件存储</router-link>┊
                  <router-link to="/ifile/ifile">IFile 对象存储</router-link>┊
                  <a href="http://sc.chinaz.com/?ng/weixing/" title="新游尾行汇总页" class="orange">更多>></a>
                </div>
              </div>
              <div class="nav_li">
                <div class="nav_li_l">
                  友情链接
                </div>
                <div class="nav_li_r">
                  <router-link to="/ifile/ifile">心声社区</router-link>┊
                  <router-link to="/ifile/ifile">论坛</router-link>┊
                  <router-link to="/easyshare/list">社区天地</router-link>
                </div>
              </div>
            </div>
          </div>
        </li>
        <li>
          <a href="#" class="nav_menu_hover">综合设置</a>
          <div class="nav_menu_content zhsz">
            <div class="nav_content">
              <div class="nav_li">
                <div class="nav_li_l">
                  综合设置
                </div>
                <div class="nav_li_r">
                  <a href="http://zq.sc.chinaz.com/?pc/?filter=0-17-0-0-3" title="更多" class="orange">更多>></a>
                </div>
              </div>
            </div>
          </div>
        </li>
        <li>
          <a href="#" class="nav_menu_hover">内容管理</a>
          <div class="nav_menu_content lrsz">
            <div class="nav_content">
              <div class="nav_li">
                <div class="nav_li_l">
                  配置项管理
                </div>
                <div class="nav_li_r">
                  <router-link to="/ilearning/configuration">查看配置项</router-link>┊
                  <a href="http://sc.chinaz.com/?web/" title="更多"class="orange">更多>></a>
                </div>
              </div>
            </div>
          </div>
        </li>
        <li>
          <a href="#" class="nav_menu_hover">
            <Icon type="ios-people" />
            <span v-if="loginUserName">{{loginUserName}}</span>
            <span v-else>登录</span>
          </a>
          <div class="nav_menu_content grzx">
            <div class="nav_content">
              <div class="nav_li">
                <div class="nav_li_l">
                  账号管理
                </div>
                <div class="nav_li_r">
                  <a href="javascript:;" @click="cancelUser">注销</a>┊
                  <a href="javascript:;" @click="cancelUser">切换账号</a>┊
                  <router-link to="/user/manage">个人中心</router-link>
                </div>
              </div>
            </div>
          </div>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
  import {getCookie} from '../../../tools/index'
  import {delCookie} from '../../../tools/index'

  export default {
    name: "NavMenu",
    data(){
      return {
        loginUserName:'',
      }
    },
    methods:{
      cancelUser() {
        delCookie("userName");
        this.loginUserName = "";
        window.location.href = "/api/auth/redirectToLogin/?redirectUrl=" + window.location.href;
      },
      getByClass (oParent, sClass) {
        var aEle = oParent.getElementsByTagName("*");
        var re = new RegExp("\\b" + sClass + "\\b");
        var arr = [];
        for (var i = 0; i < aEle.length; i++) {
          if (re.test(aEle[i].className)) {
            arr.push(aEle[i]);
          }
        }
        return arr;
      },
      setMainNav() {
        var oMainNav = document.getElementById("mainNav");
        var aLi = this.getByClass(oMainNav, "list")[0].getElementsByTagName("li");
        var navMenuHover = this.getByClass(oMainNav, "nav_menu_hover");
        var navMenuContent = this.getByClass(oMainNav, "nav_menu_content");
        for (var i = 0; i < navMenuHover.length; i++) {
          navMenuHover[i].index = i;
          navMenuHover[i].onmouseover = function () {
            this.className += " "+"nav_menu_hover_current";
            for (var j = 0; j < navMenuContent.length; j++) {
              navMenuContent[j].index_j = j;
              navMenuContent[j].style.display = "none";
              navMenuContent[j].onmouseover = function () {
                this.style.display = "block";
                navMenuHover[this.index_j].className += " "+"nav_menu_hover_current";
              };
              navMenuContent[j].onmouseout = function () {
                this.style.display = "none";
              };
            }
            if (navMenuContent[this.index]) {
              navMenuContent[this.index].style.display = "block";
            }
          }
        }
        for (var i = 0; i < aLi.length; i++) {
          aLi[i].index = i;
          aLi[i].onmouseout = function () {
            if (navMenuContent[this.index]) {
              navMenuContent[this.index].style.display = "none";
            }
            navMenuHover[this.index].className = "nav_menu_hover";
          }
        }
      }
    },
    mounted:function(){
      this.loginUserName = getCookie("userName");
      this.setMainNav();
    },
  }
</script>

<style scoped>
  *{margin:0;padding:0;}
  body{font:14px "微软雅黑", arial, serif;color:#333;}
  a,a:link,a:visited{color:#039;}
  img{border:0;}
  .nav{position:relative;left:4px;top:2px;width:100%;height:35px;background:#333;border-top:1px solid #444;}
  .nav .list{list-style-type:none;margin-left:15px;}
  .nav .list li{float:left;position:relative;height:35px;line-height:26px;}
  .nav .list li .nav_menu_hover{float:left;display:block;margin-top:5px;height:30px;padding:0 10px 0 10px;color:#ccc;font-weight:bold;text-decoration:none;}
  .nav .list li .nav_menu_hover_current,
  .nav .list li .nav_menu_hover:hover{background-color:#fff;color:#575757;margin-top:4px;padding:0 9px 0 9px;border:1px solid #666;border-bottom:0;}
  .nav_menu_content{display:none;position:absolute;width:auto;height:auto; top:35px;border:1px solid #666;border-top:0;border-bottom-width:2px;background:#fff;z-index:1000;}
  .nav_menu_content .nav_content{padding:15px;padding-bottom:0;}
  .nav_menu_content .nav_li{display:inline-block;width:100%;height:100%;*height:auto;*margin-top:7px;padding-bottom:5px;*padding-bottom:12px;border-bottom:1px dashed #ccc;}
  .nav_menu_content .nav_li_l{float:left;width:80px;color:#f60;font-weight:bold;}
  .nav_menu_content .nav_li_r{float:left;color:#999;font-family:"宋体";font-size:10px;line-height:26px;}
  .nav_menu_content .nav_li_r a{padding:0 1px 0 1px;color:#666;font-size:12px;text-decoration:none;}
  .nav_menu_content .nav_li_r a:hover{color:red;}
  .nav_menu_content .nav_li_r a.orange{color:#f60;}
  .nav .list .jpyy{width:720px;left:0;}
  .nav .list .zhsz{width:800px;left:0;}
  .nav .list .lrsz{width:740px;left:0;}
  .nav .list .grzx{width:610px;left:0;}
</style>
