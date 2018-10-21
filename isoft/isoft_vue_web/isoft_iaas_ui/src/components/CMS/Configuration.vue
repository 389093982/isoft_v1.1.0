<template>
  <div>
    <Row style="margin-bottom: 10px;">
      <Col span="12">
        <Button type="success" @click="showAddConfiguration = true">新增</Button>
      </Col>
      <Col span="12"><Input v-model="search" search enter-button placeholder="Enter something..." @on-search="refreshConfigurations"/></Col>
    </Row>

    <Modal
      v-model="showAddConfiguration"
      width="500"
      title="新增配置项"
      :mask-closable="false">
      <!-- 表单正文 -->
      <Form ref="formValidate" :model="formValidate" :rules="ruleValidate" :label-width="100">
        <FormItem label="父配置项 id" prop="parent_id">
          <Input v-model="formValidate.parent_id" placeholder="无父配置项 id 可不填"></Input>
        </FormItem>
        <FormItem label="配置项名称" prop="configuration_name">
          <Input v-model="formValidate.configuration_name" placeholder="请输入配置项名称"></Input>
        </FormItem>
        <FormItem label="配置项值" prop="configuration_value">
          <Input v-model="formValidate.configuration_value" placeholder="请输入配置项值"></Input>
        </FormItem>
        <FormItem>
          <Button type="primary" @click="handleSubmit('formValidate')">Submit</Button>
          <Button @click="handleReset('formValidate')" style="margin-left: 8px">Reset</Button>
        </FormItem>
      </Form>
    </Modal>

    <Table :columns="columns1" :data="configurations" size="small" height="450"></Table>
    <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
          @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>
  </div>
</template>

<script>
  import {AddConfiguration} from "../../api"
  import {FilterConfigurations} from "../../api"

  export default {
    name: "Configuration",
    data(){
      return {
        showAddConfiguration:false,
        configurations:[],
        // 当前页
        current_page:1,
        // 总页数
        total:1,
        // 每页记录数
        offset:10,
        search:"",
        columns1: [
          {
            title: 'id',
            key: 'id'
          },
          {
            title: '父级id',
            key: 'parent_id'
          },
          {
            title: '配置项名称',
            key: 'configuration_name'
          },
          {
            title: '配置项值',
            key: 'configuration_value'
          },
          {
            title: '状态',
            key: 'status'
          }
        ],
        formValidate: {
          parent_id: "",
          configuration_name: '',
          configuration_value: '',
        },
        ruleValidate: {
          configuration_name: [
            { required: true, message: '配置项名称不能为空', trigger: 'blur' }
          ],
          configuration_value: [
            { required: true, message: '配置项值不能为空', trigger: 'blur' }
          ]
        }
      }
    },
    methods:{
      refreshConfigurations:async function(){
        const result = await FilterConfigurations(this.search, this.offset, this.current_page);
        if(result.status == "SUCCESS"){
          this.total = result.paginator.totalcount;
          this.configurations = result.configurations;
        }
      },
      handleSubmit (name) {
        var _this = this;
        this.$refs[name].validate(async (valid) => {
          if (valid) {
            const result = await AddConfiguration(_this.formValidate.parent_id,
              _this.formValidate.configuration_name, _this.formValidate.configuration_value);
            if(result.status == "SUCCESS"){
              this.showAddConfiguration = false;
              _this.refreshConfigurations();
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
      handleChange(page){
        this.current_page = page;
        this.refreshConfigurations();
      },
      handlePageSizeChange(pageSize){
        this.offset = pageSize;
        this.refreshConfigurations();
      },
    },
    mounted:function () {
      this.refreshConfigurations();
    }
  }
</script>

<style scoped>

</style>
