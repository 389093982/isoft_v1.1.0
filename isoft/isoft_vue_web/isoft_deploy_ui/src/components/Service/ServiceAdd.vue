<template>
<div>
  <Button type="success" @click="showFormModal = true">新增</Button>

  <Modal
    v-model="showFormModal"
    width="850"
    title="新增/编辑服务信息"
    :footer-hide="true"
    :mask-closable="false">
    <div>
      <!-- 表单正文 -->
      <Form ref="formValidate" :model="formValidate" :rules="ruleValidate" :label-width="80">
        <Row>
          <Col span="12">
            <FormItem label="环境名称" prop="env_ids">
              <Select v-model="formValidate.env_ids" filterable multiple>
                <Option v-for="envInfo in envInfos" :value="envInfo.id" :key="envInfo.id">
                  {{ envInfo.env_name }} - [ {{ envInfo.env_ip }} ]
                </Option>
              </Select>
            </FormItem>
          </Col>
          <Col span="12">
            <FormItem label="服务名称" prop="service_name">
              <Input v-model="formValidate.service_name" placeholder="请输入服务名称"></Input>
            </FormItem>
          </Col>
        </Row>

        <Row>
          <Col span="12">
            <FormItem label="服务类型" prop="service_type">
              <Input v-model="formValidate.service_type" readonly placeholder="请输入服务类型"></Input>
            </FormItem>
          </Col>
          <Col span="12">
            <FormItem label="端口号" prop="service_port">
              <Input v-model="formValidate.service_port" placeholder="请输入端口号"></Input>
            </FormItem>
          </Col>
        </Row>

        <Row>
          <Col span="12" v-if="formValidate.service_type=='beego' || formValidate.service_type=='api'">
            <FormItem label="运行模式" prop="run_mode">
              <Input v-model="formValidate.run_mode" placeholder="请输入运行模式"></Input>
            </FormItem>
          </Col>
          <Col span="12" v-if="formValidate.service_type=='beego' || formValidate.service_type=='api'">
            <FormItem label="部署包名" prop="package_name">
              <Input v-model="formValidate.package_name" placeholder="请输入部署包名"></Input>
            </FormItem>
          </Col>
        </Row>

        <Row>
          <Col span="12" v-if="formValidate.service_type=='mysql'">
            <FormItem label="root密码" prop="mysql_root_pwd">
              <Input v-model="formValidate.mysql_root_pwd" placeholder="请输入 mysql 数据库 root 密码"></Input>
            </FormItem>
          </Col>
        </Row>

        <Row>
          <Col span="12">
            <FormItem>
              <Button type="success" @click="handleSubmit('formValidate')" style="margin-right: 8px">Submit</Button>
              <Button type="warning" @click="handleReset('formValidate')" style="margin-right: 8px">Reset</Button>
              <Button @click="showFieldDetailModal = true">不太了解,查看字段含义</Button>
            </FormItem>
          </Col>
        </Row>
      </Form>
    </div>
  </Modal>



  <Modal
    v-model="showFieldDetailModal"
    title="字段含义">
    <p>环境名称:必填</p>
    <p>服务名称:必填</p>
    <p>服务类型:必填</p>
    <p>部署包名:选填,主要指web应用类型的软件包名称</p>
    <p>运行模式:选填,指定部署时使用哪套配置文件</p>
    <p>端口号:选填,不填情况下使用软件包默认端口,单个环境端口唯一</p>
  </Modal>
</div>
</template>

<script>
  import {EnvEdit} from '../../api'
  import {mapState} from 'vuex'
  import {ServiceEdit} from '../../api'

  export default {
    data () {
      return {
        showFieldDetailModal: false,
        showFormModal: false,
        formValidate: {
          env_ids: '',
          service_name: '',
          service_type: this.$route.query.service_type,
          mysql_root_pwd: '',
          package_name: '',
          run_mode: '',
          service_port: ''
        },
        ruleValidate: {
          env_ids: [
            { required: true, type: 'array', message: '环境名称不能为空', trigger: 'blur' }
          ],
          service_name: [
            { required: true, message: '服务名称不能为空', trigger: 'blur' }
          ],
          service_type: [
            { required: true, message: '服务类型不能为空', trigger: 'blur' }
          ],
          service_port: [
            { required: true, message: '端口号不能为空', trigger: 'blur' }
          ]
        },
      }
    },
    methods: {
      handleSubmit (name) {
        var _this = this;
        this.$refs[name].validate(async (valid) => {
          if (valid) {
            const result = await ServiceEdit(_this.formValidate.env_ids.join(','),
              _this.formValidate.service_name, _this.formValidate.service_type,
              _this.formValidate.package_name, _this.formValidate.run_mode,
              _this.formValidate.service_port, _this.formValidate.mysql_root_pwd);
            if(result.status == "SUCCESS"){
              _this.$Message.success('提交成功!');
              // 关闭模态对话框
              _this.showFormModal = false;
              _this.$router.go(0);     // 页面刷新,等价于 location.reload()
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
      }
    },
    computed:{
      ...mapState(['envInfos']),
    },
    watch:{
      // 路由变化更新 service_type
      '$route' (to, from) {
        this.formValidate.service_type = this.$route.query.service_type;
      }
    }
  }
</script>

<style scoped>

</style>
