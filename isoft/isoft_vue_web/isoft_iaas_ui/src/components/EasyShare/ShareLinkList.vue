<template>
  <div>
    <div style="margin:15px;">
      <Row :gutter="16">
        <Col span="6">
          <div>col-6</div>
        </Col>
        <Col span="6">
          <div>col-6</div>
        </Col>
        <Col span="6">
          <div>col-6</div>
        </Col>
        <Col span="6">
          <div>col-6</div>
        </Col>
      </Row>
    </div>

    <div style="margin: 0 15px;background-color: #fff;border: 1px solid #e6e6e6;border-radius: 4px;">
      <Row>
        <Col span="16" style="padding: 0 0 20px;border-right: 1px solid #e6e6e6;">
          <div style="border-bottom: 1px solid #e6e6e6;padding: 20px;height: 62px;">
            <Row type="flex" justify="end" class="code-row-bg">
              <Col span="3"><a href="javascript:;">全部分享</a></Col>
              <Col span="3"><a href="javascript:;">最新分享</a></Col>
              <Col span="3"><a href="javascript:;">我的分享</a></Col>
              <Col span="3"><ShareLinkAdd/></Col>
            </Row>
          </div>
          <div style="padding-top: 20px;">
            <div v-for="shareLink in shareLinks" style="padding: 0 20px 0 20px;">
              <router-link to="">
                <Avatar size="small" src="https://i.loli.net/2017/08/21/599a521472424.jpg" />
              </router-link>
              <Tag>{{shareLink.share_type}}</Tag>
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

  export default {
    name: "ShareLinkList",
    components:{ShareLinkAdd},
    data(){
      return {
        shareLinks:[],
        // 当前页
        current_page:1,
        // 总页数
        total:1,
        // 每页记录数
        offset:10,
      }
    },
    methods:{
      refreshShareLinkList:async function () {
        const result = await FilterShareLinkList(this.offset, this.current_page);
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
