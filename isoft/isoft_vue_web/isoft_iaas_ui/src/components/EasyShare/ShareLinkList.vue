<template>
  <div>
    <!-- 热门分类 -->
    <HotShareLinkItem @chooseItem="chooseItem"/>

    <div style="margin: 0 15px;background-color: #fff;border: 1px solid #e6e6e6;border-radius: 4px;">
      <Row>
        <Col span="16" style="padding: 0 0 20px;border-right: 1px solid #e6e6e6;">
          <div style="border-bottom: 1px solid #e6e6e6;padding: 20px;height: 62px;">
            <Row>
              <Col span="4" style="text-align: center;font-size: 20px;color: #333;">
                <span v-if="share_type==='all'">全部分类</span>
                <span v-else>{{share_type}}</span>
              </Col>
              <Col span="3" offset="8" style="text-align: center;"><a href="javascript:;" @click="filterToggle('all')">全部分类</a></Col>
              <Col span="3" style="text-align: center;"><a href="javascript:;" @click="filterToggle('hot')">热门分享</a></Col>
              <Col span="3" style="text-align: center;"><a href="javascript:;" @click="filterToggle('my')">我的分享</a></Col>
              <Col span="3" style="text-align: center;"><ShareLinkAdd/></Col>
            </Row>
          </div>
          <div style="padding-top: 20px;">
            <div v-for="shareLink in shareLinks" style="padding: 0 20px 0 20px;">
              <router-link to="">
                <Avatar size="small" src="https://i.loli.net/2017/08/21/599a521472424.jpg" />
              </router-link>
              <Tag><a @click="chooseItem(shareLink.share_type)">{{shareLink.share_type}}</a></Tag>
              <a href="shareLink.link_href">{{shareLink.link_href}}</a>
              <span style="float: right;font-size: 12px;"><Time :time="shareLink.last_updated_time"/></span>
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
            <Divider />
          </Row>
          <Row>
            <Col span="8"><h6 style="color: #333;font-weight: 500;">热门用户</h6></Col>
            <Col span="4" offset="12"><a href="javascript:;">更多></a></Col>
            <Divider />
          </Row>
        </Col>
      </Row>
    </div>
  </div>
</template>

<script>
  import {FilterShareLinkList} from "../../api"
  import ShareLinkAdd from "./ShareLinkAdd.vue"
  import HotShareLinkItem from "./HotShareLinkItem.vue"

  export default {
    name: "ShareLinkList",
    components:{ShareLinkAdd,HotShareLinkItem},
    data(){
      return {
        shareLinks:[],
        // 当前页
        current_page:1,
        // 总页数
        total:1,
        // 每页记录数
        offset:10,
        share_type:'all',
      }
    },
    methods:{
      filterToggle:function(filter){
        if(filter == "all"){
          this.chooseItem("all");
        }
      },
      chooseItem:function(item_name){
        if(this.share_type != item_name){
          this.share_type = item_name;
          this.current_page = 1;
          this.refreshShareLinkList();
        }
      },
      refreshShareLinkList:async function () {
        const result = await FilterShareLinkList(this.offset, this.current_page, this.share_type);
        if(result.status == "SUCCESS"){
          this.shareLinks = result.shareLinks;
          this.total = result.paginator.totalcount;
        }
      },
      handleChange(page){
        this.current_page = page;
        this.refreshShareLinkList();
      },
      handlePageSizeChange(pageSize){
        this.offset = pageSize;
        this.refreshShareLinkList();
      },
    },
    mounted(){
      this.refreshShareLinkList();
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
