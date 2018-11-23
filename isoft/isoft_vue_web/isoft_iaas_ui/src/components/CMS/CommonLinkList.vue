<template>
  <div>
    <Row style="margin-bottom: 10px;">
      <Col span="12">
        <Button type="success" @click="showAddCommonLink = true">新增</Button>
      </Col>
      <Col span="12"><Input v-model="search" search enter-button placeholder="Enter something..." @on-search="refreshFilterCommonLinks"/></Col>
    </Row>

    <Modal
      v-model="showAddCommonLink"
      width="500"
      title="新增链接地址"
      footer-hide="true"
      :mask-closable="false">
      <!-- 表单正文 -->
      <Form ref="formValidate" :model="formValidate" :rules="ruleValidate" :label-width="100">
        <FormItem label="链接类型" prop="link_type">
          <Select v-model="formValidate.link_type">
            <Option v-for="item in link_type_list" :value="item.value" :key="item.value">{{ item.label }}</Option>
          </Select>
        </FormItem>
        <FormItem label="链接名称" prop="link_name">
          <Input v-model="formValidate.link_name" placeholder="请输入链接名称"></Input>
        </FormItem>
        <FormItem label="链接地址" prop="link_addr">
          <Input v-model="formValidate.link_addr" placeholder="请输入链接地址"></Input>
        </FormItem>
        <FormItem>
          <Button type="primary" @click="handleSubmit('formValidate')">Submit</Button>
          <Button @click="handleReset('formValidate')" style="margin-left: 8px">Reset</Button>
        </FormItem>
      </Form>
    </Modal>


    <Table :columns="columns1" :data="commonLinks" size="small" height="450"></Table>
    <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
          @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>
  </div>
</template>

<script>
  import {FilterCommonLinks} from "../../api"
  import {AddCommonLink} from "../../api"

  export default {
    name: "CommonLinkList",
    data(){
      return {
        showAddCommonLink:false,
        commonLinks:[],
        // 当前页
        current_page:1,
        // 总页数
        total:1,
        // 每页记录数
        offset:10,
        search:"",
        link_type_list:[{value: 'friend_link',label: 'friend_link 友情链接'},{value: 'hot_project',label: 'hot_project 热门项目'}],
        columns1: [
          {
            title: 'id',
            key: 'id'
          },
          {
            title: '链接类型',
            key: 'link_type'
          },
          {
            title: '链接名称',
            key: 'link_name'
          },
          {
            title: '链接地址',
            key: 'link_addr'
          }
        ],
        formValidate: {
          link_type: '',
          link_name: '',
          link_addr: '',
        },
        ruleValidate: {
          link_type: [
            { required: true, message: '链接类型不能为空', trigger: 'blur' }
          ],
          link_name: [
            { required: true, message: '链接名称不能为空', trigger: 'blur' }
          ],
          link_addr: [
            { required: true, message: '链接地址不能为空', trigger: 'blur' }
          ]
        }
      }
    },
    methods:{
      async refreshFilterCommonLinks(){
        const result = await FilterCommonLinks(this.offset,this.current_page,this.search);
        if(result.status == "SUCCESS"){
          this.total = result.paginator.totalcount;
          this.commonLinks = result.commonLinks;
        }
      },
      handleSubmit (name) {
        this.$refs[name].validate(async (valid) => {
          if (valid) {
            const result = await AddCommonLink(this.formValidate.link_type, this.formValidate.link_name, this.formValidate.link_addr);
            if(result.status == "SUCCESS"){
              this.showAddConfiguration = false;
              this.refreshFilterCommonLinks();
              this.showAddCommonLink = false;
            }else{
              this.$Message.error('提交失败!');
            }
          } else {
            this.$Message.error('验证失败!');
          }
        })
      },
      handleReset (name) {
        this.$refs[name].resetFields();
      },
      handleChange(page){
        this.current_page = page;
        this.refreshFilterCommonLinks();
      },
      handlePageSizeChange(pageSize){
        this.offset = pageSize;
        this.refreshFilterCommonLinks();
      },
    },
    mounted:function () {
      this.refreshFilterCommonLinks();
    }
  }
</script>

<style scoped>

</style>
