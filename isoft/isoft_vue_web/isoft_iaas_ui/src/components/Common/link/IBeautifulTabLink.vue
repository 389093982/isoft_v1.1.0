<template>
  <div class="container">
    <div class="price">
      <h2>{{pData.title}}</h2>
      <div class="category">
        <ul>
          <li :class="index == current_index ? 'active' : ''" v-for="(itemData, index) in pData.pItemData"
            @click="current_index = index">{{itemData.itemTitle}}</li>
        </ul>
      </div>
      <div class="cont" :class="index == current_index ? 'active' : ''" v-for="(itemData, index) in pData.pItemData">
        <ul>
          <li v-for="item in itemData.items">
            <a :href="item.item_link">
              <div class="cont_main">
                <h3 class="title">{{item.item_label}}</h3>
                <p class="desc">{{item.item_desc}}</p>
              </div>
            </a>
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    name: "IBeautifulTabLink",
    data(){
      return {
        current_index : 0,
      }
    },
    props: {
      // 带有默认值的对象
      pData: {
        type: Object,
        // 对象或数组默认值必须从一个工厂函数获取
        default: function () {
          return {
            title: '产品定价',
            pItemData:[
              {
                itemTitle:'热门',
                items:[
                  {item_label:'云服务器1',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器1',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器1',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器1',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器1',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器1',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器1',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                ],
              },
              {
                itemTitle:'计算',
                items:[
                  {item_label:'云服务器2',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器2',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器2',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器2',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器2',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器2',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器2',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                ],
              },
              {
                itemTitle:'存储',
                items:[
                  {item_label:'云服务器3',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器3',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器3',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器3',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器3',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器3',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器3',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                ],
              },
              {
                itemTitle:'视频服务',
                items:[
                  {item_label:'云服务器4',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器4',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器4',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器4',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器4',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器4',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器4',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                ],
              },
              {
                itemTitle:'数据库',
                items:[
                  {item_label:'云服务器5',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器5',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器5',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器5',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器5',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器5',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器5',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                ],
              }
            ]
          }
        }
      },
    },
    methods:{

    },
  }
</script>

<style scoped>
  * {
    margin: 0;
    padding: 0;
    font-size: 14px;
    font-family: "microsoft yahei";
  }

  ul {
    list-style: none;
  }

  a {
    text-decoration: none;
  }

  .container {
    background: #f7f8fa;
    height: auto;
    overflow: hidden;
    padding-top: 15px;
    padding-bottom: 15px;
  }

  .price {
    width: 100%;
    margin: 0 auto;
  }

  .price h2 {
    font-size: 24px;
    color: #333;
    line-height: 1;
    font-weight: 400;
    padding-bottom: 10px;
    padding-left: 10px;
  }

  .price .category {
    width: 100%;
    overflow: hidden;
    position: relative;
    overflow: hidden;
    margin-bottom: 20px;
    margin-left: 10px;
  }

  .price .category:after {
    content: "";
    position: absolute;
    width: 100%;
    left: 0;
    top: 0;
    height: 1px;
    background-color: #EDEDED;
    z-index: 2;
  }

  .price .category:before {
    content: "";
    position: absolute;
    width: 100%;
    left: 0;
    bottom: 0;
    height: 1px;
    background-color: #EDEDED;
    z-index: 2;
  }

  .price .category a:after {
    content: "";
    display: inline-block;
    width: 0;
    height: 100%;
    vertical-align: middle;
  }

  .price .category ul li {
    display: table-cell;
    white-space: nowrap;
    font-size: 14px;
    text-align: center;
    border-bottom: 3px solid transparent;
    cursor: pointer;
    position: relative;
    padding: 18px 25px;
  }

  .price .category ul li.active {
    border-bottom-color: #00a4ff;
    color: #00a4ff;
  }

  .price .cont {
    display: none;
  }

  .price .cont.active {
    display: block;
  }

  .price .cont ul li {
    width: 280px;
    float: left;
    margin: 0 10px;
    box-sizing: border-box;
  }

  .price .cont ul li:nth-child(4n) {
    margin-right: 0px!important;
  }

  .price .cont ul li a {
    display: block;
    border: 1px solid #E5E5E5;
    background-color: #fff;
    margin-top: 20px;
    padding: 13px 20px 15px;
  }

  .price .cont ul li a:hover {
    box-shadow: 0 3px 7px 0 rgba(0, 0, 0, .1);
    border-color: rgba(229, 229, 229, .73);
  }

  .price .cont ul li a:hover .cont_main h3.title {
    color: #00a4ff;
  }

  .price .cont ul li a .cont_main {
    display: table-cell;
    vertical-align: top;
  }

  .price .cont ul li a .cont_main h3.title {
    font-size: 18px;
    font-weight: 400;
    color: #333;
    line-height: 32px;
    white-space: nowrap;
    overflow: hidden;
    display: block;
    text-overflow: ellipsis;
  }

  .price .cont ul li a .cont_main p.desc {
    font-size: 14px;
    line-height: 1.5;
    color: #666;
    height: 42px;
    margin-top: -2px;
    display: -webkit-box;
    -webkit-box-orient: vertical;
    -webkit-line-clamp: 2;
    overflow: hidden;
  }
</style>
