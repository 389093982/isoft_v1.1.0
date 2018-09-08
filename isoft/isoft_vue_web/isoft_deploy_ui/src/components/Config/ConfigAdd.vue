<template>
<div>
  <Button type="success" @click="modal1 = true">新增</Button>
  <Modal
    v-model="modal1"
    title="新增/编辑配置信息"
    :footer-hide="true"
    :mask-closable="false">       <!-- 是否允许点击遮罩层关闭 -->
    <div>
      <!-- 表单正文 -->
      <Form ref="formValidate" :model="formValidate" :rules="ruleValidate" :label-width="80">
        <FormItem label="环境名称" prop="env_ids">
          <Select v-model="formValidate.env_ids" filterable multiple>
            <Option v-for="envInfo in envInfos" :value="envInfo.id" :key="envInfo.id">
              {{ envInfo.env_name }} - [ {{ envInfo.env_ip }} ]
            </Option>
          </Select>
        </FormItem>
        <FormItem label="环境变量" prop="env_property">
          <Input v-model="formValidate.env_property" placeholder="环境变量"></Input>
        </FormItem>
        <FormItem label="变量值" prop="env_value">
          <Input v-model="formValidate.env_value" placeholder="变量值"></Input>
        </FormItem>
        <FormItem>
          <Button type="primary" @click="handleSubmit('formValidate')">Submit</Button>
          <Button @click="handleReset('formValidate')" style="margin-left: 8px">Reset</Button>
        </FormItem>
      </Form>
    </div>
  </Modal>
</div>
</template>

<script>
  import {mapState} from 'vuex'
  import {ConfigEdit} from '../../api'

  export default {
    data () {
      return {
        modal1: false,    // 遮罩层
        formValidate: {
          env_ids: '',
          env_property: '',
          env_value: ''
        },
        ruleValidate: {
          env_ids: [
            { required: true, type: 'array', message: '环境名称不能为空', trigger: 'blur' }
          ],
          env_property: [
            { required: true, message: '环境变量不能为空', trigger: 'blur' }
          ],
          env_value: [
            { required: true, message: '变量值不能为空', trigger: 'blur' }
          ]
        }
      }
    },
    methods: {
      // 关闭模态对话框
      closeModalDialog (){
        this.modal1 = false;
      },
      handleSubmit (name) {
        var _this = this;
        this.$refs[name].validate(async (valid) => {
          if (valid) {
            const result = await ConfigEdit(_this.formValidate.env_ids.join(','),
              _this.formValidate.env_property,_this.formValidate.env_value);
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
        });
      },
      handleReset (name) {
        this.$refs[name].resetFields();
      }
    },
    computed:{
      ...mapState(['envInfos']),
    }
  }
</script>

<style scoped>

</style>
