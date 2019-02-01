<template>
  <LeftMenu>
    <div style="margin: 10px;">
      <ISimpleLeftRightRow>
        <!-- left 插槽部分 -->
        <!-- 按钮触发模态框 -->
        <!-- ref 的作用是为了在其它地方方便的获取到当前子组件 -->
        <ISimpleBtnTriggerModal ref="triggerModal" slot="left" btn-text="新增" modal-title="新增/编辑系统地址信息" :modal-width="600">
          <!-- 表单添加系统注册信息 -->
          <Form ref="formValidate" :model="formValidate" :rules="ruleValidate" :label-width="80">
            <Row :gutter="6">
              <Col span="12">
                <FormItem label="注册地址" prop="app_address">
                  <Input v-model="formValidate.app_address" placeholder="请输入系统注册地址"></Input>
                </FormItem>
              </Col>
              <Col span="12">
                <Button type="success" @click="handleSubmit('formValidate')" style="margin-right: 6px">Submit</Button>
                <Button type="warning" @click="handleReset('formValidate')" style="margin-right: 6px">Reset</Button>
              </Col>
            </Row>
          </Form>
        </ISimpleBtnTriggerModal>
        <!-- right 插槽部分 -->
        <ISimpleSearch slot="right" @handleSimpleSearch="handleSearch"/>
      </ISimpleLeftRightRow>

      <Table :columns="columns1" :data="appRegisters" size="small"></Table>
      <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
            @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>
    </div>
  </LeftMenu>
</template>

<script>
  import {AppRegisterList} from "../../api"
  import {AddAppRegister} from "../../api"
  import LeftMenu from "./LeftMenu"
  import ISimpleBtnTriggerModal from "../Common/modal/ISimpleBtnTriggerModal"
  import ISimpleLeftRightRow from "../Common/layout/ISimpleLeftRightRow"
  import ISimpleSearch from "../Common/search/ISimpleSearch"

  export default {
    name: "AppRegist",
    components: {LeftMenu,ISimpleLeftRightRow,ISimpleSearch,ISimpleBtnTriggerModal},
    data(){
      return {
        // 当前页
        current_page:1,
        // 总页数
        total:1,
        // 每页记录数
        offset:10,
        // 搜索条件
        search:"",
        appRegisters: [],
        showFormModal:false,
        formValidate: {
          app_address: '',
        },
        ruleValidate: {
          app_address: [
            { required: true, message: '系统注册地址不能为空', trigger: 'blur' }
          ]
        },
        columns1: [
          {
            title: '注册地址',
            key: 'app_address',
            width:300
          },
          {
            title: '创建人',
            key: 'created_by',
            width:120
          },
          {
            title: '创建时间',
            key: 'created_time',
            width:200
          },
          {
            title: '修改人',
            key: 'last_updated_by',
            width:120
          },
          {
            title: '修改时间',
            key: 'last_updated_time',
          },
        ],
      }
    },
    methods:{
      refreshAppRegistList: async function(){
        const result = await AppRegisterList(this.offset, this.current_page, this.search);
        if(result.status=="SUCCESS"){
          this.appRegisters = result.appRegisters;
          this.total = result.paginator.totalcount;
        }
      },
      handleChange(page){
        this.current_page = page;
        this.refreshAppRegistList();
      },
      handlePageSizeChange(pageSize){
        this.offset = pageSize;
        this.refreshAppRegistList();
      },
      handleSubmit (name) {
        var _this = this;
        this.$refs[name].validate(async (valid) => {
          if (valid) {
            const result = await AddAppRegister(_this.formValidate.app_address);
            if(result.status == "SUCCESS"){
              _this.$Message.success('提交成功!');
              // 调用子组件隐藏 modal (this.refs.xxx.子组件定义的方法())
              this.refs.triggerModal.hideModal();
              // 关闭模态对话框
              _this.showFormModal = false;
              _this.refreshAppRegistList();
            }else{
              _this.$Message.error('提交失败!');
            }
          } else {
            _this.$Message.error('验证失败!');
          }
        })
      },
      handleReset (name) {
        this.$refs[name].resetFields();
      },
      handleSearch(data){
        this.offset = 10;
        this.current_page = 1;
        this.search = data;
        this.refreshAppRegistList();
      }
    },
    mounted: function () {
      this.refreshAppRegistList();
    },
  }
</script>

<style scoped>

</style>
