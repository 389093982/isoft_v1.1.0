<template>
  <div>
    <!-- 热门分类 -->
    <HotShareItem @chooseItem="chooseItem"/>

    <div style="margin: 0 15px;background-color: #fff;border: 1px solid #e6e6e6;border-radius: 4px;">
      <Row>
        <Col span="16" style="padding: 0 0 20px;border-right: 1px solid #e6e6e6;">
          <div style="border-bottom: 1px solid #e6e6e6;padding: 20px;height: 62px;">
            <Row>
              <Col span="4" style="text-align: center;font-size: 20px;color: #333;">
                <span v-if="search_type==='_all'">全部分类</span>
                <span v-else-if="search_type==='_hot'">热门分享</span>
                <span v-else-if="search_type==='_personal'">我的分享</span>
                <span v-else>{{search_type}}</span>
              </Col>
              <Col span="3" offset="8" style="text-align: center;"><a href="javascript:;" @click="chooseItem('_all')">全部分类</a></Col>
              <Col span="3" style="text-align: center;"><a href="javascript:;" @click="chooseItem('_hot')">热门分享</a></Col>
              <Col span="3" style="text-align: center;"><a href="javascript:;" @click="chooseItem('_personal')">我的分享</a></Col>
              <Col span="3" style="text-align: center;"><router-link to="/share/add">新增</router-link></Col>
            </Row>
          </div>
          <div style="padding-top: 20px;">
            <div v-for="share in shares" style="padding: 0 20px 0 20px;">
              <router-link to="">
                <Avatar size="small" src="https://i.loli.net/2017/08/21/599a521472424.jpg" />
              </router-link>
              <Tag><a @click="chooseItem(share.share_type)">{{share.share_type}}</a></Tag>
              <a @click="$router.push({path: '/share/detail', query: {share_id: share.id}})">{{share.share_desc}}</a>
              <div style="font-size: 12px;">
                <router-link :to="{path:'/iblog/author',query:{author:share.author}}">{{share.author}}</router-link>
                发布于:<Time :time="share.created_time" style="color:red;"/>&nbsp;
                更新于:<Time :time="share.last_updated_time" style="color:red;"/>&nbsp;
              </div>
              <span style="float: right;font-size: 12px;"><Time :time="share.last_updated_time"/></span>
              <Divider />
            </div>
            <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'right','margin-top': '10px'}"
                  @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>
          </div>
        </Col>
        <Col span="8" style="padding: 20px;">
          <Row>
            <Col span="8"><h6 style="color: #333;font-weight: 500;">热门分享</h6></Col>
            <Col span="4" offset="12"><a href="javascript:;">更多></a></Col>
            <TopNShare/>
            <Divider />
          </Row>
          <Row>
            <Col span="8"><h6 style="color: #333;font-weight: 500;">热门用户</h6></Col>
            <Col span="4" offset="12"><a href="javascript:;">更多></a></Col>
            <TopNUser/>
            <Divider />
          </Row>
        </Col>
      </Row>
    </div>
  </div>
</template>

<script>
  import {FilterShareList} from "../../api"
  import HotShareItem from "./HotShareItem"
  import TopNShare from "./TopNShare"
  import TopNUser from "./TopNUser"

  export default {
    name: "ShareList",
    components:{HotShareItem,TopNShare,TopNUser},
    data(){
      return {
        shares:[],
        // 当前页
        current_page:1,
        // 总页数
        total:1,
        // 每页记录数
        offset:10,
        search_type:'_all',
      }
    },
    methods:{
      chooseItem:function(item_name){
        if(this.search_type != item_name){
          this.search_type = item_name;
          this.current_page = 1;
          this.refreshShareList();
        }
      },
      refreshShareList: async function () {
        const result = await FilterShareList(this.offset, this.current_page, this.search_type);
        if(result.status == "SUCCESS"){
          this.shares = result.shares;
          this.total = result.paginator.totalcount;
        }
      },
      handleChange(page){
        this.current_page = page;
        this.refreshShareList();
      },
      handlePageSizeChange(pageSize){
        this.offset = pageSize;
        this.refreshShareList();
      },
    },
    mounted(){
      this.refreshShareList();
    }
  }
</script>

<style scoped>
  a{
    color: #155faa;
  }
  a:hover{
    color: #6cb0ca;
  }
</style>
